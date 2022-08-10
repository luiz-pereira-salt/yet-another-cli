package main

import (
	"fmt"
    "os"
	// "os/exec"
	// "encoding/json"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
    "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
type Plugin struct { 
	Version			string 					`yaml:"version"`
	Name			string					`yaml:"name"`
	Desc			string					`yaml:"description"`
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

// func (c Command) Execute() {
// 	execCommand := m.choices[m.cursor]
				
// 				cmd := exec.Command("/bin/sh", execCommand.Command)
// 				cmd.Stdout = os.Stdout
// 				cmd.Stderr = os.Stderr
// 				err := cmd.Run()
// 				if err != nil {
// 				fmt.Printf("error %s", err)
// 				}
// }

func (p Plugin) FilterValue() string { return p.Name }
func (p Plugin) Title() string       { return p.Name }
func (p Plugin) Description() string { return fmt.Sprintf("%s is %v", p.Desc, p.Commands)}


type model struct {
	list 		list.Model
}

func initialModel(choices []Plugin) model {
	items := make([]list.Item, len(choices))
	
	for k, v := range choices {
		fmt.Printf("Comands %s: %s\n", k, v)
		items = append(items, v)
   	}	

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Fave Things"
	
	fmt.Printf("Items %v\n", m)

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}


func (m model) View() string {
	return docStyle.Render(m.list.View())
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

	fmt.Printf("Plugin%v\n", plugin)

	var choices = []Plugin{ plugin }

     for k, v := range plugin.Commands {

          fmt.Printf("Comands %s: %s\n", k, v)
     }

	p := tea.NewProgram(initialModel(choices))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	} 
}