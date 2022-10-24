### go-jq ![Tests](https://github.com/dbackowski/go-jq/actions/workflows/test.yml/badge.svg)

Very simple command-line JSON processor written in Go.

## Usage

* clone the repo
* go to cloned repo directory and run:

```sh
go build
```

```sh
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}" | ./go-jq '.repo'
```

![screenshot](https://i.imgur.com/4B2KorZ.png)

```sh
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}" | ./go-jq '.repo.type'

```

![screenshot](https://i.imgur.com/k9prcfH.png)


```sh
echo "{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}" | ./go-jq '.test[0]'
```

![screenshot](https://i.imgur.com/mAbqeBW.png)
