# bsass (Basically SASS)
`bsass` is a simple application that will render SCSS files into precompile CSS files.


# Why `bsass` and not sass.
bsass was created to be a micro scss compiler without any requirements, all you need is the precompiled binary for your operating system.

## Commands
```
bsass scss/base.scss css/base.css
```

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

.container {
    width: 780rem;
    padding: 5rem;
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

