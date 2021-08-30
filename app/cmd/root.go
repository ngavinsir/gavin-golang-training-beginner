package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/golangtraining/inquiries"
	"github.com/ngavinsir/golangtraining/internal/jobs"
	postgresRepository "github.com/ngavinsir/golangtraining/internal/postgres"
	"github.com/ngavinsir/golangtraining/internal/sqs"
	"github.com/ngavinsir/golangtraining/internal/users"
	"github.com/ngavinsir/golangtraining/paymentcodes"
	"github.com/ngavinsir/golangtraining/payments"
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
	inquiriesRepository    *postgresRepository.InquiriesRepository
	inquiriesService       *inquiries.InquiriesService
	paymentsRepository     *postgresRepository.PaymentsRepository
	paymentsService        *payments.PaymentsService
	expirePaymentCodesJob  *jobs.ExpirePaymentCodesJob
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
	sqsPublisher := initSQSPublisher()

	paymentCodesRepository = postgresRepository.NewPaymentCodesRepository(dbConn)
	users := users.NewUsersClient(httpClient)
	paymentCodesService = paymentcodes.NewService(paymentCodesRepository, users)

	inquiriesRepository = postgresRepository.NewInquiriesRepository(dbConn)
	inquiriesService = inquiries.NewService(inquiriesRepository, paymentCodesService)

	paymentsRepository = postgresRepository.NewPaymentsRepository(dbConn)
	paymentsService = payments.NewService(inquiriesService, paymentsRepository, paymentCodesService, sqsPublisher)

	expirePaymentCodesJob = jobs.NewExpirePaymentCodesJob(paymentCodesService)
}

func initDB() *sql.DB {
	host := mustHaveEnv("POSTGRES_HOST")
	portStr := mustHaveEnv("POSTGRES_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err, "POSTGRES_PORT is not well set ")
	}
	user := mustHaveEnv("POSTGRES_USER")
	password := mustHaveEnv("POSTGRES_PASSWORD")
	database := mustHaveEnv("POSTGRES_DATABASE")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func initHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}

func initSQSPublisher() *sqs.Publisher {
	s := initAWSSession()
	q := mustHaveEnv("SQS_QUEUE_NAME")
	p, err := sqs.NewPublisher(s, q)
	if err != nil {
		panic(err)
	}

	return p
}

func initAWSSession() *session.Session {
	region := mustHaveEnv("SQS_AWS_REGION")
	endpoint := mustHaveEnv("SQS_ENDPOINT")

	sess, err := session.NewSession(&aws.Config{
		Region:   &region,
		Endpoint: &endpoint,
	})
	if err != nil {
		panic(err)
	}

	return sess
}

func mustHaveEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("%s is not well set", key))
	}
	return value
}
