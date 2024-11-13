package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateMiddleware(name string) error {
	caser := cases.Title(language.English, cases.Compact)
	casedName := caser.String(strings.ToLower(name))

	data := struct {
		MiddlewareName string
	}{
		MiddlewareName: casedName,
	}

	middlewareFilePath := filepath.Join("internal", "middleware", fmt.Sprintf("%s_middleware.go", ToSnakeCase(name)))
	err := os.MkdirAll(filepath.Dir(middlewareFilePath), 0755)
	if err != nil {
		return err
	}

	return CreateFileFromTemplate(middlewareFilePath, MiddlewareTemplate, data)
}

// MiddlewareTemplate is the template for generating middleware
var MiddlewareTemplate = `package middleware

import (
    "github.com/gofiber/fiber/v2"
)

func {{ .MiddlewareName }}Middleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // TODO: Middleware logic here
        return c.Next()
    }
}
`
