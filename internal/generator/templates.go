package generator

import (
	"os"
	"strings"
	"text/template"

	"github.com/theHamdiz/it"
)

func CreateFileFromTemplate(destination string, tmplContent string, data interface{}) error {
	tmpl, err := template.New("template").Funcs(template.FuncMap{
		"ToLower":     strings.ToLower,
		"ToSnakeCase": ToSnakeCase,
		"Title":       strings.Title,
	}).Parse(tmplContent)
	if err != nil {
		return err
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			it.StructuredError("failed to close file", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}

// ModelTemplate is the template for generating models
var ModelTemplate = `package models

type {{ .ModelName }} struct {
{{- range .Fields }}
    {{ .Name }} {{ .Type }} ` + "`" + `json:"{{ ToLower .Name }}" db:"{{ ToSnakeCase .Name }}"` + "`" + `
{{- end }}
}
`

// MainTemplate is the template for main.go
var MainTemplate = `package main

import (
    "{{ .ModuleName }}/internal/app"
)

func main() {
    application := app.New()
    application.Run(":3000")
}
`

// ChannelsTemplate is the template for channels.go
var ChannelsTemplate = `package utils

// Channel utilities for concurrency patterns
`

// SignalsTemplate is the template for signals.go
var SignalsTemplate = `package utils

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/gofiber/fiber/v2"
)

func SetupGracefulShutdown(app *fiber.App) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        log.Println("Gracefully shutting down...")
        _ = app.Shutdown()
    }()
}
`
