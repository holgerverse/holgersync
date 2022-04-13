package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/alexflint/go-arg"
)

var args struct {
	ActionGroup string `arg:"positional" help:"Possible values: holgerdocs"`
	ModulePath  string `arg:"required" help:"Relative path to the Terraform module you want to perform actions on."`
}

func main() {

	arg.MustParse(&args)

	absolutePath, err := filepath.Abs(args.ModulePath)
	if err != nil {
		log.Fatal(err)
	}

	if args.ActionGroup == "holgerdocs" {
		temp := (createDocs(absolutePath))

		// Collect existing README content
		existingContent := parseMarkdown(absolutePath + "/README.md")
		markdownContent := MarkdownContent{Title: existingContent["title"], Description: existingContent["description"], ExampleUsage: existingContent["example_usage"], Variables: temp["variables"], Outputs: temp["outputs"]}

		// Template rendering
		templateFilePath, err := filepath.Abs("templates/holgerdocs.tmpl")
		if err != nil {
			log.Fatal(err)
		}

		templateContent, err := ioutil.ReadFile(templateFilePath)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err := template.New("holgerdocs").Parse(string(templateContent))
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(absolutePath + "/temp.md")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(f, markdownContent)
		if err != nil {
			log.Fatal(err)
		}
	}
}
