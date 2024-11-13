package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateRoute(resourceName string) error {
	data := struct {
		ResourceName string
	}{
		ResourceName: resourceName,
	}

	routeFilePath := filepath.Join("internal", "routes", fmt.Sprintf("%s_routes.go", ToSnakeCase(resourceName)))
	os.MkdirAll(filepath.Dir(routeFilePath), 0755)

	return CreateFileFromTemplate(routeFilePath, RouteTemplate, data)
}

// RouteTemplate is the template for generating routes
var RouteTemplate = `package routes

import (
    "github.com/gofiber/fiber/v2"
    "yourmodule/internal/handlers"
)

func Register{{ .ResourceName }}Routes(app *fiber.App) {
    app.Get("/{{ ToLower .ResourceName }}s", handlers.Get{{ .ResourceName }}s)
    app.Get("/{{ ToLower .ResourceName }}s/:id", handlers.Get{{ .ResourceName }})
    app.Post("/{{ ToLower .ResourceName }}s", handlers.Create{{ .ResourceName }})
    app.Put("/{{ ToLower .ResourceName }}s/:id", handlers.Update{{ .ResourceName }})
    app.Delete("/{{ ToLower .ResourceName }}s/:id", handlers.Delete{{ .ResourceName }})
}
`
