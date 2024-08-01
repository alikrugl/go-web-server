package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	if params.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "Webhook event not supported")
		return
	}

	_, err = cfg.DB.UpdateMembershipUser(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find user")
		return
	}

	type response struct {
		Upgraded bool `json:"upgraded"`
	}

	respondWithJSON(w, http.StatusNoContent, response{
		Upgraded: true,
	})
}
