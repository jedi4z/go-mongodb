package handler

type RequestHandler interface {
	userHandler
	pingHandler
}

type Handler struct{}

func BuildRequestHandler() RequestHandler {
	return &Handler{}
}
