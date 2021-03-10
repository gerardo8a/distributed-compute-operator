package ray

import (
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"

	dcv1alpha1 "github.com/dominodatalab/distributed-compute-operator/api/v1alpha1"
	"github.com/dominodatalab/distributed-compute-operator/pkg/util"
)

var (
	defaultCmd = []string{ApplicationName}
	defaultEnv = []corev1.EnvVar{
		{
			Name: "MY_POD_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name: "MY_CPU_REQUEST",
			ValueFrom: &corev1.EnvVarSource{
				ResourceFieldRef: &corev1.ResourceFieldSelector{
					Resource: "requests.cpu",
				},
			},
		},
	}
	defaultVolumes = []corev1.Volume{
		{
			Name: sharedMemoryVolumeName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{
					Medium: corev1.StorageMediumMemory,
				},
			},
		},
	}
	defaultVolumeMounts = []corev1.VolumeMount{
		{
			Name:      sharedMemoryVolumeName,
			MountPath: "/dev/shm",
		},
	}
)

const sharedMemoryVolumeName = "dshm"

// NewDeployment generates a Deployment configured to manage Ray cluster nodes.
// The configuration is based the provided spec and the desired Component workload.
func NewDeployment(rc *dcv1alpha1.RayCluster, comp Component) (*appsv1.Deployment, error) {
	var replicas int32
	var nodeAttrs dcv1alpha1.RayClusterNode

	switch comp {
	case ComponentHead:
		replicas = 1
		nodeAttrs = rc.Spec.Head.RayClusterNode
	case ComponentWorker:
		replicas = *rc.Spec.Worker.Replicas
		nodeAttrs = rc.Spec.Worker.RayClusterNode
	default:
		return nil, fmt.Errorf("invalid ray component: %q", comp)
	}

	imageRef, err := util.ParseImageDefinition(rc.Spec.Image)
	if err != nil {
		return nil, err
	}

	args := processArgs(rc, comp)
	ports := processPorts(rc, comp)
	labels := processLabels(rc, comp, nodeAttrs.Labels)
	envVars := append(defaultEnv, rc.Spec.EnvVars...)
	volumes := append(defaultVolumes, nodeAttrs.Volumes...)
	volumeMounts := append(defaultVolumeMounts, nodeAttrs.VolumeMounts...)

	serviceAccountName := InstanceObjectName(rc.Name, ComponentNone)
	if rc.Spec.ServiceAccountName != "" {
		serviceAccountName = rc.Spec.ServiceAccountName
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      InstanceObjectName(rc.Name, comp),
			Namespace: rc.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: SelectorLabelsWithComponent(rc, comp),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: nodeAttrs.Annotations,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccountName,
					NodeSelector:       nodeAttrs.NodeSelector,
					Affinity:           nodeAttrs.Affinity,
					Tolerations:        nodeAttrs.Tolerations,
					InitContainers:     nodeAttrs.InitContainers,
					ImagePullSecrets:   rc.Spec.ImagePullSecrets,
					SecurityContext:    rc.Spec.PodSecurityContext,
					Containers: []corev1.Container{
						{
							Name:            ApplicationName,
							Command:         defaultCmd,
							Args:            args,
							Image:           imageRef,
							ImagePullPolicy: rc.Spec.Image.PullPolicy,
							Ports:           ports,
							Env:             envVars,
							VolumeMounts:    volumeMounts,
							Resources:       nodeAttrs.Resources,
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									TCPSocket: &corev1.TCPSocketAction{
										Port: intstr.FromInt(int(rc.Spec.NodeManagerPort)),
									},
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									TCPSocket: &corev1.TCPSocketAction{
										Port: intstr.FromInt(int(rc.Spec.NodeManagerPort)),
									},
								},
							},
						},
					},
					Volumes: volumes,
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
		},
	}

	return deploy, nil
}

func processArgs(rc *dcv1alpha1.RayCluster, comp Component) []string {
	cmdArgs := []string{
		"start",
		"--block",
		"--node-ip-address=$(MY_POD_IP)",
		"--num-cpus=$(MY_CPU_REQUEST)",
		fmt.Sprintf("--object-manager-port=%d", rc.Spec.ObjectManagerPort),
		fmt.Sprintf("--node-manager-port=%d", rc.Spec.NodeManagerPort),
	}

	if rc.Spec.ObjectStoreMemoryBytes != nil {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--object-store-memory=%d", *rc.Spec.ObjectStoreMemoryBytes))
	}

	// TODO: the exhaustive linter unveiled a nasty code smell in this set of functions.
	// 	i think we want to refactor this so that we only check the component "once" in the
	// 	main routine and leverage a deployment processor type that adheres to a common interface
	// 	and returns the values based on the implementation (i.e. headDeploymentProcess).
	if comp == ComponentHead {
		cmdArgs = addHeadCmdArgs(rc, cmdArgs)
	} else if comp == ComponentWorker {
		cmdArgs = addWorkerCmdArgs(rc, cmdArgs)
	}

	return cmdArgs
}

func processPorts(rc *dcv1alpha1.RayCluster, comp Component) []corev1.ContainerPort {
	ports := []corev1.ContainerPort{
		{
			Name:          "object-manager",
			ContainerPort: rc.Spec.ObjectManagerPort,
		},
		{
			Name:          "node-manager",
			ContainerPort: rc.Spec.NodeManagerPort,
		},
	}

	if comp == ComponentHead {
		ports = addHeadContainerPorts(rc, ports)
	}

	return ports
}

func processLabels(rc *dcv1alpha1.RayCluster, comp Component, extraLabels map[string]string) map[string]string {
	labels := MetadataLabelsWithComponent(rc, comp)
	if extraLabels != nil {
		labels = util.MergeStringMaps(extraLabels, labels)
	}

	return labels
}

func addHeadContainerPorts(rc *dcv1alpha1.RayCluster, ports []corev1.ContainerPort) []corev1.ContainerPort {
	redisPorts := []corev1.ContainerPort{
		{
			Name:          "redis-primary",
			ContainerPort: rc.Spec.Port,
		},
	}
	for idx, port := range rc.Spec.RedisShardPorts {
		redisPorts = append(redisPorts, corev1.ContainerPort{
			Name:          fmt.Sprintf("redis-shard-%d", idx),
			ContainerPort: port,
		})
	}

	return append(ports, redisPorts...)
}

func addHeadCmdArgs(rc *dcv1alpha1.RayCluster, args []string) []string {
	headArgs := []string{
		"--head",
		fmt.Sprintf("--ray-client-server-port=%d", rc.Spec.ClientServerPort),
		fmt.Sprintf("--port=%d", rc.Spec.Port),
		fmt.Sprintf("--redis-shard-ports=%s", strings.Join(util.IntsToStrings(rc.Spec.RedisShardPorts), ",")),
	}

	if util.BoolPtrIsTrue(rc.Spec.EnableDashboard) {
		dashArgs := []string{
			"--include-dashboard=true",
			"--dashboard-host=0.0.0.0",
			fmt.Sprintf("--dashboard-port=%d", rc.Spec.DashboardPort),
		}
		headArgs = append(headArgs, dashArgs...)
	}

	return append(args, headArgs...)
}

func addWorkerCmdArgs(rc *dcv1alpha1.RayCluster, args []string) []string {
	return append(args, fmt.Sprintf("--address=%s:%d", HeadServiceName(rc.Name), rc.Spec.Port))
}