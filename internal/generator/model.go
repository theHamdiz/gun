package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

type Field struct {
	Name string
	Type string
}

func GenerateModel(name string, fields []Field) error {
	data := struct {
		ModelName string
		Fields    []Field
	}{
		ModelName: name,
		Fields:    fields,
	}

	modelFilePath := filepath.Join("internal", "models", fmt.Sprintf("%s.go", ToSnakeCase(name)))
	err := os.MkdirAll(filepath.Dir(modelFilePath), 0755)
	if err != nil {
		return err
	}

	return CreateFileFromTemplate(modelFilePath, ModelTemplate, data)
}
