package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"templates/infrastructure"
	"templates/model"
)

type AuthorizeMiddleware interface {
	Authorizer() func(next http.Handler) http.Handler
	//GetEnforcer() *casbin.Enforcer
	GetUser() *model.User
}
type AuthenticationMeta struct {
	Token  string
	Method string
	Path   string
}
type authorizedMiddleware struct {
	Nats infrastructure.EventPublisher
}
type AuthenticationResult struct {
	Code    int
	Message string
	Data    model.User
}

var user *model.User

func (a *authorizedMiddleware) Authorizer() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			authenModel := AuthenticationMeta{
				Token:  r.Header.Get("authorization"),
				Method: r.Method,
				Path:   r.URL.Path,
			}
			res, err := a.Nats.Request(infrastructure.NATSAuthSubject, authenModel)
			if err != nil {
				infrastructure.ErrLog.Printf("Request AUTH error: %+v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			var resM AuthenticationResult

			err = json.Unmarshal(res, &resM)
			user = &resM.Data

			if err != nil {
				infrastructure.ErrLog.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			if resM.Code == 200 {
				ctx := context.WithValue(r.Context(), "user", resM.Data)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusForbidden)
				http.Error(w, http.StatusText(resM.Code), resM.Code)
				return
			}

		}

		return http.HandlerFunc(fn)
	}
}

func (a *authorizedMiddleware) GetUser() *model.User {
	return user
}

func NewAuthorizeMiddleware() AuthorizeMiddleware {
	return &authorizedMiddleware{
		Nats: infrastructure.NATSConnection,
	}
}
