# Wake On Lan Service

This application is implemented with golang to wake up the computers.

## Service Configration

By default, http server listens `:8080` address.
To change or bind different address, set *`LISTEN`* environment variable.

## Available API endpoints

You wake up the computers via web ui or if you want you can
use api endpoints.

### Interfaces

List the interfaces with IP addresses.

- Method: *`GET`*
- Endpoint: `/api/interfaces`

```json
[
    {
        "Ifname": "lo",
        "Ifaddrs": [
            "127.0.0.1/8",
            "::1/128"
        ]
    },
    {
        "Ifname": "eth0",
        "Ifaddrs": [
            "172.23.91.231/20",
            "fe80::215:5dff:fef6:b545/64"
        ]
    }
]
```

### Arptables

List the MAC addreses with details.

- Method: *`GET`*
- Endpoint: `/api/arptable`

```json
[
    {
        "IPAddress": "172.23.80.1",
        "HWType": "0x1",
        "Flags": "0x2",
        "HWAddress": "00:15:5d:f5:bd:67",
        "Mask": "*",
        "Device": "eth0"
    }
]
```

### WakeOnLan

To wake up the computers.

- Method: POST
- Endpoint: /api/wakeonlan

Example request body

```json
[
    {
        "mac": "00:15:5d:f5:bd:67",
        "if": "lo",
        "addr": "255.255.255.255",
        "port": "7"
    }
]

```

Server respond

```json
[
    {
        "RequestNo": 1,
        "Status": "packet send"
    }
]
```

## Usage of package

You can use wake on lan as golang function

```go
package main

import (
    wakeonlan "github.com/ahmetozer/wakeonlan/share"
)

func main() {
    mac, _ := net.ParseMAC("00:15:5d:f5:bd:67")
    err := wakeonlan.MagicPacket{MAC: mac, IF: "eth0", ADDR: "255.255.255.255", PORT: "7"}.SendMagicPacket()
    fmt.Printf("%v",err)
}
```
