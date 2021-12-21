**Prefer JSON RPC when multiple client technologies need to connect to your RPC server**

The json client can be anything, even curl
```
curl -X POST \
    http://localhost:1234/rpc \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{
        "method": "JSONServer.GiveBookDetail",
        "params": [{
            "Id":"1234"
        }],
        "id": 1
    }'
```

Response is
```
{"result":{"Id":"1234","Name":"In the sunburned country","Author":"Bill Bryson"},"error":null,"id":1}
```