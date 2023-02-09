package pkg

import (
	"log"

	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // import file driver for migrate
	"github.com/jmoiron/sqlx"
)

var (
	// ErrDriverInit is returned when we cannot initialize the driver.
	ErrDriverInit = errors.New("failed to initialize postgres driver")
	// ErrMigrateInit is returned when we cannot initialize the migrate driver.
	ErrMigrateInit = errors.New("failed to initialize migration driver")
	// ErrMigration is returned when we cannot run a migration.
	ErrMigration = errors.New("failed to migrate database")
)

const migrationsPath = "file:schema"

// MigratePostgres migrates the database to the latest version.
func MigratePostgres(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to run")
		} else {
			return err
		}
	}

	log.Println("migrations successfully run")

	return nil
}
