package generator

func GenerateApp(moduleName string) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: moduleName,
	}
	return CreateFileFromTemplate("internal/app/app.go", AppTemplate, data)
}
