#! /usr/bin/env bash

DAY=$1
COOKIE=$AoCCookie
DIR=day$(printf "%02d" $DAY)

curl \
  -fSL -o $DIR/input.txt \
  -H "cookie:$COOKIE" \
  https://adventofcode.com/2019/day/$DAY/input
