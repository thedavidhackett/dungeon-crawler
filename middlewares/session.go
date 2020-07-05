package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
)


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


		next.ServeHTTP(w, r)

    })
}
