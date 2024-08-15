package atpg

import (
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	dbInfo := DBInfo{
		DBString: "host=127.0.0.1 port=5500 user=postgres password=pwge24 dbname=app-absensi sslmode=disable", // Adjust with your actual credentials
		DBName:   "app-absensi",
	}
	db, err = PGConnect(dbInfo)
	if err != nil {
		panic(err)
	}

	code := m.Run()

	db.Close()

	os.Exit(code)
}

func TestInsertOneRow(t *testing.T) {
	query := `INSERT INTO users (id_role, id_penginputan, nama, username, password, email, phone_number, created_at, is_active)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id`

	createdAt := time.Now() // Use the current time for the 'created_at' field

	// Assuming id_role and id_penginputan are integers, and others are strings or boolean
	idRole := 1
	idPenginputan := 1
	nama := "Test User"
	username := "testuser"
	password := "password123"
	email := "testuser@example.com"
	phoneNumber := "1234567890"
	isActive := true

	var insertedID int
	err := db.QueryRow(query, idRole, idPenginputan, nama, username, password, email, phoneNumber, createdAt, isActive).Scan(&insertedID)
	if err != nil {
		t.Fatalf("InsertOneRow failed: %v", err)
	}

	if insertedID == 0 {
		t.Fatalf("InsertOneRow failed: no ID returned")
	}
}

func TestGetOneRow(t *testing.T) {
	query := "SELECT nama FROM users WHERE username = $1"

	// Use the same username as in the insert test to ensure it can be found
	row := db.QueryRow(query, "testuser")

	var result string
	if err := row.Scan(&result); err != nil {
		t.Fatalf("Failed to scan row: %v", err)
	}

	if result != "Test User" {
		t.Fatalf("Expected 'Test User', got %v", result)
	}
}
