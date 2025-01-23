package main

import (
	"embed"

	"github.com/elc49/copod/server"
)

var static embed.FS

func main() {
	s := server.New()
	s.NewTigrisService()
	s.Database()
	s.NewCacheService()
	s.NewPaystackService()
	s.NewIpinfoService()
	s.NewEthereumService()
	s.NewResendEmailService()
	s.NewSentryService()
	s.MountController()
	s.Start(static)
}
