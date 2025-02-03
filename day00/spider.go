package main

import (
    "github.com/alecthomas/kong"
    "fmt"
    "strings"
    "net/http"
    "os"
    "io"
    "github.com/gocolly/colly"
    "net/url"
    "path/filepath"
)

var CLI struct {
    URL string `arg:"" name:"url" help:"Target URL to spider."`
    
    Recursive bool `short:"r" help:"Recursively download images."`
    Links     int `short:"l" help:"Indicates the maximum depth level of the recursive download (default: 5)"`
    Pictures  string `short:"p" help:"Path where the downloaded files will be saved. (default: /data)"`
}


func main() {
    ctx := kong.Parse(&CLI,
        kong.Description("Web spider tool that can recursively fetch URLs and extract all images"),
        kong.UsageOnError())

    if CLI.URL == "" {
        panic("URL is required")
    }

    if CLI.Links == 0 {
        CLI.Links = 5
    }

    if CLI.Pictures == "" {
        CLI.Pictures = "./data"
    } else 
    {
        CLI.Pictures = "./" + CLI.Pictures
    }

    ctx.Validate()
    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
        colly.Async(true),
        colly.MaxDepth(CLI.Links),
    )

    links := []string{}

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnHTML("img", func(e *colly.HTMLElement) {
        link := e.Attr("src")
        links = append(links, link)
    })
    c.Visit(CLI.URL)
    c.Wait()

    fmt.Printf("Found %d links\n", len(links));
    for i := 0; i < len(links); i++ {
        // Resolve relative URLs to absolute URLs
        imgURL, err := url.Parse(links[i])
        if err != nil {
            fmt.Println("Invalid URL:", links[i], err)
            continue
        }
        
        // If imgURL is relative, resolve against base URL
        if !imgURL.IsAbs() {
            baseURL, _ := url.Parse(CLI.URL)
            imgURL = baseURL.ResolveReference(imgURL)
        }
    
        resp, err := http.Get(imgURL.String())
        if err != nil {
            fmt.Println("Error downloading:", imgURL, err)
            continue
        }
        defer resp.Body.Close()
    
        parts := strings.Split(imgURL.Path, "/")
        filename := parts[len(parts)-1]
        err = os.MkdirAll(CLI.Pictures, 0755)
        if err != nil {
            fmt.Println("Error creating directory:", err)
            continue
        }
    
        file, err := os.Create(filepath.Join(CLI.Pictures, filename))
        if err != nil {
            fmt.Println("Error creating the file:", err)
            continue
        }
        defer file.Close()
    
        _, err = io.Copy(file, resp.Body)
        if err != nil {
            fmt.Println("Error saving file:", err)
            continue
        }
    
        fmt.Printf("Downloaded: %s\n", imgURL)
    }
}