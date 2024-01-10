package file

import (
	"io"
	"os"
)

/*
	crea una ruta completa de Dir ejemplo

/ruta1/ruta2/ruta3
si no existe crea todo
*/
func (i *Jfile) CreateAllDir(CompletePath string) error {
	err := i.FileExist(CompletePath)
	if err != nil {
		return err
	}
	return os.MkdirAll(CompletePath, 0755)
}
func (i *Jfile) FileExist(dir string) error {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return nil
	}
	return i.MapError.GetError("ErrNotExist")
}

func (i *Jfile) CreateFile(DestinationDir string, Data string) error {
	return os.WriteFile(DestinationDir, []byte(Data), 0644)
}

/* Copia un archivo usando operaciones del sistema
 */
func (i *Jfile) Copy(srcFileDir string, DestFileDir string) error {
	err := i.FileExist(srcFileDir)
	if err != nil {
		return err
	}
	err1 := i.FileExist(DestFileDir)
	if err1 == nil {
		return i.MapError.GetError("ErrExist")
	}
	srcFile, err := os.Open(srcFileDir)

	if err != nil {
		return err
	}
	defer srcFile.Close()
	destFile, err := os.Create(DestFileDir) // creates if file doesn't exist

	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return err
	}
	err = destFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

/*añade al final del archivo un string
 */
func (i *Jfile) AppEndToFile(DestFileDir string, data string) error {
	osfile, err := os.OpenFile(DestFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer osfile.Close()
	_, err = osfile.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

/*
agrega al final de un archivo varios string
*/
func (i *Jfile) AppEndArrayToFile(DestFileDir string, datas []string) error {
	osfile, err := os.OpenFile(DestFileDir, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		return err
	}
	defer osfile.Close()
	for _, data := range datas {
		_, err = osfile.WriteString(data)
		if err != nil {
			return err
		}
	}
	return nil
}
