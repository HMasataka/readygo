package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run setup.go <module-name>")
		os.Exit(1)
	}

	moduleName := os.Args[1]
	fmt.Printf("Setting up Go project: %s\n", moduleName)

	if err := runCommand("mkdir", moduleName); err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.Chdir(moduleName); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize Go module
	if err := runCommand("go", "mod", "init"); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Go module initialized")

	// Initialize Git repository
	if err := runCommand("git", "init"); err != nil {
		fmt.Printf("Error initializing Git repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Git repository initialized")

	// Create README.md
	if err := createReadme(moduleName); err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ README.md created")

	// Create Hello World main.go
	if err := createMainGo(); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Hello World main.go created")

	// Download .gitignore
	if err := downloadGitignore(); err != nil {
		fmt.Printf("Error downloading .gitignore: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ .gitignore downloaded")

	fmt.Println("\n🎉 Go project setup completed successfully!")
	fmt.Println("Next steps:")
	fmt.Println("  go run main.go    # Run the hello world program")
	fmt.Println("  git add .")
	fmt.Println("  git commit -m \"Initial commit\"")
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createReadme(moduleName string) error {
	projectName := filepath.Base(moduleName)

	readmeTemplate := `# {{.ProjectName}}

A Go project.

## Getting Started

### Prerequisites

- Go 1.19 or later

### Installation

` + "```bash" + `
git clone {{.ModuleName}}
cd {{.ProjectName}}
go mod tidy
` + "```" + `

### Usage

` + "```bash" + `
go run main.go
` + "```" + `

### Build

` + "```bash" + `
go build -o {{.ProjectName}}
./{{.ProjectName}}
` + "```" + `

## License

This project is licensed under the MIT License.
`

	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse readme template: %v", err)
	}

	file, err := os.Create("README.md")
	if err != nil {
		return fmt.Errorf("failed to create README.md: %v", err)
	}
	defer file.Close()

	data := struct {
		ProjectName string
		ModuleName  string
	}{
		ProjectName: projectName,
		ModuleName:  moduleName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute readme template: %v", err)
	}

	return nil
}

func createMainGo() error {
	mainGoTemplate := `package main

import "fmt"

func main() {
	fmt.Println("{{.Message}}")
}
`

	tmpl, err := template.New("main").Parse(mainGoTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse main.go template: %v", err)
	}

	file, err := os.Create("main.go")
	if err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}
	defer file.Close()

	data := struct {
		Message string
	}{
		Message: "Hello, World!",
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute main.go template: %v", err)
	}

	return nil
}

func downloadGitignore() error {
	url := "https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download .gitignore: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download .gitignore: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(".gitignore")
	if err != nil {
		return fmt.Errorf("failed to create .gitignore file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write .gitignore file: %v", err)
	}

	return nil
}
