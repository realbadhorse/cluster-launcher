package main

import (
    "fmt"
    "os"
    "context"

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
    apiToken := os.Getenv("DO_API_TOKEN") 

    if apiToken == "" {
        log.Printf("error: env var DO_API_TOKEN is empty")
        return nil
    }

    client := digitalocean.CreateClient(apiToken)

    ctx := context.Background()

    digitalocean.ListProjects(ctx, client)

    return nil
}
