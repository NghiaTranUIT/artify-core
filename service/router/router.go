package router

import (
	"github.com/NghiaTranUIT/artify-core/resources"
	"github.com/labstack/echo"
)

type Router struct {
	R *resources.Resource
}

func (r *Router) ApplyRoutes(e *echo.Echo) {
	r.ApplyHomeRoute(e)
	r.ApplyPhotoRoute(e)
	r.ApplyAuthorRoute(e)
	r.ApplyGeneratorRoute(e)
	r.ApplyVersionRoute(e)
}
