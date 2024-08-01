package main

import (
	"encoding/json"
	"net/http"

	"github.com/alikrugl/go-web-server/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	token, err := auth.GetApiKeyToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find Api Token")
		return
	}

	if token != cfg.polkaSecret {
		respondWithError(w, http.StatusUnauthorized, "Wrong ApiKey")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
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
