provider "aws" {
  #access_key = ""
  #secret_key = ""
  region = "us-east-2"
}

module "aws-network-configs" {
  source = "./network-configs"
}