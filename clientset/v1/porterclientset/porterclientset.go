package porterclientset

import (
	v1 "get.porter.sh/operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewPorterClientSet(config *rest.Config) (client.WithWatch, error) {
	scheme := runtime.NewScheme()
	v1.AddToScheme(scheme)
	porterClient, err := client.NewWithWatch(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}
	return porterClient, nil
}
