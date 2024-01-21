#  pipecolor

A pipeline utility that colore words passed to it as arguments. Windows compatible.

## Install 

Clone the repository and build it with Go.
```
> git clone git@github.com:jalmuedo/pipecolor.git
> cd pipecolor
> go build .
```


## Usage

The binary receives the text through Stdin and as an argument a list of words separated by spaces and returns in Stdout the same text with the colored words.

```
> top | pipecolor word1 word2
```

## Licence

This project is [MIT licenced](license.txt)

