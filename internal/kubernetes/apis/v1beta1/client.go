package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type IngressMapClient struct {
	Client dynamic.Interface
}

func (repository *IngressMapClient) List(namespace string, options v1.ListOptions) []*IngressConfig {

	ingressMapsAPI := repository.Client.Resource(schema.GroupVersionResource{
		Group:    "kubeswag.io",
		Version:  "v1beta1",
		Resource: "ingressmaps",
	})

	ingressMaps, _ := ingressMapsAPI.List(options)
	result := []*IngressConfig{}

	for _, ingressMap := range ingressMaps.Items {
		metadata := ingressMap.UnstructuredContent()["metadata"].(map[string]interface{})
		spec := ingressMap.UnstructuredContent()["spec"].(map[string]interface{})
		backendSpec := spec["backend"].(map[string]interface{})

		ingressConfig := &IngressConfig{
			UID:       metadata["uid"].(string),
			Name:      metadata["name"].(string),
			Namespace: metadata["namespace"].(string),
			Host:      spec["host"].(string),
			Backend: BackendSpec{
				ServiceName: backendSpec["serviceName"].(string),
				ServicePort: backendSpec["servicePort"].(string),
			},
			AdditionalLabels:      toStringMap(spec["additionalLabels"]),
			AdditionalAnnotations: toStringMap(spec["additionalAnnotations"]),
			SpecURL:               spec["specURL"].(string),
			PathSelector:          toStringMap(spec["pathSelector"]),
		}

		result = append(result, ingressConfig)
	}

	return result
}

func toStringMap(mapToConvert interface{}) map[string]string {
	result := make(map[string]string)

	if mapToConvert == nil {
		return result
	}

	for k, v := range mapToConvert.(map[string]interface{}) {
		result[k] = v.(string)
	}

	return result
}
