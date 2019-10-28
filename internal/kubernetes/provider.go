package kubernetes

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/santiagopoli/kubeswag/internal/kubernetes/apis/v1beta1"
	"github.com/santiagopoli/kubeswag/internal/kubernetes/infrastructure"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig = getKubeConfig()

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func getKubeConfig() *string {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	return kubeconfig
}

func getKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

func getKubernetesDynamicClient() (dynamic.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	return dynamic.NewForConfig(config)
}

func IngressMapClient() *v1beta1.IngressMapClient {
	kubernetesDynamicClient, _ := getKubernetesDynamicClient()
	return &v1beta1.IngressMapClient{Client: kubernetesDynamicClient}
}

func IngressOutputGenerator() *infrastructure.IngressOutputGenerator {
	kubernetsClient, _ := getKubernetesClient()
	return &infrastructure.IngressOutputGenerator{Client: kubernetsClient}
}
