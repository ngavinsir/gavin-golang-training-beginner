package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/ngavinsir/golangtraining/internal/jobs"
	postgresRepository "github.com/ngavinsir/golangtraining/internal/postgres"
	"github.com/ngavinsir/golangtraining/internal/users"
	"github.com/ngavinsir/golangtraining/paymentcodes"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Application",
	}
)

var (
	paymentCodesRepository *postgresRepository.PaymentCodesRepository
	paymentCodesService    *paymentcodes.PaymentCodesService
	expirePaymentCodesJob  *jobs.ExpirePaymentCodesJob
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "golangbeginner"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	dbConn := initDB()
	httpClient := initHttpClient()

	paymentCodesRepository = postgresRepository.NewPaymentCodesRepository(dbConn)
	users := users.NewUsersClient(httpClient)
	paymentCodesService = paymentcodes.NewService(paymentCodesRepository, users)

	expirePaymentCodesJob = jobs.NewExpirePaymentCodesJob(paymentCodesService)
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
