# ISO20022 Client

## Testing
For testing, we need to open three command lines:

1. To run the sending party:
```bash
./bin/iso20022-client start --server-addr=':2843' --key-name=iso20022-client --keyring-backend=test --cache-path=/tmp/iso20022-client-1
```

2. To run the receiving party:
```bash
./bin/iso20022-client start --server-addr=':2844' --key-name=iso20022-client-2 --keyring-backend=test --cache-path=/tmp/iso20022-client-2
```

3. To send the message:
```bash
./bin/iso20022-client message send /path/to/pacs008-request.xml --server-addr=':2843'
```

And then receive it:
```bash
./bin/iso20022-client message receive /path/to/pacs008-response.xml --server-addr=':2844'
```
