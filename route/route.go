package route

import (
	"github.com/DQNEO/go-webapp-example/handler"
	"github.com/DQNEO/go-webapp-example/router"
)

func Register(rt *router.Router) *router.Router {

	rt.Get("/", handler.GetIndex)

	rt.Get("/hello", handler.GetHello)
	rt.Post("/hello", handler.PostHello)
	rt.Put("/hello", handler.PutHello)
	rt.Delete("/hello", handler.DeleteHello)

	rt.Get("/issues", handler.GetAllIssues)
	rt.Get("/issues/search", handler.SearchIssues)
	rt.Get("/issues/{id}", handler.GetIssue)
	rt.Post("/issues", handler.CreateIssue)
	rt.Put("/issues/{id}", handler.UpdateIssue)
	rt.Delete("/issues/{id}", handler.DeleteIssue)
	return rt
}
