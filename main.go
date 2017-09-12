package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/kevingimbel/colog"
)

var Log colog.Colog
var version = "1.0.1"

func init() {
	config := make(map[string]string)
	config["LogFormat"] = "[%s][%s] - %s"
	config["TimeFormat"] = "2006-Jan-2 15:04:05"
	Log = colog.NewColog(config)
}

func showhelp() {
	fmt.Fprintf(os.Stderr, `USAGE: %[1]s <url>
Fetch a file by URL and save it.

OPTIONS:
url - HTTP or HTTPS URL to the file which should be downloaded

EXAMPLE:
%[1]s https://i.kevingimbel.me/fget/sample.png

Version: %s
`, os.Args[0], version)
}

func main() {
	// not enough arguments
	if len(os.Args) != 2 {
		showhelp()
		os.Exit(1)
	}

	// match the file name
	r := regexp.MustCompile(`[^\\/]+$`)
	filename := r.FindStringSubmatch(os.Args[1])

	// Create http client and prepare the request
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", os.Args[1], nil)

	if err != nil {
		fmt.Println(Log.Error(fmt.Sprintf("Request failed: %s", err)))
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "fget")
	response, err := httpClient.Do(req)

	if response.StatusCode > 400 {
		fmt.Println(Log.Error(fmt.Sprintf("Unable to download file. Response: %s", response.Status)))
		os.Exit(2)
	}

	defer response.Body.Close()
	img, _ := os.Create(filename[0])
	size, err := io.Copy(img, response.Body)
	if err != nil {
		fmt.Println(Log.Error(fmt.Sprintf("Unable to save file. %s", err)))
		os.Exit(3)
	}

	fmt.Println(Log.Success(fmt.Sprintf("Downloaded file %s (%v bytes)", filename[0], size)))
}
