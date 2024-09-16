package digitalocean

import (
    "github.com/digitalocean/godo"
)

func CreateClient(token string) *godo.Client {
    client := godo.NewFromToken(token)
    return client
}
