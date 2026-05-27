package rest

import "github.com/alpardfm/e-commerce/src/middlewares"

func (r *rest) Register() {
	// Health check
	r.http.GET("/ping", r.Ping)
	r.registerSwaggerRoutes()
	r.registerPlatformRoutes()

	// Public routes
	r.http.POST("/api/loginDashboard", r.LoginDashboard)

	// Protected routes (JWT required)
	protected := r.http.Group("/api")
	protected.Use(middlewares.JWTAuth(r.conf))
	{
		// Categories
		protected.GET("/pagination/categories", r.GetListCategoriesDashboard)
		protected.GET("/categories/:id", r.GetDetailCategories)
		protected.POST("/categories", r.CreateCategories)
		protected.PUT("/categories/:id", r.UpdateCategories)
		protected.DELETE("/categories/:id", r.DeleteCategories)

		// Location
		protected.GET("/pagination/location", r.GetListLocationDashboard)
		protected.GET("/location/:id", r.GetDetailLocation)
		protected.POST("/location", r.CreateLocation)
		protected.PUT("/location/:id", r.UpdateLocation)
		protected.DELETE("/location/:id", r.DeleteLocation)

		// Role
		protected.GET("/pagination/role", r.GetListRoleDashboard)
		protected.GET("/role/:id", r.GetDetailRole)
		protected.POST("/role", r.CreateRole)
		protected.PUT("/role/:id", r.UpdateRole)
		protected.DELETE("/role/:id", r.DeleteRole)

		// Products
		protected.GET("/pagination/products", r.GetListProductsDashboard)
		protected.GET("/products/:id", r.GetDetailProducts)
		protected.POST("/products", r.CreateProducts)
		protected.PUT("/products/:id", r.UpdateProducts)
		protected.DELETE("/products/:id", r.DeleteProducts)

		// Cart
		protected.GET("/cart", r.GetListCart)
		protected.POST("/cart", r.CreateCart)
		protected.PUT("/cart/:id", r.UpdateCart)
		protected.DELETE("/cart/:id", r.DeleteCart)

		// Orders
		protected.POST("/orders", r.CreateOrder)
		protected.GET("/orders", r.GetListOrders)
		protected.PUT("/orders/:id/status", r.UpdateOrderStatus)
	}
}
