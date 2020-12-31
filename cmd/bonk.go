package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
	kubectl "k8s.io/kubectl/pkg/cmd"
	apiserver "k8s.io/kubernetes/cmd/kube-apiserver/app"
	manager "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	proxy "k8s.io/kubernetes/cmd/kube-proxy/app"
	scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/cmd"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var cmd *cobra.Command

	switch path.Base(os.Args[0]) {
	case "kubelet":
		cmd = kubelet.NewKubeletCommand()
	case "kubeadm":
		cmd = kubeadm.NewKubeadmCommand(os.Stdin, os.Stdout, os.Stderr)
	case "kubectl":
		cmd = kubectl.NewDefaultKubectlCommand()
	case "kube-scheduler":
		cmd = scheduler.NewSchedulerCommand()
	case "kube-proxy":
		cmd = proxy.NewProxyCommand()
	case "kube-controller-manager":
		cmd = manager.NewControllerManagerCommand()
	case "kube-apiserver":
		cmd = apiserver.NewAPIServerCommand()
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", os.Args[0])
		os.Exit(2)
	}

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
