
provider "aws" {
  
    ApiKey = "string"
	Endpoint =  "apigateway.eu-west-1.amazonaws.com"
	Timeout   =  60
	MaxRetries = 5

}

resource "aws_organization" "org" {

  arn = "",
  "id" = ""

}
