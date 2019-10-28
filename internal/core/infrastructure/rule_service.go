package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/santiagopoli/kubeswag/internal/core/domain"
)

type RuleService struct{}

func getAPISpec(URL string) (*openapi3.Swagger, error) {
	response, _ := http.Get(URL)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return openapi3.NewSwaggerLoader().LoadSwaggerFromData(body)
}

func (service *RuleService) GetRulesFrom(URL string, pathSelector map[string]string) []*domain.Rule {
	apiSpec, _ := getAPISpec(URL)
	var rules []*domain.Rule
	for key, path := range apiSpec.Paths {
		shouldAppend := true

		for filterKey, filterValue := range pathSelector {
			extension := path.Extensions[filterKey]
			if extension != nil {
				var extensionValue string
				json.Unmarshal(extension.(json.RawMessage), &extensionValue)
				if filterValue != extensionValue {
					shouldAppend = false
				}
			} else {
				shouldAppend = false
			}
		}

		if shouldAppend {
			rules = append(rules, &domain.Rule{Path: key})
		}
	}
	return rules
}
