package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	postgresRepository "github.com/ngavinsir/golangtraining/internal/postgres"
	"github.com/ngavinsir/golangtraining/internal/users"
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
	httpClient := initHttpClient()

	paymentCodesRepository = postgresRepository.NewPaymentCodesRepository(dbConn)
	users := users.NewUsersClient(httpClient)
	paymentCodesService = paymentcodes.NewService(paymentCodesRepository, users)
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

func initHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}
