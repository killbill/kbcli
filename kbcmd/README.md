# kbcmd - Kill bill command line

kbcmd is a command line client for killbill. kbcmd uses go client library
to talk to killbill. 

## Installing
```bash
cd $GOPATH
go get -u src/github.com/killbill/kbcli/kbcmd
go install src/github.com/killbill/kbcli/kbcmd

# For bash completion
source src/github.com/killbill/kbcli/kbcmd/bash-autocomplete

# If you haven't added Go bin path to PATH
export PATH=$PWD/bin:$PATH
```


### Running the command
```bash
kbcmd -h
# (or)
kbcmd <command> -h
```
