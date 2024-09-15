package main

import (
    "fmt"
    "os"
    "context"

    "github.com/realbadhorse/cluster-launcher/internal/config"
    "github.com/realbadhorse/cluster-launcher/internal/digitalocean"
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    } else {
        os.Exit(0)
    }
}

func run() error {
    config := config.Config{
        APIToken: os.Getenv("DO_API_TOKEN"),
    }

    if config.APIToken == "" {
        fmt.Println("env var DO_API_TOKEN is empty")
        return nil
    }

    client := digitalocean.CreateClient(config)

    ctx := context.Background()

    digitalocean.ListProjects(ctx, client)

    return nil
}
