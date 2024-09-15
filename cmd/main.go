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

    fmt.Println("API Token: ", config.APIToken)

    if config.APIToken == "" {
        fmt.Println("env var DO_API_TOKEN is empty")
        return nil
    }

    client := digitalocean.CreateClient(config)

    fmt.Println("Client returned from func: ", client)
    fmt.Println()

    ctx := context.Background()

    digitalocean.ListProjects(ctx, client)

    return nil
}
