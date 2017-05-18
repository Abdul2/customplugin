package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

type homeofficeclient struct {
	ApiKey     string
	Endpoint   string
	Timeout    int
	MaxRetries int
}

type aws_organization struct {
	arn string
	id  int
}

func (m *aws_organization) Id() string {
	return "id-" + m.arn + "!"
}

func (c *homeofficeclient) Createaws_organization(m *aws_organization) error {
	return nil
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

// List of supported configuration fields for aws provider.

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "API Key used to authenticate with the service provider",
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the API",
		},
		"timeout": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Max. wait time we should wait for a successful connection to the API",
		},
		"max_retries": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The max. amount of times we will retry to connect to the API",
		},
	}
}

// List of supported resources and their configuration fields.

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"aws_organization": &schema.Resource{
			SchemaVersion: 1,
			Create:        createFunc,
			Read:          readFunc,
			Update:        updateFunc,
			Delete:        deleteFunc,
			Schema: map[string]*schema.Schema{
				"arn": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"id": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
	}
}

// This is the function used to fetch the configuration params given

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := homeofficeclient{
		ApiKey:     d.Get("api_key").(string),
		Endpoint:   d.Get("endpoint").(string),
		Timeout:    d.Get("timeout").(int),
		MaxRetries: d.Get("max_retries").(int),
	}

	return &client, nil
}

// The methods defined below will get called for each resource that needs to
// get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).

func createFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*homeofficeclient)
	aws_organization := aws_organization{
		Name: d.Get("arn").(string),
		CPUs: d.Get("id").(int),
	}

	err := client.Createaws_organization(&aws_organization)
	if err != nil {
		return err
	}

	d.SetId(aws_organization.Id())

	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
