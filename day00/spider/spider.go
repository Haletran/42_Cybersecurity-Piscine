package main

import (
    "github.com/alecthomas/kong"
    "fmt"
    "github.com/gocolly/colly"
    "github.com/schollz/progressbar/v3"
)

type CLI struct {
    URL string `arg:"" name:"url" help:"Target URL to spider."`
    
    Recursive bool `short:"r" help:"Recursively download images."`
    DepthLevel     int `short:"l" help:"Indicates the maximum depth level of the recursive download (default: 5)"`
    Folder  string `short:"p" help:"Path where the downloaded files will be saved. (default: /data)"`
}

func setupTags(c *colly.Collector, cli *CLI, links *[]string, visited map[string]bool, wrongFormat *int, currentDepth *int, bar1 *progressbar.ProgressBar) {
    c.OnHTML("img", func(e *colly.HTMLElement) {
        link := e.Attr("src")
		bar1.Add(1)
		if checkFormat(link) == true { 
			*links = append(*links, link) 
		} else {
			(*wrongFormat)++
		}
    })


    c.OnHTML("link", func(e *colly.HTMLElement) {
        href := e.Attr("href")
        bar1.Add(1)
        if checkFormat(href) == true { 
            *links = append(*links, href) 
        } else {
            (*wrongFormat)++
        }
    })


    c.OnHTML("meta", func(e *colly.HTMLElement) {
        href := e.Attr("content")
        bar1.Add(1)
        if checkFormat(href) == true { 
            *links = append(*links, href) 
        } else {
            (*wrongFormat)++
        }
    })

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        href := e.Attr("href")
		if checkFormat(href) == true { 
			*links = append(*links, href)
		} else { 
			(*wrongFormat)++ 
		}
        if _, ok := visited[href]; !ok && *currentDepth < cli.DepthLevel {
            visited[href] = true
            (*currentDepth)++
            c.Visit(e.Request.AbsoluteURL(href))
        }
    })
}

func main() {
    var cli CLI
    links := []string{}
    visited := map[string]bool{}
    wrongFormat := 0
    var currentDepth int

    ctx := kong.Parse(&cli,
        kong.Description("Web spider tool that can recursively fetch URLs and extract all images"),
        kong.UsageOnError())

    parseInput(&cli)
    ctx.Validate()

    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
        colly.MaxDepth(cli.DepthLevel),
    )
    bar1 := progressbar.Default(-1)
    setupTags(c, &cli, &links, visited, &wrongFormat, &currentDepth, bar1)
    c.Visit(cli.URL)
    c.Wait()

    fmt.Println("\n=== Spider Results ===")
    fmt.Printf("Total images found: %d\n", len(links))
    fmt.Printf("Images with invalid format: %d\n", wrongFormat)
    fmt.Printf("Saving to directory: %s\n", cli.Folder)
    fmt.Printf("Recursiveness: %d\n", cli.DepthLevel)
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
        extractImgs(links[i], cli.Folder, cli.URL)
    }

    fmt.Printf("\n\n✓ Successfully downloaded %d images to %s\n", len(links), cli.Folder)
}
