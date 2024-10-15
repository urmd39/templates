package infrastructure

import (
	"log"
	"os"
	"time"

	guuid "github.com/google/uuid"
)

const (
	// POSTGRE Database
	POSTGREHOST   = "DB_POSTGRE_HOST"
	POSTGREPORT   = "DB_POSTGRE_PORT"
	USER          = "DB_USER"
	POSTGREDBNAME = "DB_POSTGRE_DBNAME"
	PASSWORD      = "DB_PASSWORD"
	SSLMODE       = "DB_SSL_MODE"
	INITDB        = "INIT_DB"
	//API
	APIHOSTPORT = "HOST_PORT"

	CERTIFICATEFILE = "CERTIFICATE_FILE"
	PRIVATEKEYFILE  = "PRIVATE_KEY_FILE"
	CAFILE          = "CA_FILE"
	HTTPSWAGGER     = "HTTP_SWAGGER"
	PATH_PUBLIC_KEY = "JWT_PUBLIC_KEY_PATH"

	NATSHOSTPORT         = "NATS_HOST_PORT"
	NATSCLUSTERID        = "NATS_CLUSTERID"
	NATSPRODUCTQUEUENAME = "NATS_PRODUCT_QUEUE_NAME"

	NATSCA        = "NATS_CA"
	NATSCLIENTPEM = "NATS_CLIENT_PEM"
	NATSCLIENTKEY = "NATS_CLIENT_KEY"

	NATSCLIENTID = "NATS_CLIENT_ID"

	NATSAUTHSUBJECT           = "NATS_AUTH_SUBJECT"
	NATSPRODUCTSUBJECT        = "NATS_PRODUCT_SUBJECT"
	NATSSTATIONSUBJECT        = "NATS_STATION_SUBJECT"
	NATSSUPPLIERSUBJECT       = "NATS_SUPPLIER_SUBJECT"
	NATSCATEGORYSUBJECT       = "NATS_CATEGORY_SUBJECT"
	NATSORDERSUBJECT          = "NATS_ORDER_SUBJECT"
	NATSORDERGROUP            = "NATS_ORDER_GROUP"
	NATSPRODUCTGROUP          = "NATS_PRODUCT_GROUP"
	NATSAUTHQUEUENAME         = "NATS_AUTH_QUEUE_NAME"
	NATSLOGGINGSUBJECT        = "NATS_LOGGING_SUBJECT"
	NATSWAREHOUSESUBJECT      = "NATS_WAREHOUSE_SUBJECT"
	NATSPRODUCTSTATIONSUBJECT = "NATS_PRODUCT_STATION_SUBJECT"

	NATS_WALLET_SUBJECT  = "NATS_WALLET_SUBJECT"
	NATS_WALLET_GROUP    = "NATS_WALLET_GROUP"
	DURABLE_WALLET_EVENT = "NATS_WALLET_DURABLE"

	NATS_TRADE_SUBJECT  = "NATS_TRADE_SUBJECT"
	NATS_TRADE_GROUP    = "NATS_TRADE_GROUP"
	DURABLE_TRADE_EVENT = "DURABLE_TRADE_EVENT"

	NATS_CUSTOMER_SUBJECT  = "NATS_CUSTOMER_SUBJECT"
	NATS_CUSTOMER_GROUP    = "NATS_CUSTOMER_GROUP"
	DURABLE_CUSTOMER_EVENT = "NATS_CUSTOMER_DURABLE"

	NATS_WALLET_REQUEST_RESPONSE_SUBJECT = "NATS_WALLET_REQUEST_RESPONSE_SUBJECT"
	//
	VietNamTimeZone    = "Asia/Vientiane"
	BARCODEIMAGEDIR    = "BARCODE_IMAGE_DIR"
	BARCODEIMAGEDOMAIN = "BARCODE_IMAGE_DOMAIN"
	//Durable name
	DURABLEPRODUCTEVENT        = "DURABLE_PRODUCT_EVENT"
	DURABLESUPPLIEREVENT       = "DURABLE_SUPPLIER_EVENT"
	DURABLESTATIONEVENT        = "DURABLE_STATION_EVENT"
	DURABLEORDEREVENT          = "DURABLE_ORDER_EVENT"
	DURABLEPRODUCTSTATIONEVENT = "DURABLE_PRODUCT_STATION_EVENT"
	// DO SPACES
	SPACESKEY         = "SPACES_KEY"
	SPACESSECRET      = "SPACES_SECRET"
	SPACESIMAGEFOLDER = "SPACES_IMAGE_FOLDER"
)

var (
	// develop environment
	DevelopEnv string
	//Postgresql database
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SSLMode  string
	InitDB   string

	//API
	APIHostport string

	CertificateFile string
	PrivateKeyFile  string
	CAFile          string
	DBAuthMechanism string
	HostPort        string
	Relication      string
	HttpSwagger     string
	PathPublicKey   string
	InfoLog, ErrLog *log.Logger

	NATSHostport              string
	NATSClusterID             string
	NATSAntradeHostport       string
	NATSAntradeClusterID      string
	NATSProductSubject        string
	NATSAuthSubject           string
	NATSStationSubject        string
	NATSSupplierSubject       string
	NATSCategorySubject       string
	NATSOrderSubject          string
	NATSLoggingSubject        string
	NATSWarehouseSubject      string
	NATSProductStationSubject string

	NATSWalletSubject  string
	NATSWalletGroup    string
	DurableWalletEvent string

	NATSTradeSubject  string
	NATSTradeGroup    string
	DurableTradeEvent string

	NATSCustomerSubject  string
	NATSCustomerGroup    string
	DurableCustomerEvent string

	NATSWalletRequestResponseSubject string

	//Durable name setting
	DurableProductEvent        string
	DurableSupplierEvent       string
	DurableStationEvent        string
	DurableCategoryEvent       string
	DurableOrderEvent          string
	DurableProductStationEvent string
	//group
	NATSOrderGroup   string
	NATSProductGroup string

	NATSAuthQueueName string
	NATSQueue         string

	NATSClientID string
	CaFileNats   string
	CertFileNats string
	KeyFileNats  string

	VNLocation *time.Location
	//
	ImagesDir   string
	ImageDomain string

	SpacesKey        string
	SpacesSecret     string
	SpaceImageFolder string
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
	// PostgreSQL
	Host = getStringEnvParameter(POSTGREHOST, "dev.ubofood.com") // host dev:dev.ubofood.com
	Port = getStringEnvParameter(POSTGREPORT, "43099")           // post databse local:5432  dev:60013
	User = getStringEnvParameter(USER, "admin")                  //   local: postgres dev:admin
	DbName = getStringEnvParameter(POSTGREDBNAME, "ubofood")
	Password = getStringEnvParameter(PASSWORD, "ubofood") // local:antrade dev: cn40.wallet
	SSLMode = getStringEnvParameter(SSLMODE, "disable")
	InitDB = getStringEnvParameter(INITDB, "true")

	//API
	APIHostport = getStringEnvParameter(APIHOSTPORT, "localhost:4000")

	CertificateFile = getStringEnvParameter(CERTIFICATEFILE, "./infrastructure/config/client.mongo.crt")
	PrivateKeyFile = getStringEnvParameter(PRIVATEKEYFILE, "./infrastructure/config/client.mongo.key") ///m/ssl/mongo.antrade.key "F:\\working-backup\\DaihocHaNoi\\ANVITA\\antrade\\key\\jwt.key"
	CAFile = getStringEnvParameter(CAFILE, "./infrastructure/config/mongoCA.crt")
	HttpSwagger = getStringEnvParameter(HTTPSWAGGER, "http://localhost:4000/api/v1/wallet/swagger/doc.json")
	PathPublicKey = getStringEnvParameter(PATH_PUBLIC_KEY, "./infrastructure/config/public.pem")

	NATSHostport = getStringEnvParameter(NATSHOSTPORT, "tls://nats1.anvita.com.vn:43085,tls://nats2.anvita.com.vn:43086,tls://nats2.anvita.com.vn:43087")
	NATSClusterID = getStringEnvParameter(NATSCLUSTERID, "ubofood-cluster") //ubofood-cluster
	NATSQueue = getStringEnvParameter(NATSPRODUCTQUEUENAME, "test_queue")
	NATSAuthQueueName = getStringEnvParameter(NATSAUTHQUEUENAME, "AuthQueue")
	NATSAuthSubject = getStringEnvParameter(NATSAUTHSUBJECT, "AuthSubject")
	NATSProductSubject = getStringEnvParameter(NATSPRODUCTSUBJECT, "ProductSubject")
	NATSProductGroup = getStringEnvParameter(NATSPRODUCTGROUP, "product-group")
	NATSSupplierSubject = getStringEnvParameter(NATSSUPPLIERSUBJECT, "SupplierSubject")
	NATSOrderSubject = getStringEnvParameter(NATSORDERSUBJECT, "OrderSubject")
	NATSOrderGroup = getStringEnvParameter(NATSORDERGROUP, "order-group")
	NATSCategorySubject = getStringEnvParameter(NATSCATEGORYSUBJECT, "CategorySubject")
	NATSStationSubject = getStringEnvParameter(NATSSTATIONSUBJECT, "station")
	NATSClientID = getStringEnvParameter(NATSCLIENTID, "cn40-wallet-"+guuid.New().String())
	NATSLoggingSubject = getStringEnvParameter(NATSLOGGINGSUBJECT, "logging")
	NATSWarehouseSubject = getStringEnvParameter(NATSWAREHOUSESUBJECT, "WarehouseTopic")
	NATSProductStationSubject = getStringEnvParameter(NATSPRODUCTSTATIONSUBJECT, "ProductStationTopicCU-01")

	NATSWalletSubject = getStringEnvParameter(NATS_WALLET_SUBJECT, "WalletSubject")
	NATSWalletGroup = getStringEnvParameter(NATS_WALLET_GROUP, "wallet-group")
	DurableWalletEvent = getStringEnvParameter(DURABLE_WALLET_EVENT, "wallet-0211245709-m1")

	NATSTradeSubject = getStringEnvParameter(NATS_TRADE_SUBJECT, "TradeTopic-05")
	NATSTradeGroup = getStringEnvParameter(NATS_TRADE_GROUP, "trade-group")
	DurableTradeEvent = getStringEnvParameter(DURABLE_TRADE_EVENT, "trade-197")

	NATSCustomerSubject = getStringEnvParameter(NATS_CUSTOMER_SUBJECT, "CustomerSubject")
	NATSCustomerGroup = getStringEnvParameter(NATS_CUSTOMER_GROUP, "CustomerTopicQueue")

	DurableCustomerEvent = getStringEnvParameter(DURABLE_CUSTOMER_EVENT, "customer-durable-16-5")
	NATSWalletRequestResponseSubject = getStringEnvParameter(NATS_WALLET_REQUEST_RESPONSE_SUBJECT, "WalletRequestResponseSubject")

	//image
	ImagesDir = getStringEnvParameter(BARCODEIMAGEDIR, "E:\\log") // /m/static/images  E:\\log
	ImageDomain = getStringEnvParameter(BARCODEIMAGEDOMAIN, "https://static.ubofood.vn/images")

	CaFileNats = getStringEnvParameter(NATSCA, "./infrastructure/config/nats/ca.pem")
	CertFileNats = getStringEnvParameter(NATSCLIENTPEM, "./infrastructure/config/nats/client.pem")
	KeyFileNats = getStringEnvParameter(NATSCLIENTKEY, "./infrastructure/config/nats/client-key.pem")

	//Durable name setting
	DurableProductEvent = getStringEnvParameter(DURABLEPRODUCTEVENT, "product-010")
	DurableSupplierEvent = getStringEnvParameter(DURABLESUPPLIEREVENT, "supplier-durable-9")
	DurableStationEvent = getStringEnvParameter(DURABLESTATIONEVENT, "station-durable-30")
	DurableCategoryEvent = getStringEnvParameter(DURABLESTATIONEVENT, "category-durable-10")
	DurableOrderEvent = getStringEnvParameter(DURABLEORDEREVENT, "order-durable-manh-32")
	DurableProductStationEvent = getStringEnvParameter(DURABLEPRODUCTSTATIONEVENT, "product-station-7")

	// spaces
	SpacesKey = getStringEnvParameter(SPACESKEY, "ZCBJIZWILIN72GSV5AVF")
	SpacesSecret = getStringEnvParameter(SPACESSECRET, "jM/sYI8rhlM4QosKDDE7udOh/XOdjuCrWxI08yfXRoU")
	SpaceImageFolder = getStringEnvParameter(SPACESIMAGEFOLDER, "https://anvita.sgp1.cdn.digitaloceanspaces.com/dev")
}
