# contrib-tracker

A lightweight CLI tool for logging and viewing your personal open-source contributions, built with Go.

## Features

- Add contributions with a title, repository, date, and optional URL
- List all logged contributions in a clean, readable format
- Stores data locally in a JSON file (`contributions.json`)
- Auto-fills today's date if none is provided
- Safe concurrent writes using mutex locking and atomic file rename

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

## Installation

Clone the repository and navigate into it:

```bash
git clone https://github.com/ikennarichard/contrib-tracker.git
cd contrib-tracker
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

## Data Storage

Contributions are persisted locally in a `contributions.json` file in the project root. This file is created automatically on the first `add` command. You can back it up, version-control it, or share it as needed.

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on getting started, making changes, and submitting a pull request.

Some ideas for contributions:

- `delete` command to remove a contribution by index
- `edit` command to update an existing entry
- Filter support for `list` (e.g. `--repo=owner/repo`)
- Colorized and table-formatted output
- Export to CSV or Markdown

## License

[MIT](./LICENSE)