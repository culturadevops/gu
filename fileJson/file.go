package fileJson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"os"

	"github.com/culturadevops/gu/HandleError"
)

type Jjson struct {
	MapError *HandleError.HandleError
}

func New(newMapError ...*HandleError.HandleError) *Jjson {

	if len(newMapError) > 0 {
		return &Jjson{
			MapError: newMapError[0],
		}
	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("ErrNotExist", errors.New("Archivo no existe"))
	return &Jjson{
		MapError: MapError,
	}
}
func (i *Jjson) ReadJsonFileAndGetGenericStruct(fileConfigName string) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonFile, err := os.Open(fileConfigName)
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (i *Jjson) ReadJsonFile(jsonFileName string, Struct interface{}) error {
	data, err := os.ReadFile(jsonFileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, Struct)
	if err != nil {
		return err
	}
	return nil
}

func (i *Jjson) UpdateAValueOnFileJson(archivo string, KeyAbuscar string, valorAremplazar interface{}) error {
	// Lee el archivo JSON
	data, err := os.ReadFile(archivo)
	if err != nil {
		return err
	}
	// Mapa para almacenar los datos del archivo JSON
	var jsonData map[string]interface{}

	// Decodifica el archivo JSON en el mapa
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	// Verifica si la clave existe en el mapa
	if _, ok := jsonData[KeyAbuscar]; ok {
		// Modifica el valor asociado a la clave
		jsonData[KeyAbuscar] = valorAremplazar

		// Codifica el mapa modificado de vuelta a JSON
		modifiedData, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			return err
		}

		// Escribe los datos modificados de vuelta al archivo
		if err := os.WriteFile(archivo, modifiedData, 0644); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("clave no encontrada: %s", KeyAbuscar)
}
func ConvertToJsonString(item map[string]string) string {
	bs1, _ := json.Marshal(item)
	return string(bs1)
}
