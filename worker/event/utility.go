package event

import (
	"encoding/json"
	"templates/infrastructure"

	"github.com/nats-io/stan.go"
)

func convertData(input interface{}, ouput interface{}) error {
	dataByte, err := json.Marshal(input)
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	err = json.Unmarshal(dataByte, &ouput)
	if err != nil {
		infrastructure.ErrLog.Print(err)
	}
	return err
}

type Action interface {
	create(objectCreate interface{}) (err error)
	update(objectUpdate interface{}, code string) (err error)
	updateStatus(objectUpdate interface{}, code string) (err error)
	GetSubject() string
	GetGroupName() string
	GetDurableName() string
	Callback(m *stan.Msg)
}

type ConfigSyn struct {
	subject     string
	groupName   string
	durableName string
}

func (s ConfigSyn) GetSubject() string {
	return s.subject
}

func (s ConfigSyn) GetGroupName() string {
	return s.groupName
}

func (s ConfigSyn) GetDurableName() string {
	return s.durableName
}
