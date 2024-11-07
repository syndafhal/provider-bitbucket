package hook

import "github.com/crossplane/upjet/pkg/config"

const hookShortGroup = "hook"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bitbucket_hook", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = hookShortGroup
		r.Kind = "Hook"
	})
}