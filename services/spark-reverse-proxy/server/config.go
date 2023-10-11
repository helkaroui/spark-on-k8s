package server

type Config struct {
	Port                      int
	SparkApplicationNamespace string
	SparkUIServiceUrl         string
	ModifyRedirectUrl         bool
	ProxyBaseUri              string
	TemplatesPath             string
}
