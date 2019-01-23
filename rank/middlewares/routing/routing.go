package routing

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/controller"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/delivery"
)

// Router sets up routing for Rank app.
func Router(controllers *controller.Controllers) *gin.Engine {
	router := setup()

	router.Use(cors.Default())

	endpoints(router, controllers)

	return router
}

// endpoints receives endpoints from each entity from Delivery layer.
func endpoints(router *gin.Engine, controllers *controller.Controllers) {
	v1 := router.Group("/api/v1")
	{
		delivery.SetGameEndpoints(v1, controllers.Game)
	}
}

// setup sets router with Gin middlewares and returns
// its default engine. It also sets up a response to the
// /hello GET request.
func setup() *gin.Engine {
	r := gin.Default()

	return r
}
