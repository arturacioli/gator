package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arturacioli/gator/internal/database"
	"github.com/google/uuid"
)
func HandlerRegister(s* State, cmd command) error{
	args := cmd.arguments
	if len(cmd.arguments) == 0{
		return fmt.Errorf("Not enough arguments")
	}
	userName := args[0]

	_,err := s.db.GetUser(context.Background(), userName)
	if err == nil{
		fmt.Println("user already registered")
		return errors.New("user already registered") 
	}
	user := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	}
	createdUser,err := s.db.CreateUser(context.Background(),user)
	
	if err != nil{
		return err
	}

	s.cfg.SetUser(createdUser.Name)
	if err != nil{
		return err
	}

	fmt.Printf("User %s with id: %d, was created sucessfully",createdUser.Name, createdUser.ID)
	return nil
}
