package model

type Model struct {
	Service  string     `json:"service"`
	NodeList []NodeList `json:"nodeList"`
}

type NodeList struct {
	Name       string      `json:"name"`
	Ip         string      `json:"ip"`
	Port       string      `json:"port"`
	Path       string      `json:"path"`
	LbMap      []WorkerMap `json:"lbMap"`
	PodList    []PodList   `json:"podList"`
	JenkinsURL JenkinsURL  `json:"jenkinsURL"`
}

type PodList struct {
	Name       string      `json:"name"`
	ExcludeMap []WorkerMap `json:"excludeMap"`
	LogPath    string      `json:"logPath"`
	WebappPath string      `json:"webappPath"`
	FileName   string      `json:"fileName"`
}

type WorkerMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JenkinsURL struct {
	BasicURL   string `json:"basicURL"`
	MiddleURL  string `json:"middleURL"`
	JobName    string `json:"jobName"`
	GroupId    string `json:"groupId"`
	ArtifactId string `json:"artifactId"`
	Version    string `json:"version"`
}
