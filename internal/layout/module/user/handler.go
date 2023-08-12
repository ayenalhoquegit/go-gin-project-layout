package user


type Handler struct{
	service  *Service
}
func NewHandler(s *Service) *Handler{
	h:= new(Handler)
	h.service=s
	return h

}