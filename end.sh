#!/bin/bash
if [ $(which docker) ]; then
    if [ $(which docker-compose) ]; then
        echo "Docker and Docker compose installed"
        docker-compose down

    else
        echo "Please Install Docker read this https://docs.docker.com"
    fi
else
    echo "Please Install Docker read this https://docs.docker.com"
fi