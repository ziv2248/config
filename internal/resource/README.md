
## Preparation

### 1. Install godotenv tool

```bash
$ go get -v -t github.com/joho/godotenv/cmd/godotenv 
```

### 2. Configurate Your `.env` File from `.env.sample`

```bash
$ cp .env.sample .env
```
> ⚠️ Don't modify the `.env.sample` as your `go test` source

```bash
$ vim .env
```


--------
## Run Test

```bash
$ godotenv -f .env go test
```
