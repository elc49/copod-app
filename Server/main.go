package main

import (
	"embed"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/server"
)

var static embed.FS

func main() {
	s := server.New()
	s.NewTigrisService()
	s.Database(config.C.Database.Rdbms)
	s.NewCacheService()
	s.NewPaystackService()
	s.NewIpinfoService()
	s.MountController()
	s.NewEthereumService()
	s.NewResendEmailService()
	s.Start(static)
}
