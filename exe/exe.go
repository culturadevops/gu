package exe

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"os/exec"

	"github.com/culturadevops/gu/HandleError"
	"github.com/culturadevops/gu/StackError"
)

type Jexe struct {
	Arg        []string
	Cmd        *exec.Cmd
	Executable string
	FinalPath  string
	MapError   *HandleError.HandleError
	StackError *StackError.StackError
}

func New(Executable, FinalPath string, newMapError ...*HandleError.HandleError) *Jexe {
	StackError := StackError.New("jexe")
	if len(newMapError) > 0 {
		return &Jexe{
			Arg:        []string{},
			Cmd:        &exec.Cmd{},
			Executable: Executable,
			FinalPath:  FinalPath,
			MapError:   newMapError[0],
			StackError: StackError,
		}
	}

	MapError := HandleError.New(StackError.AddInternalError(errors.New("No existe ese codigo")))
	MapError.SetError("NoExecutable", StackError.AddInternalError(errors.New("Falta Nombre del Ejecutable")))
	return &Jexe{
		Arg:        []string{},
		Cmd:        &exec.Cmd{},
		Executable: Executable,
		FinalPath:  FinalPath,
		MapError:   MapError,
		StackError: StackError,
	}
}

func (i *Jexe) CommandInternal(withArgument bool) error {
	Err := i.Command(i.Executable, withArgument)
	if Err != nil {
		return i.StackError.AddExternalError(Err)
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
	fmt.Println(i.Cmd)
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
	if err != nil {

		err1 := errors.New(fmt.Sprintf("%v:%v", strings.TrimRight(string(stderr.Bytes()), "\n"), err.Error()))
		return string(stdout.Bytes()), string(stderr.Bytes()), i.StackError.AddExternalError(err1)
	}
	return string(stdout.Bytes()), string(stderr.Bytes()), nil
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
