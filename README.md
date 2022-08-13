### go-jq

Very simple command-line JSON processor written in Go.

## Usage

* clone the repo
* go to cloned repo directory and run:

```
go build
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"}}" | ./go-jq '.repo.type'
```

![screenshot](https://i.imgur.com/kYT6V3u.png)
