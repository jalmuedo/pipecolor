#  pipecolor

A pipeline utility that colore words passed to it as arguments. Windows compatible.

## Install

You can download the binary generated in each release or build it from the sources.

### Download the binary

On the [releases page](https://github.com/jalmuedo/pipecolor/releases), download the latest available binary. Note that `pipecolor` artifact is for Linux and `pipecolor.exe` artifact is for Windows.

### Build from sources

Clone the repository and build it with Go.
```
> git clone git@github.com:jalmuedo/pipecolor.git
> cd pipecolor
> go build .
```

### Create an alias (optional)

You can use the binary directly, but it is much more convenient to use aliases for the shell.

On **Linux** you can create an alias to your binary by modifying the `~/.bashrc` file.

```
color () {
  ~/.pipecolor "$@"
}
```

On **Windows PowerShell** you can add it to your profile.
```
function color {
    C:\Users\foo\pipecolor.exe @args
}
```

## Usage

The binary receives the text through Stdin and as an argument a list of words separated by spaces and returns in Stdout the same text with the colored words.

![> ps | pipecolor word1 word2](https://github.com/jalmuedo/pipecolor/blob/main/misc/img/usage-windows.png?raw=true)

## Licence

This project is [MIT licenced](license.txt)

[![Go](https://github.com/jalmuedo/pipecolor/actions/workflows/go.yml/badge.svg)](https://github.com/jalmuedo/pipecolor/actions/workflows/go.yml)