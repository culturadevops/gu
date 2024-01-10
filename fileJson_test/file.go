package fileJson_test

import (
	"os"
	"testing"

	"github.com/culturadevops/gu/fileJson"
)

func TestReadJsonFileAndGetGenericStruct(t *testing.T) {
	fileContent := `{"key1": "value1", "key2": "value2"}`
	fileName := "testFile.json"

	// Crear un archivo temporal para la prueba
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creando archivo de prueba: %s", err)
	}
	defer os.Remove(fileName)

	jjson := fileJson.New()
	result, err := jjson.ReadJsonFileAndGetGenericStruct(fileName)
	if err != nil {
		t.Errorf("Error al leer archivo JSON: %s", err)
	}

	expected := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	// Verificar si el resultado coincide con lo esperado
	for key, val := range expected {
		if result[key] != val {
			t.Errorf("Se esperaba '%v' para la clave '%s', pero se obtuvo '%v'", val, key, result[key])
		}
	}
}

func TestUpdateAValueOnFileJson(t *testing.T) {
	fileContent := `{"key1": "value1", "key2": "value2"}`
	fileName := "testFile.json"

	// Crear un archivo temporal para la prueba
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creando archivo de prueba: %s", err)
	}
	defer os.Remove(fileName)

	jjson := fileJson.New()
	key := "key2"
	newValue := "new_value"

	// Actualizar un valor en el archivo JSON
	err = jjson.UpdateAValueOnFileJson(fileName, key, newValue)
	if err != nil {
		t.Errorf("Error al actualizar valor en archivo JSON: %s", err)
	}

	// Leer el archivo actualizado y verificar si el valor se actualizó correctamente
	result, err := jjson.ReadJsonFileAndGetGenericStruct(fileName)
	if err != nil {
		t.Errorf("Error al leer archivo JSON actualizado: %s", err)
	}

	if result[key] != newValue {
		t.Errorf("Se esperaba '%s' para la clave '%s', pero se obtuvo '%v'", newValue, key, result[key])
	}
}
