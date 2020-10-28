package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"parking/repository"
	"parking/service"
)

func main() {
	repo, err := repository.New(context.Background(), "mongodb://localhost:27017", "test", "parking")
	if err != nil {
		panic(err)
	}

	srv := service.New(repo)

	var pathFile string
	flag.StringVar(&pathFile, "opt", "", "Usage")
	flag.Parse()

	command, err := srv.OpenFile(pathFile)
	if err == nil && len(command) != 0 {
		srv.RunCommand(command)
	}

	var action, stParam, ndParam string
	for {
		length, _ := fmt.Scanln(&action, &stParam, &ndParam)
		result := srv.Stdin(service.CommandInput{
			LengthOfParam: length,
			Action:        action,
			StParam:       stParam,
			NdParam:       ndParam,
		})
		if result == "exit" {
			_ = repo.DB.Drop(context.Background())
			os.Exit(0)
		}
		fmt.Println(result)
	}
}
