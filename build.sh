#!/bin/sh

echo 'clean dist'
rm -rf ./dist/*

echo 'set basic dir & files'

mkdir ./dist/mac
mkdir ./dist/linux

cp -r ./init_files/* ./dist/mac/
cp -r ./init_files/* ./dist/linux/

echo  'build web'
cd ./NKNMining/web/
npm install
npm run build

mkdir -p ../../dist/mac/web/
mkdir -p ../../dist/linux/web/

cp -r ./dist/* ../../dist/mac/web/
cp -r ./dist/* ../../dist/linux/web/

cd ../src/NKNMining/

echo 'install golang package'
glide install

echo 'build MAC version'
GOPATH=$GOPATH:$PWD/../../ GOOS=darwin GOARCH=amd64 go build
mv ./NKNMining ../../../dist/mac/

echo 'build linux version'
GOPATH=$GOPATH:$PWD/../../ GOOS=linux GOARCH=amd64 go build
mv ./NKNMining ../../../dist/linux/

echo 'done! ^_T'
