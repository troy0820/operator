package controllers

import (
	"context"

	installationv1 "get.porter.sh/porter/gen/proto/go/porterapis/installation/v1alpha1"
	"google.golang.org/grpc"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

const (
	PorterGRPCName        = "porter-grpc-service"
	PorterDeploymentImage = "ghcr.io/getporter/server:v1.1.0"
)

type PorterClient interface {
	ListInstallations(ctx context.Context, in *installationv1.ListInstallationsRequest, opts ...grpc.CallOption) (*installationv1.ListInstallationsResponse, error)
	ListInstallationLatestOutputs(ctx context.Context, in *installationv1.ListInstallationLatestOutputRequest, opts ...grpc.CallOption) (*installationv1.ListInstallationLatestOutputResponse, error)
}

type ClientConn interface {
	Close() error
}

var GrpcDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name:      PorterGRPCName,
		Namespace: operatorNamespace,
		Labels: map[string]string{
			"app": PorterGRPCName,
		},
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: ptr.To(int32(1)),
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": PorterGRPCName,
			},
		},

		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": PorterGRPCName,
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  PorterGRPCName,
						Image: PorterDeploymentImage,

						Ports: []corev1.ContainerPort{
							{

								Name:          "grpc",
								ContainerPort: 3001,
							},
						},

						Args: []string{"api-server", "run"},
						VolumeMounts: []corev1.VolumeMount{
							{
								MountPath: "/porter-config",
								Name:      "porter-grpc-service-config-volume",
							},
						},
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("2000m"),
								corev1.ResourceMemory: resource.MustParse("512Mi"),
							},
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("100m"),
								corev1.ResourceMemory: resource.MustParse("32Mi"),
							},
						},
					},
				},
				Volumes: []corev1.Volume{
					{
						Name: "porter-grpc-service-config-volume",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{

								LocalObjectReference: corev1.LocalObjectReference{
									Name: "porter-grpc-service-config",
								},
								Items: []corev1.KeyToPath{
									{
										Key:  "config",
										Path: "config.yaml",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var GrpcService = &corev1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      PorterGRPCName,
		Namespace: operatorNamespace,
		Labels: map[string]string{
			"app": "porter-grpc-service",
		},
	},
	Spec: corev1.ServiceSpec{

		Ports: []corev1.ServicePort{
			{
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(3001),
				Port:       int32(3001),
			},
		},
		Selector: map[string]string{"app": "porter-grpc-service"},
		Type:     corev1.ServiceTypeClusterIP,
	},
}

var GrpcConfigMap = &corev1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "porter-grpc-service-config",
		Namespace: operatorNamespace,
	},
	Data: map[string]string{
		// FILLED out during resolve PorterConfig
		"config": ConfigmMapConfig,
	},
}

var ConfigmMapConfig string
