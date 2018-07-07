# bsass (Basically SASS)
`bsass` is a simple application that will render SCSS files into precompile CSS files.


# Why bsass is bsass
bsass was created to be a micro scss compiler without any requirements, all you need is the precompiled binary for your operating system.

# No Requirements
Unlike other sass compilers, bsass has 0 requirements. Download the latest release for your operating system and your good to go.

## Commands
```bash
bsass version
//bsass v1.34
```
```bash
bsass scss/base.scss css/base.css
```
##### loads `base.scss` and exports a compiled css to `base.css`

## Variables
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
```
```css
.math {
   padding-left: 700rem;
   padding-right: 800rem;
   padding-bottom: 7800rem;
   padding-top: 156rem;
}
```

## Importing
```scss
@import 'reset'
@import 'variables'
@import 'https://assets.statup.io/remote.css'

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
```scss
$box-color: #5cd338;
```
```scss
.darken_me {
  background-color: darken($box-color, 30%);
}
```
```css
.darken_me {
  background-color: #1c3f11
}
```

## Mixins
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

