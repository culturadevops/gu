package exe

import (
	"bytes"
	"errors"

	"os/exec"

	"github.com/culturadevops/gu/HandleError"
)

type Jexe struct {
	Arg        []string
	Cmd        *exec.Cmd
	Executable string
	FinalPath  string
	MapError   *HandleError.HandleError
}

func New(Executable, FinalPath string, newMapError ...*HandleError.HandleError) *Jexe {
	if len(newMapError) > 0 {
		return &Jexe{
			Arg:        []string{},
			Cmd:        &exec.Cmd{},
			Executable: Executable,
			FinalPath:  FinalPath,
			MapError:   newMapError[0],
		}
	}

	MapError := HandleError.New(errors.New("No existe ese codigo"))
	MapError.SetError("NoExecutable", errors.New("Falta Nombre del Ejecutable"))
	return &Jexe{
		Arg:        []string{},
		Cmd:        &exec.Cmd{},
		Executable: Executable,
		FinalPath:  FinalPath,

		MapError: MapError,
	}
}

func (i *Jexe) CommandInternal(withArgument bool) error {
	Err := i.Command(i.Executable, withArgument)
	if Err != nil {
		return Err
	}
	return nil
}

func (i *Jexe) Command(exectuble string, withArgument bool) error {
	var err error
	if exectuble == "" {
		return i.MapError.GetError("NoExecutable")
	}
	if withArgument {
		i.Cmd = exec.Command(exectuble, i.Arg...)
	} else {
		i.Cmd = exec.Command(exectuble)
	}
	if i.FinalPath != "" {
		i.Cmd.Dir = i.FinalPath
	}
	return err
}

func (i *Jexe) Execute(arg ...string) (string, string, error) {
	var withArgument bool
	if len(arg) > 0 {
		i.Arg = arg
		withArgument = true
	}
	Err := i.CommandInternal(withArgument)
	if Err != nil {
		return "", "", Err
	}
	return i.Run()
}
func (i *Jexe) ExecuteWithArgAndData(data string, arg ...string) (string, string, error) {
	i.Arg = arg
	Err := i.CommandInternal(true)
	if Err != nil {
		return "", "", Err
	}
	return i.Run(data)
}
func (i *Jexe) Run(data ...string) (string, string, error) {

	var stdout, stderr bytes.Buffer

	if len(data) > 0 {
		buffer := bytes.Buffer{}
		buffer.Write([]byte(data[0]))
		i.Cmd.Stdin = &buffer
	}

	i.Cmd.Stdout = &stdout
	i.Cmd.Stderr = &stderr

	err := i.Cmd.Run()
	return string(stdout.Bytes()), string(stderr.Bytes()), err
}

func (i *Jexe) AddKeyAndValue(Index string, Value string) []string {
	i.Arg = append([]string{Index, Value}, i.Arg...)
	return i.Arg
}
func (i *Jexe) Addflag(flag string) []string {
	i.Arg = append([]string{flag}, i.Arg...)
	return i.Arg
}
func (i *Jexe) AddArgs(arg ...string) []string {
	i.Arg = append(arg, i.Arg...)
	return i.Arg
}
func (i *Jexe) CommandAndRun(withArgument bool) (string, string, error) {
	Err := i.CommandInternal(withArgument)
	if Err != nil {
		return "", "", Err
	}
	return i.Run()
}
