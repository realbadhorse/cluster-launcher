package digitalocean

import (
    "fmt"
    "context"
    "log"


    "github.com/digitalocean/godo"
)

// func GetResources() {
//    
//}

func ListProjects(ctx context.Context, client *godo.Client) {
    fmt.Println("Projects:")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        projects, resp, err := client.Projects.List(ctx, opt)
        if err != nil {
            log.Printf("error fetching projects: ", err)
            return
        }
        for _, project := range projects {
            fmt.Printf("- ID: %s, Name: %s\n", project.ID, project.Name)
        }
        if resp.Links.IsLastPage() {
            break
        }
        page, err := resp.Links.CurrentPage()
        if err != nil {
            log.Printf("error fetching current page: ", err)
            return
        }
        opt.Page = page + 1
    }
    fmt.Println()
}
