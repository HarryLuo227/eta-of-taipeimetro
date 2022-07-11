package routes

import (
	v1 "eta-of-taipeimetro/controllers/v1"

	"github.com/gin-gonic/gin"
)

func RoutesHandler(superRouter *gin.Engine) {
	/**
	 * Routes extend for /api/v1
	 */
	apiGroup := superRouter.Group("/api/v1")
	apiRoutes(apiGroup)

	/**
	 * Routes extend for /distance/v1
	 */
	durationGroup := superRouter.Group("/duration")
	durationRoutes(durationGroup)

}

func apiRoutes(superRoute *gin.RouterGroup) {
	lineTransferRouter := superRoute.Group("/LineTransfer")
	{
		lineTransferRouter.GET("/", v1.QueryAllLineTransfer)
		// lineTransferRouter.POST()
	}

	S2STravelTimeRouter := superRoute.Group("/S2STravelTime")
	{
		S2STravelTimeRouter.GET("/", v1.QueryAllS2STravelTime)
	}
}

func durationRoutes(superRoute *gin.RouterGroup) {
	superRoute.GET("/:startStation/:endStation", v1.DurationTesting)
}
