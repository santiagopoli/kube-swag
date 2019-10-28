package domain

type Rule struct {
	Path string
}

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

type RuleService interface {
	GetRulesFrom(URL string, pathSelector map[string]string) []*Rule
}

type OutputGenerator interface {
	Generate(rules []*Rule, ingressMap *IngressConfig)
}
