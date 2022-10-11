package main

import (
	"github.com/Arvintian/go-utils/cmdutil/flagtools"
	"github.com/Arvintian/go-utils/cmdutil/pflagenv"
	cli "github.com/rancher/wrangler-cli"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var Version = "0.0.0-dev"

type RootCommand struct {
	Host string `name:"host" usage:"bind host" default:"0.0.0.0"`
	Port int    `name:"port" usage:"bind port" default:"8000"`
}

func (r *RootCommand) Run(cmd *cobra.Command, args []string) error {
	pflagenv.ParseSet(pflagenv.Prefix, cmd.Flags())
	klog.Infof("Run on %s:%d", r.Host, r.Port)
	<-cmd.Context().Done()
	return nil
}

func main() {
	root := cli.Command(&RootCommand{}, cobra.Command{
		Long: "{{cookiecutter.description}}",
	})
	flagtools.BindFlags(root.Flags())
	cli.Main(root)
}
