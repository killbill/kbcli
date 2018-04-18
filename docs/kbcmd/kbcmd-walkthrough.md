# kbcmd walk through: Creating new subscription and generating invoice

## Prerequisites
This walkthrough assumes that killbill is running
in localhost:8080 with username/password set to 'username' and 'password'. If your
killbill installation has different values, you can set the values through exporting
the appropriate KB_* variables. (For details run `kbcmd -h`)

## Step 1: Create new tenant
```sh
# Set the api key and secret so future calls will use this.
# These exports ensure kbcmd uses them for future calls.
# Note: You can also pass these in explicitly to each command through flags.
export KB_API_KEY=tenant1key
export KB_API_SECRET=tenant1secret

kbcmd tenants create tenant1 $KB_API_KEY $KB_API_SECRET
```
This command will print the output when it runs

```
EXTERNAL_KEY API_KEY    API_SECRET
tenant2      tenant2key tenant2secret
```

This is the default output of kbcmd. To change formatting option, you can specify -f option.

## Step 2: Upload catalog
Upload a sample catalog to the created tenant.
```sh
kbcmd catalog upload docs/samples/simple-catalog.xml
```

Now the catalog has been uploaded.

## Step 3: Create new Account
Let's create a new killbill account.

```sh
kbcmd acc create Name="John Doe" ExternalKey=johndoe Email=johndoe@gmail.com Company="Stark" Currency=USD
```

This command will print the created account
```sh
NAME     EXTERNAL_KEY ACCOUNT_ID                           EMAIL             BALANCE CURRENCY
John Doe johndoe      e4f47a6b-9975-4922-9c5a-baae7ef4d03c johndoe@gmail.com <nil>   USD
```

If you want to get help for specific command, just specify `-h` option. For ex.,
`kbcmd acc create -h` will print help for create command.

## Step 4: Create new subscription
```sh
kbcmd subscriptions create ExternalKey=bundle1 Account=johndoe ProductName=simple PlanName=simple-monthly PriceList=default
```
This will create a new bundle and print the output. 
```
EXTERNAL_KEY BUNDLE_ID
bundle1      447ac056-182f-40a0-a672-8f7a4daa3851
```

## Step 5: Generate invoices
```sh
# Set the date one month in advance to current date
DATE=2018-05-01
kbcmd invoices dry-run johndoe $DATE
```

## Invoking other APIs

kbcmd -h will print all the supported commands. You can also specify `-h` for a specific command to usage.
For ex., to print usage for `accounts create` command,

```sh
kbcmd accounts create -h
```