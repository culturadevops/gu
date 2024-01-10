package file

import (
	"errors"

	"github.com/culturadevops/gu/HandleError"
)

type Jfile struct {
	Map      map[string]string
	MapError *HandleError.HandleError
}

func New(newMapError ...*HandleError.HandleError) *Jfile {

	if len(newMapError) > 0 {
		return &Jfile{
			Map:      map[string]string{},
			MapError: newMapError[0],
		}

	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("ErrNotExist", errors.New("Archivo no existe"))
	//	MapError.SetError("ErrInvalid ", errors.New("invalid argument"))
	//	MapError.SetError("ErrPermission", errors.New("permission denied"))
	MapError.SetError("ErrExist", errors.New("file already exists"))

	//	MapError.SetError("ErrClosed", errors.New("file already closed"))

	//ejemplo:MapError.SetError("NoExecutable", errors.New("Falta Nombre del Ejecutable"))
	return &Jfile{
		Map:      map[string]string{},
		MapError: MapError,
	}
}
func (i *Jfile) AddMap(index string, value string) {
	if i.Map == nil {
		i.Map = make(map[string]string)
	}
	i.Map[index] = value
}
