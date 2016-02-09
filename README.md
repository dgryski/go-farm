# go-farm

*Google's FarmHash hash functions implemented in Go language*

[![Master Branch](https://img.shields.io/badge/-master:-gray.svg)](https://github.com/cognitivelogic/go-farm/tree/master)
[![Master Build Status](https://secure.travis-ci.org/cognitivelogic/go-farm.png?branch=master)](https://travis-ci.org/cognitivelogic/go-farm?branch=master)
[![Master Coverage Status](https://coveralls.io/repos/cognitivelogic/go-farm/badge.svg?branch=master&service=github)](https://coveralls.io/github/cognitivelogic/go-farm?branch=master)
*
[![Develop Branch](https://img.shields.io/badge/-develop:-gray.svg)](https://github.com/cognitivelogic/go-farm/tree/develop)
[![Develop Build Status](https://secure.travis-ci.org/cognitivelogic/go-farm.png?branch=develop)](https://travis-ci.org/cognitivelogic/go-farm?branch=develop)
[![Develop Coverage Status](https://coveralls.io/repos/cognitivelogic/go-farm/badge.svg?branch=develop&service=github)](https://coveralls.io/github/cognitivelogic/go-farm?branch=develop)


## Description

FarmHash, a family of hash functions.

This is a (mechanical) translation of the non-SSE4/non-AESNI hash functions from Google's FarmHash (https://github.com/google/farmhash).


FarmHash provides hash functions for strings and other data.
The functions mix the input bits thoroughly but are not suitable for cryptography.

All members of the FarmHash family were designed with heavy reliance on previous work by Jyrki Alakuijala, Austin Appleby, Bob Jenkins, and others.

For more information please consult https://github.com/google/farmhash


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
