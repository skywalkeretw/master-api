package kubernetes

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/skywalkeretw/master-api/app/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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
func GetKubernetesPods(namespace string) ([]corev1.Pod, error) {

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
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

func GetOrCreateNamespace(namespace string) error {
	var err error
	// Check if the namespace exists
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err == nil {
		// Namespace exists, no need to create
		fmt.Printf("Namespace %s already exists\n", namespace)
		return nil
	}

	// If the error is not nil, it means something went wrong during the get operation
	// Create the namespace since it doesn't exist
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Namespace %s created\n", namespace)
	return nil
}

type FunctionModes struct {
	HTTPSync       bool `json:"httpsync"`
	HTTPAsync      bool `json:"httpasync"`
	MessagingSync  bool `json:"messagingsync"`
	MessagingAsync bool `json:"messagingasync"`
}

// CreateKubernetesDeployment creates a new deployment in the Kubernetes cluster.
func CreateKubernetesDeployment(name, namespace, imageName, description, tags string, modes FunctionModes, replicas int) error {

	err := GetOrCreateNamespace(namespace)
	if err != nil {
		return fmt.Errorf("failed to create or retrieve namespace: %v", err)
	}

	// Define the deployment object
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(replicas), // Set the number of replicas as needed
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"run": name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"run": name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:            name,
						Image:           imageName,
						ImagePullPolicy: v1.PullAlways,
						Ports: []v1.ContainerPort{{
							Name:          "http",
							Protocol:      v1.ProtocolTCP,
							ContainerPort: 8080, // Specify the container port to open
						}},
						Env: []v1.EnvVar{
							{
								Name:  "DESCRIPTION",
								Value: description,
							},
							{
								Name:  "TAGS",
								Value: tags,
							},
							{
								Name:  "HTTPSYNC",
								Value: strconv.FormatBool(modes.HTTPSync),
							},
							{
								Name:  "HTTPASYNC",
								Value: strconv.FormatBool(modes.HTTPAsync),
							},
							{
								Name:  "MESSAGINGSYNC",
								Value: strconv.FormatBool(modes.MessagingSync),
							},
							{
								Name:  "MESSAGINGASYNC",
								Value: strconv.FormatBool(modes.MessagingAsync),
							},
							{
								Name: "RABBITMQ_USERNAME",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{Name: "hello-world-default-user"},
										Key:                  "username",
									},
								},
							},
							{
								Name: "RABBITMQ_PASSWORD",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{Name: "hello-world-default-user"},
										Key:                  "password",
									},
								},
							},
							{
								Name: "RABBITMQ_HOST",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{Name: "hello-world-default-user"},
										Key:                  "host",
									},
								},
							},
							{
								Name: "RABBITMQ_PORT",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{Name: "hello-world-default-user"},
										Key:                  "port",
									},
								},
							},
						},
					}},
				},
			},
		},
	}

	// Create the deployment
	_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create deployment: %v", err)
	}
	fmt.Printf("Deployment %s created successfully in namespace %s\n", name, namespace)

	_, err = CreateKubernetesService(name, namespace)
	return err
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

func CreateKubernetesService(name, namespace string) (*corev1.Service, error) {
	// Define your service object
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP, // Specify the service type (e.g., ClusterIP, NodePort, LoadBalancer)
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       8080,                             // Specify the port number
					TargetPort: intstr.IntOrString{IntVal: 8080}, // Specify the target port number
				},
			},
			Selector: map[string]string{
				"run": name, // Specify the labels to match for selecting pods
			},
			// Add any additional specifications as needed
		},
	}

	// Call the Kubernetes API to create the service
	createdService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	// Wait for the service to be ready
	// err = waitForServiceReady(namespace, name)
	// if err != nil {
	// 	return nil, err
	// }

	return createdService, nil
}

// GetKubernetesServices retrieves a list of all services in the Kubernetes cluster within the specified namespace.
func GetKubernetesServices(namespace string) ([]corev1.Service, error) {
	// Retrieve the list of services in the specified namespace
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return services.Items, nil
}

// GetKubernetesService retrieves information about a specific service in the Kubernetes cluster.
func GetKubernetesService(name, namespace string) (*corev1.Service, error) {
	service, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// DeleteKubernetesService deletes a specific service in the Kubernetes cluster.
func DeleteKubernetesService(name, namespace string) error {
	err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

// UpdateKubernetesService updates a specific service in the Kubernetes cluster.
func UpdateKubernetesService(name, namespace string, updateOptions metav1.UpdateOptions) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Fetch the latest service object
		service, err := GetKubernetesService(name, namespace)
		if err != nil {
			return err
		}
		// Modify the service as needed
		// service.Spec.Ports = newPorts // For example, change the ports configuration

		// Update the service
		_, updateErr := clientset.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("update failed: %v", retryErr)
	}
	return nil
}
