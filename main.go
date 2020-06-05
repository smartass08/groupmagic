package main

import (
	"fmt"
	"github.com/urfave/cli"
	"helpers"
	"log"
	"os"
)

func SingleAdd(c *cli.Context) error {
	srv, err := helpers.GetService()
	if err != nil {
		log.Fatalf("Unable to create service, Error is : %v", err)
	}
	helpers.Addemails(c.String("email"), c.String("group"), srv)
	fmt.Printf("Added Email %v to the group" , c.String("email"))
	return nil
}


func MultipleAdd(c *cli.Context) error {
	fmt.Println("Adding Emails!")
	helpers.JsonAdd(c.String("SaPath"), c.String("group"))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Google Group Utility"
	app.Usage = "A minimal program written in GoLang to add members to gsuite groups."
	app.Authors = []*cli.Author{
		{Name:"Shubham Dubey",},
	}
	multiple := []cli.Flag{
		&cli.StringFlag{
			Name:"SaPath",
			Value:"accounts",
		},
		&cli.StringFlag{
			Name:"group",
			Value:"",
		},
	}
	single:= []cli.Flag{
		&cli.StringFlag{
			Name:"email",
			Value:"",
		},
		&cli.StringFlag{
			Name:"group",
			Value:"",
		},
	}

	app.Commands = cli.Commands{
		{
			Name:"single",
			Usage:"Add single email to your google group.",
			Flags:single,
			Action:SingleAdd,
		},
		{
			Name:"multiple",
			Usage:"Add multiple emails from jsons to your google group.",
			Flags:multiple,
			Action:MultipleAdd,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}