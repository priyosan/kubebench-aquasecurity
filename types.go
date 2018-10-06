package main

type nodeType string

const (
	// MASTER a master node
	MASTER nodeType = "master"
	// NODE a node
	NODE nodeType = "node"
	// FEDERATED a federated deployment.
	FEDERATED nodeType = "federated"
)

type kubeComponent struct {
	Optional bool   `yaml:"optional"`
	Bin      string `yaml:"bin"`
	Conf     string `yaml:"conf"`
}

type MasterConfig struct {
	Components        []string       `yaml:"components"`
	Kubernetes        *kubeComponent `yaml:"kubernetes"`
	APIServer         *kubeComponent `yaml:"apiserver"`
	Scheduler         *kubeComponent `yaml:"scheduler"`
	ControllerManager *kubeComponent `yaml:"controllermanager"`
	Etcd              *kubeComponent `yaml:"etcd"`
	Flanneld          *kubeComponent `yaml:"flanneld"`
}

type NodeConfig struct {
	Components []string
	Kubernetes *kubeComponent
	Kubelet    *kubeComponent
	Proxy      *kubeComponent
}

type FederatedConfig struct {
	Components           []string
	FedAPIServer         *kubeComponent
	FedControllerManager *kubeComponent
}

type config struct {
	MasterConfig `yaml:"master"`
	//	NodeConfig      `yaml:"node"`
	//	FederatedConfig `yaml:"federated"`
}
