package internal

func LoadTarget() {

}

func LoadContextMetadata() {}

// type Target struct {
// 	Name           string
// 	Platform       []string
// 	Domain         string `yaml:"dns,omitempty"`
// 	Overwrite      bool   `yaml:",omitempty"`
// 	Registry       string `yaml:",omitempty"`
// 	RegistryUserId string `yaml:"registry_user,omitempty"`

// 	// Image is a setting to override the docker image name (app/slug name)
// 	Image   string `yaml:",omitempty"`
// 	Compose string `yaml:"compose,omitempty"`

// 	// Actions are the ordered list of actions the target will execute
// 	// A Compose file can have several services, in case order is important
// 	// this parameter allows you to define the right order
// 	Actions     []string `yaml:"actions,omitempty"`
// 	DockerBuild bool     `yaml:"docker_build"`
// 	Paused      bool     `yaml:"paused,omitempty"`
// }

// type Component struct {
// 	Name      string   `yaml:",omitempty"`
// 	Slug      string   `yaml:",omitempty"`
// 	Port      int      `yaml:",omitempty"`
// 	Version   string   `yaml:",omitempty"`
// 	Workspace string   `yaml:",omitempty"`
// 	Src       string   `yaml:",omitempty"`
// 	Lang      string   `yaml:",omitempty"`
// 	Cmd       string   `yaml:",omitempty"`
// 	Overwrite bool     `yaml:",omitempty"`
// 	Config    Conf     `yaml:",omitempty"`
// 	Targets   []Target `yaml:",omitempty"`
// 	Secrets   Secrets  `yaml:",omitempty"`
// 	Path      string   `yaml:",omitempty"`
// }
