This server will serve image to text and text to image api, it will have history , users , and prompt save.

# INSTALLATION

## First step of the installation for postgres

Clone the repository then run , cd into it:

```bash
# For running postgresql in separate terminal
docker compose up
```

## Second step ,the installation for go

!!! If you don't have go , go has a pretty straight forward way to install in their website

```bash
go mod init ,
go mod tidy
go mod vendor
```

## Third step , migration of database for postgres

```bash
go run ./migration/migration.go
```

# USAGE

For usage , just run

```bash
go mod run
```

And that's it
