package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBStore struct {
	Db      *mongo.Database
	Session *mongo.Client
}

func NewDatastore(caFile string, certificateFile string, privateKeyFile string, dbAuthMechanism string, replication string) *MongoDBStore {
	var mongoDBStore *MongoDBStore
	db, session := connectMongoDatabase(caFile, certificateFile, privateKeyFile, dbAuthMechanism, replication)
	if db != nil && session != nil {
		mongoDBStore = new(MongoDBStore)
		mongoDBStore.Db = db
		mongoDBStore.Session = session
		return mongoDBStore
	}
	log.Fatal("Datastore not create")
	return nil
}

func connectMongoDatabase(caFile string, certificateFile string, privateKeyFile string, dbAuthMechanism string, replication string) (*mongo.Database, *mongo.Client) {
	var connectOne sync.Once
	var db *mongo.Database
	var session *mongo.Client
	var err error
	connectOne.Do(func() {
		opt := configureConnectionInformation(caFile, certificateFile, privateKeyFile, dbAuthMechanism, replication)
		session, err = mongo.Connect(context.TODO(), opt)
		if err != nil {
			ErrLog.Fatal(err)
		}
		err = session.Ping(context.TODO(), nil)
		db = session.Database(DBMongoName)
	})

	return db, session
}

func configureConnectionInformation(caFile string, certificateFile string, privateKeyFile string, dbAuthMechanism string, relication string) *options.ClientOptions {
	if runtime.GOOS == "windows" {
		opt := options.Client().
			ApplyURI("mongodb://" + DBMongoHostPort).
			SetConnectTimeout(10 * time.Second)
		return opt
	}
	cer, err := tls.LoadX509KeyPair(certificateFile, privateKeyFile)
	if err != nil {
		ErrLog.Fatal(err)
	}
	certs, err := os.ReadFile(caFile)
	if err != nil {
		ErrLog.Fatal(err)
	}
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Fatal("No certs appended, using system certs only")

	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		RootCAs:      rootCAs,
	}
	cert := options.Credential{
		AuthMechanism: dbAuthMechanism,
		Username:      DBMongoName,
	}
	rp, err := readpref.New(readpref.PrimaryMode)
	if err != nil {
		log.Print(err)
	}
	opt := options.Client().
		ApplyURI("mongodb://" + DBMongoHostPort).
		SetTLSConfig(config).
		SetAuth(cert).SetMaxPoolSize(10).
		SetConnectTimeout(10 * time.Second).
		SetReplicaSet(relication).
		SetReadPreference(rp)
	return opt
}
