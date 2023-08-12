package flags

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/stdlib"
)

type PostgresFlags struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

func (p *PostgresFlags) InitDB(logger *log.Logger) (*sql.DB, error) {
	logger.Debug("POSTGRES! Start init postgreSQL", "user", p.User, "DBName", p.DBName,
		"host", p.Host, "port", p.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		p.User, p.DBName, p.Password,
		p.Host, p.Port)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("POSTGRES! Error in method ping")
		return nil, err
	}

	db.SetMaxOpenConns(10)

	logger.Info("POSTGRES! Successfully init postgreSQL")
	return db, nil
}
