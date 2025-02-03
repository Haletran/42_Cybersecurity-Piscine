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
    "github.com/schollz/progressbar/v3"
)

var CLI struct {
    URL string `arg:"" name:"url" help:"Target URL to spider."`
    
    Recursive bool `short:"r" help:"Recursively download images."`
    Links     int `short:"l" help:"Indicates the maximum depth level of the recursive download (default: 5)"`
    Pictures  string `short:"p" help:"Path where the downloaded files will be saved. (default: /data)"`
}


func extractImgs(link string, folder string, base_url string) {
     imgURL, err := url.Parse(link)
     if err != nil {
         fmt.Println("Invalid URL:", link, err)
         return
     }
     
     if !imgURL.IsAbs() {
         baseURL, _ := url.Parse(base_url)
         imgURL = baseURL.ResolveReference(imgURL)
     }
 
     resp, err := http.Get(imgURL.String())
     if err != nil {
         fmt.Println("Error downloading:", imgURL, err)
         return 
     }
     defer resp.Body.Close()
 
     parts := strings.Split(imgURL.Path, "/")
     filename := parts[len(parts)-1]
     err = os.MkdirAll(folder, 0755)
     if err != nil {
         fmt.Println("Error creating directory:", err)
         return 
     }
 
     file, err := os.Create(filepath.Join(folder, filename))
     if err != nil {
         fmt.Println("Error creating the file:", err)
        return         
     }
     defer file.Close()
 
     _, err = io.Copy(file, resp.Body)
     if err != nil {
         fmt.Println("Error saving file:", err)
        return         
     }
}


func main() {
    ctx := kong.Parse(&CLI,
        kong.Description("Web spider tool that can recursively fetch URLs and extract all images"),
        kong.UsageOnError())

    if CLI.URL == "" {
        panic("URL is required")
    }

    if CLI.Recursive {
        if CLI.Links == 0 {
            CLI.Links = 5
        }
    } else 
    {
        CLI.Links = 0
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
        colly.MaxDepth(CLI.Links),
    )

    links := []string{}
    visited := map[string]bool{}

    bar1 := progressbar.Default(-1)

    wrongFormat := 0
    c.OnHTML("img", func(e *colly.HTMLElement) {
        link := e.Attr("src")
        format := [5]string{".png", ".gif", ".jpg", ".bmp", ".jpeg"}
        hasValidFormat := false
        for i := 0; i < len(format); i++ {
            if strings.HasSuffix(strings.ToLower(link), format[i]) {
                hasValidFormat = true
                break
            }
        }
        if !hasValidFormat {
            wrongFormat++
            return
        }
        links = append(links, link)
        bar1.Add(1)
    })
    var currentDepth int

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        href := e.Attr("href")
        if _, ok := visited[href]; !ok && currentDepth < CLI.Links {
            visited[href] = true
            currentDepth++
            c.Visit(e.Request.AbsoluteURL(href))
        }
    })

    c.Visit(CLI.URL)
    c.Wait()

    fmt.Println("\n=== Spider Results ===")
    fmt.Printf("Total images found: %d\n", len(links))
    fmt.Printf("Images with invalid format: %d\n", wrongFormat)
    fmt.Printf("Saving to directory: %s\n", CLI.Pictures)
    fmt.Printf("Recursiveness : %d\n", CLI.Links)
    fmt.Println("====================")
    
    bar := progressbar.NewOptions(len(links),
        progressbar.OptionSetDescription("Downloading images..."),
        progressbar.OptionShowCount(),
    )
    
    if len(links) == 0 {
        fmt.Println("No imgs found...")
    }

    for i := 0; i < len(links); i++ {
        bar.Add(1)
        extractImgs(links[i], CLI.Pictures, CLI.URL)
    }
    
    fmt.Printf("\n\n✓ Successfully downloaded %d images to %s\n", len(links), CLI.Pictures)
}