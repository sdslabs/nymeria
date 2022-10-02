package middleware

import (
	"context"
	"errors"
	"net/http"
)

// func isAdmin(id string) (bool, error) {

// }

func (k *kratosMiddleware) validateAdmin(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		return false, err
	}
	if cookie == nil {
		return false, errors.New("no session found in cookie")
	}
	_, _, err = k.client.V0alpha2Api.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return false, err
	}

	admin, err := true, nil //isAdmin(resp.Identity.Id)

	if err != nil {
		return false, err
	}

	// if admin != true {
	// 	return false, errors.New("user is not an admin")
	// }

	return admin, nil
}
