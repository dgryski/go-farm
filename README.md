# go-farm

*Google's FarmHash implemented in Go language*

[![Master Branch](https://img.shields.io/badge/-master:-gray.svg)](https://github.com/cognitivelogic/go-farm/tree/master)
[![Master Build Status](https://secure.travis-ci.org/cognitivelogic/go-farm.png?branch=master)](https://travis-ci.org/cognitivelogic/go-farm?branch=master)
[![Master Coverage Status](https://coveralls.io/repos/cognitivelogic/go-farm/badge.svg?branch=master&service=github)](https://coveralls.io/github/cognitivelogic/go-farm?branch=master)
*
[![Develop Branch](https://img.shields.io/badge/-develop:-gray.svg)](https://github.com/cognitivelogic/go-farm/tree/develop)
[![Develop Build Status](https://secure.travis-ci.org/cognitivelogic/go-farm.png?branch=develop)](https://travis-ci.org/cognitivelogic/go-farm?branch=develop)
[![Develop Coverage Status](https://coveralls.io/repos/cognitivelogic/go-farm/badge.svg?branch=develop&service=github)](https://coveralls.io/github/cognitivelogic/go-farm?branch=develop)


## Description

This is a (mechanical) translation of the non-SSE4/non-AESNI hash functions from Google's FarmHash.

For more information on FarmHash, please see https://github.com/google/farmhash

For a cgo library wrapping the C++ one, please see https://github.com/dgryski/go-farmhash


## Getting started

This application is written in GO language, please refere to the guides in https://golang.org for getting started.

This project include a Makefile that allows you to test and build the project with simple commands.
To see all available options:
```bash
make help
```

## Running all tests

Before committing the code, please check if it passes all tests using
```bash
make qa
```
