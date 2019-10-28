package action

import "github.com/santiagopoli/kubeswag/internal/core/domain"

type GenerateOutputsAction func(ingressMap *domain.IngressConfig)

func NewGenerateOutputsAction(ruleService domain.RuleService, outputGenerators []domain.OutputGenerator) GenerateOutputsAction {
	return func(ingressMap *domain.IngressConfig) {
		rules := ruleService.GetRulesFrom(ingressMap.SpecURL, ingressMap.PathSelector)
		for _, outputGenerator := range outputGenerators {
			outputGenerator.Generate(rules, ingressMap)
		}
	}
}
