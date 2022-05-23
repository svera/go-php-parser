# Simple AST generation in PHP vs Go

This is a small test to check performance of solutions to generate ASTs of PHP code in Go and PHP.

## Requirements

* PHP 7.4 or up.
* Composer.
* Go 1.18 or up.

## Go

1. Build a binary with `go build -mod=readonly`.
2. Run it with `./go-php-parser <absolute path to source PHP file>`.

## PHP

1. Install dependencies with `composer install`.
2. Run `php index.php <absolute path to source PHP file>`
