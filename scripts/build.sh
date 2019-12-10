#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
outputPath="build"
app="gql-server"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"

printf "\nBuilding: $app\n"
time env GOOS=linux GOARCH=amd64 go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"
