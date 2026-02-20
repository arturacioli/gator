package main

import (
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
			cmds : make(map[string]func(*State, command) error),
	}

	commands.register("login",HandlerLogin)
	commands.register("register",HandlerRegister)
	commands.register("reset", HandlerReset)	
	commands.register("users", HandlerUsers)
	commands.register("agg", HandlerAgg)
	commands.register("addfeed", HandlerAddFeed)
	commands.register("feeds", HandlerFeeds)
	commands.register("follow", HandlerFollow)
	commands.register("following", HandlerFollowing)


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
