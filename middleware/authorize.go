package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"templates/infrastructure"
	"templates/model"

	"github.com/google/uuid"
)

var (
	AuthenticationNats Authentication
)

func init() {
	authen := NewAuthentication()
	AuthenticationNats = authen
}

type Authentication interface {
	Authorizer() func(next http.Handler) http.Handler
	GetUser() (user model.User)
}

type authentication struct {
	Nats infrastructure.EventPublisher
	User model.User
}

type AuthenticationMeta struct {
	Token  string
	Method string
	Path   string
}

type AuthenticationResult struct {
	Code    int
	Message string
	Data    model.User
}

func (au *authentication) Authorizer() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			authenModel := AuthenticationMeta{
				Token:  r.Header.Get("authorization"),
				Method: r.Method,
				Path:   r.URL.Path,
			}
			res, err := au.Nats.Request(infrastructure.NATSAuthSubject, authenModel)
			if err != nil {
				infrastructure.ErrLog.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			var resM AuthenticationResult
			err = json.Unmarshal(res, &resM)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			if resM.Code == 200 {
				ctx := context.WithValue(r.Context(), "user", resM.Data)
				//save user
				au.User = resM.Data
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

func (au *authentication) GetUser() (user model.User) {
	return au.User
}

func NewAuthentication() Authentication {
	//init Nats
	authenClientID := "authen-service-" + uuid.NewString()
	nats := infrastructure.NewNatsPublisher(infrastructure.NATSHostport, infrastructure.NATSClusterID, authenClientID)
	return &authentication{
		Nats: nats,
	}
}
