package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateHandler(resourceName string) error {
	data := struct {
		ResourceName string
	}{
		ResourceName: resourceName,
	}

	handlerFilePath := filepath.Join("internal", "handlers", fmt.Sprintf("%s_handler.go", ToSnakeCase(resourceName)))
	err := os.MkdirAll(filepath.Dir(handlerFilePath), 0755)
	if err != nil {
		return err
	}

	return CreateFileFromTemplate(handlerFilePath, HandlerTemplate, data)
}

// HandlerTemplate is the template for generating handlers
var HandlerTemplate = `package handlers

import (
    "github.com/gofiber/fiber/v2"
    "{{ .ModuleName }}/internal/models"
)

func Get{{ .ResourceName }}s(c *fiber.Ctx) error {
    // TODO: Implement logic to retrieve list of {{ .ResourceName }}s
    return c.JSON(fiber.Map{"message": "List of {{ .ResourceName }}s"})
}

func Get{{ .ResourceName }}(c *fiber.Ctx) error {
    // TODO: Implement logic to retrieve a single {{ .ResourceName }}
    return c.JSON(fiber.Map{"message": "Get {{ .ResourceName }}"})
}

func Create{{ .ResourceName }}(c *fiber.Ctx) error {
    var item models.{{ .ResourceName }}
    if err := c.BodyParser(&item); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    // TODO: Save item to database
    return c.JSON(item)
}

func Update{{ .ResourceName }}(c *fiber.Ctx) error {
    // TODO: Implement logic to update {{ .ResourceName }}
    return c.JSON(fiber.Map{"message": "Update {{ .ResourceName }}"})
}

func Delete{{ .ResourceName }}(c *fiber.Ctx) error {
    // TODO: Implement logic to delete {{ .ResourceName }}
    return c.JSON(fiber.Map{"message": "Delete {{ .ResourceName }}"})
}
`
