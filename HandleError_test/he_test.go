package HandleError_test

import (
	"errors"
	"testing"

	"github.com/culturadevops/gu/HandleError"
)

func TestHandleError(t *testing.T) {
	// Creamos una nueva instancia de HandleError
	errHandler := HandleError.New(errors.New("Error: No Code Found"))

	// Agregamos un error personalizado
	errHandler.SetError("CustomCode", errors.New("Error: Custom Code"))

	// Probamos la función GetError para un código existente
	expected := "Error: Custom Code"
	if err := errHandler.GetError("CustomCode"); err.Error() != expected {
		t.Errorf("Se esperaba el mensaje de error '%s' pero se obtuvo '%s'", expected, err.Error())
	}

	// Probamos la función GetError para un código inexistente
	expected = "Error: No Code Found"
	if err := errHandler.GetError("InvalidCode"); err.Error() != expected {
		t.Errorf("Se esperaba el mensaje de error '%s' pero se obtuvo '%s'", expected, err.Error())
	}
}
