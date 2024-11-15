package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

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

	// Use a WaitGroup to handle asynchronous tasks
	var wg sync.WaitGroup
	var errChan = make(chan error, 4)

	// Create the base project directory with the given project name
	baseDir := strings.ToLower(proj.Name)

	// Create directories asynchronously
	if err := createDirectories(baseDir); err != nil {
		it.Errorf("❌ Failed to create directories : %v", err)
		return err
	}

	// Initialize go.mod asynchronously
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := initializeGoMod(baseDir, proj.ModuleName); err != nil {
			it.Errorf("❌ Failed to initialize go.mod: %v", err)
			errChan <- err
		}
	}()

	// Generate main.go in cmd/http/server/
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := createServerMainFile(baseDir, proj); err != nil {
			it.Errorf("❌ Failed to create main.go: %v", err)
			errChan <- err
		}
	}()

	// Generate apiv1.go in cmd/http/api/v1/
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := createAPIV1File(baseDir, proj); err != nil {
			it.Errorf("❌ Failed to create apiv1.go: %v", err)
			errChan <- err
		}
	}()

	// Create utility files if channels or signals are enabled
	if proj.WithChannels || proj.WithSignals {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := createUtils(baseDir, proj.WithChannels, proj.WithSignals); err != nil {
				it.Errorf("❌ Failed to create utils: %v", err)
				errChan <- err
			}
		}()
	}

	// Setup styling configuration based on user choice
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := setupStyling(baseDir, proj); err != nil {
			it.Errorf("❌ Failed to set up styling: %v", err)
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	// Check for errors during setup
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	it.Info("✅️ Project scaffolding complete.")
	return nil
}

func createDirectories(baseDir string) error {
	it.Infof("✅️ Creating project in %s", baseDir)
	// Attempt to create the base directory. If it exists, MkdirAll will not throw an error.
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		it.Errorf("❌ Failed to create base directory %s: %w", baseDir, err)
		return fmt.Errorf("❌ Failed to create base directory %s: %w", baseDir, err)
	}
	dirs := []string{
		filepath.Join(baseDir, "cmd/http/server"),
		filepath.Join(baseDir, "cmd/http/api/v1"),
		filepath.Join(baseDir, "internal/app"),
		filepath.Join(baseDir, "internal/db"),
		filepath.Join(baseDir, "internal/models"),
		filepath.Join(baseDir, "internal/handlers"),
		filepath.Join(baseDir, "internal/routes"),
		filepath.Join(baseDir, "internal/middleware"),
		filepath.Join(baseDir, "internal/views"),
		filepath.Join(baseDir, "internal/utils"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			it.LogErrorWithStack(err)
			return err
		}
	}
	return nil
}

func initializeGoMod(baseDir, moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = baseDir // Set the directory for go.mod initialization
	it.Infof("✅️ Initializing go.mod in %s", baseDir)
	return cmd.Run()
}

func createServerMainFile(baseDir string, proj Project) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: proj.ModuleName,
	}

	mainFilePath := filepath.Join(baseDir, "cmd", "http", "server", "main.go")
	it.Infof("✅️ Creating main.go in %s", mainFilePath)
	return generator.CreateFileFromTemplate(mainFilePath, generator.ServerMainTemplate, data)
}

func createAPIV1File(baseDir string, proj Project) error {
	data := struct {
		ModuleName string
	}{
		ModuleName: proj.ModuleName,
	}

	apiFilePath := filepath.Join(baseDir, "cmd", "http", "api", "v1", "apiv1.go")
	return generator.CreateFileFromTemplate(apiFilePath, generator.APIV1Template, data)
}

func createUtils(baseDir string, withChannels, withSignals bool) error {
	utilsDir := filepath.Join(baseDir, "internal", "utils")
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

func setupStyling(baseDir string, proj Project) error {
	switch proj.Style {
	case "tailwind":
		return setupTailwind(baseDir)
	case "shadcn":
		return setupShadcnUI(baseDir)
	case "both":
		if err := setupTailwind(baseDir); err != nil {
			return err
		}
		return setupShadcnUI(baseDir)
	default:
		it.Warn("🚧 No styling framework selected.")
		return nil
	}
}

func setupTailwind(baseDir string) error {
	it.Info("✅️ Setting up Tailwind CSS with deno...")
	cmd := exec.Command("deno", "init", "-y")
	cmd.Dir = baseDir

	// Install Node.js dependencies
	if err := cmd.Run(); err != nil {
		it.LogErrorWithStack(err)
		return err
	}
	if err := exec.Command("deno", "i", "-D", "tailwindcss", "postcss", "autoprefixer").Run(); err != nil {
		it.LogErrorWithStack(err)
		return err
	}
	if err := exec.Command("deno", "tailwindcss", "init", "-p").Run(); err != nil {
		it.LogErrorWithStack(err)
		return err
	}

	// Create Tailwind CSS input file
	err := os.MkdirAll(filepath.Join("assets", "css"), 0755)
	if err != nil {
		it.LogErrorWithStack(err)
		return err
	}
	inputCSSPath := filepath.Join("assets", "css", "input.css")
	inputCSSContent := `@tailwind base;
@tailwind components;
@tailwind utilities;`
	if err := os.WriteFile(inputCSSPath, []byte(inputCSSContent), 0644); err != nil {
		it.LogErrorWithStack(err)
		return err
	}

	/* Todo:
	   Update tailwind.config.js content if needed
	   Add build script to package.json
	   (Additional code to modify package.json can be added here) */

	it.Info("✅️ Tailwind CSS setup complete.")
	return nil
}

func setupShadcnUI(baseDir string) error {
	it.Info("✅️ Setting up shadcn/ui...")

	// Ensure Tailwind CSS is set up first
	if err := setupTailwind(baseDir); err != nil {
		it.LogErrorWithStack(err)
		return err
	}

	// Install shadcn/ui components (hypothetical example)
	cmd := exec.Command("deno", "i", "@shadcn/ui")
	cmd.Dir = baseDir
	if err := cmd.Run(); err != nil {
		it.LogErrorWithStack(err)
		return err
	}

	// Copy shadcn/ui component files into the project
	// (Additional code to handle shadcn/ui setup)
	it.Info("✅️ Shadcn/ui setup complete.")
	return nil
}
