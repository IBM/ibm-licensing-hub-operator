module github.com/ibm/ibm-licensing-hub-operator

go 1.14

require (
	github.com/go-logr/logr v0.1.0
    github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
    github.com/operator-framework/operator-sdk v0.17.0
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)
