package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ikennarichard/contrib-tracker/internal/postgres"
	"github.com/ikennarichard/contrib-tracker/internal/server"
	"github.com/ikennarichard/contrib-tracker/internal/service"
	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file", err)
    }

    runServer := flag.Bool("server", false, "Run as HTTP server")
    port := flag.String("port", "8080", "Port for HTTP server")
    flag.Parse()

    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    repo, err := postgres.New(connStr)
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer repo.Close()

    svc := service.New(repo)

        if *runServer {
        fmt.Printf("Starting http server on http://localhost:%s\n", *port)
        srv := server.NewServer(svc, *port)
        log.Fatal(srv.Start())
    } else {
        runCLI(svc)
    }
}

func runCLI(svc *service.ContributionService) {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    switch os.Args[1] {
    case "add":
        addCmd(os.Args[2:], svc)
    case "list":
        listCmd(svc)
    case "filter":
        filterCmd(os.Args[2:], svc)
    default:
        fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
        printUsage()
        os.Exit(1)
    }
}

func addCmd(args []string, svc *service.ContributionService) {
    fs := flag.NewFlagSet("add", flag.ExitOnError)
    title := fs.String("title", "", "Contribution title")
    repo  := fs.String("repo", "", "Repository (owner/repo)")
    date  := fs.String("date", "", "Date (YYYY-MM-DD), defaults to today")
    url   := fs.String("url", "", "URL to PR, issue or commit")
    fs.Parse(args)

    if err := svc.Add(*title, *repo, *date, *url); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }

    fmt.Println("Contribution added.")
}

func listCmd(svc *service.ContributionService) {
    contribs, err := svc.List()
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

func filterCmd(args []string, svc *service.ContributionService) {
    fs := flag.NewFlagSet("filter", flag.ExitOnError)
    repo := fs.String("repo", "", "Filter by repository (owner/repo)")
    fs.Parse(args)

    if *repo == "" {
        fmt.Fprintln(os.Stderr, "Error: --repo is required")
        os.Exit(1)
    }

    contribs, err := svc.FindByRepo(*repo)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }

    if len(contribs) == 0 {
        fmt.Printf("No contributions found for %s\n", *repo)
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
    fmt.Println("Usage:")
    fmt.Println("  go run main.go <command> [options]")
    fmt.Println()
    fmt.Println("Commands:")
    fmt.Println("  add     Add a new contribution")
    fmt.Println("  list    List all contributions")
    fmt.Println("  filter  Filter contributions by repo")
    fmt.Println()
    fmt.Println("Examples:")
    fmt.Println(`  go run main.go add --title "Fix bug" --repo "owner/repo"`)
    fmt.Println(`  go run main.go list`)
    fmt.Println(`  go run main.go filter --repo "owner/repo"`)
}