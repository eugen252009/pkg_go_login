package login

import (
	"context"
	"net/http"

	log "github.com/eugen252009/pkg_go_logger/logger"
)

func Register(fn http.Handler) http.Handler {
	contextdata := &AuthInfo{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Log(log.LogLevelError, LogLoginRegister, ErrorParseForm.Message)
			http.Error(w, ErrorParseForm.Message, ErrorParseForm.Code)
			return
		}

		contextdata.Email = r.Form.Get("email")
		contextdata.Password = r.Form.Get("password")

		ctx := context.WithValue(r.Context(), contextdata, contextdata)

		fn.ServeHTTP(w, r.WithContext(ctx))
	})
}
