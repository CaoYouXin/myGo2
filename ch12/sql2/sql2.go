package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/OHIS_dev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	s := &service{
		db:       db,
		healthz:  regexp.MustCompile("^/healthz/?"),
		findname: regexp.MustCompile("^/findname/?"),
	}

	http.ListenAndServe("localhost:8080", s)
}

type service struct {
	db       *sql.DB
	healthz  *regexp.Regexp
	findname *regexp.Regexp
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case s.healthz.MatchString(r.URL.Path):
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		err := db.PingContext(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	case s.findname.MatchString(r.URL.Path):
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("request form can not be parsed: %v", err), http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to parse form value: %v", err), http.StatusInternalServerError)
			return
		}

		var name string
		err = db.QueryRowContext(ctx,
			fmt.Sprintf("select p.real_name from user as p where p.id = %d;", id),
		).Scan(&name)
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to execute search query: %v", err), http.StatusFailedDependency)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(struct {
			Name string `json:"name"`
		}{name})
		return
	}
}
