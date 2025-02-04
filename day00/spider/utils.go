package main

import (
    "fmt"
    "strings"
    "net/http"
    "os"
    "io"
    "net/url"
    "path/filepath"
)


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

func checkFormat(link string) bool {
	format := [5]string{".png", ".gif", ".jpg", ".bmp", ".jpeg"}
	for i := 0; i < 5; i++ {
		if strings.HasSuffix(strings.ToLower(link), format[i]) {
			return true
		}
	}
	return false
}

func parseInput(cli *CLI){
	if cli.URL == "" {
		panic("URL is required")
	}

	if cli.Recursive {
		if cli.DepthLevel == 0 {
			cli.DepthLevel = 5
		}
	} else {
		cli.DepthLevel = 0
	}

	if cli.Folder == "" {
		cli.Folder = "./data"
	} else {
		cli.Folder = "./" + cli.Folder
	}
}