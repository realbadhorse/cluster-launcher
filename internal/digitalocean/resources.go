package digitalocean

import (
    "fmt"
    "context"
    "log"

    "github.com/digitalocean/godo"
)

type Droplets struct {
    
}

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

func ListDroplets(ctx context.Context, client *godo.Client) {
    
    var allDroplets []godo.Droplet

    fmt.Println("Droplets:")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        droplets, resp, err := client.Droplets.List(ctx, opt)
        if err != nil {
            log.Printf("error fetching droplets: ", err)
            return
        }
       
        allDroplets = append(allDroplets, droplets...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }
        
        opt.Page++

    }

    for _, droplet := range allDroplets {
        fmt.Printf("- ID: %d, Name: %s, Size: %s, Image: %s, Tags: %s, VPC UUID: %s\n", droplet.ID, droplet.Name, droplet.SizeSlug, droplet.Image.Name, droplet.Tags[0], droplet.VPCUUID)
    }
    fmt.Println()
}

func ListVPCs(ctx context.Context, client *godo.Client) {
    
    var allVPCs []*godo.VPC

    fmt.Println("VPCs:")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        vpcs, resp, err := client.VPCs.List(ctx, opt)
        if err != nil {
            log.Printf("error fetching VPCs: ", err)
        }

        allVPCs = append(allVPCs, vpcs...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }

        opt.Page++
    }

    for _, vpc := range allVPCs {
        fmt.Printf("- ID: %s, Name: %s, IP Range: %s\n", vpc.ID, vpc.Name, vpc.IPRange)
    }
    fmt.Println()
}

func ListLBs(ctx context.Context, client *godo.Client) {

    var allLBs []godo.LoadBalancer

    fmt.Println("LoadBalancers: ")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        lbs, resp, err := client.LoadBalancers.List(ctx, opt)
        if err != nil {
            log.Printf("error fetching loadbalancers: ", err)
        }

        allLBs = append(allLBs, lbs...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }

        opt.Page++
    }

    for _, lb := range allLBs {
        fmt.Printf("- ID: %s, Name: %s, IP address: %s\n", lb.ID, lb.Name, lb.IP)
    }
    fmt.Println()
}

func ListVolumes(ctx context.Context, client *godo.Client) {
    
    var allVolumes []godo.Volume

    fmt.Println("Volumes: ")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    params := &godo.ListVolumeParams{
        Region:         "",
        Name:           "",       
        ListOptions:    opt,
    }
    for {
        volumes, resp, err := client.Storage.ListVolumes(ctx, params)
        if err != nil {
            log.Printf("error fetching volumes: ", err)
        }

        allVolumes = append(allVolumes, volumes...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }

        opt.Page++
    }

    for _, volume := range allVolumes {
        fmt.Printf("- ID: %s, Name: %s, Droplet ID: %d\n", volume.ID, volume.Name, volume.DropletIDs[0])
    }
    fmt.Println()
}

func ListFirewalls(ctx context.Context, client *godo.Client) {
    
    var allFirewalls []godo.Firewall

    fmt.Println("Firewalls: ")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        firewalls, resp, err := client.Firewalls.List(ctx, opt)
        if err != nil {
            log.Printf("error fetching firewalls: ", err)
        }
        
        allFirewalls = append(allFirewalls, firewalls...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }

        opt.Page++
    }

    for _, firewall := range allFirewalls {
        fmt.Printf("- ID: %s, Name: %s, Droplet IDs: %d\n", firewall.ID, firewall.Name, firewall.DropletIDs)
    }
    fmt.Println()
}

func ListImages(ctx context.Context, client *godo.Client) {
    
    var allImages []godo.Image
    tag := "Custom"

    fmt.Println("Images: ")
    opt := &godo.ListOptions{Page: 1, PerPage: 200}
    for {
        images, resp, err := client.Images.ListByTag(ctx, tag, opt)
        if err != nil {
            log.Printf("error fetching images: ", err)
        }

        allImages = append(allImages, images...)

        if resp.Links == nil || resp.Links.IsLastPage() {
            break
        }

        opt.Page++
    }

    for _, image := range allImages {
        fmt.Printf("- ID: %d, Name: %s, Regions: %s\n", image.ID, image.Name, image.Regions)
    }
    fmt.Println()
}
