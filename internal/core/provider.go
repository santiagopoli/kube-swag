package core

import (
	"github.com/santiagopoli/kubeswag/internal/core/action"
	"github.com/santiagopoli/kubeswag/internal/core/domain"
	"github.com/santiagopoli/kubeswag/internal/core/infrastructure"
	"github.com/santiagopoli/kubeswag/internal/kubernetes"
	"github.com/santiagopoli/kubeswag/internal/kubernetes/apis/v1beta1"
)

func ruleService() domain.RuleService {
	return &infrastructure.RuleService{}
}

func generators() []domain.OutputGenerator {
	return []domain.OutputGenerator{kubernetes.IngressOutputGenerator()}
}

func IngressMapClient() *v1beta1.IngressMapClient {
	return kubernetes.IngressMapClient()
}

func GenerateOutputsAction() action.GenerateOutputsAction {
	return action.NewGenerateOutputsAction(ruleService(), generators())
}
