package main

import (
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/jpeach/bonk/pkg/cli"

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

// Progname ...
const Progname = "bonk"

func main() {
	rand.Seed(time.Now().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	components := map[string]*cobra.Command{
		"kubelet":                 kubelet.NewKubeletCommand(),
		"kubeadm":                 kubeadm.NewKubeadmCommand(os.Stdin, os.Stdout, os.Stderr),
		"kubectl":                 kubectl.NewDefaultKubectlCommand(),
		"kube-scheduler":          scheduler.NewSchedulerCommand(),
		"kube-proxy":              proxy.NewProxyCommand(),
		"kube-controller-manager": manager.NewControllerManagerCommand(),
		"kube-apiserver":          apiserver.NewAPIServerCommand(),
	}

	// Check if arg0 is a component name. In that case, we are being
	// invoked from a link, so just exec that component.
	if cmd, ok := components[path.Base(os.Args[0])]; ok {
		cli.Execute(Progname, cmd.Execute)
	}

	root := &cobra.Command{
		Use:   "bonk",
		Short: "bonk is a Kubernetes",
	}

	for _, c := range components {
		root.AddCommand(c)
	}

	cli.Execute(Progname, root.Execute)
}
