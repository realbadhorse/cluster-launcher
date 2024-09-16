package digitalocean

import (
    "github.com/digitalocean/godo"
)

type Token string

func CreateClient(token string) *godo.Client {
    client := godo.NewFromToken(token)
    return client
}
