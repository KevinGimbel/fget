# !!! THIS IS A WORK IN PROGRESS AND JUST A FIRST DRAFT. !!!


---

# fget
> Download files using Go `net/http` and `io`

"Dumb" file downloader for the command line which will just download whatever
file you request. Probably has lots of issues, e.g. it just dumps the content it receives into a file with the same name. For example, requesting a `index.php` file will create a `index.php` file with the content of the rendered page on your local machine. Consider the following example: 

Content of http://example.com/index.php 
```php
<?php
echo "Hello world"
```
Running the script against this url with 
``` 
$ go run main.go http://example.com/index.php 
```

Will result in the download of a file named `index.php` to your local machine with the content

```txt
Hello world
```

Downloading images or tar, zip, binary files works. 

### Usage

Clone the repo and execute the `main.go` file.

```
$ go run main.go https://i.kevingimbel.me/fget/sample.png
```

Will download `sample.png` to the current directory.
