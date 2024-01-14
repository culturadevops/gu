package fileYml

import (
	"errors"

	"os"

	"github.com/culturadevops/gu/HandleError"
	"gopkg.in/yaml.v2"
)

type JYml struct {
	MapError *HandleError.HandleError
}

func New(newMapError ...*HandleError.HandleError) *JYml {

	if len(newMapError) > 0 {
		return &JYml{
			MapError: newMapError[0],
		}
	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("ErrNotExist", errors.New("fileName no existe"))
	return &JYml{
		MapError: MapError,
	}
}
func (i *JYml) ReadFileAndGetGenericStruct(fileConfigName string) (map[string]interface{}, error) {
	var result map[string]interface{}
	File, err := os.Open(fileConfigName)
	if err != nil {
		return result, err
	}
	defer File.Close()

	// Parsear el contenido YAML en el map
	err = yaml.Unmarshal(File, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (i *JYml) ReadFile(jsonFileName string, Struct interface{}) error {
	data, err := os.ReadFile(jsonFileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Struct)
	if err != nil {
		return err
	}
	return nil
}
func (i *JYml) WriteFile(fileName string, structure interface{}) error {

	// Convertir la estructura a formato YAML
	data, err := yaml.Marshal(&structure)
	if err != nil {
		return err
	}

	// Escribir el YAML en un fileName
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil

}
func (i *JYml) UpdateAValueOnFile(fileName string, KeyAbuscar string, valorAremplazar interface{}) error {
	// Lee el fileName
	data, err := i.ReadFileAndGetGenericStruct(fileName)
	if err != nil {
		return err
	}
	// Verifica si la clave existe en el mapa
	if _, ok := data[KeyAbuscar]; ok {
		// Modifica el valor asociado a la clave
		data[KeyAbuscar] = valorAremplazar
		err = i.WriteFile(fileName, &data)
		if err != nil {
			return err
		}
	}
	return nil
}
func ConvertYamlToString(item map[string]string) (string, error) {
	data, err := yaml.Marshal(&item)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
