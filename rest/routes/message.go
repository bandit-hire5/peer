package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bandit/blockchain-core"

	mw "github.com/bandit/peer/rest/middlewares"
)

func NewMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	server := mw.GetServer(r)

	var data core.Message

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	data.Type = "message"

	err := server.SendMessage(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, nil)
}
