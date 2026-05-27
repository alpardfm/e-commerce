package rest

func (r *rest) Register() {
	// server health and testing purpose
	r.http.GET("/ping", r.Ping)
	r.registerSwaggerRoutes()
	r.registerPlatformRoutes()

	//Auth
	r.http.POST("/api/loginDashboard", r.LoginDashboard)

	//Dashboard
	r.http.GET("/api/pagination/categories", r.GetListCategoriesDashboard)
	r.http.GET("/api/categories/:id", r.GetDetailCategories)
	r.http.POST("/api/categories", r.CreateCategories)
	r.http.PUT("/api/categories/:id", r.UpdateCategories)
	r.http.DELETE("/api/categories/:id", r.DeleteCategories)

	r.http.GET("/api/pagination/location", r.GetListLocationDashboard)
	r.http.GET("/api/location/:id", r.GetDetailLocation)
	r.http.POST("/api/location", r.CreateLocation)
	r.http.PUT("/api/location/:id", r.UpdateLocation)
	r.http.DELETE("/api/location/:id", r.DeleteLocation)

	r.http.GET("/api/pagination/role", r.GetListRoleDashboard)
	r.http.GET("/api/role/:id", r.GetDetailRole)
	r.http.POST("/api/role", r.CreateRole)
	r.http.PUT("/api/role/:id", r.UpdateRole)
	r.http.DELETE("/api/role/:id", r.DeleteRole)

	// Products
	r.http.GET("/api/pagination/products", r.GetListProductsDashboard)
	r.http.GET("/api/products/:id", r.GetDetailProducts)
	r.http.POST("/api/products", r.CreateProducts)
	r.http.PUT("/api/products/:id", r.UpdateProducts)
	r.http.DELETE("/api/products/:id", r.DeleteProducts)

	// Cart
	r.http.GET("/api/cart", r.GetListCart)
	r.http.POST("/api/cart", r.CreateCart)
	r.http.PUT("/api/cart/:id", r.UpdateCart)
	r.http.DELETE("/api/cart/:id", r.DeleteCart)

	// Orders
	r.http.POST("/api/orders", r.CreateOrder)
	r.http.GET("/api/orders", r.GetListOrders)
	r.http.PUT("/api/orders/:id/status", r.UpdateOrderStatus)
}
