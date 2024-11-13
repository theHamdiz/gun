package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/theHamdiz/gun/internal/generator"
)

func CreateProject(name string, withChannels, withSignals bool) error {
	proj := Project{
		Name:         name,
		ModuleName:   name,
		WithChannels: withChannels,
		WithSignals:  withSignals,
	}

	// Use proj.ModuleName for go.mod initialization
	if err := initializeGoMod(proj.ModuleName); err != nil {
		return err
	}

	// Use proj in other functions
	if err := createMainFile(proj); err != nil {
		return err
	}

	// Install dependencies
	if err := installDependencies(); err != nil {
		return err
	}

	// Include channels and signals if requested
	if withChannels || withSignals {
		if err := createUtils(withChannels, withSignals); err != nil {
			return err
		}
	}

	fmt.Println("Project scaffolding complete.")
	return nil
}

func initializeGoMod(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	return cmd.Run()
}

func createMainFile(proj Project) error {
	data := struct {
		ProjectName string
		ModuleName  string
	}{
		ProjectName: proj.Name,
		ModuleName:  proj.ModuleName,
	}

	return generator.CreateFileFromTemplate("main.go", generator.MainTemplate, data)
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
