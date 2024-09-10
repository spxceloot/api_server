package database

import (
	"context"
	"database/sql"
	"fmt"
	"github/luqxus/spxce/types"
	"os"

	_ "github.com/lib/pq"
)

// Database configurations
type DatabaseConfig struct {
	Host     string // host address | localhost is default
	Port     int    // database engine port number
	User     string // database user
	Password string // database user password
	DBName   string // database name
}

// database interface
type Database interface {
	// counts number of rows matching provided email
	CountEmail(ctx context.Context, email string) (int64, error)

	// create a new user from provided data
	CreateUser(ctx context.Context, data *types.User) error

	// get user matching email
	GetUser(ctx context.Context, email string) (*types.User, error)
}

// PGDatabase struct is a postgres implementation
// of the Database interface
type PGDatabase struct {
	db *sql.DB // database instance
}

// create a new PGDatabase database implementation
// this function connects to posgres database engine
// returns new PGDatabase or error if failure
func NewPGDatabase(config DatabaseConfig) (*PGDatabase, error) {

	// database connection string
	// from values in the database configuration
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	// connect to postgres database using connStr [connection string]
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// return error if failure
		return nil, err
	}

	// create tables
	err = createTables(db)
	if err != nil {
		// return error on create tables failure
		return nil, err
	}

	// return PGDatabase on successful connection
	return &PGDatabase{
		db: db,
	}, nil
}

// create a new user entity in the database
// return nil or error on failure
func (db *PGDatabase) CreateUser(ctx context.Context, data *types.User) error {
	// query string
	sql := "INSERT INTO Users (uid, username, email, password) VALUES ($1,$2,$3,$4)"

	// execute query with context
	_, err := db.db.ExecContext(ctx, sql, data.UID, data.Username, data.Email, data.Password)

	// return error which equals nil on success
	return err
}

// count number of rows that match given email
// returns int64(count result) and nil or -1 and error on failure
func (db *PGDatabase) CountEmail(ctx context.Context, email string) (int64, error) {
	// query string
	query := "SELECT COUNT(email) FROM Users WHERE email=$1"

	// execute query with context | result expected to have on row
	result := db.db.QueryRowContext(ctx, query, email)

	// result count
	var count int64

	// scan result into count
	if err := result.Scan(&count); err != nil {
		// return -1 and error on failure
		return -1, err
	}

	// return count and nil on success
	return count, nil
}

// gets user that matches email
// returns user pointer and nil or
// nil and error on failure
func (db *PGDatabase) GetUser(ctx context.Context, email string) (*types.User, error) {
	// query string
	query := "SELECT uid, username, email, password, created_at FROM Users WHERE email=$1"

	// execute query with context | pass emai filter
	result := db.db.QueryRowContext(ctx, query, email)

	// resulted row result
	var user types.User

	// scan result row into user
	if err := result.Scan(
		&user.UID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		// return nil and error on failure
		return nil, err
	}

	// return user pinter and nil on success
	return &user, nil
}

func readSchema() ([]byte, error) {

	// open schema files
	file, err := os.Open("./schema.sql")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// get schema file stats
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// new []byte with file size
	b := make([]byte, stats.Size())

	// read file bytes into b
	file.Read(b)

	// return schema []byte and nil on success
	return b, nil
}

func createTables(db *sql.DB) error {

	// get schema []bytens
	b, err := readSchema()
	if err != nil {
		return err
	}

	// parse scheme []byte to string
	sql := string(b)

	// execute schema string
	_, err = db.Exec(sql)

	// return error | error == nil on success
	return err

}
