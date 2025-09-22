package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

//go:embed templates
var templatesFS embed.FS

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

	// Create .gitignore
	if err := createGitignore(); err != nil {
		fmt.Printf("Error creating .gitignore: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ .gitignore created")

	// Create Taskfile.yml
	if err := createTaskfile(moduleName); err != nil {
		fmt.Printf("Error creating Taskfile.yml: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Taskfile.yml created")

	if err := runCommand("git", "add", "-A"); err != nil {
		fmt.Printf("Warning: failed to stage files for initial commit: %v\n", err)
		os.Exit(1)
	}

	if err := runCommand("git", "commit", "-m", "Initial commit"); err != nil {
		fmt.Printf("Warning: failed to create initial commit: %v\n", err)
		fmt.Println("Hint: ensure git user.name and user.email are configured.")
		os.Exit(1)
	}
	fmt.Println("✅ Initial commit created")

	fmt.Println("\n🎉 Go project setup completed successfully!")
	fmt.Println("Next steps:")
	fmt.Println("  go run main.go    # Run the hello world program")
	fmt.Println("  task --list       # List available tasks")
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createReadme(moduleName string) error {
	projectName := filepath.Base(moduleName)

	tmpl, err := template.ParseFS(templatesFS, "templates/readme.tmpl")
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
	tmpl, err := template.ParseFS(templatesFS, "templates/main.tmpl")
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

func createTaskfile(moduleName string) error {
	projectName := filepath.Base(moduleName)

	tmpl, err := template.ParseFS(templatesFS, "templates/taskfile.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse taskfile template: %v", err)
	}

	file, err := os.Create("Taskfile.yml")
	if err != nil {
		return fmt.Errorf("failed to create Taskfile.yml: %v", err)
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
		return fmt.Errorf("failed to execute taskfile template: %v", err)
	}

	return nil
}

func createGitignore() error {
	tmpl, err := template.ParseFS(templatesFS, "templates/gitignore.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse gitignore template: %v", err)
	}

	file, err := os.Create(".gitignore")
	if err != nil {
		return fmt.Errorf("failed to create .gitignore: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, nil)
	if err != nil {
		return fmt.Errorf("failed to execute gitignore template: %v", err)
	}

	return nil
}
