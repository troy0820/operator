package fake

import (
	v1 "get.porter.sh/operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func NewFakePorterClientSetWithObjects(objs ...runtime.Object) client.WithWatch {
	scheme := runtime.NewScheme()
	v1.AddToScheme(scheme)
	clientgoscheme.AddToScheme(scheme)
	fakePorterClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(objs...).Build()
	return fakePorterClient
}

func NewFakePorterClient() client.WithWatch {
	scheme := runtime.NewScheme()
	v1.AddToScheme(scheme)
	clientgoscheme.AddToScheme(scheme)
	fakePorterClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	return fakePorterClient
}
