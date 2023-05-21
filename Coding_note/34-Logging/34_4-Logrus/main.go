package main

// go get github.com/sirupsen/logrus
// JSON Logging

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")
}
