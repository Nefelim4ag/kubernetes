package calico

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/golang/glog"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
)

const (
	calicoctlDefaultPath = "/opt/bin/calicoctl-wrapper"
)

type endpointMetadata struct {
	Workload, Name, Node string
}

func ReleasePodNetwork(namespace, podName, ip string) error {
	metadata, err := getPodEndpointMetadata(namespace, podName)
	if err != nil {
		return err
	}

	_, err = execCalicoctl("delete", "workloadEndpoint",
		"--workload", metadata.Workload,
		"-n", metadata.Node,
		"--orchestrator", "k8s", metadata.Name)

	if err != nil {
		return err
	}

	if ip != "" {
		_, err = execCalicoctl("ipam", "release", "--ip", ip)
		if err != nil {
			return err
		}

		glog.Infof("calico IP %s successfully released", ip)
	}

	return nil
}

func getPodEndpointMetadata(namespace, podName string) (metadata endpointMetadata, err error) {
	out, err := execCalicoctl("get", "workloadEndpoint", "-o", "json")
	if err != nil {
		return
	}

	endpoints := make([]struct{ Metadata endpointMetadata }, 0)
	if err = json.Unmarshal(out, &endpoints); err != nil {
		return
	}

	podWorkload := fmt.Sprintf("%s.%s", namespace, podName)
	for _, v := range endpoints {
		if podWorkload == v.Metadata.Workload {
			return v.Metadata, nil
		}
	}

	return endpointMetadata{}, errors.New("pod workload not found")
}

func execCalicoctl(args ...string) (out []byte, err error) {
	calicoctlPath := os.Getenv("CALICOCTL_PATH")
	if calicoctlPath == "" {
		calicoctlPath = calicoctlDefaultPath
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	c := exec.Cmd{
		Path:   calicoctlPath,
		Args:   append([]string{calicoctlPath}, args...),
		Stdin:  nil,
		Stdout: stdout,
		Stderr: stderr,
	}

	if err := c.Run(); err != nil {
		glog.Errorf("Error running calicoctl: %v\n %s \n %s", err, stdout.Bytes(), stderr.Bytes())
		return nil, err
	}

	return stdout.Bytes(), nil
}
