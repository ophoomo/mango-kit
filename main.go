package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Mango-Kit CLI"
	app.Usage = "mango-kit is tool generate code for golang"
	app.Author = "Thanaphoom Aunchit"
	app.Version = "Demo"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create project golang",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					newProject(c.Args()[0])
				} else {
					fmt.Print("Enter Name Project")
				}
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate code in project",
			Action: func(c *cli.Context) {
				if len(c.Args()) >= 2 {
					generateCode(c.Args()[0], c.Args()[1])
				} else {
					showOptionsGenerateCode()
				}
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func showOptionsGenerateCode() {
	fmt.Printf("Hello")
}

func generateCode(options string, name string) {
	if options == "res" {
		createDirectory(name)
		generateRes(name)
	} else {
		showOptionsGenerateCode()
	}
}

func generateRes(name string) {
	defaultFile := []string{"Controller.tmpl", "Repository.tmpl", "Model.tmpl", "Interface.tmpl", "Service.tmpl", "Dto.tmpl"}
	var file *os.File
	for _, s := range defaultFile {
		vars := make(map[string]interface{})
		vars["name"] = strings.Title(name)
		vars["package"] = name
		tmpl, _ := template.ParseFiles("templates/" + s)

		split := strings.Split(s, ".")
		if split[len(split)-1] == "tmpl" {
			file, _ = os.Create(name + "/" + name + split[0] + ".go")
		} else {
			file, _ = os.Create(name + "/" + s)
		}
		defer file.Close()

		err := tmpl.Execute(file, vars)
		if err != nil {
			fmt.Printf("\033[1;31m CREATE\t\t %s \t\tError \033[0m \n", name+split[0]+".go")
		} else {
			fmt.Printf("\033[1;32m CREATE\t\t %s \t\tSUCCESS \033[0m \n", name+split[0]+".go")
		}
	}
}

func newProject(name string) {
	createDirectory(name)
	createDirectory(name + "/routes")
	createFileDefault(name)
	fmt.Printf("\033[1;32m CREATE PROJECT\t\t %s \t\tSUCCESS \033[0m \n", name)
}

func createFileDefault(directory string) {
	defaultFile := []string{"main.tmpl", "route.tmpl", ".gitignore", "cloudbuild.yaml", "dockerfile", ".env"}
	var file *os.File
	for _, s := range defaultFile {
		tmpl, _ := template.ParseFiles("templates/" + s)

		split := strings.Split(s, ".")
		if split[len(split)-1] == "tmpl" {
			if s == "route.tmpl" {
				file, _ = os.Create(directory + "/routes/route.go")
			} else {
				file, _ = os.Create(directory + "/main.go")
			}
		} else {
			file, _ = os.Create(directory + "/" + s)
		}
		defer file.Close()

		tmpl.Execute(file, nil)
	}
}

func createDirectory(directory string) {
	s := strings.Split(directory, "/")
	if len(s) > 1 {
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := os.Mkdir(directory, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
