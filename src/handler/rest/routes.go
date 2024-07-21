package rest

func (r *rest) Register() {
	// server health and testing purpose
	r.http.GET("/ping", r.Ping)
	r.registerSwaggerRoutes()
	r.registerPlatformRoutes()
}
