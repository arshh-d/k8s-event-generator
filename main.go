package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	logf.SetLogger(zap.New())
	log.Println("Starting event generator")

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error creating client config: %v", err)
	}

	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatalf("Error creating manager: %v", err)
	}

	recorder := mgr.GetEventRecorderFor("event-generator")

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-pod",
			Namespace: "events",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
				},
			},
		},
	}

	// Create the pod using client
	err = mgr.GetClient().Create(context.Background(), pod)
	if err != nil {
		log.Fatalf("Error creating pod: %v", err)
	}

	log.Printf("Pod created: %s/%s", pod.Namespace, pod.Name)

	count := 1
	for {
		log.Println(fmt.Sprintf("Event count: %d", count))
		recorder.Event(pod, v1.EventTypeNormal, fmt.Sprintf("Event count: %d", count), "The pod is now running")
		count++
		// event generation interval
		time.Sleep(time.Second)
	}
}
