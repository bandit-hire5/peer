package middlewares

import (
	"context"
	"net/http"

	ws "github.com/bandit/peer/websocket"
)

const ctxLedgerKey = "server"

func AddContextServer(server *ws.Server, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxLedgerKey, server)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetServer(r *http.Request) *ws.Server {
	return r.Context().Value(ctxLedgerKey).(*ws.Server)
}
