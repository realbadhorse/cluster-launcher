package digitalocean

import (
    "github.com/digitalocean/godo"
    "github.com/realbadhorse/cluster-launcher/internal/config"
)

func CreateClient(config config.Config) *godo.Client {
    client := godo.NewFromToken(cfg.APIToken)
    return client
}
