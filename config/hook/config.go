package hook

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("bitbucket_hook", func(r *config.Resource) {
        // Set a custom group for the hook resource
        r.ShortGroup = "hook"

        // Define a reference to the repository in which this hook will be created
        r.References["repository"] = config.Reference{
            Type: "github.com/upbound/provider-bitbucket/apis/repository/v1alpha1.Repository",
        }
    })
}
