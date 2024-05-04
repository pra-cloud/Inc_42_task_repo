#!/bin/bash

docker login -u $username -p $password

docker pull pm55/next-js-app:latest
