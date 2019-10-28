package operator

import (
	"time"

	"github.com/santiagopoli/kubeswag/internal/core"
	"github.com/santiagopoli/kubeswag/internal/kubernetes/mapper"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Start() {
	ingressMapsV1Beta1 := core.IngressMapClient()
	generateOutputs := core.GenerateOutputsAction()

	for range time.Tick(time.Second * 5) {
		ingressMapsK8S := ingressMapsV1Beta1.List("default", v1.ListOptions{})
		for _, ingressMapK8S := range ingressMapsK8S {
			ingressMap := mapper.ToIngressMap(ingressMapK8S)
			generateOutputs(ingressMap)
		}
	}
}
