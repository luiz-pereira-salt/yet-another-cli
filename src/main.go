package main

import (
	"fmt"
    "os"
	"os/exec"
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
}

type CommandArg struct {
	Arg			string		`yaml:"name"`
	Description	string		`yaml:"description"`
}

func (c Command) Execute() {
	fmt.Printf("Exeucting %s:", c.Name)
				
				// cmd := exec.Command("/bin/sh", execCommand.Command)
				// cmd.Stdout = os.Stdout
				// cmd.Stderr = os.Stderr
				// err := cmd.Run()
				// if err != nil {
				// fmt.Printf("error %s", err)
				// }
}

type item struct {
	title, desc string
	cmd			Command
}


func (p item) FilterValue() string { return p.title }
func (p item) Title() string       { return p.title }
func (p item) Description() string { return p.desc}
func (p item) Command() Command { return p.cmd}


// func (p Plugin) FilterValue() string { return p.Name }
// func (p Plugin) Title() string       { return p.Name }
// func (p Plugin) Description() string { return "Desc"}

// func (c Command) FilterValue() string { return c.Name }
// func (c Command) Title() string       { return c.Name }
// func (c Command) Description() string { return fmt.Sprintf("%v", c)}


type model struct {
	list 		map[string]list.Model
	focused		string
	selected	list.Model
}

func initialModel(plugins []Plugin) model {

	items := make([]list.Item, 0, len(plugins))
	all := make(map[string]list.Model)
	
	for _, v := range plugins {
		fmt.Printf("Plugin %v\n", v)
		commands := make([]list.Item, 0, len(v.Commands))

		for _, y := range v.Commands {
			// fmt.Printf("Comands %s: %s\n", x, y)
			commands = append(commands, item{title: y.Name, desc: v.Desc, cmd: y})	
		}
		all[v.Name] = list.New(commands, list.NewDefaultDelegate(), 20, 14) 

		items = append(items, item{title: v.Name, desc: v.Desc})
   	}
	   fmt.Printf("Plugin %v\n", items)

	pluginList := list.New(items, list.NewDefaultDelegate(), 20, 14)
	pluginList.Title = "My Plugins"
	all["plugins"] = pluginList
	
	m := model{list: all, focused: "plugins", selected: all["plugins"]}
	
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
		} else if msg.String() == "enter" {
			i, ok := m.selected.SelectedItem().(item)

			if ok {
				m.focused = i.title
				m.selected = m.list[m.focused]

				if i.cmd != (Command{}){
					fmt.Printf("Executing command %s\n", i.cmd.Command)
					cmdString := fmt.Sprintf("./plugins/web/%s", i.cmd.Command)
					cmd := exec.Command("/bin/sh", cmdString)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
					fmt.Printf("error %s", err)
					}
					// return m, tea.Quit	
				}
			}
			
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		list := m.selected
		list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd

	m.selected, cmd = m.selected.Update(msg)
	return m, cmd
}


func (m model) View() string {

	return m.selected.View()
}

func main() {
	yfile, err := ioutil.ReadFile("./plugins/web/plugin.yml")

     if err != nil {
          log.Fatal(err)
     }

     var plugin Plugin

     err2 := yaml.Unmarshal(yfile, &plugin)

     if err2 != nil {

          log.Fatal(err2)
     }


	var choices = []Plugin{ plugin }


	p := tea.NewProgram(initialModel(choices))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	} 
}