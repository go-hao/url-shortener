package main

import (
	"flag"
	"fmt"

	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/config"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/handler"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/svc"

	"github.com/go-hao/zero/xvalidator"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/urlshortener-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// set custom validator
	httpx.SetValidator(xvalidator.New())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
