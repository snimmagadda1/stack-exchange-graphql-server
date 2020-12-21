#!/bin/sh
echo "WARNING: downloading and unzipping data for: $1";
curl -O "https://archive.org/download/stackexchange/${1}.stackexchange.come.7z"
mkdir "${1}.stackexchange.come.7z/"
unzip -d "${1}.stackexchange.come.7z/" "${1}.stackexchange.come.7z" 