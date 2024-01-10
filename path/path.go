package path

import (
	"path/filepath"
	"strings"
)

func FileNameWithoutExtensionFromPath(absolutePath string) string {
	filename := filepath.Base(absolutePath)
	nameOnly := strings.TrimSuffix(filename, filepath.Ext(filename))
	return nameOnly
}
