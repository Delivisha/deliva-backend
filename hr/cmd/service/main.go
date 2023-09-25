package main

import (
	"database/sql"
	"deliva/hr"
	"deliva/hr/migrations"
	"deliva/internal/config"
	"deliva/internal/system"
	"deliva/internal/web"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("human resources exited abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}
	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			return
		}
	}(s.DB())
	if err = s.MigrateDB(migrations.FS); err != nil {
		return err
	}
	s.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))

	if err = hr.Root(s.Waiter().Context(), s); err != nil {
		return err
	}
	fmt.Println("started human resources service")
	defer fmt.Println("stopped human resources service")

	s.Waiter().Add(
		s.WaitForWeb,
		s.WaitForRPC,
		s.WaitForStream,
	)
	return s.Waiter().Wait()
}
