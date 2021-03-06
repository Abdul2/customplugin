package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{

			"Id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"Arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Email,JoinedMethod,JoinedTimeStamp,Status

		},
	}

}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
