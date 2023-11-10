package server

import (
	"application/internal/utils"
	"encoding/json"
	"errors"
	"net/http"
	"runtime/debug"
)

func (m *GorillaMuxMiddleware) RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Handle the panic here, log the error, or send an appropriate response to the client.
				m.logger.Error("Panic", "err", err, "ctx", utils.LogContext(r.Context()), "trace", debug.Stack())
				apperror := utils.NewAppError(http.StatusInternalServerError, 2000, "Internal Server Error")
				switch v := err.(type) {
				case error:
					apperror = apperror.AddError(v)
				case string:
					apperror = apperror.AddError(errors.New(v))
				}
				w.WriteHeader(http.StatusInternalServerError)
				bt, err := json.Marshal(apperror)
				if err != nil {
					m.logger.Error("Panic", "err", err, "ctx", utils.LogContext(r.Context()), "trace", debug.Stack())
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				http.Error(w, string(bt), http.StatusInternalServerError)

			}
		}()
		next.ServeHTTP(w, r)
	})
}
