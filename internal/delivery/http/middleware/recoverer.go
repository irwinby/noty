package middleware

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/hereisajvi/noty/internal/delivery/http/status"
)

type Recoverer struct{}

func NewRecoverer() *Recoverer {
	return &Recoverer{}
}

func (r *Recoverer) Recover() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err, _ := recover().(error)
				if err != nil {
					zap.L().Error("panic caught with recover", zap.Error(err))
					status.Error(w, http.StatusInternalServerError, err.Error())

					return
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
