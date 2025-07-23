package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos_sample/app/gateway/internal/conf"
	"kratos_sample/app/gateway/internal/router"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, routerR *router.Router) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.GatewayHttp.Network != "" {
		opts = append(opts, http.Network(c.GatewayHttp.Network))
	}
	if c.GatewayHttp.Addr != "" {
		opts = append(opts, http.Address(c.GatewayHttp.Addr))
	}
	if c.GatewayHttp.Timeout != nil {
		opts = append(opts, http.Timeout(c.GatewayHttp.Timeout.AsDuration()))
	}
	//srv := http.NewServer(opts...)
	//v1.RegisterGatewayHTTPServer(srv, greeter)

	routerR.SetupRoutes()
	//printRoutes(routerR.Engine)
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", routerR.Engine)

	return srv
}

func printRoutes(r *gin.Engine) {
	fmt.Println("Registered Routes:")
	for _, route := range r.Routes() {
		fmt.Printf("Method: %s, Path: %s, Handler: %s\n", route.Method, route.Path, route.Handler)
	}
}
