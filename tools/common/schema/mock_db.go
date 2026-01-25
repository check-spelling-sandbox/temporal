package schema

import (
	"errors"
)

type (
	mockSQLDB struct {
	}
)

// Exec executes a cql statement
func (db *mockSQLDB) Exec(stmt string, args ...interface{}) error {
	return errors.New("unimplemented")
}

// DropAllTables drops all tables
func (db *mockSQLDB) DropAllTables() error {
	return errors.New("unimplemented")
}

// CreateSchemaVersionTables sets up the schema version tables
func (db *mockSQLDB) CreateSchemaVersionTables() error {
	return errors.New("unimplemented")
}

// ReadSchemaVersion returns the current schema version for the keyspace
func (db *mockSQLDB) ReadSchemaVersion() (string, error) {
	return "", errors.New("unimplemented")
}

// UpdateSchemaVersion updates the schema version for the keyspace
func (db *mockSQLDB) UpdateSchemaVersion(newVersion string, minCompatibleVersion string) error {
	return errors.New("unimplemented")
}

// WriteSchemaUpdateLog adds an entry to the schema update history table
func (db *mockSQLDB) WriteSchemaUpdateLog(oldVersion string, newVersion string, manifestMD5 string, desc string) error {
	return errors.New("unimplemented")
}

// Close gracefully closes the client object
func (db *mockSQLDB) Close() {}

// Type gives the type of db (e.g. "cassandra", "sql")
func (db *mockSQLDB) Type() string {
	return "sql"
}
