package domain

import (
    "errors"
    "strings"
    "time"
)

// Contribution is the core business entity.
type Contribution struct {
    ID         int    `json:"id,omitempty"`
    Title      string `json:"title"`
    Repo string `json:"repository"`
    Date       string `json:"date"`
    URL        string `json:"url,omitempty"`
}

func NewContribution(title, repo, date, url string) (*Contribution, error) {
    c := &Contribution{
        Title: title,
        Repo:  repo,
        Date:  date,
        URL:   url,
    }

    if err := c.Validate(); err != nil {
        return nil, err
    }

    return c, nil
}

func (c *Contribution) Validate() error {
    if strings.TrimSpace(c.Title) == "" {
        return errors.New("title is required")
    }
    if strings.TrimSpace(c.Repo) == "" {
        return errors.New("repo is required")
    }
    if !strings.Contains(c.Repo, "/") {
        return errors.New("repo must be in owner/repo format")
    }
    if c.Date == "" {
        c.Date = time.Now().Format("2006-01-02")
    }
    if _, err := time.Parse("2006-01-02", c.Date); err != nil {
        return errors.New("date must be in YYYY-MM-DD format")
    }
    return nil
}