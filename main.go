package main

import (
	"c3/internal"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

const (
	filenameAnnotation   = "compose.docker.io/filename"
	targetNameAnnotation = "target.docker.io/name"
	targetCMDAnnotation  = "target.docker.io/cmd"
	DEBUG                = true
)

var (
	namespace = "default"
	target    *Target
)

func outClusterConfig() (*rest.Config, error) {
	loadRules := clientcmd.NewDefaultClientConfigLoadingRules()

	cfg, err := loadRules.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	clientConfig, err := clientcmd.NewDefaultClientConfig(
		*cfg,
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create client config: %v", err)
	}
	return clientConfig, nil
}

func main() {

	ns := os.Getenv("NAMESPACE")
	if ns != "" {
		namespace = ns
	}
	cfg, err := rest.InClusterConfig()
	// config, err := outClustxerConfig()
	if err != nil {
		cfg, err = outClusterConfig()
		if err != nil {
			klog.Fatalf("failed to build config from flags: %v", err)
		}
	}

	clientset := kubernetes.NewForConfigOrDie(cfg)
	cm, err := clientset.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to get configmaps: %v", err)
	}
	fmt.Printf("There are %d configmaps in this cluster.\n", len(cm.Items))
	fmt.Printf("Configmaps: %v\n", cm.Items)

	watcher, err := clientset.CoreV1().ConfigMaps("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to get configmaps: %v", err)
	}

	for ev := range watcher.ResultChan() {
		cm, ok := ev.Object.(*v1.ConfigMap)
		if !ok {
			klog.Errorf("not a corev1.ConfigMap, instead got %T\n", ev.Object)
		}
		switch ev.Type {
		case watch.Added:
			fmt.Printf("event %s cm: %#v\n", ev.Type, cm)
			composefile, ok := cm.Annotations[filenameAnnotation]
			if !ok {
				continue
			}
			target = getDefaultTarget(composefile)

			// annotation value and key match.
			if _, ok := cm.Data[composefile]; ok {
				// composeToK8S(composefile, cm.BinaryData[composefile])
				// if err != nil {
				// 	klog.Fatalf("Error converting compose to K8S: %v", err)
				// 	continue
				// }
				if _, hasTarget := cm.Annotations[targetNameAnnotation]; hasTarget {
					target := &Target{
						Name: cm.Annotations[targetNameAnnotation],
						Seed: cm.Name,
					}

					if _, hasTargetCMD := cm.Annotations[targetCMDAnnotation]; hasTargetCMD {
						target.CMD = strings.Split(cm.Annotations[targetCMDAnnotation], " ")
					}

				}
				runTarget(cm.Name, target, clientset)

				// cm.Annotations[procAnnotation] = time.Now().Format("2017-09-07 17:06:04.000000")
				// _, err = client.CoreV1().ConfigMaps("default").Update(context.Background(), cm, metav1.UpdateOptions{})
				// if err != nil {
				// 	klog.Fatalf("CM not updated: %v", err)
				// 	continue
				// }

			} else {
				debug("Compose file name mismatch in annotation")
			}
		case watch.Modified:
			debug("CM modified " + cm.Name + " - " + cm.Namespace)
		default:
			continue
		}
	}

	watcher.Stop()
}

func getDefaultTarget(compose string) *Target {
	envT := os.Getenv("DEFAULT_TARGET")
	defaultImage := "ipedrazas/ktools:latest" //"busybox" //"ipedrazas/c2-target-kompose"
	t := &Target{
		Name: defaultImage,
		CMD:  internal.DefaultTargetCMD(compose),
		// CMD:  internal.SleepCMD(),
		Seed: "c2-default",
	}
	if envT != "" {
		t.Name = envT
	}
	return t

}

func runTarget(cmname string, target *Target, clientset *kubernetes.Clientset) {
	jobName := target.Seed + "-" + internal.RandomString(5)
	jobs := clientset.BatchV1().Jobs(namespace)
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{

			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					ServiceAccountName: "c3",
					Containers: []v1.Container{
						{
							Name:    jobName,
							Image:   target.Name,
							Command: target.CMD,
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "compose-vol",
									MountPath: "/data",
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
					Volumes: []v1.Volume{
						{
							Name: "compose-vol",
							VolumeSource: v1.VolumeSource{
								ConfigMap: &v1.ConfigMapVolumeSource{
									LocalObjectReference: v1.LocalObjectReference{
										Name: cmname,
									},
								},
							},
						},
					},
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}
	debug(jobSpec)

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		debug(err)
		log.Fatalln("Failed to create K8s job.")
	}
}

func debug(msg ...any) {
	if DEBUG {
		fmt.Println(msg...)
	}
}

type Target struct {
	Name string
	CMD  []string
	Seed string
}
