This server will serve image to text and text to image api, it will have history , users , and prompt save.
# REQUIREMENTS
You need to have go install

# INSTALLATION

## First step of the installation for postgres

Clone the repository then run , cd into it:

Create .env files - In this .env file you should have the following :

```bash
API_TOKEN=<your-api-token>
URL_TEXT_TO_IMAGE=https://api-inference.huggingface.co/models/stabilityai/stable-diffusion-xl-base-1.0
URL_IMAGE_TO_TEXT=https://api-inference.huggingface.co/models/nlpconnect/vit-gpt2-image-captioning
URL_IMAGE_CLASSIFICATION=https://api-inference.huggingface.co/models/google/vit-base-patch16-224
URL_TEXT_SUMMERIZATION=https://api-inference.huggingface.co/models/facebook/bart-large-cnn
URL_TEXT_CLASSIFICATION=https://api-inference.huggingface.co/models/distilbert/distilbert-base-uncased-finetuned-sst-2-english
SERVER_PORT=8080
DB_USER=pguser
DB_PASSWORD=pgpasswd
DB_NAME=ai_database
DB_HOST=localhost
DB_PORT=5432
```
Then run the following command : 

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
go mod build .
go run .
```

And that's it
