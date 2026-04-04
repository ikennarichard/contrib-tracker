package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type Contribution struct {
	Title string `json:"title"`
	Repo  string `json:"repo"`
	Date  string `json:"date"`
	URL   string `json:"url"`
}

const dataFile = "contributions.json"

// streaming + safe empty file handling
func loadContributions() ([]Contribution, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Contribution{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var contribs []Contribution
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&contribs); err != nil {
		if errors.Is(err, io.EOF) {
			return []Contribution{}, nil
		}
		return nil, err
	}

	return contribs, nil
}

func saveContributions(contribs []Contribution) error {
	tmpFile := dataFile + ".tmp"

	f, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(contribs); err != nil {
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

	if *title == "" || *repo == "" {
		fmt.Fprintln(os.Stderr, "Error: --title and --repo are required")
		fs.Usage()
		os.Exit(1)
	}

	if *date == "" {
		*date = time.Now().Format("2006-01-02")
	}

	contribs, err := loadContributions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading contributions:", err)
		os.Exit(1)
	}

	newContrib := Contribution{
		Title: *title,
		Repo:  *repo,
		Date:  *date,
		URL:   *url,
	}

	contribs = append(contribs, newContrib)

	if err := saveContributions(contribs); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving:", err)
		os.Exit(1)
	}

	fmt.Println("Contribution added")
}

func listCmd(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.Parse(args)

	contribs, err := loadContributions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if len(contribs) == 0 {
		fmt.Println("No contributions yet.")
		return
	}

	for i, c := range contribs {
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
	fmt.Println("contrib-cli — Go CLI for tracking contributions")
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