package main

import (
	"flag"
	server "github.com/datapunchorg/spark-ui-reverse-proxy/pkg/server"
	"github.com/golang/glog"
	"log"
)

var (
	namespace         = flag.String("namespace", "", "The Kubernetes namespace where Spark applications are running.")
	port              = flag.Int("port", 8080, "Server port for this reverse proxy.")
	sparkUIServiceUrl = flag.String("spark-ui-service-url", "http://{{$appName}}-ui-svc.{{$appNamespace}}.svc.cluster.local:4040", "Spark UI Service URL, this should point to the Spark driver service which provides Spark UI inside that driver.")
	modifyRedirectUrl = flag.Bool("modify-redirect-url", false, "Whether to modify redirect url in the HTTP response returned from the Spark UI.")
)

func main() {
	flag.Parse()

	log.Printf("Starting server on port %d, application namespace: %s", *port, *namespace)

	if *namespace == "" {
		log.Printf("WARNING: application namespace is empty, the reverse proxy will not work properly. Please add argument like -namespace=your-spark-application-namepace")
	}

	config := server.Config{
		Port: *port,
		SparkApplicationNamespace: *namespace,
		SparkUIServiceUrl: *sparkUIServiceUrl,
		ModifyRedirectUrl: *modifyRedirectUrl,
	}
	server.Run(config)

	glog.Info("Shutting down the server")
}