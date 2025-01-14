package main

import (
	"embed"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/server"
)

var static embed.FS

func main() {
	s := server.New()
	s.TigrisService()
	s.Database(config.C.Database.Rdbms)
	s.CacheService()
	s.PaystackService()
	s.IpinfoService()
	s.MountController()
	s.NewEthereumService()
	s.Start(static)
}
