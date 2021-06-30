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
	"log"
	"os"

	"github.com/evankanderson/periscope/pkg/remoteproxy"
	"github.com/spf13/cobra"
)

// Flags
var (
	httpPort *int
	grpcPort *int
)

var InnerCmd = &cobra.Command{
	Use:   "inner",
	Short: "Periscope inner proxy",
	Long:  "Inner proxy that receives requests from the outer periscope instance",
	Run: func(cmd *cobra.Command, args []string) {
		proxy, err := remoteproxy.NewLocalProxy(*httpPort, *grpcPort)
		if err != nil {
			log.Printf("Failed to initialize: %s", err)
			os.Exit(2)
		}
		log.Printf("HTTP Proxy on :%d, GRPC server on %d", *httpPort, *grpcPort)
		if err := proxy.Start(); err != nil {
			log.Printf("Failed to start services: %s", err)
			os.Exit(100)
		}
	},
}

func init() {
	httpPort = InnerCmd.Flags().IntP("port", "p", 8080, "Local HTTP proxy port out of the cluster")
	grpcPort = InnerCmd.Flags().IntP("server", "s", 5000, "GRPC proxy service port for connection")

	//	RootCmd.AddCommand(InnerCmd)
}
