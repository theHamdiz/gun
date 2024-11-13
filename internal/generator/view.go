package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateViews(resourceName string, fields []Field) error {
	views := []string{"index", "show", "edit", "new"}
	viewsDir := filepath.Join("internal", "views", ToSnakeCase(resourceName))
	err := os.MkdirAll(viewsDir, 0755)
	if err != nil {
		return err
	}

	for _, view := range views {
		data := struct {
			ResourceName string
			Fields       []Field
		}{
			ResourceName: resourceName,
			Fields:       fields,
		}

		viewFilePath := filepath.Join(viewsDir, fmt.Sprintf("%s.html", view))
		var templateContent string

		switch view {
		case "index":
			templateContent = IndexTemplate
		case "show":
			templateContent = ShowTemplate
		case "edit":
			templateContent = EditTemplate
		case "new":
			templateContent = NewTemplate
		}

		err := CreateFileFromTemplate(viewFilePath, templateContent, data)
		if err != nil {
			return err
		}
	}
	return nil
}

// IndexTemplate is the template for index.html
var IndexTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .ResourceName }} List</title>
    <link href="/css/tailwind.css" rel="stylesheet">
</head>
<body>
    <h1>{{ .ResourceName }} List</h1>
    <table>
        <thead>
            <tr>
            {{- range .Fields }}
                <th>{{ .Name }}</th>
            {{- end }}
            </tr>
        </thead>
        <tbody>
            {{ "{{ range .Items }}" }}
            <tr>
            {{- range .Fields }}
                <td>{{ "{{ .{{ .Name }} }}" }}</td>
            {{- end }}
            </tr>
            {{ "{{ end }}" }}
        </tbody>
    </table>
    <a href="/{{ ToLower .ResourceName }}s/new">Create New {{ .ResourceName }}</a>
</body>
</html>
`

// ShowTemplate is the template for show.html
var ShowTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .ResourceName }} Details</title>
    <link href="/css/tailwind.css" rel="stylesheet">
</head>
<body>
    <h1>{{ .ResourceName }} Details</h1>
    {{- range .Fields }}
    <p>{{ .Name }}: {{ "{{ ." }}{{ .Name }}{{ " }}" }}</p>
    {{- end }}
    <a href="/{{ ToLower .ResourceName }}s/{{ "{{ ." }}ID{{ " }}/edit" }}">Edit {{ .ResourceName }}</a>
    <a href="/{{ ToLower .ResourceName }}s">Back to List</a>
</body>
</html>
`

// EditTemplate is the template for edit.html
var EditTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Edit {{ .ResourceName }}</title>
    <link href="/css/tailwind.css" rel="stylesheet">
</head>
<body>
    <h1>Edit {{ .ResourceName }}</h1>
    <form method="POST" action="/{{ ToLower .ResourceName }}s/{{ "{{ ." }}ID{{ " }}" }}">
    {{- range .Fields }}
        <label>{{ .Name }}</label>
        <input type="text" name="{{ .Name }}" value="{{ "{{ ." }}{{ .Name }}{{ " }}" }}">
    {{- end }}
        <button type="submit">Update</button>
    </form>
</body>
</html>
`

// NewTemplate is the template for new.html
var NewTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Create {{ .ResourceName }}</title>
    <link href="/css/tailwind.css" rel="stylesheet">
</head>
<body>
    <h1>Create {{ .ResourceName }}</h1>
    <form method="POST" action="/{{ ToLower .ResourceName }}s">
    {{- range .Fields }}
        <label>{{ .Name }}</label>
        <input type="text" name="{{ .Name }}">
    {{- end }}
        <button type="submit">Create</button>
    </form>
</body>
</html>
`
