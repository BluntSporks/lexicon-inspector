# lexicon-inspector
Golang program to print data about lexicons in the lexicon repository

## Status
Ready to use

## Environment variables
### LANGUTIL_DATA
Set the LANGUTIL_DATA directory to the location of the data files for this project, for example:
export LANGUTIL_DATA=~/go/src/github.com/BluntSporks/language-utilities/data

You should have create, read, and append privileges to the files in the LANGUTIL_DATA directory.

### LEXICON_DATA
Set the LEXICON_DATA directory to the location of the lexicon data files, for example:
export LEXICON_DATA=~/go/src/github.com/BluntSporks/lexicon/data

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/lexicon-inspector

## Usage
The program runs in one of two modes:
* cntchars: Counts characters in a lexicon.
* cntstrs: Counts substrings in a lexicon.

Usage:

    lexicon-inspector [options]

Options:

    lang=LANG                 Name of language file to inspect
    strlen=LEN                Number of runes to include in each substring
    lexdir=DIR                Location of lexicon data directory
    mode=(cntchars|cntwords)  Mode of program execution [default: cntchars]
