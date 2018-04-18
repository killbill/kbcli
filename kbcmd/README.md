# kbcmd - Kill bill command line

kbcmd is a command line client for killbill. kbcmd uses go client library
to talk to killbill.

## Prerequisites
For users new to go, you need to have go compiler installed. If not,

1. Install [go from here](https://golang.org/dl/)
2. Follow the [instructions here](https://golang.org/doc/install) to setup your environment.
3. [Tests your installation](https://golang.org/doc/install#testing) and verify everything works.

## Installing kbcmd
```bash
go get -u src/github.com/killbill/kbcli/kbcmd
go install src/github.com/killbill/kbcli/kbcmd

# For bash completion
source $GOPATH/src/github.com/killbill/kbcli/kbcmd/bash-autocomplete
```

## Running kbcmd
```bash
kbcmd -h
# (or)
kbcmd <command> -h
```

## Walkthrough: Create subscription and invoices
The following walkthrough will walk you through the steps to create new account and subscription
and then generate invoice for it.

[kbcmd walkthrough](../docs/kbcmd/kbcmd-walkthrough.md)