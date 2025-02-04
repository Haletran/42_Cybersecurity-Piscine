package main


import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
    "github.com/alecthomas/kong"
	"github.com/rwcarlsen/goexif/tiff"
    "log"
    "os"
    "strings"
)

type CLI struct {
    File []string `arg:"" name:"url" help:"Path of the images"`
}

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
    fmt.Printf("%s: %s\n", name, tag)
    return nil
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

func main() {
	var cli CLI
	ctx := kong.Parse(&cli,
        kong.Description("Parse imgs and extract EXIF data"),
        kong.UsageOnError())

	if cli.File[0] == "" {
		panic("One file is required...")
	}
    ctx.Validate()

    for i := 0; i < len(cli.File); i++ {
        fmt.Printf("\033[1;34m=> Processing file: %s\033[0m\n", cli.File[i])
        
        if checkFormat(cli.File[i]) == false {
            log.Fatal("Wrong img extensions")
        }

        f, err := os.Open(cli.File[i])
        if err != nil {
            log.Fatal("Error opening file:", err)
        }
        defer f.Close()
        
        x, err := exif.Decode(f)
        if err != nil {
            log.Fatal("Error decoding EXIF:", err)
        }
        lat, long, _ := x.LatLong()
        fmt.Println("lat, long: ", lat, ", ", long)

        var p Printer
        err = x.Walk(p)
        if err != nil {
            log.Fatal("Error walking EXIF data:", err)
        }
    }
}