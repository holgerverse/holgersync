package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/alexflint/go-arg"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

var args struct {
	ActionGroup string `arg:"positional" help:"Possible values: holgerdocs"`
	ModulePath  string `arg:"required" help:"Relative path to the Terraform module you want to perform actions on."`
}

type MarkdownContent struct {
	Title        string
	Description  string
	ExampleUsage string
	Variables    []map[string]string
	Outputs      []map[string]string
}

type Variable struct {
	Name        string         `hcl:",label"`
	Description string         `hcl:"description,optional"`
	Sensitive   bool           `hcl:"sensitive,optional"`
	Type        *hcl.Attribute `hcl:"type,optional"`
	Default     *hcl.Attribute `hcl:"default,optional"`
	Options     hcl.Body       `hcl:",remain"`
}

type Output struct {
	Name        string   `hcl:",label"`
	Description string   `hcl:"description,optional"`
	Sensitive   bool     `hcl:"sensitive,optional"`
	Value       string   `hcl:"value,optional"`
	Options     hcl.Body `hcl:",remain"`
}

type Resource struct {
	Name    string   `hcl:"name,label"`
	Options hcl.Body `hcl:",remain"`
}

type Config struct {
	Outputs   []*Output   `hcl:"output,block"`
	Variables []*Variable `hcl:"variable,block"`
	Resources []*Resource `hcl:"resource,block"`
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

func filesInDirectory(hclPath string) []fs.FileInfo {

	var terraformFiles []fs.FileInfo

	// Read all files in the terraform directory
	files, err := ioutil.ReadDir(hclPath)
	if err != nil {
		log.Fatal(err)
	}

	// Only safe *.tf files
	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".tf") {
			terraformFiles = append(terraformFiles, v)
		}
	}

	return terraformFiles
}

func createDocs(hclPath string) map[string][]map[string]string {

	var variables []map[string]string
	var outputs []map[string]string

	parsedConfig := make(map[string][]map[string]string)
	hclConfig := make(map[string][]byte)

	c := &Config{}

	// Iterate all Terraform files and safe the contents in the hclConfig map
	for _, file := range filesInDirectory(hclPath) {
		fileContent, err := os.ReadFile(hclPath + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		hclConfig[file.Name()] = fileContent
	}

	// Iterate all file contents
	for k, v := range hclConfig {

		parsedConfig, diags := hclsyntax.ParseConfig(v, k, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			log.Fatal(diags)
		}

		diags = gohcl.DecodeBody(parsedConfig.Body, nil, c)
		if diags.HasErrors() {
			log.Fatal(diags)
		}
	}

	for _, v := range c.Variables {

		var variableType string
		var variableDefault string

		if v.Type != nil {
			variableType = (v.Type.Expr).Variables()[0].RootName()
		}

		if v.Default != nil {
			variableDefault = (v.Default.Expr).Variables()[0].RootName()
		}

		variables = append(variables, map[string]string{
			"name":        v.Name,
			"description": v.Description,
			"sensitive":   strconv.FormatBool(v.Sensitive),
			"type":        variableType,
			"default":     variableDefault,
		})
	}

	for _, v := range c.Outputs {
		outputs = append(outputs, map[string]string{
			"name":        v.Name,
			"description": v.Description,
			"sensitive":   strconv.FormatBool(v.Sensitive),
			"value":       v.Value,
		})
	}

	parsedConfig["variables"], parsedConfig["outputs"] = variables, outputs

	return parsedConfig

}

func parseMarkdown(markdownPath string) map[string]string {

	// Read in the content of the existing markdown file
	fileContent, err := ioutil.ReadFile(markdownPath)
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string]string)

	// Create custom Markdown parser
	extensions := parser.FencedCode | parser.Tables
	parser := parser.NewWithExtensions(extensions)

	// Parse Markdown
	temp := markdown.Parse(fileContent, parser)

	for _, child := range temp.AsContainer().Children {

		// Check if the child node is of type Heading
		if _, ok := child.(*ast.Heading); ok {

			/*
			  Copy the heading text and safe it to results if the heading level is 1
			  Also check if the next node is of type paragraph and safe it to results, it
			  is going to be used as description.
			*/
			if (child.(*ast.Heading)).Level == 1 {
				results["title"] = string(child.GetChildren()[0].AsLeaf().Literal)

				description := ast.GetNextNode(child)
				if _, ok := description.(*ast.Paragraph); ok {
					results["description"] = string(description.GetChildren()[0].AsLeaf().Literal)
				}
			}

			/*
			  Check if the heading is level 2 and says 'Example usage'. If so check if the next node is of
			  type *ast.CodeBlock and safe it to results. && string(child.GetChildren()[0].AsLeaf().Literal) == "Example usage"
			*/
			if (child.(*ast.Heading)).Level == 2 && string(child.GetChildren()[0].AsLeaf().Literal) == "Example Usage" {
				exampleUsageCodeBlock := ast.GetNextNode(child)
				if _, ok := exampleUsageCodeBlock.(*ast.CodeBlock); ok {
					results["example_usage"] = string(exampleUsageCodeBlock.AsLeaf().Literal)
				}
			}

		}
	}

	return results
}
