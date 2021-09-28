package pubsub

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/zakiafada32/retail/business/user"
)

func Publish(data interface{}) {
	urls := nats.DefaultURL
	options := []nats.Option{nats.Name("users service")}
	nc, err := nats.Connect(urls, options...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var subject string
	var message []byte
	switch dataType := data.(type) {
	case user.User:
		subject = "new-user-registered"
		message, err = json.Marshal(user.User{
			Email:   dataType.Email,
			Name:    dataType.Name,
			Address: dataType.Address,
		})
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Println("dont know about type", dataType)
	}

	nc.Publish(subject, []byte(message))
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal()
	}

	log.Printf("Published [%s] : '%s'\n", subject, message)
	log.SetFlags(log.LstdFlags)

}
