package agility

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
		},

		//define the supported resources and point to their respective .go classes
		ResourcesMap: map[string]*schema.Resource{
			"agility_compute":			resourceAgilityCompute(),
			"agility_blueprint":		resourceAgilityBlueprint(),
			"agility_project":          resourceAgilityProject(),
			"agility_environment":      resourceAgilityEnvironment(),
		},
	}
}

var descriptions map[string]string