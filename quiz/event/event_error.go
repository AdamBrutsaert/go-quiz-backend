package event

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (Error) Kind() string {
	return "error"
}

var (
	ErrMalformedCommand  = Error{Code: "malformed_command", Message: "Malformed command"}
	ErrInvalidCommand    = Error{Code: "invalid_command", Message: "Invalid command"}
	ErrInvalidName       = Error{Code: "invalid_name", Message: "Invalid name"}
	ErrNameAlreadyTaken  = Error{Code: "name_taken", Message: "Name already taken"}
	ErrAlreadyRegistered = Error{Code: "already_registered", Message: "Already registered"}
	ErrNotOwner          = Error{Code: "not_owner", Message: "Not the lobby owner"}
)
