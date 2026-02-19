package main

import (
	"context"
	"fmt"
)
func HandlerLogin(s* State, cmd command) error{
	args := cmd.arguments
	if len(cmd.arguments) == 0{
		return fmt.Errorf("Not enough arguments")
		
	}
	userName := args[0]

	foundUser,err := s.db.GetUser(context.Background(),userName)
	if err != nil {
		fmt.Println("User does not exist")
		return err
	}
	
	err = s.cfg.SetUser(foundUser.Name)

	if err != nil{
		return err
	}	
	fmt.Printf("User %s has been set!\n",userName)
	return nil

}
