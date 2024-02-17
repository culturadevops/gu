package StackError

import "fmt"

type StackError struct {
	Object string
}

type SetError struct {
	Object  string
	Message error
	Where   string //internal-external

}

func New(Object string) *StackError {
	return &StackError{
		Object: Object,
	}
}
func (he *StackError) SetError(Where string, Message error) error {
	return &SetError{
		Object:  he.Object,
		Message: Message,
		Where:   Where,
	}
}
func (he *StackError) AddInternalError(Message error) error {
	return he.SetError("internal", Message)
}
func (he *StackError) AddExternalError(Message error) error {
	return he.SetError("External", Message)
}
func (he *SetError) Error() string {
	return fmt.Sprintf("%v:%v:%v", he.Object, he.Where, he.Message)
}
