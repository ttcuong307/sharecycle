package migrations

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"net/url"
	"sharecycle/configs"
	"sharecycle/pkg/logger"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	//go:embed version.txt
	version string
	//go:embed sqls/*.sql
	embedMigrations embed.FS

	src = MigrationSource{
		Migrations: embedMigrations,
		Dir:        "sql", //embedded dir name
	}

	// Parse version from version.txt as int64
	DesiredVersion = func() int64 {
		v, err := strconv.ParseInt(strings.Trim(version, "\n"), 10, 64)
		if err != nil {
			panic(fmt.Errorf("Unable to parse version from version.txt: %w", err))
		}
		return v
	}()
)

// MigrationSource wraps the embedded migrations filesystem and the directory name (it's a convenience thing)
type MigrationSource struct {
	Migrations embed.FS
	Dir        string // Name of the embedded migrations directory
}

// Config holds the information needed to run the migrations
type Config struct {
	User             string
	Password         string
	Host             string
	Port             int
	Name             string
	AdditionalParams map[string]string
	LoggerOverride   goose.Logger
	DryRun           bool
}

// Migrate connects to the database described in the Config,
// and migrations it up or down based on if the current database version is different from the passed
// in desiredVersion
func Migrate(ctx context.Context, cfg Config) error {
	desiredVersion := DesiredVersion

	dsn := func() string {
		q := make(url.Values)
		for k, v := range cfg.AdditionalParams {
			q.Add(k, v)
		}

		//	This is required when writing multiple queries separated by ';' characters in a single sql file.
		q.Set("sslmode", "disable")

		return fmt.Sprintf(configs.MigrateDnsFormat,
			cfg.User,
			cfg.Password,
			cfg.Name,
			"parseTime=true")
	}()

	fmt.Println("DSN: ", dsn)
	db, err := goose.OpenDBWithDriver("mysql", dsn) // This is just runs sql.Open and tells goose which driver to use
	if err != nil {
		return fmt.Errorf("unable to open database for migrating: %w", err)
	}
	defer db.Close()

	// Tell Goose which filesystem to use
	goose.SetBaseFS(src.Migrations)

	// Use our custom logger for goose logging.
	var l goose.Logger
	if cfg.LoggerOverride != nil {
		goose.SetLogger(cfg.LoggerOverride)
		l = cfg.LoggerOverride
	} else {
		l = log.Default()
	}

	// grab current migration version of the database
	dbVersion, err := goose.EnsureDBVersion(db)
	if err != nil {
		return fmt.Errorf("Unable to get current database version: %w", err)
	}

	if cfg.DryRun {
		l.Printf("DRY RUN MODE: Current DB Version: %s.  Desired DB Version: %s", dbVersion, desiredVersion)
		return nil
	}

	if dbVersion > desiredVersion {
		// migrate down to desired version
		return goose.DownTo(db, src.Dir, desiredVersion)
	}

	// migrate up to desired version
	// if dbVersion == desiredVersion, goose will log "no migration needed"
	return goose.UpTo(db, src.Dir, desiredVersion)
}

func WrapLogger(arl logger.Logger) *GooseLogger {
	return &GooseLogger{arl}
}

var _ goose.Logger = &GooseLogger{}

type GooseLogger struct {
	logger.Logger
}

func (l *GooseLogger) Print(v ...interface{}) {
	l.Info(v...)
}

func (l *GooseLogger) Println(v ...interface{}) {
	l.Info(v...)
	l.Infof("\n", nil)
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	l.Infof(format, v...)
}
