package main

import (
	"log"

	"github.com/santiagopoli/kubeswag/internal/kubernetes/operator"
)

func main() {
	log.Println("Started Kubeswag")
	operator.Start()
}
