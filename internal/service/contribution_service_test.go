package service_test

import (
    "errors"
    "testing"

    "github.com/ikennarichard/contrib-tracker/internal/repository"
    "github.com/ikennarichard/contrib-tracker/internal/service"
)

func newTestService() (*service.ContributionService, *repository.MockContributionRepo) {
    mock := &repository.MockContributionRepo{}
    svc := service.New(mock)
    return svc, mock
}

func TestAdd_Valid(t *testing.T) {
    svc, mock := newTestService()

    err := svc.Add("Fix bug", "owner/repo", "2025-04-01", "")
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if len(mock.Contributions) != 1 {
        t.Errorf("expected 1 contribution in store, got %d", len(mock.Contributions))
    }
}

func TestAdd_InvalidTitle(t *testing.T) {
    svc, _ := newTestService()

    err := svc.Add("", "owner/repo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected validation error for empty title, got nil")
    }
}

func TestAdd_InvalidRepo(t *testing.T) {
    svc, _ := newTestService()

    err := svc.Add("Fix bug", "badrepo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected validation error for bad repo, got nil")
    }
}

func TestAdd_RepoError(t *testing.T) {
    svc, mock := newTestService()
    mock.AddError = errors.New("db is down")

    err := svc.Add("Fix bug", "owner/repo", "2025-04-01", "")
    if err == nil {
        t.Fatal("expected error when repo fails, got nil")
    }
}

func TestList_Empty(t *testing.T) {
    svc, _ := newTestService()

    contribs, err := svc.List()
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if len(contribs) != 0 {
        t.Errorf("expected empty list, got %d items", len(contribs))
    }
}

func TestList_ReturnsSaved(t *testing.T) {
    svc, _ := newTestService()

    _ = svc.Add("Fix bug",       "owner/repo", "2025-04-01", "")
    _ = svc.Add("Add dark mode", "owner/repo", "2025-04-02", "")

    contribs, err := svc.List()
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if len(contribs) != 2 {
        t.Errorf("expected 2 contributions, got %d", len(contribs))
    }
}

func TestList_RepoError(t *testing.T) {
    svc, mock := newTestService()
    mock.ListError = errors.New("db is down")

    _, err := svc.List()
    if err == nil {
        t.Fatal("expected error when repo fails, got nil")
    }
}

func TestFindByRepo_Match(t *testing.T) {
    svc, _ := newTestService()

    _ = svc.Add("Fix bug", "owner/repo",  "2025-04-01", "")
    _ = svc.Add("Add CI",  "other/repo",  "2025-04-02", "")

    contribs, err := svc.FindByRepo("owner/repo")
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if len(contribs) != 1 {
        t.Errorf("expected 1 match, got %d", len(contribs))
    }
    if contribs[0].Repo != "owner/repo" {
        t.Errorf("expected repo 'owner/repo', got '%s'", contribs[0].Repo)
    }
}

func TestFindByRepo_NoMatch(t *testing.T) {
    svc, _ := newTestService()

    _ = svc.Add("Fix bug", "owner/repo", "2025-04-01", "")

    contribs, err := svc.FindByRepo("nobody/norepo")
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }
    if len(contribs) != 0 {
        t.Errorf("expected 0 matches, got %d", len(contribs))
    }
}

func TestFindByRepo_EmptyRepo(t *testing.T) {
    svc, _ := newTestService()

    _, err := svc.FindByRepo("")
    if err == nil {
        t.Fatal("expected error for empty repo, got nil")
    }
}

func TestFindByRepo_RepoError(t *testing.T) {
    svc, mock := newTestService()
    mock.FindError = errors.New("db is down")

    _, err := svc.FindByRepo("owner/repo")
    if err == nil {
        t.Fatal("expected error when repo fails, got nil")
    }
}