#!/bin/sh
git pull origin master
sudo docker stop golang-api
sudo docker rm golang-api
sudo docker build -t golang-api .
sudo docker run -d --name golang-api --network="host" -p 8080:8080 -tid golang-api