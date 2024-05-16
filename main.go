package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	rHttps := gin.Default()
	r.GET("/", index1)
	imageInit.xx(r)
	cron.InitCron()
	go fxx.UploadDefaultImage()
	rHttps.Use(TlsHandler())

	server := &http.Server{Addr: ":8443", Handler: rHttps, TLSConfig: tlsConf}
	go server.ListenAndServeTLS(xx, xx)

	//go rHttps.RunTLS("0.0.0.0:8443", CertCrtDir, CertKeyDir)
	routes.Route(rHttps)
	routes.RouteV1(rHttps)
	r.Run("0.0.0.0:8080")
}
