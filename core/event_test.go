package core_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sttronn/usage-billing/core"
	"github.com/sttronn/usage-billing/db"

	_ "github.com/go-sql-driver/mysql"
)

func TestSumEventAggregate(t *testing.T) {
	// Create a test database connection
	db_driver, err := sql.Open("mysql", "root:password@/go_billing_test?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	defer db_driver.Close()

	// Create the event table if it does not exist
	_, err = db_driver.Exec(`
        CREATE TABLE IF NOT EXISTS event (
            id INT NOT NULL AUTO_INCREMENT,
            customer_id VARCHAR(255) NOT NULL,
            code VARCHAR(255) NOT NULL,
            timestamp DATETIME NOT NULL,
						value FLOAT NOT NULL,
            PRIMARY KEY (id)
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	// Insert test data into the event table if it does not exist
	_, err = db_driver.Exec(`
        INSERT IGNORE INTO event (customer_id, code, timestamp, value)
        VALUES
            ('customer1', 'txn_event', '2019-01-01 00:00:00', '1.0'),
            ('customer1', 'txn_event', '2019-01-01 00:00:00', '20.0');
    `)
	if err != nil {
		t.Fatal(err)
	}

	plan := core.Plan{
		CustomerId: "customer1",
		BillingMetric: &core.BillingMetric{
			Code:        "txn_event",
			Aggregation: core.SUM,
		},
	}

	// Create a test db.Queries object
	queries := db.New(db_driver)

	// Call the Aggregate method and check the result
	result, err := core.Aggregate(plan, queries)
	assert.NoError(t, err)
	assert.Equal(t, 21.0, result.Value)
}

func TestCountEventAggregate(t *testing.T) {
	// Create a test database connection
	db_driver, err := sql.Open("mysql", "root:password@/go_billing_test?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	defer db_driver.Close()

	// Create the event table if it does not exist
	_, err = db_driver.Exec(`
        CREATE TABLE IF NOT EXISTS event (
            id INT NOT NULL AUTO_INCREMENT,
            customer_id VARCHAR(255) NOT NULL,
            code VARCHAR(255) NOT NULL,
            timestamp DATETIME NOT NULL,
						value FLOAT NOT NULL,
            PRIMARY KEY (id)
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	// Insert test data into the event table if it does not exist
	_, err = db_driver.Exec(`
        INSERT IGNORE INTO event (customer_id, code, timestamp, value)
        VALUES
            ('customer1', 'txn_event', '2019-01-01 00:00:00', '1.0'),
            ('customer1', 'txn_event', '2019-01-01 00:00:00', '20.0');
    `)
	if err != nil {
		t.Fatal(err)
	}

	plan := core.Plan{
		CustomerId: "customer1",
		BillingMetric: &core.BillingMetric{
			Code:        "txn_event",
			Aggregation: core.COUNT,
		},
	}

	// Create a test db.Queries object
	queries := db.New(db_driver)

	// Call the Aggregate method and check the result
	result, err := core.Aggregate(plan, queries)
	assert.NoError(t, err)
	assert.Equal(t, 21.0, result.Value)
}

func TestEventCreate(t *testing.T) {
	// Create a test database connection
	db_driver, err := sql.Open("mysql", "root:password@/go_billing_test?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	defer db_driver.Close()

	// Create a test database schema
	_, err = db_driver.Exec(`
        CREATE TABLE IF NOT EXISTS event (
            id INT NOT NULL AUTO_INCREMENT,
            customer_id VARCHAR(255) NOT NULL,
            code VARCHAR(255) NOT NULL,
            timestamp DATETIME NOT NULL,
						value FLOAT NOT NULL,
            PRIMARY KEY (id)
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test event object
	event := core.Event{
		CustomerID: "customer1",
		Code:       "txn_event",
		Timestamp:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		Value:      1.0,
	}

	// Create a test db.Queries object
	queries := db.New(db_driver)

	// Call the Create method and check the result
	lastInsertId, err := event.Create(queries)
	assert.NoError(t, err)
	assert.NotEqual(t, -1, lastInsertId)
}
