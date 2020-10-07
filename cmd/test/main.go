package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test-buyer-experience/internal/app/test/getter/prodgetter"
	"test-buyer-experience/internal/app/test/sender/prodsender"
	"test-buyer-experience/internal/app/test/server"
	"test-buyer-experience/internal/app/test/store/sqlstore"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config.yml", "path to config file")
}

func main() {
	flag.Parse()

	if err := Start(); err != nil {
		log.Fatalf("Error while running application: %s", err)
	}
}

func Start() error {
	cfg := viper.New()
	cfg.AutomaticEnv()
	cfg.SetConfigFile(configPath)
	if err := cfg.ReadInConfig(); err != nil {
		return err
	}

	logger, _ := zap.NewDevelopment()
	//noinspection GoUnhandledErrorResult
	defer logger.Sync()

	db, err := newDB(cfg.GetString("database.url"))
	if err != nil {
		return err
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	srv := server.NewServer(sqlstore.New(db), logger, prodgetter.NewGetter(cfg), prodsender.NewSender(logger, cfg))
	srv.BackgroundTask()
	return srv.Router.Run(":" + cfg.GetString("server.port"))
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := runMigrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
