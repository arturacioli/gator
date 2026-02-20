package main

import (
	"context"
	"fmt"
)
func HandlerUsers(s* State, cmd command) error{
	foundUsers,err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("User does not exist")
		return err
	}
		
	for _,user := range foundUsers{
		if user.Name == s.cfg.Username{
			fmt.Printf("* %s (current)\n",user.Name)
		}else{
			fmt.Printf("* %s\n",user.Name)
		}
		
	}

	return nil
}
