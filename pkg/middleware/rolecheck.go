package middleware

import (
    "context"
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    client "github.com/ory/kratos-client-go"
)

func (k *kratosMiddleware) validateAdmin(r *http.Request) (*client.Session, error) {
    cookie, err := r.Cookie("ory_kratos_session")
    if err != nil {
        return nil, err
    }
    if cookie == nil {
        return nil, errors.New("no session found in cookie")
    }
    resp, _, err := k.client.V0alpha2Api.ToSession(context.Background()).Cookie(cookie.String()).Execute()
    if err != nil {
        return nil, err
    }
    return resp, nil
}

