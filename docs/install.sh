#!/usr/bin/env bash
OS=osx
ARCH=x64
REPO=github.com/hunterlong/bsass
VERSION=$(curl -s "https://$REPO/releases/latest" | grep -o 'tag/[v.0-9]*' | awk -F/ '{print $2}')
if [ `getconf LONG_BIT` = "64" ]
then
    ARCH=x64
else
    ARCH=x32
fi
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     OS=linux;;
    Darwin*)    OS=osx;;
    CYGWIN*)    OS=windows;;
    MINGW*)     OS=windows;;
    *)          OS="UNKNOWN:${unameOut}"
esac
printf "Installing $VERSION for $OS $ARCH...\n"
FILE="https://$REPO/releases/download/$VERSION/bsass-$OS-$ARCH.tar.gz"
printf "Downloading latest version URL: $FILE\n"
curl -L -sS $FILE -o bsass.tar.gz && tar xzf bsass.tar.gz && rm bsass.tar.gz
chmod +x bsass
mv bsass /usr/local/bin/
printf "bsass $VERSION has been successfully installed in /usr/local/bin/bsass\nTry 'bsass version' to check it!\n"