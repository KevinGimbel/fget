# fget
> Download files using Go `net/http` and `io`

"Dumb" file downloader for the command line which will just download whatever
file you request from an URL. Probably has lots of issues, e.g. it just dumps the content it receives into a file with the same name. For example, requesting a `index.php` file will create a `index.php` file with the content of the rendered page on your local machine. Consider the following example:

Content of http://example.com/index.php
```php
<?php
echo "Hello world"
```
Running the script against this url with
```
$ fget http://example.com/index.php
```

Will result in the download of a file named `index.php` to your local machine with the content

```txt
Hello world
```

Downloading images or tar, zip, binary files works, too because they are "just" copied into a local file without any checking or whatsoever.

### Usage

Clone the repo and execute the `main.go` file. OR install with `go get github.com/kevingimbel/fget`.

```
# Execute main.go directly
$ go run main.go https://i.kevingimbel.me/fget/sample.png
# Install and run from everywhere (`go get github.com/kevingimbel/fget`)
$ fget https://i.kevingimbel.me/fget/sample.png
```

Both of the above will download the `sample.png` image to the current directory.
