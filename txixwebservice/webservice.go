package webservice

import (
	"context"

	"github.com/txix-open/isp-kit/http"

	"github.com/txix-open/isp-kit/http/endpoint"
	"github.com/txix-open/isp-kit/http/router"
	"github.com/txix-open/isp-kit/log"
)

type serviceController interface {
	Route(wrapper *endpoint.Wrapper, router *router.Router)
}

type Server struct {
	http        *http.Server
	wrapper     endpoint.Wrapper
	router      *router.Router
	controllers []serviceController
	address     string
}

func New(ctx context.Context, logger *log.Adapter, address string) *Server {
	s := Server{
		http:        http.NewServer(logger),
		wrapper:     endpoint.DefaultWrapper(logger),
		router:      router.New(),
		controllers: make([]serviceController, 0, 1),
		address:     address,
	}
	s.http.Upgrade(s.router)
	return &s
}

func (s *Server) Add(controller ...serviceController) {
	s.controllers = append(s.controllers, controller...)
}

func (s *Server) Run(ctx context.Context) error {
	for _, controller := range s.controllers {
		controller.Route(&s.wrapper, s.router)
	}
	s.wrapper.Logger.Info(ctx, "Listen on address "+s.address)
	return s.http.ListenAndServe(s.address)
}

func (s *Server) Close() error {
	return nil
}
