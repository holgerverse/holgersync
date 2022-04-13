package main

import (
	"log"
	"path/filepath"

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
		holgerdocs(absolutePath, "terraform")
	}
}
