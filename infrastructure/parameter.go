package infrastructure

import (
	"log"
	"os"
	"time"
)

const (
	// Develop environment
	DEVELOP_ENV = "DEVELOP_ENV"

	DEVELOP_ENV_DEV = "dev" // dev mode
	DEVELOP_ENV_PRD = "prd" // production mode

	// API
	BASE_PATH     = "BASE_PATH"
	API_HOST_PORT = "HOST_PORT"
	HTTP_SWAGGER  = "HTTP_SWAGGER"

	// Database Postgres
	DB_POSTGRES_HOST     = "DB_POSTGRES_HOST"
	DB_POSTGRES_PORT     = "DB_POSTGRES_PORT"
	DB_POSTGRES_NAME     = "DB_POSTGRES_NAME"
	DB_POSTGRES_USERNAME = "DB_POSTGRES_USERNAME"
	DB_POSTGRES_PASSWORD = "DB_POSTGRES_PASSWORD"
	DB_POSTGRES_SSL_MODE = "DB_POSTGRES_SSL_MODE"
	INITDB               = "INIT_DB"

	// Database Mongo
	DB_MONGO_HOST_PORT        = "DB_MONGO_HOST_PORT"
	DB_MONGO_NAME             = "DB_MONGO_NAME"
	DB_MONGO_USERNAME         = "DB_MONGO_USERNAME"
	DB_MONGO_CERTIFICATE_FILE = "DB_MONGO_CERTIFICATE_FILE"
	DB_MONGO_PRIVATE_KEY_FILE = "DB_MONGO_PRIVATE_KEY_FILE"
	DB_MONGO_CA_FILE          = "DB_MONGO_CA_FILE"
	DB_MONGO_AUTH_MECHANISM   = "DB_MONGO_AUTH_MECHANISM"
	DB_MONGO_REPLICATION      = "DB_MONGO_REPLICATION"

	// JWT
	PATH_PUBLIC_KEY = "JWT_PUBLIC_KEY_PATH"

	// Collection

	// NATS
	NATS_HOST_PORT    = "NATS_HOST_PORT"
	NATS_CLUSTER_ID   = "NATS_CLUSTER_ID"
	NATS_DURABLE_NAME = "NATS_DURABLE_NAME"
	NATS_CA           = "NATS_CA"
	NATS_CLIENT_PEM   = "NATS_CLIENT_PEM"
	NATS_CLIENT_KEY   = "NATS_CLIENT_KEY"
	NATS_CLIENT_ID    = "NATS_CLIENT_ID"

	// NATS Subject
	NATS_AUTH_SUBJECT                       = "NATS_AUTH_SUBJECT"
	NATS_HTTP_REQUEST_HANDLE_LOGGER_SUBJECT = "NATS_HTTP_REQUEST_HANDLE_LOGGER_SUBJECT"

	// DO SPACES
	SPACESKEY           = "SPACES_KEY"
	SPACESSECRET        = "SPACES_SECRET"
	SPACES_EXCEL_FOLDER = "SPACES_EXCEL_FOLDER"
	EXCELSTORAGEPATH    = "EXCEL_STORAGE_PATH"
	EXCELDOMAIN         = "EXCEL_DOMAIN"

	VietNamTimeZone = "Asia/Vientiane"
)

var (
	// develop environment
	DevelopEnv string

	// Logger
	InfoLog, ErrLog *log.Logger

	// API
	APIHostport string
	HttpSwagger string
	BasePath    string

	// JWT
	PathPublicKey string

	// Postgres
	DBPostgresHost     string
	DBPostgresPort     string
	DBPostgresName     string
	DBPostgresUsername string
	DBPostgresPassword string
	DBPostgresSSLMode  string
	DBInitDB           string

	// Mongo
	DBMongoHostPort        string
	DBMongoName            string
	DBMongoUsername        string
	DBMongoCertificateFile string
	DBMongoPrivateKeyFile  string
	DBMongoCAFile          string
	DBMongoAuthMechanism   string
	DBMongoReplication     string

	// Collection

	// NATS
	NATSHostport  string
	NATSClusterID string
	NATSClientID  string
	CaFileNATS    string
	CertFileNATS  string
	KeyFileNATS   string

	// NATS Subject
	NATSAuthSubject                    string
	NATSHttpRequestHandleLoggerSubject string

	// DO SPACES
	SpacesKey        string
	SpacesSecret     string
	SpaceExcelFolder string
	ExcelStoragePath string
	ExcelDomain      string
	DOSpaceConn      DOSpaceConnection

	// Setting
	VNLocation *time.Location
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	var value string
	if val, ok := os.LookupEnv(envParam); ok {
		value = val
	} else {
		value = defaultValue
	}
	InfoLog.Printf("%s: %s", envParam, value)
	return value

}

func loadEnvParameters() {
	// Develop environment
	DevelopEnv = getStringEnvParameter(DEVELOP_ENV, DEVELOP_ENV_DEV)

	// API
	BasePath = getStringEnvParameter(BASE_PATH, "/api/v1/backend")
	APIHostport = getStringEnvParameter(API_HOST_PORT, "")
	HttpSwagger = getStringEnvParameter(HTTP_SWAGGER, "")

	// JWT
	PathPublicKey = getStringEnvParameter(PATH_PUBLIC_KEY, "")

	// Postgres
	DBPostgresHost = getStringEnvParameter(DB_POSTGRES_HOST, "")
	DBPostgresPort = getStringEnvParameter(DB_POSTGRES_PORT, "")
	DBPostgresName = getStringEnvParameter(DB_POSTGRES_NAME, "")
	DBPostgresUsername = getStringEnvParameter(DB_POSTGRES_USERNAME, "")
	DBPostgresPassword = getStringEnvParameter(DB_POSTGRES_PASSWORD, "")
	DBPostgresSSLMode = getStringEnvParameter(DB_POSTGRES_SSL_MODE, "disable")
	DBInitDB = getStringEnvParameter(INITDB, "true")

	// Mongo
	DBMongoHostPort = getStringEnvParameter(DB_MONGO_HOST_PORT, "")
	DBMongoName = getStringEnvParameter(DB_MONGO_NAME, "")
	DBMongoUsername = getStringEnvParameter(DB_MONGO_USERNAME, "")
	DBMongoCertificateFile = getStringEnvParameter(DB_MONGO_CERTIFICATE_FILE, "")
	DBMongoPrivateKeyFile = getStringEnvParameter(DB_MONGO_PRIVATE_KEY_FILE, "")
	DBMongoCAFile = getStringEnvParameter(DB_MONGO_CA_FILE, "")
	DBMongoAuthMechanism = getStringEnvParameter(DB_MONGO_AUTH_MECHANISM, "")
	DBMongoReplication = getStringEnvParameter(DB_MONGO_REPLICATION, "")

	// NATS
	NATSHostport = getStringEnvParameter(NATS_HOST_PORT, "")
	NATSClusterID = getStringEnvParameter(NATS_CLUSTER_ID, "")
	NATSClientID = getStringEnvParameter(NATS_CLIENT_ID, "")
	CaFileNATS = getStringEnvParameter(NATS_CA, "")
	CertFileNATS = getStringEnvParameter(NATS_CLIENT_PEM, "")
	KeyFileNATS = getStringEnvParameter(NATS_CLIENT_KEY, "")

	// NATS Subject
	NATSAuthSubject = getStringEnvParameter(NATS_AUTH_SUBJECT, "AuthSubject")
	NATSHttpRequestHandleLoggerSubject = getStringEnvParameter(NATS_HTTP_REQUEST_HANDLE_LOGGER_SUBJECT, "HttpRequestHandleLoggerSubject")

	// Collection

	// Spaces
	SpacesKey = getStringEnvParameter(SPACESKEY, "")
	SpacesSecret = getStringEnvParameter(SPACESSECRET, "")
	SpaceExcelFolder = getStringEnvParameter(SPACES_EXCEL_FOLDER, "")
	ExcelStoragePath = getStringEnvParameter(EXCELSTORAGEPATH, "")
	ExcelDomain = getStringEnvParameter(EXCELDOMAIN, "")
}
