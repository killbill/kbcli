# Kill Bill go client library and kill bill command line
This repository contains killbill go client library (kbclient)
and killbill command line tool (kbcmd)

## Versions

| KB Version | KBCli Version |
|------------|---------------|
| 0.20.x     | 1.x.y         |
| 0.22.x     | 2.x.y         |
| 0.24.x     | 3.x.y         |



## Kill bill go client library
Kill bill go client library is a go package that can be used to connect to Kill Bill server.

### Install
```bash
go get -u github.com/killbill/kbcli/v3
```

### Creating new client
```go
    trp := httptransport.New("127.0.0.1:8080", "", nil)
    // Add text/xml producer which is not handled by openapi runtime.
    trp.Producers["text/xml"] = runtime.TextProducer()
    // Set this to true to dump http messages
    trp.Debug = false
    // Authentication
    authWriter := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
        encoded := base64.StdEncoding.EncodeToString([]byte("admin"/*username*/ + ":" + "password" /**password*/))
        if err := r.SetHeaderParam("Authorization", "Basic "+encoded); err != nil {
            return err
        }
        if err := r.SetHeaderParam("X-KillBill-ApiKey", apiKey); err != nil {
            return err
        }
        if err := r.SetHeaderParam("X-KillBill-ApiSecret", apiSecret); err != nil {
            return err
        }
        return nil
    })
    client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})
```

Look at the [complete example here](examples/listaccounts/main.go).
For more examples, look at [kbcmd tool](kbcmd/README.md).

### Wrapper client

We also provide a client wrapper (a higher level and easier to use api) that is built (not generated) on top of the generated code.
See the package `killbill` package under the `wrapper` directory.

There is a suite of test `client_test.go` that shows how to use it.

### Client code generation

We use a forked version of the `go-swagger` client hosted under https://github.com/killbill/go-swagger.
Every so often we rebase our fork from upstream to keep it up-to-date. Given a version of a `swagger` binary, the
client can be regenarated using:

`swagger  generate client -f kbswagger.yaml -m kbmodel -c kbclient --default-scheme=http`


### Generating dev extensions
We also have dev extension APIs (like clock etc), that are in swagger-dev.json. To generate,
run the following.

```bash
# Regenerate the tool
swagger generate client -f swagger-dev.json -m kbmodel -c kbclient --default-scheme=http

# Delete the client file.
rm kbclient/kill_bill_dev_client.go
```

## Kill bill command line tool (kbcmd)
kbcmd is a command line tool that uses the go client library. This tool can do many of the
kill bill operations. More details are [available here in README](kbcmd/README.md).