# Wake On Lan Service

This application is implemented with golang to wake up the computers.

## Service Configration

By default, http server listens `:8080` address.
To change or bind different address, set *`LISTEN`* environment variable.

## Container

```bash
docker run -it --rm --network host ghcr.io/ahmetozer/wakeonlan
```

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
        "Device": "lo",
        "IPAddr": [
            "127.0.0.1/8",
            "::1/128"
        ]
    }
]
```

### ArpEntries

List the MAC addreses with details.

- Method: *`GET`*
- Endpoint: `/api/arpentries`

```json
[
    {
        "IPAddr": "172.23.80.1",
        "HWType": "0x1",
        "Flags": "0x2",
        "HWAddr": "00:15:5d:f5:bd:67",
        "Mask": "*",
        "Device": "eth0"
    }
]
```

### WakeOnLan

To wake up the computers.

- Method: *`POST`*
- Endpoint: `/api/wakeonlan`

Example request body

```json
[
    {
        "mac": "00:15:5d:f5:bd:67",
        "if": "lo",
        "addr": "255.255.255.255",
        "port": "7"
    },
    {
        "mac": "00:15:5d:f5:bd:67",
        "if": "eth0",
        "addr": "fe80::900d:cafe:900d:c0de",
        "port": "7"
    },
    {
        "mac": "00:15:5d:f5:bd:67",
        "if": "eth0",
        "addr": "2001:db8:900d:cafe:900d:c0de:900d:11fe",
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
    packet := wakeonlan.MagicPacket{HWAddr: mac, Device: "eth0", IPAddr: "255.255.255.255", Port: "7"}
    err := packet.SendMagicPacket()
    if err != nil {
    fmt.Printf("%v", err)
    }
}
```
