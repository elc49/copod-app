package ip

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	copodCache "github.com/elc49/copod/cache"
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/util"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
	"github.com/sirupsen/logrus"
)

var ip *ipClient

type IP interface {
	GetIpinfo(context.Context, string) (*Ipinfo, error)
}

type ipClient struct {
	log   *logrus.Logger
	c     *ipinfo.Client
	cache copodCache.Cache
	http  *http.Client
}

func New() {
	log := logger.GetLogger()

	cache := cache.NewInMemory().WithExpiration(time.Hour * 24)
	client := ipinfo.NewClient(nil, ipinfo.NewCache(cache), config.C.Ipinfo.ApiKey)

	ip = &ipClient{log, client, copodCache.GetCache(), &http.Client{}}
}

func GetIpService() IP {
	return ip
}

func (c *ipClient) GetIpinfo(ctx context.Context, ip string) (*Ipinfo, error) {
	ipinfo := &Ipinfo{}

	cValue, err := c.cache.Get(ctx, copodCache.IpCacheKey(ip), ipinfo)
	if err != nil {
		c.log.WithError(err).Errorf("ip: reach ipinfo from cache")
		return nil, err
	} else if cValue != nil {
		return (cValue).(*Ipinfo), nil
	}

	ipapi := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	req, err := http.NewRequest("GET", ipapi, nil)
	if err != nil {
		c.log.WithError(err).Errorf("ip: http.NewRequest")
		return nil, err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")

	res, err := c.http.Do(req)
	if err != nil {
		c.log.WithError(err).Errorf("ip: http.Do")
		return nil, err
	}

	if err := util.DecodeHttp(res.Body, &ipinfo); err != nil {
		c.log.WithError(err).Errorf("ip: DecodeHttp")
		return nil, err
	}

	secondaryIpinfo, err := c.c.GetIPInfo(net.ParseIP(ip))
	if err != nil {
		c.log.WithError(err).WithFields(logrus.Fields{"ip": ip}).Errorf("ip: GetIpInfo")
		return nil, err
	}
	ipinfo.CountryFlagURL = secondaryIpinfo.CountryFlagURL
	ipinfo.Gps = secondaryIpinfo.Location

	go func() {
		if err := c.cache.Set(ctx, copodCache.IpCacheKey(ip), ipinfo, time.Hour*24); err != nil {
			c.log.WithError(err).WithFields(logrus.Fields{"ipinfo": ipinfo}).Errorf("ip: cache ipinfo")
			return
		}
	}()

	return ipinfo, nil
}
