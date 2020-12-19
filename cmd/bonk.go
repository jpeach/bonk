package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
	kubectl "k8s.io/kubectl/pkg/cmd"
	manager "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	proxy "k8s.io/kubernetes/cmd/kube-proxy/app"
	scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/cmd"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
)

// NewRootCmd ...
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:                   "bonk",
		Short:                 "🛠",
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	root := NewRootCmd()

	root.AddCommand(
		kubelet.NewKubeletCommand(),
		kubeadm.NewKubeadmCommand(os.Stdin, os.Stdout, os.Stderr),
		kubectl.NewDefaultKubectlCommand(),
		// API server fails with this error:
		//
		// ../../go/pkg/mod/k8s.io/kubernetes@v1.20.1/cmd/kube-apiserver/app/server.go:477:70: undefined: "k8s.io/kubernetes/pkg/generated/openapi".GetOpenAPIDefinitions
		//
		// apiserver.NewAPIServerCommand(),
		manager.NewControllerManagerCommand(),
		proxy.NewProxyCommand(),
		scheduler.NewSchedulerCommand(),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
