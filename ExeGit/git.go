package ExeGit

import (
	"fmt"

	"github.com/culturadevops/gu/HandleError"
	"github.com/culturadevops/gu/exe"
	"github.com/culturadevops/gu/path"
)

type ExeGit struct {
	Exe *exe.Jexe
}

func New(newMapError ...*HandleError.HandleError) *ExeGit {

	if len(newMapError) > 0 {
		return &ExeGit{
			Exe: exe.New("git", "", newMapError[0]),
		}
	}
	return &ExeGit{
		Exe: exe.New("git", ""),
	}
}

func (g *ExeGit) CloneB(Branch, RepoName string, FinalPath string) (string, string, error) {
	fmt.Println("cloneB")
	fmt.Println(path.FileNameWithoutExtensionFromPath(RepoName))
	if FinalPath == "" {
		FinalPath = path.FileNameWithoutExtensionFromPath(RepoName)
	}
	a, b, e := g.Exe.Execute("clone", "-b", Branch, RepoName, FinalPath)
	g.Exe.FinalPath = FinalPath
	return a, b, e

}
func (g *ExeGit) Clone(RepoName string, FinalPath string) (string, string, error) {
	if FinalPath != "" {
		return g.Exe.Execute("clone", RepoName, FinalPath)
	} else {
		return g.Exe.Execute("clone", RepoName)
	}
}
func (g *ExeGit) CloneBwithSSH(RepoName, Branch, FinalPath string) (string, string, error) {
	return g.CloneB("git@github.com:/"+RepoName, Branch, FinalPath)
}
func (g *ExeGit) CloneSSH(RepoName string, FinalPath string) (string, string, error) {
	return g.Clone("git@github.com:/"+RepoName, FinalPath)
}
func (g *ExeGit) RemoteAdd(NameOrigin, RepoName string) (string, string, error) {
	return g.Exe.Execute("remote", "add", NameOrigin, RepoName)
}
func (g *ExeGit) RemoteRm(NameOrigin string) (string, string, error) {

	return g.Exe.Execute("remote", "rm", NameOrigin)
}
func (g *ExeGit) Add(path string) (string, string, error) {
	return g.Exe.Execute("add", path)
}
func (g *ExeGit) AddAll() (string, string, error) {
	return g.Add(".")
}
func (g *ExeGit) Commit(comentario string) {
	g.Exe.Execute("commit", "-m", comentario)
}
func (g *ExeGit) Push(origin, branch string, arg ...string) (string, string, error) {
	if len(arg) > 0 {
		g.Exe.AddArgs("push", origin, branch)
		return g.Exe.Execute(arg...)
	}
	return g.Exe.Execute("push", origin, branch)
}
func (g *ExeGit) Git(arg ...string) {
	g.Exe.Execute(arg...)
}
func (g *ExeGit) PushAll() {
	g.Exe.Execute("push", "origin")
}
func (g *ExeGit) Checkout(branch string) {
	g.Exe.Execute("checkout", branch)
}
func (g *ExeGit) CheckoutB(branch string) {
	g.Exe.Execute("checkout", "-b", branch)
}
