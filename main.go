package main

import (
	"fmt"
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
	content := fmt.Sprintf("# %s\n\nA Go project.\n\n## Getting Started\n\n### Prerequisites\n\n- Go 1.19 or later\n\n### Installation\n\n```bash\ngit clone %s\ncd %s\ngo mod tidy\n```\n\n### Usage\n\n```bash\ngo run main.go\n```\n\n### Build\n\n```bash\ngo build -o %s\n./%s\n```\n\n## License\n\nThis project is licensed under the MIT License.\n", projectName, moduleName, projectName, projectName, projectName)

	return os.WriteFile("README.md", []byte(content), 0644)
}

func createMainGo() error {
	content := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	return os.WriteFile("main.go", []byte(content), 0644)
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
