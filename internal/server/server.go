package server

type Server struct {
	port string
}

func New(port string) (*Server, error) {
	return &Server{}, nil
}
