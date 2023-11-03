package kubernets

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

var clientset kubernetes.Interface

func init() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = home + "/.kube/config"
	} else {
		kubeconfig = "/.kube/config"
	}

	// use the in-cluster configuration if it exists, otherwise use the kubeconfig file
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	// create a Kubernetes clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
}

// GetKubernetesPods retrieves a list of all pods in the Kubernetes cluster
func GetKubernetesPods() ([]corev1.Pod, error) {

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}

// GetKubernetesDeployments retrieves a list of all deployments in the Kubernetes cluster
func GetKubernetesDeployments(namespace string) ([]appsv1.Deployment, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments.Items, nil
}

// GetKubernetesDeployment retrieves information about a specific deployment in the Kubernetes cluster.
func GetKubernetesDeployment(name, namespace string) (*appsv1.Deployment, error) {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

// DeleteKubernetesDeployment deletes a specific deployment in the Kubernetes cluster.
func DeleteKubernetesDeployment(name, namespace string) error {
	// Delete the deployment
	deletePolicy := metav1.DeletePropagationForeground
	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdateKubernetesDeployment updates a specific deployment in the Kubernetes cluster.
func UpdateKubernetesDeployment(name, namespace string, updateOptions metav1.UpdateOptions) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Fetch the latest deployment object
		deployment, err := GetKubernetesDeployment(name, namespace)
		if err != nil {
			return err
		}
		// Modify the deployment as needed
		//deployment.Spec.Replicas = int32Ptr(3) // For example, change the number of replicas

		// Update the deployment
		_, updateErr := clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("update failed: %v", retryErr)
	}
	return nil
}
