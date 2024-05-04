#!/bin/bash

docker login -u $username -p $password

docker pull pm55/go-app:latest
