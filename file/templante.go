package file

import (
	"regexp"
	"strings"
)

func (i *Jfile) ReplaceTextInFile(TemplateName string, MapForReplace map[string]string) (string, error) {
	input, err := i.ReadFile(TemplateName)
	if err != nil {
		return "", err
	}
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input, nil
}
func (i *Jfile) ReplaceContentAndCreateNewFile(TemplateName, NewName string, MapForReplace map[string]string) error {
	data, err := i.ReplaceTextInFile(TemplateName, MapForReplace)
	if err != nil {
		return err
	}
	err = i.CreateFile(NewName, data)
	if err != nil {
		return err
	}
	return nil
}

/* Remplaza terminos en un archivo con data, los terminos a remplazar estan en el MapForReplace
 */
func (i *Jfile) ReplaceContentAndOverwriteFile(TemplateFileName, DestinationFileName string, MapForReplace ...map[string]string) error {
	if DestinationFileName == "" {
		DestinationFileName = TemplateFileName
	}

	if len(MapForReplace) > 0 {
		return i.ReplaceContentAndCreateNewFile(TemplateFileName, DestinationFileName, MapForReplace[0])
	}
	return i.ReplaceContentAndCreateNewFile(TemplateFileName, DestinationFileName, i.Map)
}

/*
agrega un valor a un archivo remplazando lo que consiga en "index" por el valor de "value"
*/
func (i *Jfile) AddValueToFile(index string, value string, Destinationfile string) error {
	MapForReplace := make(map[string]string)
	MapForReplace[index] = value
	err := i.ReplaceContentAndOverwriteFile(Destinationfile, Destinationfile, MapForReplace)
	if err != nil {
		return err
	}
	return nil
}

/*
Add content to file if not exist match string
*/
func (i *Jfile) AddContentIfNotExist(DestineName string, MatchString string, ContentToAdd string) error {
	word, err := i.ReadFile(DestineName)
	if err != nil {
		return err
	}
	existe, _ := regexp.MatchString(MatchString, word)
	if !existe {
		err = i.AppEndToFile(DestineName, ContentToAdd)
		if err != nil {
			return err
		}
	}
	return nil
}
