package echo

import (
	"fmt"
	"information-collector/cmd/api-server/web/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

var defaultMiddleware = []echo.MiddlewareFunc{
	middlewares.TransFormContext,
}

type WebServer struct {
	routeGroups []*RouteGroups
	server      *echo.Echo
	listenPort  string
}

// NewServer - new a echo server
func NewServer() *WebServer {
	return &WebServer{
		routeGroups: []*RouteGroups{},
		server:      echo.New(),
	}
}

// WithValidator - set validator for echo server
func (s *WebServer) WithValidator(validator echo.Validator) *WebServer {
	s.server.Validator = validator
	return s
}

// WithListenPort - set address for echo server
func (s *WebServer) WithListenPort(listenPort string) *WebServer {
	s.listenPort = listenPort
	return s
}

// WithRouteGroups - set routeGroups for echo server
func (s *WebServer) WithRouteGroups(routeGroups []*RouteGroups) *WebServer {
	for _, routeGroup := range routeGroups {
		s.routeGroups = append(s.routeGroups, routeGroup)
	}
	return s
}

// Register - register new route groups to array
func (s *WebServer) Register(registerRoutes *RouteGroups) {
	s.routeGroups = append(s.routeGroups, registerRoutes)
}

// Register - register new route groups to array
func (s *WebServer) GetEchoServer() *echo.Echo {
	return s.server
}

// Start - starts an HTTP server
func (s *WebServer) Start() error {
	return s.server.Start(
		s.listenPort,
	)
}

// Start - closes an HTTP server
func (s *WebServer) Close() error {
	return s.server.Close()
}

// Initial - initial echo server
func (s *WebServer) Initial() error {

	// add handler and route into echo server
	for _, routeGroup := range s.routeGroups {
		group := s.server.Group(routeGroup.Prefix)

		group.Use(defaultMiddleware...)
		group.Use(routeGroup.Mideleware...)

		fmt.Printf("Route Group: %s \n", routeGroup.Prefix)
		for _, route := range routeGroup.Routes {
			if err := addRoute(group, route); err != nil {
				fmt.Printf("Failed to load route: %v", route.Path)
				continue
			}
			fmt.Printf("load path: [%s] %s\n", route.Method, route.Path)
		}
		fmt.Printf("\n")
	}
	return nil
}

func addRoute(server *echo.Group, route *Route) error {
	switch route.Method {
	case http.MethodGet:
		server.GET(route.Path, route.Handler)
	case http.MethodPost:
		server.POST(route.Path, route.Handler)
	case http.MethodPut:
		server.PUT(route.Path, route.Handler)
	case http.MethodDelete:
		server.DELETE(route.Path, route.Handler)
	default:
		return fmt.Errorf("unknow method: %s", route.Method)
	}
	return nil
}
