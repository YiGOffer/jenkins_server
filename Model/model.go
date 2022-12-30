package Model

type MainConfig struct {
	JenkinsPort string `json:"jenkins_port"`
	JenkinsAddress string `json:"jenkins_address"`
	Token string `json:"token"`
}
type DebBuildRequest struct {
    Branch string `json:"branch"`
	Business string `json:"business"`
}