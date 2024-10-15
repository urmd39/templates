package infrastructure

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type EventPublisher interface {
	Publish(subject string, e interface{}) error
	Reply(queue string, subject string, callback func(msg *nats.Msg))
	Request(queue string, data interface{}) ([]byte, error)
	Consumer(subject string, durableName string, f stan.MsgHandler) error
}

type natsPublisher struct {
	natsHostport string
	clusterID    string
	clientID     string
	sc           *nats.Conn
	stan         stan.Conn
}

func (n natsPublisher) Consumer(subject string, durableName string, f stan.MsgHandler) error {
	InfoLog.Println("Consume NATS subject: ", subject)
	queueName := fmt.Sprintf("InventoryManagement2-QueueName-%s", subject)
	_, err := n.stan.QueueSubscribe(subject, queueName, f, stan.SetManualAckMode(), stan.DeliverAllAvailable(), stan.AckWait(10*time.Second), stan.DurableName(durableName))
	if err != nil {
		ErrLog.Fatal(err)
		return err
	}
	return err
}

func (n natsPublisher) Request(subject string, i interface{}) ([]byte, error) {
	rep, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	res, err := n.sc.Request(subject, rep, 10*time.Second)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (n natsPublisher) Publish(subject string, e interface{}) error {
	eventData, err := json.Marshal(e)
	if err != nil {
		ErrLog.Printf("Can not marshal e data %v\n", e)
		return err
	}
	if err = n.stan.Publish(subject, eventData); err != nil {
		ErrLog.Printf("Publish to nats at subject %s get Error: %v", subject, err)
	}

	return err
}

func (n natsPublisher) Reply(subject string, queue string, callback func(msg *nats.Msg)) {
	_, err := n.sc.QueueSubscribe(subject, queue, callback)
	ErrLog.Print(err)
}

func (n *natsPublisher) setupConnection() {
	connOpts := n.setupConnOptions(CaFileNATS, CertFileNATS, KeyFileNATS)
	natsCon, err := nats.Connect(NATSHostport, connOpts...)
	if err != nil {
		ErrLog.Fatal(err)
	}
	n.sc = natsCon
	sc, err := stan.Connect(n.clusterID, n.clientID, stan.NatsConn(natsCon),
		stan.SetConnectionLostHandler(func(_ stan.Conn, err error) {
			ErrLog.Fatalf("Connection lost, reason: %+v\n", err)
		}),
		stan.Pings(stan.DefaultPingInterval, 100),
	)
	if err != nil {
		ErrLog.Fatal(err)

	}
	n.stan = sc
}

func (n *natsPublisher) setupConnOptions(caFileNats string, certFileNats string, keyFileNats string) []nats.Option {
	var opts = make([]nats.Option, 0)

	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	ca := nats.RootCAs(caFileNats)
	cert := nats.ClientCert(certFileNats, keyFileNats)

	opts = append(opts, ca)
	opts = append(opts, cert)
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		ErrLog.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		ErrLog.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		ErrLog.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func NewNatsPublisher(natsHostport string, natsClusterID string, natsClientID string) EventPublisher {

	publisher := natsPublisher{
		natsHostport: natsHostport,
		clusterID:    natsClusterID,
		clientID:     natsClientID,
	}

	publisher.setupConnection()

	return publisher
}
