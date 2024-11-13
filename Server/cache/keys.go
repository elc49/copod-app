package cache

import "fmt"

const (
	PAYMENT_UPDATED_CHANNEL = "payment_updated"
)

func IpCacheKey(ip string) string {
	return fmt.Sprintf("ip:%s", ip)
}
