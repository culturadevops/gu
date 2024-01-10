package exe_test

import (
	"testing"

	"github.com/culturadevops/gu/exe"
)

func TestJexe_ExecuteEmpty(t *testing.T) {
	j := exe.New("", "")
	_, _, err := j.Execute("-l")
	if err != j.MapError.GetError("NoExecutable") {
		t.Errorf("Error en ExecuteEmpty: %v", err)
	}
}
func TestJexe_ExecuteWithArg(t *testing.T) {
	j := exe.New("ls", "") // Cambia los parámetros según tu caso
	out, outerr, err := j.Execute("-l")

	if err != nil {
		t.Errorf("Error en ExecuteWithArg: %v outerror=%v", err, outerr)
	}
	if out == "" {
		t.Errorf("ExecuteWithArg no devolvió ninguna salida")
	}
	// Puedes agregar más comprobaciones aquí según lo que espere tu función
}
func TestAddKeyAndValue(t *testing.T) {
	j := exe.New("command", "/path")
	initialArgs := []string{"initialArg1", "initialArg2"}

	j.Arg = initialArgs
	key := "key"
	value := "value"

	expectedArgs := []string{"key", "value", "initialArg1", "initialArg2"}

	result := j.AddKeyAndValue(key, value)

	for i, v := range result {
		if v != expectedArgs[i] {
			t.Errorf("Expected %s at index %d, got %s", expectedArgs[i], i, v)
		}
	}
}

func TestExecuteWithArgAndData(t *testing.T) {
	j := exe.New("echo", "/path")

	expectedStdout := "Hello, World!\n"
	expectedStderr := ""
	expectedError := error(nil)

	stdout, stderr, err := j.ExecuteWithArgAndData("Hello, World!")

	if stdout != expectedStdout {
		t.Errorf("Expected stdout: %s, got: %s", expectedStdout, stdout)
	}

	if stderr != expectedStderr {
		t.Errorf("Expected stderr: %s, got: %s", expectedStderr, stderr)
	}

	if err != expectedError {
		t.Errorf("Expected error to be nil, got: %s", err)
	}
}

func TestCommandInternal(t *testing.T) {
	// Configuración inicial
	executable := "echo"
	finalPath := "/path"

	// Crear una instancia de Jexe
	j := exe.New(executable, finalPath)

	// Ejecutar CommandInternal
	err := j.CommandInternal(true)

	// Verificar si se generó un error
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Verificar si el comando está configurado correctamente
	if j.Cmd.Path != "/usr/bin/echo" {
		t.Errorf("Expected executable: %s, got: %s", executable, j.Cmd.Path)
	}

	if j.Cmd.Dir != finalPath {
		t.Errorf("Expected finalPath: %s, got: %s", finalPath, j.Cmd.Dir)
	}
}
