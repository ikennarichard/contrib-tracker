package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/ikennarichard/contrib-tracker/internal/domain"
)

const dataFile = "contributions.json"
var mu sync.Mutex

// streaming + safe empty file handling
func loadContributions() ([]*domain.Contribution, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []*domain.Contribution{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var contributions []*domain.Contribution
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&contributions); err != nil {
		if errors.Is(err, io.EOF) {
			return []*domain.Contribution{}, nil
		}
		return nil, err
	}

	return contributions, nil
}

func saveContributions(contributions []*domain.Contribution) error {
	tmpFile := dataFile + ".tmp"

	f, err := os.Create(tmpFile)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(contributions); err != nil {
		f.Close()
		return err
	}
	
	if err := f.Close(); err != nil {
		return err
	}

	return os.Rename(tmpFile, dataFile)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd(os.Args[2:])
	case "list":
		listCmd(os.Args[2:])
	case "help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func addCmd(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	title := fs.String("title", "", "Contribution title")
	repo := fs.String("repo", "", "Repository (owner/repo)")
	date := fs.String("date", "", "Date (YYYY-MM-DD)")
	url := fs.String("url", "", "URL")

	fs.Parse(args)

	mu.Lock()
	defer mu.Unlock() // no over lap or data loss

	contributions, err := loadContributions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading contributions:", err)
		os.Exit(1)
	}

 contrib, err := domain.NewContribution(*title, *repo, *date, *url)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Validation error:", err)
        os.Exit(1)
    }

	contributions = append(contributions, contrib)

	if err := saveContributions(contributions); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving:", err)
		os.Exit(1)
	}

	fmt.Println("Contribution added")
}

func listCmd(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.Parse(args)

	contributions, err := loadContributions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if len(contributions) == 0 {
		fmt.Println("No contributions yet.")
		return
	}

	for i, c := range contributions {
		fmt.Printf("%d. %s\n", i+1, c.Title)
		fmt.Printf("   Repo: %s\n", c.Repo)
		fmt.Printf("   Date: %s\n", c.Date)
		if c.URL != "" {
			fmt.Printf("   URL:  %s\n", c.URL)
		}
		fmt.Println()
	}
}

func printUsage() {
	fmt.Println("Go CLI for tracking contributions")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run main.go <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add     Add a new contribution")
	fmt.Println("  list    List all contributions")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go add --title \"Fix bug\" --repo \"owner/repo\"")
	fmt.Println("  go run main.go list")
}