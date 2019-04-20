package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func dataSourceKubernetesNamespace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesNamespaceRead,

		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema("namespace", false),
			"name": {
				Type:        schema.TypeString,
				Description: "Namespace Name",
				Computed:    true,
			},
		},
	}
}
func dataSourceKubernetesNamespaceExists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesNamespaceExistsRead,
		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema("namespace", false),
			"exists": {
				Type:        schema.TypeBool,
				Description: "Is Namespace exists",
				Computed:    true,
			},
		},
	}
}

func dataSourceKubernetesNamespaceExistsRead(d *schema.ResourceData, meta interface{}) error {
	om := meta_v1.ObjectMeta{
		Name: d.Get("metadata.0.name").(string),
	}
	d.SetId(om.Name)
	return resourceKubernetesNamespaceExistsRead(d, meta)
}

func dataSourceKubernetesNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	om := meta_v1.ObjectMeta{
		Name: d.Get("metadata.0.name").(string),
	}
	d.SetId(om.Name)
	return resourceKubernetesNamespaceRead(d, meta)
}
