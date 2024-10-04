#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -dbname=bookings-go -dbuser=postgres -cache=false -production=false