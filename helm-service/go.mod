module github.com/keptn/keptn/helm-service

go 1.13

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/docker/docker v1.4.2-0.20200203170920-46ec8731fbce
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/keptn/go-utils v0.6.3-0.20201008173953-84d1386b6a35
	github.com/keptn/kubernetes-utils v0.2.1-0.20201028073840-0750ebb3521b
	github.com/kinbiko/jsonassert v1.0.1
	github.com/stretchr/testify v1.5.1
	gotest.tools v2.2.0+incompatible
	helm.sh/helm/v3 v3.1.2
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/cli-runtime v0.17.2
	k8s.io/client-go v0.17.2
	k8s.io/kubectl v0.17.2
	sigs.k8s.io/yaml v1.2.0
)

// Transitive requirement from Helm: See https://github.com/helm/helm/blob/v3.1.2/go.mod
replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
)
