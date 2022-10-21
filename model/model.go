package model

type Model struct {
	Service  string     `json:"service"`
	NodeList []NodeList `json:"nodeList"`
}

type NodeList struct {
	Name    string      `json:"name"`
	Ip      string      `json:"ip"`
	Port    string      `json:"port"`
	Path    string      `json:"path"`
	LbMap   []WorkerMap `json:"lbMap"`
	PodList []PodList   `json:"podList"`
}

type PodList struct {
	Name       string      `json:"name"`
	ExcludeMap []WorkerMap `json:"excludeMap"`
	LogPath    string      `json:"logPath"`
}

type WorkerMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
