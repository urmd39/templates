package infrastructure

import (
	"log"
	"os"
	"time"

	guuid "github.com/google/uuid"
)

var (
	NATSConnection EventPublisher
	VNLocation     *time.Location
)

func init() {
	InfoLog = log.New(os.Stdout, "\u001B[1;34m[INFO]\u001B[0m ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stderr, "\033[1;31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Llongfile)
	LogSysterm = LogController{
		InfoLog: InfoLog,
		ErrLog:  ErrLog,
	}
	//VN time
	location, err := time.LoadLocation(VietNamTimeZone)
	if err != nil {
		ErrLog.Printf("Error load Vietname time zone %v", err)
		os.Exit(1)
	}
	VNLocation = location

	loadEnvParameters()

	authenClientID := "authen-service-" + guuid.New().String()
	NATSConnection = NewNatsPublisher(NATSHostport, NATSClusterID, authenClientID)
	DOSpaceConn = setupDOSpaceConnection()
}
