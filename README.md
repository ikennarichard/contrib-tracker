# contrib-tracker

A lightweight CLI tool for logging and viewing your personal open-source contributions, built with Go.

## Features

- Add contributions with a title, repository, date, and optional URL
- List all logged contributions ordered by date
- Filter contributions by repository
- Auto-fills today's date if none is provided
- Input validation on the domain entity
- Persistent storage in PostgreSQL using prepared statements
- Clean architecture — domain, repository, service, and CLI layers fully separated

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher
- [Docker](https://docker.com/) (for running PostgreSQL locally)

## Installation

Clone the repository and navigate into it:

```bash
git clone https://github.com/ikennarichard/contrib-tracker.git
cd contrib-tracker
go mod tidy
```

### Database Setup

Start a PostgreSQL container with Docker:

```bash
docker run -p 5435:5432 --name contrib-db \
  -e POSTGRES_PASSWORD=yourpassword \
  -e POSTGRES_DB=contrib_tracker \
  -d postgres
```

If the container already exists from a previous run, just start it:

```bash
docker start contrib-db
```

### Environment Variables

```bash
DATABASE_URL="postgres://postgres:yourpassword@localhost:5435/contrib_tracker?sslmode=disable"
```

## Usage

### Add a contribution

```bash
go run main.go add --title "Fix bug" --repo "owner/repo"
```

**All flags:**

| Flag | Required | Description |
|------|----------|-------------|
| `--title` | ✅ | A short title describing the contribution |
| `--repo` | ✅ | Repository in `owner/repo` format |
| `--date` | ❌ | Date in `YYYY-MM-DD` format (defaults to today) |
| `--url` | ❌ | Link to the PR, issue, or commit |

**Example:**

```bash
go run main.go add \
  --title "Add dark mode support" \
  --repo "vercel/next.js" \
  --date "2025-04-01" \
  --url "https://github.com/vercel/next.js/pull/12345"
```

### List all contributions

```bash
go run main.go list
```

**Example output:**

```
1. Add dark mode support
   Repo: vercel/next.js
   Date: 2025-04-01
   URL:  https://github.com/vercel/next.js/pull/12345

2. Fix typo in README
   Repo: golang/go
   Date: 2025-03-28
```

### Help

```bash
go run main.go help
```

## How to run server

```bash
go run main.go -server

go run main.go -server -port 3000 # use custom port
```

## How to TS client

```bash
tsx contrib.ts add --title "Fixed login bug" --repo "ikennarichard/contrib-tracker" --url "https://github.com/ikennarichard/contrib-tracker/pull/5"

# List contributions
tsx contrib.ts list
tsx contrib.ts list --repo "ikennarichard/contrib-tracker"
```

## Running Tests

Tests use only the standard testing package and an in-memory mock repository — no database or Docker required.

Run all tests:

```bash
go test ./...
```

Run specific package

```bash
go test ./internal/domain/...
go test ./internal/service/...
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on getting started, making changes, and submitting a pull request.

Some ideas for contributions:

- `delete` command to remove a contribution by index
- `edit` command to update an existing entry
- Colorized and table-formatted output
- Export to CSV or Markdown

## License

[MIT](./LICENSE)