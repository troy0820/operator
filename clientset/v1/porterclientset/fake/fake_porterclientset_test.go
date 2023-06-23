package fake

import (
	"context"
	"testing"

	v1 "get.porter.sh/operator/api/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// This test demonstrates on how to use the fake client set to get/delete
// things with the porterclientset with objects.  These objects can be
// preloaded to allow the test to utilize the functions within the client set.
func TestCreateInstallWithFakeClientWithObjects(t *testing.T) {
	ctx := context.TODO()
	install := &v1.Installation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-install",
			Namespace: "fake-namespace",
		},
	}
	pclient := NewFakePorterClientSetWithObjects(install)
	gotInstall := &v1.Installation{}
	err := pclient.Get(ctx, types.NamespacedName{Name: "fake-install", Namespace: "fake-namespace"}, gotInstall)
	assert.NoError(t, err)
	assert.Equal(t, "fake-install", gotInstall.Name)
	assert.Equal(t, "fake-namespace", gotInstall.Namespace)
}

func TestCreateInstallWithFakeClient(t *testing.T) {
	ctx := context.TODO()
	install := &v1.Installation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-install",
			Namespace: "fake-namespace",
		},
	}
	pclient := NewFakePorterClient()
	err := pclient.Create(ctx, install)
	assert.NoError(t, err)
	gotInstall := &v1.Installation{}
	err = pclient.Get(ctx, types.NamespacedName{Name: "fake-install", Namespace: "fake-namespace"}, gotInstall)
	assert.NoError(t, err)
	assert.Equal(t, "fake-install", gotInstall.Name)
}
