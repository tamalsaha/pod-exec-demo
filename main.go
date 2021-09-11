package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/client-go/tools/exec"
)

func main() {
	masterURL := ""
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	dc2 := dynamic.NewForConfigOrDie(config)

	gvrNode := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "nodes",
	}
	nodes, err := dc2.Resource(gvrNode).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, obj := range nodes.Items {
		fmt.Printf("%+v\n", obj.GetName())
	}
	nodes, err = dc2.Resource(gvrNode).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, obj := range nodes.Items {
		fmt.Printf("%+v\n", obj.GetName())
	}

	pod := types.NamespacedName{
		Namespace: "demo",
		Name:      "voyager-community-dbcc6b476-5scx6",
	}
	out, err := exec.Exec(config, pod, exec.Container("haproxy"), exec.Command("ls", "-l"))
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
