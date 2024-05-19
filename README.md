# Exchange Rates Servise
This project can send you USD to UAH exchange rate via API and subscribe you for everyday exchange rate update.

## Used services and libraries
### Services
- [Mailersend](https://www.mailersend.com/)
- [Currency API](https://www.jsonapi.co/public-api/Currency-api)
### Libraries
- [mailersend-go](https://github.com/mailersend/mailersend-go)
- [gorm](https://gorm.io/)
- [cron](https://github.com/robfig/cron)

## API
Send get request to `{{your_url}}/api/rate`.

Retrieve response as string currency rate.

Example response:
``` 
39.523
```  


## How to run

### Setup enviroment
Copy `.env.template` to `.env`.
```sh
cp .env.template.env
```
Set enviroment variables for your configuration in `.env`.
``` sh
DB_USER=DB_USER
DB_PASSWORD=DB_PASSWORD
DB_NAME=DB_NAME
MAILERSEND_API_KEY=MAILERSEND_API_KEY
MAILERSEND_INFO_MAIL=MAILERSEND_INFO_MAIL
```

### Docker
Run docker to build and start containers.
``` sh
docker-compose build
docker-compose run
```