# Go CLI Simple
[![Build Status](https://travis-ci.org/samuskitchen/go-cli-simple.svg?branch=master)](https://travis-ci.org/samuskitchen/go-cli-simple)


This project is purely formative, it has been created to illustrate a simple CLI based on the article of a [blog](https://blog.friendsofgo.tech/posts/crear-tu-primer-cli-en-go/).

This cli use the Banc Ipsum API for generate random texts and print in the output.

## Getting started

To install the library, run

```sh
go get -u github.com/samuskitchen/go-cli-simple/cmd/bacon-ipsum
```

## Use the cli

```sh
$ bacon-ipsum --help
usage: bacon-ipsum [OPTIONS]
        bacon-ipsum is a simple tool to generate random text based on a bacon ipsum API

  -paras int
        number of paragraphs (default 5)
  -sentences int
        number of sentences (this overrides paragraphs)
  -type string
        Type of the text to generate (Required) [Valid options: all-meat, meat-and-filler]
  -withLorem
        if it is true the first paragraph start with 'Bacon dolor sit amet (Carrots with bacon)'
```