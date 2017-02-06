package main

import (
	"fmt"
	"github.com/kevingimbel/colog"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var Log colog.Colog

func init() {
	config := make(map[string]string)
	config["LogFormat"] = "[%s](%s) - %s"
	config["TimeFormat"] = "2006-Jan-2"
	Log = colog.NewColog(config)
}

func showhelp() {
	fmt.Fprintf(os.Stderr, `USAGE: %[1]s <url>
Fetch a file by URL and save it.
	
OPTIONS:
url			-			HTTP or HTTPS URL to the file which should be downloaded

EXAMPLE:
%[1]s https://i.kevingimbel.me/fget/sample.png
`, os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		showhelp()
		os.Exit(1)
	}

	// match the file name
	r := regexp.MustCompile(`[^\\/]+$`)
	filename := r.FindStringSubmatch(os.Args[1])

	response, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	img, _ := os.Create(filename[0])
	b, err := io.Copy(img, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Log.Success(fmt.Sprintf("Downloaded file %s (%v bytes)", filename[0], b)))
}
