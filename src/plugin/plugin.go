package plugin

type Plugin struct {
	Version		string
	Name		string
	Commands	[]Command
}

type Command struct {
	Name		string
	Args		[]string
	Command		string
}