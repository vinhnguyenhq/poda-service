package main

import (
	log "github.com/vinhnguyenhq/poda-service/internal/logger"
	"github.com/vinhnguyenhq/poda-service/internal/orm"
	"github.com/vinhnguyenhq/poda-service/pkg/server"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	// Send: ORM instance
	server.Run(orm)
}
