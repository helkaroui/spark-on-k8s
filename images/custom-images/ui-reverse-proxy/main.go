package main

import (
	"flag"
	"github.com/golang/glog"
	"log"
	"reverse-proxy/server"
)

var (
	namespace         = flag.String("namespace", "default", "The Kubernetes namespace where Spark applications are running.")
	port              = flag.Int("port", 8000, "Server port for this reverse proxy.")
	sparkUIServiceUrl = flag.String("spark-ui-service-url", "http://{{$appName}}-ui-svc.{{$appNamespace}}.svc.cluster.local:4040", "Spark UI Service URL, this should point to the Spark driver service which provides Spark UI inside that driver.")
	modifyRedirectUrl = flag.Bool("modify-redirect-url", false, "Whether to modify redirect url in the HTTP response returned from the Spark UI.")
	proxyBaseUri      = flag.String("proxy-base-uri", "proxy", "the proxy base uri.")
)

func main() {
	flag.Parse()

	log.Printf("Starting server on port %d, application namespace: %s", *port, *namespace)

	config := server.Config{
		Port:                      *port,
		SparkApplicationNamespace: *namespace,
		SparkUIServiceUrl:         *sparkUIServiceUrl,
		ModifyRedirectUrl:         *modifyRedirectUrl,
		ProxyBaseUri:              *proxyBaseUri,
	}
	server.Run(config)

	glog.Info("Shutting down the server")
}
