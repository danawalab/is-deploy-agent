package utils

//type DynamicType struct {
//	FileType   string `yaml:"fileType"`
//	DeployType string `yaml:"deployType"`
//}

type Apache struct {
	Name        string       `yaml:"name"`
	Agent       Agent        `yaml:"agent"`
	Path        string       `yaml:"path"`
	LbMap       []UriMap     `yaml:"lbMap"`
	TomcatLists []TomcatList `yaml:"podList"`
	JenkinsURL  *JenkinsURL  `yaml:"jenkinsURL,omitempty"`
}

type TomcatList struct {
	Name       string   `yaml:"name"`
	LbMap      []UriMap `yaml:"lbMap"`
	LogPath    string   `yaml:"logPath"`
	ShPath     string   `yaml:"shPath,omitempty"`
	WebappPath string   `yaml:"webappPath,omitempty"`
	FileName   string   `yaml:"fileName,omitempty"`
}

type UriMap struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Agent struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type JenkinsURL struct {
	BasicURL   string `yaml:"basicURL"`
	MiddleURL  string `yaml:"middleURL"`
	JobName    string `yaml:"jobName"`
	GroupId    string `yaml:"groupId"`
	ArtifactId string `yaml:"artifactId"`
	Version    string `yaml:"version"`
}
