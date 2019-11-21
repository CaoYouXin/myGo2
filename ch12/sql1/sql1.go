package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	id := flag.Int64("id", 82, "person ID to find")
	dsn := flag.String("dsn", "root@/OHIS_dev", "connection data source name")

	flag.Parse()
	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}
	var err error

	db, err = sql.Open("mysql", *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	ping(ctx)

	query(ctx, *id)
}

func ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("unable to connect to database", err)
	}
}

func query(ctx context.Context, id int64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var name sql.NullString
	err := db.QueryRowContext(ctx,
		"select p.real_name from user as p where p.id = ?;",
		id,
	).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	if name.Valid {
		log.Println("name=", name.String)
	} else {
		log.Println("该用户没有说明")
	}
}
