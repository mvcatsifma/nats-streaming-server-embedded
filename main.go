package main

import (
	"fmt"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"log"
)

func main() {
	clusterId := "mystreamingserver"
	s, err := stand.RunServer(clusterId)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := stan.Connect(clusterId, "test-client")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Subscribe("test-channel", func(msg *stan.Msg) {
		log.Printf("%+v", msg)
	}, stan.StartAt(pb.StartPosition_NewOnly))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", conn.NatsConn().Status())

	s.Shutdown()
}
