package users

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/m-nt/gomod/users/application"
	"github.com/m-nt/gomod/users/infrastructure"
	grpciface "github.com/m-nt/gomod/users/interfaces/grpc"
	httpiface "github.com/m-nt/gomod/users/interfaces/http"
	wsiface "github.com/m-nt/gomod/users/interfaces/ws"
)

func RegisterHTTP(r gin.IRouter, dsn string, m ...gin.HandlerFunc) {
	client := infrastructure.NewEntClient(dsn)
	svc := application.New(client)
	h := httpiface.New(svc)
	httpiface.Register(r, h, m...)
}

func RegisterGRPC(server *grpc.Server, dsn string) {
	client := infrastructure.NewEntClient(dsn)
	svc := application.New(client)
	grpciface.RegisterUsersServiceServer(server, grpciface.New(svc))
}

func RegisterWS(r gin.IRouter, reg *wsiface.Registry) {
	wsiface.Register(r, reg)
}
