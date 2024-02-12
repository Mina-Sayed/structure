// pkg/api/routes/routes.go

package routes

import (
    "github.com/gin-gonic/gin"

    "structure/pkg/api/handlers"
    "structure/pkg/api/middleware"
)

// SetupRouter sets up the router for the API
func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Initialize middleware
    authMiddleware := middleware.AuthMiddleware()

    // Initialize handlers
    orgHandler := handlers.NewOrganizationHandler()

    v1 := r.Group("/api/v1")
    {
        // Public routes
        v1.POST("/signup", handlers.SignUpHandler)
        v1.POST("/signin", handlers.SignInHandler)

        // Protected routes
        v1.Use(authMiddleware)
        {
            v1.POST("/refresh-token", handlers.RefreshTokenHandler)
            v1.POST("/organization", orgHandler.CreateOrganizationHandler)
            v1.GET("/organization/:organization_id", orgHandler.GetOrganizationByIDHandler)
            v1.GET("/organization", orgHandler.GetAllOrganizationsHandler)
            v1.PUT("/organization/:organization_id", orgHandler.UpdateOrganizationHandler)
            v1.DELETE("/organization/:organization_id", orgHandler.DeleteOrganizationHandler)
            v1.POST("/organization/:organization_id/invite", orgHandler.InviteUserToOrganizationHandler)
        }
    }

    return r
}
