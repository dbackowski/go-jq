### go-jq

Very simple command-line JSON processor written in Go.

## Usage

* clone the repo
* go to cloned repo directory and run:

```sh
go build
```

```sh
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"}}" | ./go-jq '.repo'
```

![screenshot](https://i.imgur.com/DdHli1l.png)

```sh
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"}}" | ./go-jq '.repo.type'
```

![screenshot](https://i.imgur.com/nYVdSFK.png)
