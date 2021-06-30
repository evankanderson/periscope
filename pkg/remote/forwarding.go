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

package remote

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os/exec"
	"time"
)

func EnsureTools() error {
	if _, err := exec.LookPath("kubectl"); err != nil {
		return fmt.Errorf("Unable to locate `kubectl` on your PATH.")
	}
	if out, err := exec.Command("kubectl", "version").Output(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok && out == nil {
			out = exit.Stderr
		}
		return fmt.Errorf("Unable to connect to cluster:\n%s", string(out))
	}
	return nil
}

func StartForward(ctx context.Context, podname string, port int) (string, error, func() error) {
	if port == 0 {
		port = 5000
	}
	cmd := exec.CommandContext(ctx, "kubectl", "port-forward", "pod/"+podname, fmt.Sprint(port))
	err := cmd.Start()
	// Give port-forward time to get started.
	time.Sleep(500 * time.Millisecond)
	return fmt.Sprintf("localhost:%d", port), err, cmd.Wait
}

//go:embed pod-config.yaml
var manifest []byte

func EnsureForwarder() error {
	cmd := exec.Command("kubectl", "apply", "-f", "-")
	cmd.Stdin = bytes.NewReader(manifest)
	if out, err := cmd.Output(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok && out == nil {
			out = exit.Stderr
		}
		return fmt.Errorf("Unable to create remote:\n%s", out)
	}

	if out, err := exec.Command("kubectl",
		"wait",
		"--for=condition=Ready",
		"pod/periscope-remote-proxy").Output(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok && out == nil {
			out = exit.Stderr
		}
		return fmt.Errorf("Error waiting for pods to become ready:\n%s", out)
	}

	return nil
}
