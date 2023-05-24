package main

import (
	server "AustemServer"
	"AustemServer/migrations"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

}

func main() {
	migrations.Migrator()
	server.StartAPI()
}
