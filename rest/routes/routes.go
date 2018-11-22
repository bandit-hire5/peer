package routes

import (
	"encoding/json"
	"net/http"

	ws "github.com/bandit/peer/websocket"
	mw "github.com/bandit/peer/rest/middlewares"

	"github.com/gorilla/mux"
)

func Router(server *ws.Server) *mux.Router {
	r := mux.NewRouter()

	//r.HandleFunc("/block", m.AddContext(context, NewBlock)).Methods("POST")
	r.Handle("/message", mw.AddContextServer(server, http.HandlerFunc(NewMessage))).Methods("POST")

	return r
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
