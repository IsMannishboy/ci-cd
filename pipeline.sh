#!/bin/bash 
channged=$(git diff --name-only | cut -d'/' -f2 | uniq)
echo "Changed files: $channged"
services=()
for dir in $channged;do 
    if [ -f "$dir/Dockerfile" ];then
        services+=("$dir")
    fi
done
echo "services:$services"

# Передаём в GitHub Actions как output
echo "services=$services" >> "$GITHUB_OUTPUT"
