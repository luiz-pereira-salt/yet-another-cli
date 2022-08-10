package plugin

import (
	"os"
	"os/exec"
)

type Plugin struct { 
	Version			string 					`yaml:"version"`
	Name			string					`yaml:"name"`
	Description		string					`yaml:"description"`
	Commands		[]Command				`yaml:"commands"`
}

type Command struct {
	Name		string			`yaml:"name"`
	Command		string			`yaml:"command"`
	Executer	string			`yaml:"executer"`
	Args		[]CommandArg	`yaml:"args"`
}

type CommandArg struct {
	Arg			string		`yaml:"name"`
	Description	string		`yaml:"description"`
}

func (c Command) Execute() {
	execCommand := m.choices[m.cursor]
				
				cmd := exec.Command("/bin/sh", execCommand.Command)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
				fmt.Printf("error %s", err)
				}
}