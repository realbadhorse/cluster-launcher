package main

import (
    "fmt"
    "os"

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

    client := CreateClient(cfg)

    fmt.Println("Client returned from func: ", client)

    if err := func(); err != nil {
        return fmt.Errorf("Error running func called from run(): %w", err)
    } 
    return nil
}
