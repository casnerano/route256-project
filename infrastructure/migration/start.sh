#!/bin/bash

goose -dir ./migrations/cart postgres postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/${POSTGRES_DB_CART}?sslmode=disable up
goose -dir ./migrations/loms postgres postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/${POSTGRES_DB_LOMS}?sslmode=disable up