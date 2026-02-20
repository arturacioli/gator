package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/arturaciolii/gator/internal/config"
	"github.com/arturaciolii/gator/internal/database"
	_ "github.com/lib/pq"
)

func main(){

	conf,err := config.Read()
	if err != nil{
		log.Fatalf("error reading config: %v", err)
		os.Exit(1)	
	}

	db, err := sql.Open("postgres", conf.Url)
	if err != nil{
		log.Fatalf("error connecting to database: %v", err)
		os.Exit(1)	
	}
	dbQueries := database.New(db)

	state := State{
		db: dbQueries,
		cfg: &conf,
	}
	commands := Commands{
			cmds : make(map[string]func(*State, command, *database.User) error),
	}

	commands.register("login",   middlewareLoggedIn(HandlerLogin)
	commands.register("register",middlewareLoggedIn(HandlerRegister)
	commands.register("reset",   middlewareLoggedIn(HandlerReset)	
	commands.register("users",   middlewareLoggedIn(HandlerUsers)
	commands.register("agg",  	 middlewareLoggedIn(HandlerAgg)
	commands.register("addfeed", middlewareLoggedIn(HandlerAddFeed)
	commands.register("feeds", middlewareLoggedIn(HandlerFeeds))
	commands.register("follow", middlewareLoggedIn(HandlerFollow))
	commands.register("following", middlewareLoggedIn(HandlerFollowing))


	if len(os.Args) < 2{
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	args := os.Args[2:]
	
	command := command{
		name: commandName,
		arguments: args,
	}	

	err = commands.run(&state,command)

	if err != nil{
		fmt.Print(err)
		os.Exit(1)
	}

}

func middlewareLoggedIn(handler func(s *State, cmd command, user database.User) error) func(*State, command) error{
	return func(s *State, cmd command) error{
		user, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			return err
		}
		return handler(s,cmd,user)
	}
}
