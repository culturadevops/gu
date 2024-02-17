package isValid

import (
	"errors"
	"net"
	"regexp"

	"github.com/culturadevops/gu/HandleError"
)

type isValid struct {
	MapError *HandleError.HandleError
}

func New(newMapError ...*HandleError.HandleError) *isValid {

	if len(newMapError) > 0 {
		return &isValid{
			MapError: newMapError[0],
		}
	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("ErrNotExist", errors.New("fileName no existe"))
	return &isValid{
		MapError: MapError,
	}
}

// Función para validar una dirección MAC
func (i *isValid) MacAddress(macAddress string) bool {
	// El patrón de expresión regular para validar la dirección MAC
	macPattern := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	/*
		// Eliminar cualquier carácter no alfanumérico de la dirección MAC
		cleanMacAddress := strings.Map(func(r rune) rune {
			if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'F') || (r >= 'a' && r <= 'f') {
				return r
			}
			return -1
		}, macAddress)*/

	// Verificar el formato usando la expresión regular
	return macPattern.MatchString(macAddress)
}

// Función para validar una dirección IP
func (i *isValid) IPAddress(ipAddress string) bool {
	// Utilizamos la función ParseIP de la biblioteca net
	// Devolverá una dirección IP si la cadena es válida, o nil si no lo es
	return net.ParseIP(ipAddress) != nil
}
