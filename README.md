# Studapi

## Description

It is a test project. It's an api, which has such endpoints:

- /places for getting autocomplete predictions
- /shipping/prices for getting quoted prices from hey.innovative360
- /orders for getting list of orders
- /shipping/order - for creating order with shipping type and details

## Configure

check the `config.json` and change to your vars

## Install

```
git clone git@github.com:macseem/studapi.git
cd studapi
# you can add .env file to init database with your creds
docker-compose up -d
psql -h localhost -U postgres postgres < dmp.sql # or use your credentials from .env file
export GO111MODULE=on
go build
./studapi
```
