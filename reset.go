package main

import (
	"context"
	"errors"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "reset is only allowed in dev environment", errors.New("cannot do this in other then dev env"))
		return
	}

	err := cfg.db.Reset(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't reset database", err)
		return
	}

	cfg.fileserverHits.Store(0)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state\n"))

}
