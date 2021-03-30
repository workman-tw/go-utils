package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres -  db instance
type Postgres struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
	ssl      string
	db       *gorm.DB
}

const defaultTimeZone = "Asia/Taipei"
const postgresDrive = "postgres"

// New - new postgres instance
func New() *Postgres {
	return &Postgres{}
}

// WithHost - set db's host
func (p *Postgres) WithHost(host string) *Postgres {
	p.host = host
	return p
}

// WithUser - set db's user
func (p *Postgres) WithUser(user string) *Postgres {
	p.user = user
	return p
}

// WithPassword - set db's password
func (p *Postgres) WithPassword(password string) *Postgres {
	p.password = password
	return p
}

// WithDBName - set db's name
func (p *Postgres) WithDBName(dbName string) *Postgres {
	p.dbName = dbName
	return p
}

// WithPort - set db's port
func (p *Postgres) WithPort(port string) *Postgres {
	p.port = port
	return p
}

// WithSSLmode - set db's ssl mode
func (p *Postgres) WithSSLmode(enable bool) *Postgres {
	if enable {
		p.ssl = "enable"
	} else {
		p.ssl = "disable"
	}
	return p
}

// GetDB - get db instance
func (p *Postgres) GetDB() *gorm.DB {
	return p.db
}

// Connect - connect to db
func (p *Postgres) Connect() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			p.host,
			p.user,
			p.password,
			p.dbName,
			p.port,
			p.ssl,
			defaultTimeZone,
		),
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		return err
	}
	p.db = db
	return nil
}
