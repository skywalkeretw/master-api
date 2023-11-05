package kubernets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetKubernetesDeployments(t *testing.T) {
	var tests = []struct {
		mockClientset kubernetes.Interface
	}{
		{mockClientset: testclient.NewSimpleClientset(
			&appsv1.DeploymentList{
				Items: []appsv1.Deployment{
					{ObjectMeta: metav1.ObjectMeta{Name: "deploymnet-1"}},
					{ObjectMeta: metav1.ObjectMeta{Name: "deploymnet-2"}},
					{ObjectMeta: metav1.ObjectMeta{Name: "deploymnet-3"}},
				},
			})},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			deployments, err := GetKubernetesDeployments("")
			assert.Nil(t, err)
			assert.Equal(t, len(deployments), 3)
		})
	}
}

func TestGetKubernetesDeployment(t *testing.T) {
	var tests = []struct {
		name          string
		mockClientset kubernetes.Interface
	}{
		{
			name: "deployment-1",
			mockClientset: testclient.NewSimpleClientset(
				&appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "deployment-1"}},
			),
		},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			deployment, err := GetKubernetesDeployment(tt.name, "")
			assert.Nil(t, err)
			assert.Equal(t, deployment.Name, tt.name)
		})

	}
}

func TestDeleteKubernetesDeployment(t *testing.T) {
	var tests = []struct {
		name          string
		mockClientset kubernetes.Interface
	}{
		{
			name: "deployment-1",
			mockClientset: testclient.NewSimpleClientset(
				&appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "deployment-1"}},
			),
		},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			err := DeleteKubernetesDeployment(tt.name, "")
			assert.Nil(t, err)

		})
	}
}

func TestUpdateKubernetesDeployment(t *testing.T) {
	var tests = []struct {
		name          string
		updateOptions metav1.UpdateOptions
		mockClientset kubernetes.Interface
	}{
		{
			name:          "deployment-1",
			updateOptions: metav1.UpdateOptions{TypeMeta: metav1.TypeMeta{Kind: appsv1.DefaultDeploymentUniqueLabelKey}},
			mockClientset: testclient.NewSimpleClientset(
				&appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "deployment-1"}},
			),
		},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			err := UpdateKubernetesDeployment(tt.name, "", tt.updateOptions)
			assert.Nil(t, err)
		})
	}
}

func TestGetKubernetesService(t *testing.T) {
	var tests = []struct {
		name          string
		namespace     string
		mockClientset kubernetes.Interface
	}{
		{
			name:      "service-1",
			namespace: "default",
			mockClientset: testclient.NewSimpleClientset(
				&corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "service-1", Namespace: "default"}},
			),
		},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			service, err := GetKubernetesService(tt.name, tt.namespace)
			assert.Nil(t, err)
			assert.Equal(t, service.Name, tt.name)
		})
	}
}

func TestDeleteKubernetesService(t *testing.T) {
	var tests = []struct {
		name      string
		namespace string
	}{
		{
			name:      "service-1",
			namespace: "default",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			err := DeleteKubernetesService(tt.name, tt.namespace)
			assert.Nil(t, err)
		})
	}
}

func TestUpdateKubernetesService(t *testing.T) {
	var tests = []struct {
		name          string
		namespace     string
		updateOptions metav1.UpdateOptions
		mockClientset kubernetes.Interface
	}{
		{
			name:          "service-1",
			namespace:     "default",
			updateOptions: metav1.UpdateOptions{},
			mockClientset: testclient.NewSimpleClientset(
				&corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "service-1", Namespace: "default"}},
			),
		},
	}

	for i, tt := range tests {
		clientset = tt.mockClientset
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			fmt.Println(i)
			err := UpdateKubernetesService(tt.name, tt.namespace, tt.updateOptions)
			assert.Nil(t, err)
		})
	}
}
