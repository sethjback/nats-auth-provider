terraform {
  required_providers {
    natsauth = {
      source = "github.com/sethjback/nats-auth-provider"
    }
  }
}

provider "natsauth" {}

resource "natsauth_operator" "example" {
  name = "test"
}
