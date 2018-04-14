# Kill Bill go client library and kill bill command line
This repository contains killbill go client library (kbclient)
and killbill command line tool (kbcmd)

## Kill bill go client library
Kill bill go client library is a go package that can be used to connect to
kill bill.

### Install
```bash
go get -u src/github.com/killbill/kbcli/kbclient
```

### Creating new client
```go
    trp := httptransport.New("127.0.0.1:8080", "", nil)
    // Add text/xml producer which is not handled by openapi runtime.
    trp.Producers["text/xml"] = runtime.TextProducer()
    // Set this to true to dump http messages
    trp.Debug = false
    authWriter := httptransport.BasicAuth("admin"/*username*/, "password" /**password*/)
    client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})

```

Look at the [complete example here](examples/listaccounts/main.go).
For more examples, look at [kbcmd tool](kbcmd/README.md).

## Kill bill command line tool (kbcmd)
kbcmd is a command line tool that uses the go client library. This tool can do many of the
kill bill operations. More details are [available here in README](kbcmd/README.md).