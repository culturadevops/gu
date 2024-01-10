package ExeMysqlDump

import (
	"github.com/culturadevops/gu/HandleError"
	"github.com/culturadevops/gu/exe"
)

type ExeMysqlDump struct {
	Exe  *exe.Jexe
	Dns  string
	User string
	Pass string
}

func New(newMapError ...*HandleError.HandleError) *ExeMysqlDump {

	if len(newMapError) > 0 {
		return &ExeMysqlDump{
			Exe: exe.New("mysqldump", "", newMapError[0]),
		}
	}
	return &ExeMysqlDump{
		Exe: exe.New("mysqldump", ""),
	}
}

func (g *ExeMysqlDump) Dump(DataBase string, Table string, filename string) (string, string, error) {
	return g.Exe.Execute("-u", g.User, "-p"+g.Pass, "-h", g.Dns, DataBase, Table, ">", filename)
}
