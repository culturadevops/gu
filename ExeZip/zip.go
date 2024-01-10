package ExeZip

import (
	"github.com/culturadevops/gu/HandleError"
	"github.com/culturadevops/gu/exe"
)

type ExeZip struct {
	Exe *exe.Jexe
}

func New(newMapError ...*HandleError.HandleError) *ExeZip {

	if len(newMapError) > 0 {
		return &ExeZip{
			Exe: exe.New("zip", "", newMapError[0]),
		}
	}
	return &ExeZip{
		Exe: exe.New("zip", ""),
	}
}

func (g *ExeZip) ZipFolder(NameZip string, FolderZip string) (string, string, error) {
	return g.Exe.Execute("-r", NameZip, FolderZip)
}
