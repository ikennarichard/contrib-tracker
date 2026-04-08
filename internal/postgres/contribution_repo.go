package postgres

import (
    "database/sql"
    "fmt"

    "github.com/ikennarichard/contrib-tracker/internal/domain"
    _ "github.com/lib/pq" // registers the postgres driver
)

type ContributionRepo struct {
    db      *sql.DB
    stmtAdd         *sql.Stmt
    stmtList        *sql.Stmt
    stmtFindByRepo  *sql.Stmt
}


func New(connStr string) (*ContributionRepo, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("postgres: open: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("postgres: ping: %w", err)
    }

    repo := &ContributionRepo{db: db}

    if err := repo.migrate(); err != nil {
        return nil, err
    }

    if err := repo.prepareStatements(); err != nil {
        return nil, err
    }

    return repo, nil
}

func (r *ContributionRepo) migrate() error {
    query := `
        CREATE TABLE IF NOT EXISTS contributions (
            id    SERIAL PRIMARY KEY,
            title TEXT        NOT NULL,
            repo  TEXT        NOT NULL,
            date  DATE        NOT NULL,
            url   TEXT        NOT NULL DEFAULT ''
        );
    `
    if _, err := r.db.Exec(query); err != nil {
        return fmt.Errorf("postgres: migrate: %w", err)
    }
    return nil
}

func (r *ContributionRepo) prepareStatements() error {
    var err error

    r.stmtAdd, err = r.db.Prepare(`
        INSERT INTO contributions (title, repo, date, url)
        VALUES ($1, $2, $3, $4)
    `)
    if err != nil {
        return fmt.Errorf("postgres: prepare add: %w", err)
    }

    r.stmtList, err = r.db.Prepare(`
        SELECT title, repo, date, url
        FROM contributions
        ORDER BY date DESC
    `)
    if err != nil {
        return fmt.Errorf("postgres: prepare list: %w", err)
    }

    r.stmtFindByRepo, err = r.db.Prepare(`
        SELECT title, repo, date, url
        FROM contributions
        WHERE repo = $1
        ORDER BY date DESC
    `)
    if err != nil {
        return fmt.Errorf("postgres: prepare findByRepo: %w", err)
    }

    return nil
}

func (r *ContributionRepo) Add(c *domain.Contribution) error {
    _, err := r.stmtAdd.Exec(c.Title, c.Repo, c.Date, c.URL)
    if err != nil {
        return fmt.Errorf("postgres: add: %w", err)
    }
    return nil
}

func (r *ContributionRepo) List() ([]*domain.Contribution, error) {
    rows, err := r.stmtList.Query()
    if err != nil {
        return nil, fmt.Errorf("postgres: list: %w", err)
    }
    defer rows.Close()

    return scanRows(rows)
}

func (r *ContributionRepo) FindByRepo(repo string) ([]*domain.Contribution, error) {
    rows, err := r.stmtFindByRepo.Query(repo)
    if err != nil {
        return nil, fmt.Errorf("postgres: findByRepo: %w", err)
    }
    defer rows.Close()

    return scanRows(rows)
}

func scanRows(rows *sql.Rows) ([]*domain.Contribution, error) {
    var results []*domain.Contribution

    for rows.Next() {
        c := &domain.Contribution{}
        if err := rows.Scan(&c.Title, &c.Repo, &c.Date, &c.URL); err != nil {
            return nil, fmt.Errorf("postgres: scan: %w", err)
        }
        results = append(results, c)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("postgres: rows error: %w", err)
    }

    return results, nil
}

// Close cleans up prepared statements and the connection pool.
func (r *ContributionRepo) Close() error {
    for _, stmt := range []*sql.Stmt{r.stmtAdd, r.stmtList, r.stmtFindByRepo} {
        if stmt != nil {
            stmt.Close()
        }
    }
    return r.db.Close()
}