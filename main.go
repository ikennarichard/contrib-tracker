package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Contribution struct {
    Title string `json:"title"`
    Repo  string `json:"repo"`
    Date  string `json:"date"`
    URL   string `json:"url"`
}

// in memory store
var contributions []Contribution

func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    switch os.Args[1] {
    case "add":
        addContribution()
    case "list":
        listContributions()
    default:
        fmt.Printf("Unknown command: %s\n\n", os.Args[1])
        printUsage()
        os.Exit(1)
    }
}

func printUsage() {
    fmt.Println("contrib-cli — Modern Go CLI for tracking contributions")
    fmt.Println()
    fmt.Println("Usage:")
    fmt.Println("  go run main.go add --title \"Title\" --repo \"owner/repo\" --date \"2025-04-01\" --url \"https://...\"")
    fmt.Println("  go run main.go list")
}

func addContribution() {
    cmd := flag.NewFlagSet("add", flag.ExitOnError)
    title := cmd.String("title", "", "Contribution title (required)")
    repo := cmd.String("repo", "", "Repository (e.g. owner/repo) (required)")
    dateStr := cmd.String("date", "", "Date in YYYY-MM-DD (required)")
    url := cmd.String("url", "", "Pull request / issue URL (required)")

    cmd.Parse(os.Args[2:])

    if *title == "" || *repo == "" || *dateStr == "" || *url == "" {
        fmt.Println("All flags are required.")
        cmd.PrintDefaults()
        os.Exit(1)
    }

    // date validation
    if _, err := time.Parse("2006-01-01", *dateStr); err != nil {
        fmt.Println("Invalid date format. Use YYYY-MM-DD (e.g. 2025-04-01)")
        os.Exit(1)
    }

    c := Contribution{
        Title: *title,
        Repo:  *repo,
        Date:  *dateStr,
        URL:   *url,
    }

    contributions = append(contributions, c)

    fmt.Printf("Contribution added!\n")
    fmt.Printf("Title: %s\n", *title)
    fmt.Printf("Repo: %s\n", *repo)
    fmt.Printf("Date: %s\n", *dateStr)
    fmt.Printf("URL: %s\n", *url)
}

func listContributions() {
    if len(contributions) == 0 {
        fmt.Println("No contributions yet. Add some with:")
        fmt.Println("   go run main.go add --title \"...\" --repo \"...\" --date \"...\" --url \"...\"")
        return
    }

    fmt.Printf("Your contributions (%d total)\n\n", len(contributions))

    for i, c := range contributions {
        fmt.Printf("%d. %s\n", i+1, c.Title)
        fmt.Printf(" Repo: %s\n", c.Repo)
        fmt.Printf(" Date: %s\n", c.Date)
        fmt.Printf(" URL: %s\n\n", c.URL)
    }
}