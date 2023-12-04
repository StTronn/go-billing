package core

import (
	"context"
	"database/sql"

	"github.com/sttronn/usage-billing/db"

	_ "github.com/go-sql-driver/mysql"
)

type Event db.Event

type AggregatedEvent struct {
	Value float64
}

func createQueries() (*db.Queries, error) {
	db_driver, err := sql.Open("mysql", "root:password@/go_billing?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db.New(db_driver), nil
}

func (event Event) Create(queries *db.Queries) (int64, error) {
	ctx := context.Background()

	result, err := queries.CreateEvent(ctx, db.CreateEventParams{
		CustomerID: event.CustomerID,
		Code:       event.Code,
		Timestamp:  event.Timestamp,
		Value:      event.Value,
	})
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()

}
