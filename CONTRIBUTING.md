# Contributing to contrib-tracker

Thanks for your interest in contributing!

## Getting Started

1. Clone the repo:

   ```bash
   git clone <repo-url>
   cd contrib-tracker
   ```

2. Clone the repo:

   ```bash
   go run main.go list
   ```

## How to contribute

1. Create a branch:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes
   Keep code simple and readable
   Follow Go best practices
   Handle errors properly

3. Test your changes

```bash
   go run main.go add --title "Test" --repo "owner/repo"
   go run main.go list
```

1. Commit, Push and open a PR

   ```bash
   git commit -m "feat: add inverted commit"
   git push origin feature/your-feature-name
   ```

## Suggestions

You can contribute by:

- Adding new commands (e.g. delete, edit)
- Improving CLI output (colors, tables)
- Adding filters (list --repo=...)
- Improving performance or structure

## Reporting Issues

If you find a bug:

- Describe what happened
- Include steps to reproduce
- Share expected vs actual behavior

Thanks for contributing ❤️
