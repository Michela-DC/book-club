package db

import (
	"context"
	"database/sql"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/uuid"

	"github.com/Michela-DC/book-club/internal/domain"

	_ "github.com/mattn/go-sqlite3" // init sql driver
)

// ErrorNotFound is the sentinel error when no rows are found.
var ErrorNotFound = errors.New("not found")

// SQLiteBookRepository provides access to book data stored in a SQLite database.
// It implements [domain.BookRepository].
type SQLiteBookRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

// NewSQLiteBookRepository creates a new SQLiteBookRepository using the provided database
// file path and logger. It opens the SQLite connection but does not apply migrations.
func NewSQLiteBookRepository(dbPath string, logger *slog.Logger) (*SQLiteBookRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.With("error", err).Error("unable to open db connection")
		return nil, err
	}

	return &SQLiteBookRepository{
		db:     db,
		logger: logger,
	}, nil
}

// ApplyMigrations executes all .sql migration files in the given directory
// that have not already been applied. It records applied migrations in a
// dedicated migrations table to ensure idempotency.
func (repo *SQLiteBookRepository) ApplyMigrations(ctx context.Context, migrationsPath string) error {
	_, err := repo.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS migrations(
		name TEXT PRIMARY KEY,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		repo.logger.With("error", err).Error("failed to create migration table")
		return err
	}

	appliedMigrations := make(map[string]struct{}, 0)
	rows, err := repo.db.QueryContext(ctx, `SELECT name FROM migrations`)
	if err != nil {
		repo.logger.With("error", err).Error("failed to read applied migrations")
		return err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			repo.logger.With("error", err).Error("failed to close rows")
		}
	}()
	for rows.Next() {
		var n string
		err = rows.Scan(&n)
		if err != nil {
			repo.logger.With("error", err).Error("failed to scan migration")
			return err
		}

		appliedMigrations[n] = struct{}{}
	}
	if err = rows.Err(); err != nil {
		repo.logger.With("error", err).Error("failed to read migrations")
		return err
	}

	migrationFiles := make([]string, 0)
	err = filepath.WalkDir(migrationsPath, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".sql") {
			migrationFiles = append(migrationFiles, d.Name())
		}
		return nil
	})
	if err != nil {
		repo.logger.With("error", err).Error("failed to walk migration folder")
		return err
	}

	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		if _, applied := appliedMigrations[file]; applied {
			continue
		}

		//nolint:gosec // we control the file
		content, err := os.ReadFile(filepath.Join(migrationsPath, file))
		if err != nil {
			repo.logger.With("error", err, "filename", file).Error("failed to read migration file")
			return err
		}

		repo.logger.With("filename", file).Info("applying migration")

		tx, err := repo.db.BeginTx(ctx, nil)
		if err != nil {
			repo.logger.With("error", err, "filename", file).Error("failed to start transaction")
			return err
		}
		_, err = tx.ExecContext(ctx, string(content))
		if err != nil {
			repo.logger.With("error", err, "filename", file).Error("failed to apply migration")
			return err
		}

		_, err = tx.ExecContext(ctx, `INSERT INTO migrations(name) VALUES (?)`, file)
		if err != nil {
			repo.logger.With("error", err, "filename", file).Error("failed to store migration")
			return err
		}

		err = tx.Commit()
		if err != nil {
			repo.logger.With("error", err, "filename", file).Error("failed to commit transaction")
			return err
		}

		repo.logger.With("filename", file).Info("migration completed")
	}

	repo.logger.Info("all migrations were applied successfully!")

	return nil
}

// Create inserts a new book record into the database. If the book has no ID,
// a new UUID is generated automatically.
func (repo *SQLiteBookRepository) Create(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book.ID == "" {
		book.ID = uuid.NewString()
	}
	_, err := repo.db.ExecContext(ctx,
		`INSERT INTO books (id, title, author, genre, published_year, status)
		 VALUES (?, ?, ?, ?, ?, ?);`, book.ID, book.Title, book.Author, book.Genre, book.PublishedYear, book.Status,
	)
	if err != nil {
		repo.logger.With("error", err).Error("failed to insert new record")
		return nil, err
	}

	return book, nil
}

// List retrieves books matching the provided filters. This method is not yet implemented.
func (*SQLiteBookRepository) List(context.Context, *domain.BookFilters) ([]*domain.Book, error) {
	return nil, errors.New("not yet implemented")
}

// Update modifies an existing book record in the database.
func (repo *SQLiteBookRepository) Update(ctx context.Context, book *domain.Book) error {
	if book.ID == "" {
		return errors.New("book id cannot be empty")
	}
	_, err := repo.db.ExecContext(ctx,
		`UPDATE books
		SET title = ?, 
		author = ?, 
		genre = ?, 
		published_year = ?, 
		status = ?
		where id = ?;`,
		book.Title, book.Author, book.Genre, book.PublishedYear, book.Status, book.ID,
	)
	if err != nil {
		repo.logger.With("error", err, "id", book.ID).Error("failed to update book")
		return err
	}

	return nil
}

// Delete removes a book record identified by its ID from the database.
func (repo *SQLiteBookRepository) Delete(ctx context.Context, bookID string) error {
	res, err := repo.db.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookID)

	if err != nil {
		repo.logger.With("error", err, "id", bookID).Error("failed to delete book")
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		repo.logger.With("error", err).Error("unable to get rows affected by the delete")
		return err
	}
	if count == 0 {
		repo.logger.With("count", count).Error("no rows affected")
		return ErrorNotFound
	}

	return nil
}
