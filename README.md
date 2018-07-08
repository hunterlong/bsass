<img src="https://img.cjx.io/bsasslogo.png">
<p align="center">
    <code>bsass</code> - basically sass<br>
    <a href="https://github.com/hunterlong/bsass/releases/latest">Linux</a> | <a href="https://github.com/hunterlong/bsass/releases/latest">Mac</a> | <a href="https://github.com/hunterlong/bsass/releases/latest">Windows</a> | <a href="https://github.com/hunterlong/bsass/releases/latest">Alpine</a> | <a href="https://github.com/hunterlong/bsass/releases/latest">Other</a> | <a href="https://bsass.app">bsass.app</a>
</p>

# bsass (Basically SASS)
`bsass` is a simple application that will compile `.scss` files into `.css` files using the normal `sass` parameters and protocol. 

# Why bsass is bsass
bsass was created to be a micro scss compiler without any requirements, all you need is the precompiled binary for your operating system. I personally needed something that could compile sass on [Alpine linux](https://github.com/hunterlong/bsass/releases/latest) for my Docker project called [Statup](https://github.com/hunterlong/statup).

# No Requirements
Unlike other sass compilers, bsass has 0 requirements. Download the latest release for your operating system and your good to go. `bsass` is less than 3mb after extracting! No need to install node, ruby, go, C, or anything else! 

<p align="center"><img width="75%" src="https://img.cjx.io/bsassrun.gif"></p>

# Install
###### MacOS install with brew
```bash
brew tap hunterlong/bsass
brew install bsass
```
###### Linux install with bash/curl
```bash
bash <(curl -s https://bsass.app/install.sh)
statup version
```
###### Docker snippet
```bash
FROM alpine:latest
ENV VERSION=v0.12
RUN apk --no-cache add libstdc++ ca-certificates
RUN wget -q https://github.com/hunterlong/bsass/releases/download/$VERSION/bsass-linux-alpine.tar.gz && \ 
      tar -xvzf bsass-linux-alpine.tar.gz && \ 
      chmod +x bsass && \ 
      mv bsass /usr/local/bin/bsass
```

## Commands
```bash
bsass version
//bsass v1.34
```
Compile a file by using `bsass <scss file> <output css>`.
```bash
bsass theme.scss theme.css
```
##### loads `base.scss` and exports a compiled css to `base.css`

## Variables
Variable's are just like normal `$foobar: 12px;`.
```scss
$container: 780rem;
$container-padding: 5rem;
```
```scss
.container {
    width: $container;
    padding: $container-padding;
}
```
```css
.container {
    width: 780rem;
    padding: 5rem;
}
```

## Math
You can do math, and more complex math of other objects.
```scss
$container: 780rem;
```
```scss
.math {
   padding-left: $container - 80;
   padding-right: $container + 20;
   padding-bottom: $container * 10;
   padding-top: $container / 5;
}
.divide_multi {
  padding: $container / 4 * 10 + ((420 * 50) / 4 );
}
```
```css
.math {
   padding-left: 700rem;
   padding-right: 800rem;
   padding-bottom: 7800rem;
   padding-top: 156rem;
}
.divide_multi {
  padding: 7200rem;
}
```

## Importing
Just like normal, but you can include an http file if you need to.
```scss
@import 'reset';
@import 'variables';
@import 'https://assets.statup.io/remote.css';

.container {
    width: $container;
    padding: $container-padding;
}
```
```css
html, body, ul, ol {
  margin: 0;
  padding: 0;
}

.remote_location {
    color: #bababa;
}

.awesome_test {
    background-color: orange;
}

.container {
    width: 780rem;
    padding: 5rem;
}
```

## Functions
Function can be useful, there's a couple.
```scss
$box-color: #5cd338;
```
```scss
.darken_me {
  background-color: darken($box-color, 30%);
}
.lighten_me {
  background-color: lighten($box-color, 80%);
}
```
```css
.darken_me {
  background-color: #1c3f11
}
.lighten_me {
  background-color: #4aa92d;
}
```
- [x] `darken` darken hex code color down a percentage
- [x] `lighten` lighten hex code color up a percentage

## Mixins
Mix it up and do it like you normally would with sass.
```scss
@mixin transform($property) {
  -webkit-transform: $property;
      -ms-transform: $property;
          transform: $property;
}
```
```scss
.box {
    @include transform(rotate(30deg));
}
```
```css
.box {
  -webkit-transform: rotate(30deg);
      -ms-transform: rotate(30deg);
          transform: rotate(30deg);
}
```

## Extends
Make's your CSS easier to handle for the future.
```scss
%message-shared {
  border: 1px solid #ccc;
  padding: 10px;
  color: #333;
}
```
```scss
.message {
    @extend %message-shared;
}
```
```css
.message {
  border: 1px solid #ccc;
  padding: 10px;
  color: #333;
}
```

## Features
- [x] Variables `$container: 500px`
- [x] Import `@import 'reset'`
- [x] Mixins `@mixin transform($property)`
- [x] Extends `%extend`
- [x] Functions `darken("#bababa", 25%)`

## License
MIT - I hope `bsass` is working in your application. :information_desk_person:



