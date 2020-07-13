package middlewares

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

type contextKey int

//AuthenticatedUserKey ..
const AuthenticatedUserKey contextKey = 0

//SessionMiddleware ...
type SessionMiddleware struct {
	store *sessions.CookieStore
}

//NewSessionMiddleware ...
func NewSessionMiddleware(store *sessions.CookieStore) *SessionMiddleware {
	return &SessionMiddleware{
		store: store,
	}
}

//Middleware ...
func (sm *SessionMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sm.store.Get(r, "cookie-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user, _ := session.Values["user"]

		ctxWithUser := context.WithValue(r.Context(), AuthenticatedUserKey, user)
		//create a new request using that new context
		rWithUser := r.WithContext(ctxWithUser)
		//call the real handler, passing the new request

		next.ServeHTTP(w, rWithUser)

    })
}
