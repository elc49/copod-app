package main

import (
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/server"
)

func main() {
	s := server.New()
	s.TigrisService()
	s.Database(config.C.Database.Rdbms)
	s.PaystackService()
	s.MountController()
	s.Start()
}
