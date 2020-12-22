#!/bin/sh
set -ev

echo "Build & push"
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build . -t snimmagadda/stack-exchange-graphql-server:"$TRAVIS_BUILD_NUMBER"
docker tag snimmagadda/stack-exchange-graphql-server:"$TRAVIS_BUILD_NUMBER" snimmagadda/stack-exchange-graphql-server:latest
docker push snimmagadda/stack-exchange-graphql-server:"$TRAVIS_BUILD_NUMBER" && docker push snimmagadda/stack-exchange-graphql-server:latest