package internal

import (
	"fmt"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

func ReadCompose(filename string, content []byte) *types.Services {

	workingDir := "./data"

	var configs []types.ConfigFile

	configs = append(configs, types.ConfigFile{
		Filename: filename,
		Content:  content,
	})

	opts := []func(*loader.Options){}

	details := types.ConfigDetails{
		ConfigFiles: configs,
		WorkingDir:  workingDir,
	}

	project, err := loader.Load(details, opts...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Project loaded: " + project.Name)
	services, err := project.GetServices()
	if err != nil {
		fmt.Println(err)
	}
	for _, srv := range services {
		fmt.Println("service: " + srv.Name)
	}

	return &services

	// enabled, err := project.GetServices(services...)

}
