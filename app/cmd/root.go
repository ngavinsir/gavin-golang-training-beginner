package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	postgresRepository "github.com/ngavinsir/golangtraining/internal/postgres"
	"github.com/ngavinsir/golangtraining/paymentcodes"
)

var (
	paymentCodesRepository *postgresRepository.PaymentCodesRepository
	paymentCodesService    *paymentcodes.PaymentCodesService
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "golangbeginner"
)

func init() {
	dbConn := initDB()

	paymentCodesRepository = postgresRepository.NewPaymentCodesRepository(dbConn)
	paymentCodesService = paymentcodes.NewService(paymentCodesRepository)
}

func initDB() (db *sql.DB) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return
}
