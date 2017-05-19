
# Warning - not finished. still working on it


###  plugin to manage aws org resource using terraform 


### why
 
  - we manage aws using terraform and automate everything but there is no (unless i have missed it) a module to manage org.
  
### Organizations

   AWS Organizations is a web service that enables you to consolidate your multiple AWS accounts into an organization and
   centrally  manage  your accounts and their resources.
   
   Commands 
   
   - create-organization
   - Other (to be implemented in due course)

### How?  [to be completed]
    
    
 - Terraform providers and provisioners are provided via plugins. 
 - Plugin are services, suchor provisioner (azure or bash ). 
 - Plugins are executed as a separate process
 - Plugins communicate with terraform via RPC interface.
 - Location of binary is specified in  ~/.terraformrc



### usage [to be completed]
 
 
    > set your go enviroment and terraform envirments  
    > clone repo
    > build terraform-provider-aws
    > run terrafform as per normal
 
 
.
