package yaml_learn

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type FileSdConfig struct {
	Files []string `yaml:"files"`
}

func NewFileSdConfig(files ...string) *FileSdConfig {
	return &FileSdConfig{files}
}

type ScrapeConfig struct {
	JobName       string            `yaml:"job_name"`
	BasicAuth     map[string]string `yaml:"basic_auth"`
	StaticConfigs interface{}       `yaml:"static_configs,omitempty"`
	FileSdConfigs []*FileSdConfig   `yaml:"file_sd_configs"`
}

type PrometheusConfig struct {
	Global        interface{}     `yaml:"global"`
	Alerting      interface{}     `yaml:"alerting"`
	RuleFiles     []string        `yaml:"rule_files"`
	ScrapeConfigs []*ScrapeConfig `yaml:"scrape_configs"`
}

func NewScrapeConfig(job, username, passwd string) *ScrapeConfig {
	paths := []string{
		fmt.Sprintf("sd/file/%s/*.yaml", job),
		fmt.Sprintf("sd/file/%s/*.json", job),
	}

	return &ScrapeConfig{
		JobName:       job,
		BasicAuth:     map[string]string{"username": username, "password": passwd},
		FileSdConfigs: []*FileSdConfig{NewFileSdConfig(paths...)},
	}
}

func main() {
	f, err := os.Open("prometheus.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	var config PrometheusConfig
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", config)

	var scrape = NewScrapeConfig("mysqld", "", "")
	config.ScrapeConfigs = append(config.ScrapeConfigs, scrape)

	output, err := os.Create("prometheus_v2.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	encoder := yaml.NewEncoder(output)
	err = encoder.Encode(config)
	if err != nil {
		log.Fatal(err)
	}
}
