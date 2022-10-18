package model

type Model struct {
	Service string    `json:"service"`
	PodList []PodList `json:"PodList"`
}

type PodList struct {
	Name string `json:"name"`
	Ip         string       `json:"ip"`
	Port       string       `json:"port"`
	Path       string       `json:"path"`
	LbMap      []WorkerMap  `json:"lbMap"`
	TomcatList []TomcatList `json:"tomcatList"`
}

type TomcatList struct {
	Name       string      `json:"name"`
	ExcludeMap []WorkerMap `json:"excludeMap"`
}

type WorkerMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
