package EasyMongoGo

import (
	"log"
	"testing"
	"time"
)

var dbg *EasyMongo

type demoS struct {
	Col1 string
	Col2 string
	Col3 time.Time
}

var (
	uri = ""
	dbName = "CalculatorAPI"
)

func TestConnect(t *testing.T) {
	db, err := NewEasyMongo(uri)
	if err != nil {
		t.Fatal("Database Error:", err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatal("Database Error:", err)
	} else {
		log.Println("Database Connected")
	}

	db.UseDatabase(dbName)

	dbg = db
}