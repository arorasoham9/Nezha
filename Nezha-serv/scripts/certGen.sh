#!/bin/bash
mkdir Nezha_certs
rm Nezha_certs/*
echo "openssl req -new -nodes -x509 -out ./Nezha_certs/server.pem -keyout Nezha_certs/server.key -days 3650 -subj \"/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1\""
echo "make server cert"
openssl req -new -nodes -x509 -out ./certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"
echo "openssl req -new -nodes -x509 -out ./Nezha_certs/client.pem -keyout Nezha_certs/client.key -days 3650 -subj \"/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1\""
echo "make client cert"
openssl req -new -nodes -x509 -out ./Nezha_certs/client.pem -keyout Nezha_certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"
