package middleware

import (
	"fmt"
	"main.go/constant"
	"main.go/errors"
	"main.go/logger"
	"net/http"
)

func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.InfoLogger.Println("Error from the recovery method: ", err)
				panicErr := errors.AppError{
					Message: fmt.Sprint(err),
				}
				constant.WriteResponse(w, http.StatusInternalServerError, panicErr)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
