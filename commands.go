package main

import (
	"fmt"

	"github.com/arturaciolii/gator/internal/config"
	"github.com/arturaciolii/gator/internal/database"
)

type State struct{
	cfg *config.Config 
	db  *database.Queries
}

type command struct{
	name string
	arguments []string
}

type Commands struct{
	cmds map[string]func(*State, command) error
}



func (c *Commands) run(s *State, cmd command) error{

	handler, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command")
	}

	return handler(s,cmd)
}
func (c *Commands) register(name string, f func(*State, command) error){
	c.cmds[name] = f
}
