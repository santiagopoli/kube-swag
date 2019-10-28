package infrastructure

import (
	"github.com/santiagopoli/kubeswag/internal/core/domain"
	"k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

type IngressOutputGenerator struct {
	Client *kubernetes.Clientset
}

func (generator *IngressOutputGenerator) Generate(rules []*domain.Rule, ingressMap *domain.IngressConfig) {
	ingress := rulesToIngress(rules, ingressMap)
	createOrUpdateIngress(generator.Client, ingress.Namespace, ingress)
}

func createOrUpdateIngress(client *kubernetes.Clientset, namespace string, ingress *v1beta1.Ingress) {
	ingressAPI := client.ExtensionsV1beta1().Ingresses(namespace)
	_, err := ingressAPI.Create(ingress)
	if err != nil {
		_, _ = ingressAPI.Update(ingress)
	}
}

func getBackendsFromRules(rules []*domain.Rule, serviceName string, port intstr.IntOrString) []v1beta1.HTTPIngressPath {
	var pathsK8s []v1beta1.HTTPIngressPath
	for _, rule := range rules {
		pathsK8s = append(pathsK8s, v1beta1.HTTPIngressPath{
			Path: rule.Path,
			Backend: v1beta1.IngressBackend{
				ServiceName: serviceName,
				ServicePort: port,
			},
		})
	}

	return pathsK8s
}

func appendDefaultAnnotations(annotations map[string]string) map[string]string {
	annotations["beta.kubeswag.io/managed"] = "true"
	return annotations
}

func rulesToIngress(rules []*domain.Rule, ingressConfig *domain.IngressConfig) *v1beta1.Ingress {
	pathsK8s := getBackendsFromRules(rules, ingressConfig.Backend.ServiceName, intstr.FromString(ingressConfig.Backend.ServicePort))

	return &v1beta1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:        ingressConfig.Name,
			Namespace:   ingressConfig.Namespace,
			Annotations: appendDefaultAnnotations(ingressConfig.AdditionalAnnotations),
			Labels:      ingressConfig.AdditionalLabels,
			OwnerReferences: []v1.OwnerReference{
				v1.OwnerReference{
					APIVersion: "kubeswag.io/v1beta1",
					Kind:       "IngressMap",
					Name:       ingressConfig.Name,
					UID:        types.UID(ingressConfig.UID),
				},
			},
		},
		Spec: v1beta1.IngressSpec{
			Rules: []v1beta1.IngressRule{
				v1beta1.IngressRule{
					Host: ingressConfig.Host,
					IngressRuleValue: v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths: pathsK8s,
						},
					},
				},
			},
		},
	}
}
