package csvinmemory

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/culturadevops/gu/HandleError"
)

type CsvInMemory struct {
	InMemory []map[string]string
	Head     []string
	MapError *HandleError.HandleError
}

func New(newMapError ...*HandleError.HandleError) *CsvInMemory {
	if len(newMapError) > 0 {
		return &CsvInMemory{
			InMemory: []map[string]string{},
			Head:     []string{},
			MapError: newMapError[0],
		}
	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("NoExecutable", errors.New("Falta Nombre del Ejecutable"))
	return &CsvInMemory{
		InMemory: []map[string]string{},
		Head:     []string{},
		MapError: MapError,
	}
}
func (i *CsvInMemory) OpenFile(file string) (*csv.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	return r, nil
}
func (i *CsvInMemory) ReadHead(file *csv.Reader) error {
	var head []string
	record, err := file.Read()
	if err == io.EOF {
		return err
	}
	head = append(head, record...)
	i.Head = head
	return nil
}
func (i *CsvInMemory) CreateCsvInMemory(head []string, r *csv.Reader) error {
	var file map[string]string
	for {
		file = make(map[string]string)
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		h := 0
		for value := range record {
			if head[h] != "-" {
				file[head[h]] = record[value]
			}
			h = h + 1
		}
		i.InMemory = append(i.InMemory, file)
	}
	return nil
}

/*
ExportToCSVwithStructDirect
trabaja con cualquier estructura pero lo hace de forma directa hay otra funcion similar pero necesita algo mas generico
como un []interface en cambio esta puede usar una estrucutra creada por el usuario ejemplo

	type EC2InstanceInfo struct {
		Region        string
		NameTag       string
		InstanceState string
		InstanceID    string
		PublicIP      string
		VPCID         string
		SubnetID      string
		Cost          float64
	}

	type EIPInfo struct {
		Region            string
		PublicIP          string
		AssociationTarget string
		NameTag           string
		Cost              float64
	}

	type LoadBalancerInfo struct {
		Region          string
		Type            string
		DNSName         string
		IPCount         int
		TrafficLastWeek int
		PublicIPs       []string
		Cost            float64
	}
*/
func ExportToCSVwithStructDirect(headers []string, data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return nil // O devolver un error si no es un slice
	}

	// Escribir encabezados al archivo CSV
	if err := writer.Write(headers); err != nil {
		return err
	}

	for i := 0; i < v.Len(); i++ {
		var values []string
		item := v.Index(i)

		for _, header := range headers {
			field := item.FieldByName(header)
			if !field.IsValid() {
				values = append(values, "") // Si el campo no existe, añadir una celda vacía
				continue
			}

			var value string
			switch field.Kind() {
			case reflect.String:
				value = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(field.Int(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Slice:
				if field.Type().Elem().Kind() == reflect.String {
					for j := 0; j < field.Len(); j++ {
						values = append(values, field.Index(j).String())
					}
				}
			}
			values = append(values, value)
		}

		if err := writer.Write(values); err != nil {
			return err
		}
	}
	return nil
}

// Convierte un slice de una estructura a un slice de interfaces
/*
toma esto y lo comvierte en
type EC2InstanceInfo struct {
	Region        string
	NameTag       string
	InstanceState string
	InstanceID    string
	PublicIP      string
	VPCID         string
	SubnetID      string
	Cost          float64
}
{Region: "us-west-1", NameTag: "Server1", InstanceState: "Running", InstanceID: "i-1234567890", PublicIP: "203.0.113.1", VPCID: "vpc-1234", SubnetID: "subnet-5678", Cost: 50.5},

esto
{us-west-1 Server1 Running i-1234567890 203.0.113.1 vpc-1234 subnet-5678 50.5}
*/
func convertToInterfaceSlice(data interface{}) []interface{} {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}
	return result
}

/*
esta funcion toma un []interface{} y lo pasa a cvs es importante usar convertToInterfaceSlice para transformar cualquier
estructura a un interface generico
*/
func ExportSliceInterfaceToCSV(headers []string, data []interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir encabezados al archivo CSV
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Escribir los valores correspondientes a los encabezados en el archivo CSV
	for _, item := range data {
		var values []string

		v := reflect.ValueOf(item)
		for _, header := range headers {
			field := v.FieldByName(header)
			if !field.IsValid() {
				values = append(values, "") // Si el campo no existe, añadir una celda vacía
				continue
			}

			var value string
			switch field.Kind() {
			case reflect.String:
				value = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(field.Int(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Slice:
				if field.Type().Elem().Kind() == reflect.String {
					for i := 0; i < field.Len(); i++ {
						values = append(values, field.Index(i).String())
					}
				}
			}
			values = append(values, value)
		}

		if err := writer.Write(values); err != nil {
			return err
		}
	}

	return nil
}

func ExportMapSimpleToCSV(body map[string]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Obtener encabezados del mapa
	var headers []string
	for header := range body {
		headers = append(headers, header)
	}

	// Escribir encabezados al archivo CSV
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Crear un slice ordenado de valores basado en los encabezados
	var values []string
	for _, header := range headers {
		values = append(values, body[header])
	}

	// Escribir los valores correspondientes a los encabezados en el archivo CSV
	if err := writer.Write(values); err != nil {
		return err
	}

	return nil
}
