package main

import (
	"GoSocial/internal/db"
	"GoSocial/internal/env"
	"GoSocial/internal/store"
	"log"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/gosocial?sslmode=disable"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONS", 30),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(config.db.addr, config.db.maxOpenCons, config.db.maxIdleCons, config.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: config,
		store:  store,
	}

	mux := app.mount()

	err = app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
