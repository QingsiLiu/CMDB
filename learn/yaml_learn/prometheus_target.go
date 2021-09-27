package yaml_learn

import (
	"gopkg.in/yaml.v2"
	"os"
)

type TargetConfig struct {
	Targets []string `yaml:"targets"`
}

func NewTargetConfig(targets ...string) *TargetConfig {
	return &TargetConfig{targets}
}

func main() {
	f, _ := os.Create("mysqld.yaml")

	encoder := yaml.NewEncoder(f)
	targets := []*TargetConfig{NewTargetConfig("127.0.0.1:3360", "127.0.0.1:3360")}
	encoder.Encode(&targets)
}
