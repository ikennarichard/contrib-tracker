package repository

import "github.com/ikennarichard/contrib-tracker/internal/domain"

type ContributionRepository interface {
    Add(c *domain.Contribution) error
    List() ([]*domain.Contribution, error)
    FindByRepo(repo string) ([]*domain.Contribution, error)
}