package main

import (
	"fmt"
    "os"
	"os/exec"
	// "encoding/json"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
    tea "github.com/charmbracelet/bubbletea"
)

type Plugin struct { 
	Version		string 					`yaml:"version"`
	Name		string					`yaml:"name"`
	Commands	[]Command	`yaml:"commands"`
}

type Command struct {
	Name		string			`yaml:"name"`
	Command		string			`yaml:"command"`
	Args		[]CommandArg	`yaml:"args"`
}

type CommandArg struct {
	Arg			string		`yaml:"name"`
	Description	string		`yaml:"description"`
}

type model struct {
	choices 	[]Command
	cursor 		int
	selected	map[int]Command
}


func initialModel(choices []Command) model {
	return model{
		choices: choices,
		selected: make(map[int]Command),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: {
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

		switch msg.String() {
		case "ctrl+c",  "q":
			return m, tea.Quit
		
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.choices[m.cursor]
				execCommand := m.choices[m.cursor]
				
				cmd := exec.Command("/bin/sh", execCommand.Command)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
				fmt.Printf("error %s", err)
				}
				return m, tea.Quit
			}
			
		}
	}
	}

	return m, nil
}


func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}

	s += "\nPress q to quit\n"
	return s
}

func main() {
	yfile, err := ioutil.ReadFile("./plugin.yml")

     if err != nil {
          log.Fatal(err)
     }

     var plugin Plugin

     err2 := yaml.Unmarshal(yfile, &plugin)

     if err2 != nil {

          log.Fatal(err2)
     }

	fmt.Printf("%v", plugin)

     for k, v := range plugin.Commands {

          fmt.Printf("%s: %s\n", k, v)
     }

	// p := tea.NewProgram(initialModel(choices))
	// if err := p.Start(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// } 
}