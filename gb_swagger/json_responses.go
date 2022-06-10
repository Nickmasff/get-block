package gb_swagger

type JSONResultMessage struct {
	Message string `json:"message"`
}

type JSONResultError struct {
	Error string `json:"error"`
}
