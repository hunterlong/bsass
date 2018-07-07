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
You can do math, and more complex math of other objects
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
Just like normal, but you can include an http file.
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



