package mono

type Error struct {
	ErrorDescription string `json:"errorDescription"`
}

func (e Error) Error() string {
	return e.ErrorDescription
}
