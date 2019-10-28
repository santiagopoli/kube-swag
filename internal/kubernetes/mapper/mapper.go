package mapper

import (
	"github.com/santiagopoli/kubeswag/internal/core/domain"
	"github.com/santiagopoli/kubeswag/internal/kubernetes/apis/v1beta1"
)

func ToIngressMap(ingressMapK8S *v1beta1.IngressConfig) *domain.IngressConfig {
	return &domain.IngressConfig{
		UID:       ingressMapK8S.UID,
		Name:      ingressMapK8S.Name,
		Namespace: ingressMapK8S.Namespace,
		Host:      ingressMapK8S.Host,
		Backend: domain.BackendSpec{
			ServiceName: ingressMapK8S.Backend.ServiceName,
			ServicePort: ingressMapK8S.Backend.ServicePort,
		},
		AdditionalLabels:      ingressMapK8S.AdditionalLabels,
		AdditionalAnnotations: ingressMapK8S.AdditionalAnnotations,
		PathSelector:          ingressMapK8S.PathSelector,
		SpecURL:               ingressMapK8S.SpecURL,
	}
}
