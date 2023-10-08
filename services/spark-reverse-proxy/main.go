package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

	k8sConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	k8sClientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		panic(err.Error())
	}

	pods, err := k8sClientSet.CoreV1().Pods(*namespace).List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	log.Printf("There are %d pods", len(pods.Items))

	log.Printf("N	Name	Labels")
	for i, item := range pods.Items {
		log.Printf("%d	%s	%s", i, item.Name, item.Labels)
	}

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
