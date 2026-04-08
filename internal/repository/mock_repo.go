package repository

import "github.com/ikennarichard/contrib-tracker/internal/domain"

type MockContributionRepo struct {
    Contributions []*domain.Contribution
    AddError      error
    ListError     error
    FindError     error
}

func (m *MockContributionRepo) Add(c *domain.Contribution) error {
    if m.AddError != nil {
        return m.AddError
    }
    m.Contributions = append(m.Contributions, c)
    return nil
}

func (m *MockContributionRepo) List() ([]*domain.Contribution, error) {
    if m.ListError != nil {
        return nil, m.ListError
    }
    return m.Contributions, nil
}

func (m *MockContributionRepo) FindByRepo(repo string) ([]*domain.Contribution, error) {
    if m.FindError != nil {
        return nil, m.FindError
    }
    var results []*domain.Contribution
    for _, c := range m.Contributions {
        if c.Repo == repo {
            results = append(results, c)
        }
    }
    return results, nil
}