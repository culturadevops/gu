package file

import (
	"bufio"
	"os"
	"strings"
)

func (i *Jfile) ReadFile(TemplateName string) (string, error) {
	err := i.FileExist(TemplateName)
	if err != nil {
		return "", err
	}
	data, err1 := os.ReadFile(TemplateName)
	if err1 != nil {
		return "", err1
	}
	return string(data), nil
}
func (i *Jfile) ReadDictionaryFile(dictionary string) ([]string, error) {
	var passDict []string
	dictFile, err := i.ReadFile(dictionary)
	if err != nil {
		return passDict, err
	}
	passDict = strings.Split(dictFile, "\n")
	return passDict, nil
}

/*dado un archivo con variable=valor el devuelve un map con [variable]= valor*/
func (i *Jfile) LeerArchivoYCrearMapa(archivo string, separador string) (map[string]string, error) {
	// Abre el archivo
	file, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crea un mapa para almacenar los datos
	datos := make(map[string]string)

	// Crea un scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Divide cada línea del archivo en nombre y valor usando "=" como delimitador
		partes := strings.Split(scanner.Text(), separador)
		if len(partes) == 2 {
			// Agrega la entrada al mapa
			datos[strings.TrimSpace(partes[0])] = strings.TrimSpace(partes[1])
		}
	}

	// Verifica errores durante el escaneo del archivo
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return datos, nil
}

/* lee un archivo y luego lo copia a otro
 */
func (i *Jfile) ReadAndCopy(srcFileDir string, DestFileDir string) error {
	err := i.FileExist(srcFileDir)
	if err != nil {
		return err
	}
	err1 := i.FileExist(DestFileDir)
	if err1 == nil {
		return i.MapError.GetError("ErrExist")
	}
	b, err := os.ReadFile(srcFileDir)
	if err != nil {
		return err
	}
	err = os.WriteFile(DestFileDir, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

/* lee un archivo y pone su contenido en una linea usando el separador
 */
func (i *Jfile) ReadFileAndPutInLine(filePath string, separator string) (string, error) {

	osfileReadFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(osfileReadFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines string
	continuebool := fileScanner.Scan()
	fileLines = fileScanner.Text()
	if continuebool {
		for fileScanner.Scan() {
			fileLines = fileLines + separator + fileScanner.Text()
		}
	}

	osfileReadFile.Close()
	return fileLines, nil
}
