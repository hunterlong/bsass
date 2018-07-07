#!/usr/bin/env bash

APP="bsass"
SUBAPP="cli"
REPO="hunterlong/bsass"

# BUILD bsass GOLANG BINS
mkdir build
xgo -go 1.10.x --targets=darwin/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./cli
xgo -go 1.10.x --targets=linux/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./cli
xgo -go 1.10.x --targets=ios/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./cli
xgo -go 1.10.x --targets=android/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./cli
xgo -go 1.10.x --targets=windows-6.0/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./cli
xgo -go 1.10.x --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION -linkmode external -extldflags -static" -out alpine ./cli

cd build
ls

mv alpine-linux-amd64 $APP
tar -czvf $APP-linux-alpine.tar.gz $APP && rm -f $APP

mv $SUBAPP-darwin-10.6-amd64 $APP
tar -czvf $APP-osx-x64.tar.gz $APP && rm -f $APP

mv $SUBAPP-darwin-10.6-386 $APP
tar -czvf $APP-osx-x32.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-amd64 $APP
tar -czvf $APP-linux-x64.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-386 $APP
tar -czvf $APP-linux-x32.tar.gz $APP && rm -f $APP

mv $SUBAPP-windows-6.0-amd64.exe $APP.exe
zip $APP-windows-x64.zip $APP.exe  && rm -f $APP.exe

mv $SUBAPP-windows-6.0-386.exe $APP.exe
zip $APP-windows-x32.zip $APP.exe  && rm -f $APP.exe

mv $SUBAPP-linux-arm-7 $APP
tar -czvf $APP-linux-arm7.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-arm-6 $APP
tar -czvf $APP-linux-arm6.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-arm-5 $APP
tar -czvf $APP-linux-arm5.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-arm64 $APP
tar -czvf $APP-linux-arm64.tar.gz $APP && rm -f $APP

mv $SUBAPP-ios-5.0-armv7 $APP
tar -czvf $APP-ios-arm7.tar.gz $APP && rm -f $APP

mv $SUBAPP-android-16-arm $APP
tar -czvf $APP-android-16-arm.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-mips $APP
tar -czvf $APP-linux-mips.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-mips64 $APP
tar -czvf $APP-linux-mips-x64.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-mips64le $APP
tar -czvf $APP-linux-mips-x64le.tar.gz $APP && rm -f $APP

mv $SUBAPP-linux-mipsle $APP
tar -czvf $APP-linux-mipsle.tar.gz $APP && rm -f $APP
