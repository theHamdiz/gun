package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

func CreateProject(name, style, moduleName string, withChannels, withSignals bool) error {
	proj := Project{
		Name:         name,
		ModuleName:   moduleName,
		Style:        style,
		WithChannels: withChannels,
		WithSignals:  withSignals,
	}

	// Initialize go.mod
	if err := initializeGoMod(proj.ModuleName); err != nil {
		it.Errorf("Failed to initialize go.mod: %v", err)
		return err
	}

	// Create directories
	if err := createDirectories(); err != nil {
		it.Errorf("Failed to create directories: %v", err)
		return err
	}

	// Create main.go in cmd/http/server/
	if err := createServerMainFile(proj); err != nil {
		it.Errorf("Failed to create main.go: %v", err)
		return err
	}

	// Create apiv1.go in cmd/http/api/v1/
	if err := createAPIV1File(proj); err != nil {
		it.Errorf("Failed to create apiv1.go: %v", err)
		return err
	}

	// Create app wrapper
	if err := createAppWrapper(proj); err != nil {
		it.Errorf("Failed to create app wrapper: %v", err)
		return err
	}

	// Include channels and signals if requested
	if proj.WithChannels || proj.WithSignals {
		if err := createUtils(proj.WithChannels, proj.WithSignals); err != nil {
			it.Errorf("Failed to create utils: %v", err)
			return err
		}
	}

	// Setup styling
	if err := setupStyling(proj); err != nil {
		it.Errorf("Failed to setup styling: %v", err)
		return err
	}

	fmt.Println("Project scaffolding complete.")
	return nil
}

func setupStyling(proj Project) error {
	switch proj.Style {
	case "tailwind":
		return setupTailwind(proj)
	case "shadcn":
		return setupShadcnUI(proj)
	case "both":
		if err := setupTailwind(proj); err != nil {
			return err
		}
		return setupShadcnUI(proj)
	default:
		fmt.Println("No styling framework selected.")
		return nil
	}
}

func setupTailwind(proj Project) error {
	fmt.Println("Setting up Tailwind CSS...")

	// Install Node.js dependencies
	if err := exec.Command("deno", "init", "-y").Run(); err != nil {
		return err
	}
	if err := exec.Command("deno", "i", "-D", "tailwindcss", "postcss", "autoprefixer").Run(); err != nil {
		return err
	}
	if err := exec.Command("deno", "tailwindcss", "init", "-p").Run(); err != nil {
		return err
	}

	// Create Tailwind CSS input file
	err := os.MkdirAll(filepath.Join("assets", "css"), 0755)
	if err != nil {
		return err
	}
	inputCSSPath := filepath.Join("assets", "css", "input.css")
	inputCSSContent := `@tailwind base;
@tailwind components;
@tailwind utilities;`
	if err := os.WriteFile(inputCSSPath, []byte(inputCSSContent), 0644); err != nil {
		return err
	}

	// Update tailwind.config.js content if needed

	// Add build script to package.json
	// (Additional code to modify package.json can be added here)

	fmt.Println("Tailwind CSS setup complete.")
	return nil
}

func setupShadcnUI(proj Project) error {
	fmt.Println("Setting up shadcn/ui...")

	// Ensure Tailwind CSS is set up first
	if err := setupTailwind(proj); err != nil {
		return err
	}

	// Install shadcn/ui components (hypothetical example)
	if err := exec.Command("deno", "i", "@shadcn/ui").Run(); err != nil {
		return err
	}

	// Copy shadcn/ui component files into the project
	// (Additional code to handle shadcn/ui setup)

	fmt.Println("shadcn/ui setup complete.")
	return nil
}

func createDirectories() error {
	dirs := []string{
		"cmd/http/server",
		"cmd/http/api/v1",
		"internal/app",
		"internal/models",
		"internal/handlers",
		"internal/routes",
		"internal/middleware",
		"internal/views",
		"internal/utils",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func initializeGoMod(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	return cmd.Run()
}

func createServerMainFile(proj Project) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: proj.ModuleName,
	}

	mainFilePath := filepath.Join("cmd", "http", "server", "main.go")
	return generator.CreateFileFromTemplate(mainFilePath, generator.ServerMainTemplate, data)
}

func createAPIV1File(proj Project) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: proj.ModuleName,
	}

	apiFilePath := filepath.Join("cmd", "http", "api", "v1", "apiv1.go")
	return generator.CreateFileFromTemplate(apiFilePath, generator.APIV1Template, data)
}

func installDependencies() error {
	// Install Fiber
	cmd := exec.Command("go", "get", "github.com/gofiber/fiber/v2")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func createUtils(withChannels, withSignals bool) error {
	utilsDir := filepath.Join("internal", "utils")
	err := os.MkdirAll(utilsDir, 0755)
	if err != nil {
		return err
	}

	if withChannels {
		channelsPath := filepath.Join(utilsDir, "channels.go")
		err := generator.CreateFileFromTemplate(channelsPath, generator.ChannelsTemplate, nil)
		if err != nil {
			return err
		}
	}

	if withSignals {
		signalsPath := filepath.Join(utilsDir, "signals.go")
		err := generator.CreateFileFromTemplate(signalsPath, generator.SignalsTemplate, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func createAppWrapper(proj Project) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: proj.ModuleName,
	}

	appFilePath := filepath.Join(proj.Name, "internal", "app", "app.go")
	return generator.CreateFileFromTemplate(appFilePath, generator.AppTemplate, data)
}
