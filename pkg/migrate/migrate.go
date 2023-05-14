package migrate

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/pkg/path"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate is a struct that contains the migrate.Migrate instance.
type Migrate struct {
	mig *migrate.Migrate
}

// New creates a new Migrate instance.
func New(db *sql.DB) *Migrate {
	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: string(env.DB.Name)})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/%s", path.Root(), "migration"),
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	return &Migrate{m}
}

// Up migrates the database to the most recent version available.
func (m *Migrate) Up() {
	fmt.Println("Migrating database UP...")

	if err := m.mig.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No changes")
			return
		}

		panic(err)
	}
}

// Down migrates the database to the previous version.
func (m *Migrate) Down() {
	fmt.Println("Migrating database DOWN...")

	if err := m.mig.Down(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No changes")
			return
		}

		panic(err)
	}
}

// Drop drops all tables.
func (m *Migrate) Drop() {
	fmt.Println("Dropping all tables...")

	if err := m.mig.Drop(); err != nil {
		panic(err)
	}
}

// Force migrates the database to a specific version.
func (m *Migrate) Force(version int) {
	fmt.Printf("Migrating database to version %d...\n", version)

	if err := m.mig.Force(version); err != nil {
		panic(err)
	}
}

// Version prints the current version.
func (m *Migrate) Version() {
	version, _, err := m.mig.Version()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current database version: %d\n", version)
}

// Steps migrates the database by a number of versions.
func (m *Migrate) Steps(steps int) {
	fmt.Printf("Migrating database by %d steps...\n", steps)

	if err := m.mig.Steps(steps); err != nil {
		panic(err)
	}
}

// Reset will down then up the database.
func (m *Migrate) Reset() {
	fmt.Println("Resetting database...")

	if err := m.mig.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			panic(err)
		}
	}

	if err := m.mig.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			panic(err)
		}
	}
}
