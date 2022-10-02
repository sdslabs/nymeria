package middleware

import (
    "context"
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    client "github.com/ory/kratos-client-go"
)
type kratosMiddleware struct {
    client *client.APIClient
}

func NewMiddleware() *kratosMiddleware {
    configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4433", // Kratos Admin API
        },
    }
    return &kratosMiddleware{
        client: client.NewAPIClient(configuration),
    }
}

func NewAdminMiddleware() *client.APIClient {
    configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Public API
        },
    }

    apiClient := client.NewAPIClient(configuration)
    return apiClient
}

func (k *kratosMiddleware) Session() gin.HandlerFunc {
    return func(c *gin.Context) {
        session, err := k.validateSession(c.Request)
        if err != nil {
            c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4455/login")
            return
        }
        if !*session.Active {
            c.Redirect(http.StatusMovedPermanently, "http://your_endpoint")
            return
        }
        c.Next()
    }
}

func (k *kratosMiddleware) validateSession(r *http.Request) (*client.Session, error) {
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
