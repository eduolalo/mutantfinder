package structs

// Response es el estándar de respuestas Walabi
type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Ok Setea los campos a un estado de respusta válida
func (s *Response) Ok(mssg string) {

	s.Code = "200"
	s.Message = "ok."
	if mssg != "" {

		s.Message = mssg
	}
}

// BadRequest Setea los campos a un estado de respusta de Error de parte del cliente
func (s *Response) BadRequest(mssg string) {

	s.Code = "400"
	s.Message = mssg
}

// Forbridden Setea los campos a un estado de respusta de Error de parte del cliente
func (s *Response) Forbridden(mssg string) {

	s.Code = "403"
	s.Message = mssg
}

// NotFound Setea los campos a un estado de respusta de Error de parte del cliente
func (s *Response) NotFound(mssg string) {

	s.Code = "404"
	s.Message = mssg
}

// InternalError Setea los campos a un estado de respusta de error interno
func (s *Response) InternalError(mssg, code string) {

	s.Code = "500"
	s.Message = mssg
	if code != "" {

		s.Code = code
	}
}
