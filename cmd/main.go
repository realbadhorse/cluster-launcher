package main

import (
    "fmt"
    "os"

    "github.com/realbadhorse/cluster-launcher/config"
    "github.com/realbadhorse/cluster-launcher/digitalocean"
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

    client := CreateClient(cfg)

    fmt.Println("Client returned from func: ", client)

    if err := func(); err != nil {
        return fmt.Errorf("Error running func called from run(): %w", err)
    } 
    return nil
}
