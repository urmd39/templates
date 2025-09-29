package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"runtime"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	Database *mongo.Database
	Client   *mongo.Client
}

func NewDBStore() *MongoDBStore {
	var mongoDBStore *MongoDBStore
	db, session := connectDatabase()
	if db != nil && session != nil {
		mongoDBStore = new(MongoDBStore)
		mongoDBStore.Database = db
		mongoDBStore.Client = session
		return mongoDBStore
	}
	ErrLog.Fatal("Datastore not create")
	return nil
}

func connectDatabase() (*mongo.Database, *mongo.Client) {
	var connectOne sync.Once
	var db *mongo.Database
	var session *mongo.Client
	var err error
	connectOne.Do(func() {
		opt := configureConnectionInformation()
		session, err = mongo.NewClient(opt)
		if err != nil {
			ErrLog.Fatal(err, "client err")

		}

		err = session.Connect(context.TODO())
		//err = session.Ping(context.TODO(), nil)
		//log.Print(err)
		if err != nil {
			ErrLog.Fatal("connect err")
		}

		err = session.Ping(context.TODO(), nil)
		ErrLog.Print(err)
		db = session.Database(DBMongoName)
	})

	return db, session
}

func configureConnectionInformation() *options.ClientOptions {
	if runtime.GOOS == "windows" {
		opt := options.Client().
			ApplyURI("mongodb://" + DBMongoHostPort).
			SetConnectTimeout(10 * time.Second)
		return opt
	}
	cer, err := tls.LoadX509KeyPair(DBMongoCertificateFile, DBMongoPrivateKeyFile)
	if err != nil {
		ErrLog.Fatal(DBMongoCertificateFile, "34234")
	}
	certs, err := ioutil.ReadFile(DBMongoCertificateFile)
	if err != nil {
		ErrLog.Fatal(err)
	}
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		ErrLog.Fatal(err)
	}
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()

	}
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		ErrLog.Fatal("No certs appended, using system certs only")

	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		RootCAs:      rootCAs,
	}
	cert := options.Credential{
		AuthMechanism: DBMongoAuthMechanism,
		//AuthSource: "$external",
		Username: DBMongoName,
	}
	//rp,err:= readpref.New(readpref.PrimaryMode)
	//if err != nil{
	//	log.Print(err)
	//}
	opt := options.Client().
		ApplyURI("mongodb://" + DBMongoHostPort).
		SetTLSConfig(config).
		SetAuth(cert).SetMaxPoolSize(10).
		SetConnectTimeout(10 * time.Second).
		SetReplicaSet(DBMongoReplication)
	//SetReadPreference(rp)
	return opt

}
