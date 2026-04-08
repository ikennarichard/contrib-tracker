package domain_test

import (
    "testing"

    "github.com/ikennarichard/contrib-tracker/internal/domain"
)

// helpers

func TestNewContribution_Valid(t *testing.T) {
    c, err := domain.NewContribution("Fix bug", "owner/repo", "2025-04-01", "https://github.com/owner/repo/pull/1")
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if c.Title != "Fix bug" {
        t.Errorf("expected title 'Fix bug', got '%s'", c.Title)
    }
    if c.Repo != "owner/repo" {
        t.Errorf("expected repo 'owner/repo', got '%s'", c.Repo)
    }
}

// tests

func TestNewContribution_EmptyTitle(t *testing.T) {
    _, err := domain.NewContribution("", "owner/repo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected error for empty title, got nil")
    }
}

func TestNewContribution_WhitespaceTitle(t *testing.T) {
    _, err := domain.NewContribution("   ", "owner/repo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected error for whitespace-only title, got nil")
    }
}

func TestNewContribution_EmptyRepo(t *testing.T) {
    _, err := domain.NewContribution("Fix bug", "", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected error for empty repo, got nil")
    }
}

func TestNewContribution_RepoMissingSlash(t *testing.T) {
    _, err := domain.NewContribution("Fix bug", "badrepo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected error for repo missing slash, got nil")
    }
}


func TestNewContribution_InvalidDateFormat(t *testing.T) {
    _, err := domain.NewContribution("Fix bug", "owner/repo", "01-04-2025", "")
    if err == nil {
        t.Fatal("expected error for invalid date format, got nil")
    }
}

func TestNewContribution_EmptyDateDefaultsToToday(t *testing.T) {
    c, err := domain.NewContribution("Fix bug", "owner/repo", "", "")
    if err != nil {
        t.Fatalf("expected no error when date is empty, got: %v", err)
    }
    if c.Date == "" {
        t.Error("expected date to be set to today, got empty string")
    }
}

func TestNewContribution_URLIsOptional(t *testing.T) {
    _, err := domain.NewContribution("Fix bug", "owner/repo", "2025-04-01", "")
    if err != nil {
        t.Fatalf("expected no error when URL is empty, got: %v", err)
    }
}

func TestNewContribution_InvalidInputs(t *testing.T) {
    tests := []struct {
        name  string
        title string
        repo  string
        date  string
        url   string
    }{
        {"empty title",         "",      "owner/repo", "2025-04-01", ""},
        {"empty repo",          "title", "",           "2025-04-01", ""},
        {"repo no slash",       "title", "badrepo",    "2025-04-01", ""},
        {"bad date format",     "title", "owner/repo", "not-a-date", ""},
        {"whitespace title",    "  ",    "owner/repo", "2025-04-01", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := domain.NewContribution(tt.title, tt.repo, tt.date, tt.url)
            if err == nil {
                t.Errorf("expected validation error for case '%s', got nil", tt.name)
            }
        })
    }
}