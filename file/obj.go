package file

import (
	"errors"

	"github.com/culturadevops/gu/HandleError"
	"github.com/culturadevops/gu/StackError"
)

type Jfile struct {
	Map        map[string]string
	MapError   *HandleError.HandleError
	StackError *StackError.StackError
}

func New(newMapError ...*HandleError.HandleError) *Jfile {
	StackError := StackError.New("Jfile")
	if len(newMapError) > 0 {
		return &Jfile{
			Map:        map[string]string{},
			MapError:   newMapError[0],
			StackError: StackError,
		}

	}
	MapError := HandleError.New(StackError.AddInternalError(errors.New("No existe ese codigo")))
	MapError.SetError("ErrNotExist", StackError.AddInternalError(errors.New("Archivo no existe")))
	MapError.SetError("ErrExist", StackError.AddInternalError(errors.New("file already exists")))
	//	MapError.SetError("ErrInvalid ", errors.New("invalid argument"))
	//	MapError.SetError("ErrPermission", errors.New("permission denied"))
	//	MapError.SetError("ErrClosed", errors.New("file already closed"))
	//ejemplo:MapError.SetError("NoExecutable", errors.New("Falta Nombre del Ejecutable"))
	return &Jfile{
		Map:        map[string]string{},
		MapError:   MapError,
		StackError: StackError,
	}
}
func (i *Jfile) AddMap(index string, value string) {
	if i.Map == nil {
		i.Map = make(map[string]string)
	}
	i.Map[index] = value
}
