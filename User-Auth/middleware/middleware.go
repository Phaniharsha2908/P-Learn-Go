package middleware

import (
	"fmt"
	"net/http"
	"user-auth/model"
)

type RequireUser struct {
	*model.UserService
}

func (mw *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")

		if err != nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		user, err := mw.ByRememberToken(cookie.Value)
		fmt.Println("*****", cookie.Value)

		if err != nil {

			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		fmt.Println("User Found", user)
		next(w, r)

	})
}
