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
.addition {
   padding: $container + 20;
}
```
```css
.addition {
   padding: 800rem;
}
```
