package digitalocean

import (
    "github.com/digitalocean/godo"
    "github.com/realbadhorse/cluster-launcher/internal/config"
)

type Token string

func CreateClient(token string) *godo.Client {
    client := godo.NewFromToken(token)
    return client
}
