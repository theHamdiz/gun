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

// In the tool's internal/generator/templates.go

// AppTemplate is the template for generating internal/app/app.go in the user's project.
var AppTemplate = `package app

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "{{ .ModuleName }}/cmd/http/api/v1"
    "{{ .ModuleName }}/internal/middleware"
    "{{ .ModuleName }}/internal/routes"
    "{{ .ModuleName }}/internal/utils"
)

type App struct {
    Fiber *fiber.App
}

func New() *App {
    app := &App{
        Fiber: fiber.New(),
    }

    app.setupMiddleware()
    app.registerRoutes()
    return app
}

func (a *App) setupMiddleware() {
    a.Fiber.Use(middleware.Recover())
    a.Fiber.Use(middleware.AuthMiddleware())
}

func (a *App) registerRoutes() {
    routes.RegisterRoutes(a.Fiber)
    v1.RegisterAPIV1(a.Fiber)
}

func (a *App) Run(address string) {
    utils.SetupGracefulShutdown(a.Fiber)

    if err := a.Fiber.Listen(address); err != nil {
        log.Fatal(err)
    }
}
`

// ModelTemplate is the template for generating models
var ModelTemplate = `package models

type {{ .ModelName }} struct {
{{- range .Fields }}
    {{ .Name }} {{ .Type }} ` + "`" + `json:"{{ ToLower .Name }}" db:"{{ ToSnakeCase .Name }}"` + "`" + `
{{- end }}
}
`

// ServerMainTemplate is the template for main.go
var ServerMainTemplate = `package main

import (
    "{{ .ModuleName }}/internal/app"
)

func main() {
    application := app.New()
    application.Run(":3000")
}
`

// APIV1Template is the template for apiv1.go
var APIV1Template = `package v1

import (
    "github.com/gofiber/fiber/v2"
    "{{ .ModuleName }}/internal/handlers"
)

func RegisterAPIV1(app *fiber.App) {
    api := app.Group("/api/v1")
    // Register API routes
    api.Get("/users", handlers.GetUsers)
    api.Get("/users/:id", handlers.GetUser)
    api.Post("/users", handlers.CreateUser)
    api.Put("/users/:id", handlers.UpdateUser)
    api.Delete("/users/:id", handlers.DeleteUser)
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

var RoutesTemplate = `package routes

import (
    "github.com/gofiber/fiber/v2"
    "{{ .ModuleName }}/internal/handlers"
)

func RegisterRoutes(app *fiber.App) {
    // Web routes
    app.Get("/", handlers.HomeHandler)

    // User routes
    RegisterUserRoutes(app)
}
`
