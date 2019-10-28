package v1beta1

type BackendSpec struct {
	ServiceName string
	ServicePort string
}

type IngressConfig struct {
	UID                   string
	Name                  string
	Namespace             string
	Host                  string
	Backend               BackendSpec
	AdditionalLabels      map[string]string
	AdditionalAnnotations map[string]string
	PathSelector          map[string]string
	SpecURL               string
}
