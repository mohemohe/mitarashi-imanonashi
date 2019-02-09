package main

import (
	"github.com/mohemohe/mitarashi-imanonashi/mstdn"
	"github.com/sirupsen/logrus"
)

func main() {
	ch, err := mstdn.OnStream()
	if err != nil {
		logrus.Panic(err)
	}

	for {
		event := <- ch
		logrus.WithField("event", event).Info("onStream")
	}
}
