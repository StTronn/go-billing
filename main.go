package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"github.com/sttronn/usage-billing/db"

	_ "github.com/go-sql-driver/mysql"
)

func run() error {
	ctx := context.Background()

	db_driver, err := sql.Open("mysql", "root:password@/go_billing?parseTime=true")
	if err != nil {
		return err
	}

	queries := db.New(db_driver)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
