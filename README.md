# stdin-matcher-go

A utility that measures the time for some input to appear in stdin.

Useful when measuring some part of the process that keeps going after emitting some output

For example, measuring how long it takes for [webpack-dev-middleware](https://github.com/webpack/webpack-dev-middleware)
to finish the initial compilation, even though the process continues. 

## Example

```sh
webpack --mode development | ./stdin-matcher-go Compilation successful
```

## Building

```sh
go build
```
