package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/m-nt/gomod/auth/application"
	"github.com/m-nt/gomod/auth/infrastructure"
	httpiface "github.com/m-nt/gomod/auth/interfaces/http"
)

type Config struct {
	Secret       string
	UserProvider application.UserProvider
}

func RegisterHTTP(r gin.IRouter, cfg Config) *application.Service {

	jwtProvider := infrastructure.NewJWTProvider(cfg.Secret)

	svc := application.New(cfg.UserProvider, jwtProvider)

	handler := httpiface.New(svc)
	httpiface.Register(r, handler)

	return svc
}
