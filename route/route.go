package route

import "../router"
import "../handler"

func Register(rt *router.Router) *router.Router {

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

