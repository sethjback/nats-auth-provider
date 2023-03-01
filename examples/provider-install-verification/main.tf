terraform {
  required_providers {
    natsauth = {
      source = "github.com/sethjback/nats-auth-provider"
    }
  }
}

provider "natsauth" {}

data "natsauth_operator" "example" {}
