#!/bin/bash
echo domainfinderをビルドします...
go build -o domainfinder.exe
echo synonymsをビルドします
cd ../synonyms
go build -o ../domainfinder/lib/synonyms.exe
echo availavleをビルドします
cd ../available
go build -o ../domainfinder/lib/available.exe
echo sprinkleをビルドします
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle.exe
echo coolifyをビルドします
cd ../coolify
go build -o ../domainfinder/lib/coolify.exe
echo domainifyをビルドします
cd ../domainify
go build -o ../domainfinder/lib/domainify.exe
echo 完了