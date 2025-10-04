package docgen

type DocumentGenerator struct {
	RoleDirectory  string
	OutputFilename string
}

type AnsibleMeta struct {
	Dependencies []string              `yaml:"dependencies"`
	GalaxyInfo   AnsibleMetaGalaxyInfo `yaml:"galaxy_info"`
}

type AnsibleMetaGalaxyInfo struct {
	Author            string   `yaml:"author"`
	Description       string   `yaml:"description"`
	Company           string   `yaml:"company"`
	License           string   `yaml:"license"`
	MinAnsibleVersion float32  `yaml:"min_ansible_version"`
	GalaxyGags        []string `yaml:"galaxy_tags"`
}
