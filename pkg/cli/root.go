/*
Copyright © 2021 Evan Anderson <Evan.K.Anderson@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/evankanderson/periscope/pkg/localproxy"
	"github.com/evankanderson/periscope/pkg/remote"
	"github.com/spf13/cobra"
)

// Flags
var (
	cfgFile      string
	port         *int
	grpcServer   *string
	target       *string
	clusterSetup *bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "periscope",
	Short: "Local HTTP proxy to Kubernetes clusters",
	Long: `Local HTTP proxy to Kubernetes clusters. Simplifies calling
cluster-local services from local tests / curl:

This commands depends on your local kubeconfig environment to set up
the proxy on the cluster; it prints an HTTP_PROXY value for your shell
and continues running proxying traffic until terminated.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()
		if err := remote.EnsureTools(); err != nil {
			log.Print(err)
			os.Exit(2)
		}

		if *clusterSetup {
			log.Print("Setting up pod on remote cluster...")
			if err := remote.EnsureForwarder(); err != nil {
				log.Print(err)
				os.Exit(3)
			}
		}

		if *target == "" {
			log.Print("Connecting to pod on cluster to forward...")
			endpoint, err, done := remote.StartForward(ctx, "periscope-remote-proxy", 5000)
			if err != nil {
				log.Print(err)
				os.Exit(4)
			}
			defer done()
			*grpcServer = endpoint
		}

		if err := localproxy.StartLocalProxy(*port, *target, *grpcServer); err != nil {
			log.Printf("Failed to start proxy: %s\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.periscope.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	port = RootCmd.Flags().IntP("port", "p", 6080, "Local proxy port to listen on.")
	target = RootCmd.Flags().StringP("target", "t", "", "If set, local address to proxy requests back to")
	grpcServer = RootCmd.Flags().StringP("server", "s", "", "Remote periscope to connect to")
	clusterSetup = RootCmd.Flags().Bool("setup", false, "Set up components on the cluster")
}
