package service

import (
    "fmt"

    "github.com/ikennarichard/contrib-tracker/internal/domain"
    "github.com/ikennarichard/contrib-tracker/internal/repository"
)

type ContributionService struct {
    repo repository.ContributionRepository
}

// New creates a new ContributionService.
func New(repo repository.ContributionRepository) *ContributionService {
    return &ContributionService{repo: repo}
}

func (s *ContributionService) Add(title, repo, date, url string) error {
    contrib, err := domain.NewContribution(title, repo, date, url)
    if err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }

    if err := s.repo.Add(contrib); err != nil {
        return fmt.Errorf("could not save contribution: %w", err)
    }

    return nil
}

func (s *ContributionService) List() ([]*domain.Contribution, error) {
    contribs, err := s.repo.List()
    if err != nil {
        return nil, fmt.Errorf("could not fetch contributions: %w", err)
    }
    return contribs, nil
}

func (s *ContributionService) FindByRepo(repo string) ([]*domain.Contribution, error) {
    if repo == "" {
        return nil, fmt.Errorf("repo name is required")
    }

    contribs, err := s.repo.FindByRepo(repo)
    if err != nil {
        return nil, fmt.Errorf("could not fetch contributions for %s: %w", repo, err)
    }
    return contribs, nil
}