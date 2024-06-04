package handers

import (
	"IntraCat/cmd/server/env"
	"net/http"
)

func IndexPageGet(e *env.Env) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("INTRA-CAT"))
    }
}
