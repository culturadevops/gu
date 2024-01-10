package HandleError

type HandleError struct {
	StatusCodes map[string]error
}

func New(NoCodeError error) *HandleError {
	me := &HandleError{StatusCodes: make(map[string]error)}
	me.SetError("NoCode", NoCodeError)
	return me
}
func (he *HandleError) SetError(code string, Message error) {
	he.StatusCodes[code] = Message
}
func (he *HandleError) GetError(code string) error {
	if val, ok := he.StatusCodes[code]; ok {
		return val
	}
	return he.GetError("NoCode")
}
