package main

import (
	"context"
	"io/ioutil"

	ginats "github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/nats-io/nats.go"
)

func main() {

	config.Load()

	logrus.NewLogger()

	var err error
	var conn *nats.Conn

	conn, err = ginats.NewDefaultConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	subject := "changeme"

	var b []byte
	b, err = ioutil.ReadFile("examples/simple/client/example.json")
	if err != nil {
		log.Fatal(err)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    b,
	}

	err = conn.PublishMsg(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("published group message on subject [%s]", subject)
}
