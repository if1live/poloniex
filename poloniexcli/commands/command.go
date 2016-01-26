package commands

//Credentials are the poloniex api credentials
type Credentials struct {
	Key    string
	Secret string
}

//Command is a toplevel command to be executed by the cli's main routine
type Command interface {
	//Execute the command
	Execute() error
}
