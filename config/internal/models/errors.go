package models

// ErrorGeneric représente une erreur générique
type ErrorGeneric struct {
	Message string `json:"message"`
}

func (e *ErrorGeneric) Error() string {
	return e.Message
}

// ErrorNotFound représente une erreur quand un objet n'est pas trouvé
type ErrorNotFound struct {
	Message string `json:"message"`
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

type ErrorUnprocessableEntity struct {
	Message string `json:"message"`
}

func (e *ErrorUnprocessableEntity) Error() string {
	return e.Message
}
