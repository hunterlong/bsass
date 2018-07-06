#!/usr/bin/env bash

APP="bsass"
REPO="hunterlong/bsass"

# BUILD bsass GOLANG BINS
mkdir build
xgo -go 1.10.x --targets=darwin/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=ios/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=android/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=windows-6.0/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION -linkmode external -extldflags -static" -out alpine ./

cd build
ls

mv alpine-linux-amd64 $APP
tar -czvf $APP-linux-alpine.tar.gz $APP && rm -f $APP

mv $APP-darwin-10.6-amd64 $APP
tar -czvf $APP-osx-x64.tar.gz $APP && rm -f $APP

mv $APP-darwin-10.6-386 $APP
tar -czvf $APP-osx-x32.tar.gz $APP && rm -f $APP

mv $APP-linux-amd64 $APP
tar -czvf $APP-linux-x64.tar.gz $APP && rm -f $APP

mv $APP-linux-386 $APP
tar -czvf $APP-linux-x32.tar.gz $APP && rm -f $APP

mv $APP-windows-6.0-amd64.exe $APP.exe
zip $APP-windows-x64.zip $APP.exe  && rm -f $APP.exe

mv $APP-windows-6.0-386.exe $APP.exe
zip $APP-windows-x32.zip $APP.exe  && rm -f $APP.exe

mv $APP-linux-arm-7 $APP
tar -czvf $APP-linux-arm7.tar.gz $APP && rm -f $APP

mv $APP-linux-arm-6 $APP
tar -czvf $APP-linux-arm6.tar.gz $APP && rm -f $APP

mv $APP-linux-arm-5 $APP
tar -czvf $APP-linux-arm5.tar.gz $APP && rm -f $APP

mv $APP-linux-arm64 $APP
tar -czvf $APP-linux-arm64.tar.gz $APP && rm -f $APP

mv $APP-ios-5.0-arm64 $APP
tar -czvf $APP-ios-arm64.tar.gz $APP && rm -f $APP

mv $APP-ios-5.0-armv7 $APP
tar -czvf $APP-ios-arm7.tar.gz $APP && rm -f $APP

mv $APP-android-16-386 $APP
tar -czvf $APP-android-16-x32.tar.gz $APP && rm -f $APP

mv $APP-android-16-arm $APP
tar -czvf $APP-android-16-arm.tar.gz $APP && rm -f $APP

mv $APP-linux-mips $APP
tar -czvf $APP-linux-mips.tar.gz $APP && rm -f $APP

mv $APP-linux-mips64 $APP
tar -czvf $APP-linux-mips-x64.tar.gz $APP && rm -f $APP

mv $APP-linux-mips64le $APP
tar -czvf $APP-linux-mips-x64le.tar.gz $APP && rm -f $APP

mv $APP-linux-mipsle $APP
tar -czvf $APP-linux-mipsle.tar.gz $APP && rm -f $APP

tar -zcvf $APP-android-16-aar.tar.gz $APP-android-16-aar && rm -rf $APP-android-16-aar

tar -zcvf $APP-ios-5.0-framework.tar.gz $APP-ios-5.0-framework && rm -rf $APP-ios-5.0-framework
