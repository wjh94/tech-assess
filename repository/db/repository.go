package db

//go:generate mockgen -destination=repository_mock.go -package=db -source=repository.go

type dbResource interface {
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Repository struct {
	db dbResource
}

func NewRepository(db dbResource) *Repository {
	return &Repository{db: db}
}
