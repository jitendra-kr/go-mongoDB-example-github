package internal

// Register Routes
func (s *Server) AddRoutes() {
	s.Server.POST("/user", s.createUser)
	s.Server.GET("/user/:id", s.getUserByID)
	s.Server.DELETE("/user/:id", s.deleteUserByID)
}
