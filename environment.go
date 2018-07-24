package model

import (
	"errors"
	"log"
	"net/url"
)

type Environment struct {
	Labels
	Component

	// The environment name
	Name string
	// The environment description
	Description string
	// The version of the environment descriptor
	Version Version

	// Lagoon platform settings
	Lagoon LagoonPlatform

	// The providers where to create the environment node sets
	Providers map[string]Provider
	// The orchestrator used to manage the environment
	Orchestrator Orchestrator
	// The nodesets to create
	NodeSets map[string]NodeSet
	// The software stacks to install on the created node sets
	Stacks map[string]Stack
	// The tasks which can be ran against the environment
	Tasks map[string]Task

	Hooks struct {
		Init      Hook
		Provision Hook
		Deploy    Hook
		Undeploy  Hook
		Destroy   Hook
	}
}

func Parse(logger *log.Logger, u *url.URL) (Environment, error) {
	vErrs := ValidationErrors{}
	if hasSuffixIgnoringCase(u.Path, ".yaml") || hasSuffixIgnoringCase(u.Path, ".yml") {
		var yamlEnv yamlEnvironment
		yamlEnv, err := parseYamlDescriptor(logger, u)
		if err != nil {
			return Environment{}, err
		}
		env := createEnvironment(&vErrs, &yamlEnv)
		postValidate(&vErrs, &env)
		if vErrs.HasErrors() || vErrs.HasWarnings() {
			return env, vErrs
		} else {
			return env, nil
		}
	} else {
		return Environment{}, errors.New("unsupported file format")
	}
}

func createEnvironment(vErrs *ValidationErrors, yamlEnv *yamlEnvironment) Environment {
	var env = Environment{}
	env.Name = yamlEnv.Name
	env.Description = yamlEnv.Description
	env.Version = createVersion(vErrs, "version", yamlEnv.Version)
	env.Labels = createLabels(vErrs, yamlEnv.Labels...)
	env.Lagoon = createLagoonPlatform(vErrs, yamlEnv)
	env.Tasks = createTasks(vErrs, &env, yamlEnv)
	env.Orchestrator = createOrchestrator(vErrs, &env, yamlEnv)
	env.Providers = createProviders(vErrs, &env, yamlEnv)
	env.NodeSets = createNodeSets(vErrs, &env, yamlEnv)
	env.Stacks = createStacks(vErrs, &env, yamlEnv)
	env.Hooks.Init = createHook(vErrs, env.Tasks, "hooks.init", yamlEnv.Hooks.Init)
	env.Hooks.Provision = createHook(vErrs, env.Tasks, "hooks.provision", yamlEnv.Hooks.Provision)
	env.Hooks.Deploy = createHook(vErrs, env.Tasks, "hooks.deploy", yamlEnv.Hooks.Deploy)
	env.Hooks.Undeploy = createHook(vErrs, env.Tasks, "hooks.undeploy", yamlEnv.Hooks.Undeploy)
	env.Hooks.Destroy = createHook(vErrs, env.Tasks, "hooks.destroy", yamlEnv.Hooks.Destroy)
	return env
}

func postValidate(vErrs *ValidationErrors, env *Environment) ValidationErrors {
	validationErrors := ValidationErrors{}
	return validationErrors
}
