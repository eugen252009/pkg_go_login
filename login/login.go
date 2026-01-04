package login

import (
	"context"
	log "github.com/eugen252009/pkg_go_logger/logger"
	"net/http"
)

func Login(fn http.Handler) http.Handler {
	contextdata := &AuthInfo{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Log(log.LogLevelError, LogLoginLogin, ErrorParseForm.Message)
			http.Error(w, ErrorParseForm.Message, ErrorParseForm.Code)
			return
		}

		contextdata.Email = r.Form.Get("email")
		contextdata.Password = r.Form.Get("password")

		ctx := context.WithValue(r.Context(), contextdata, contextdata)

		fn.ServeHTTP(w, r.WithContext(ctx))
	})
}
