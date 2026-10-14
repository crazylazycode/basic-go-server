package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	const maxChirpLength = 140

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	notValidWords := []string{"kerfuffle", "sharbert", "fornax"}

	splitBody := strings.Split(params.Body, " ")

	for i, ch := range splitBody {
		for _, word := range notValidWords {
			if strings.EqualFold(ch, word) {
				splitBody[i] = "****"
			}
		}
	}

	cleanedBody := strings.Join(splitBody, " ")

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBody,
	})

}
