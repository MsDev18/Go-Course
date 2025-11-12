package migrator

import (
	"E-04/repository/mysql"
	"database/sql"
	"fmt"

	"github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name 
// TODO - add limit to Up and Down method 
func New(dbConfig mysql.Config) Migrator {
	// OR: Read migrations from folder;
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}
	return Migrator{
		dialect:    "mysql",
		dbConfig:   dbConfig,
		migrations: migrations,
	}
}

func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("can't apply migration: %v", err))

	}
	fmt.Printf("Migrate %d migrations !", n)
}
func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Sprintf("can't rollback migration: %v", err))
	}
	fmt.Printf("Rollback %d migrations !", n)
}

func (m Migrator) Status() {
	// TODO - add Status
}
