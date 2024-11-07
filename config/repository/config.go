package repository

import "github.com/crossplane/upjet/pkg/config"

const repositoryShortGroup = "repository"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bitbucket_repository", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = repositoryShortGroup
		r.Kind = "Repository"
	})
}
