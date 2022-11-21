package model

type Node struct {
	Name    string    `json:"name"`
	Agent   Agent     `json:"agent"`
	Path    string    `json:"path"`
	LbMap   []UriMap  `json:"lbMap"`
	PodList []PodList `json:"podList"`
}

type PodList struct {
	Name    string   `json:"name"`
	LbMap   []UriMap `json:"lbMap"`
	LogPath string   `json:"logPath"`
	ShPath  string   `json:"shPath"`
}

type UriMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Agent struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
