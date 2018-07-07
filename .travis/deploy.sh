#!/usr/bin/env bash

# update homebrew and cypress testing to newest version by building on travis
body='{ "request": { "branch": "master", "config": { "env": { "VERSION": "'$VERSION'" } } } }'

# notify Docker hub to built this branch
if [ "$TRAVIS_BRANCH" == "master" ]
then
     curl -s -X POST \
     -H "Content-Type: application/json" \
     -H "Accept: application/json" \
     -H "Travis-API-Version: 3" \
     -H "Authorization: token $TRAVIS_API" \
     -d "$body" \
     https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-bsass/requests
fi
