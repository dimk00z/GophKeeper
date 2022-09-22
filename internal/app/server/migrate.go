package server

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/dimk00z/GophKeeper/pkg/logger"
	"github.com/golang-migrate/migrate/v4"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
	_sslMode         = "?sslmode=disable"
)

func doMigrations(databaseURL string, l *logger.Logger) {
	if !strings.Contains(databaseURL, _sslMode) {
		databaseURL += _sslMode
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		l.Debug("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		l.Fatal("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		l.Debug("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
