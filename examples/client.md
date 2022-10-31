# Creating a Client

When using this SDK you're going to need to set up a client and establish a connection to the remote AWX/Ansible Automation
Platform instance.

There are some parameters you can provide:

* Your AWX/Tower DNS hostname or IP address (remembering to include the `http://` too)
* Basic Auth:
    * The username you wish to authenticate as
    * The password you wish to authenticate with
* Token Auth:
    * The token you wish to authenticate with
* And an optional `*http.Client` you can use to custom how the SDK communicates with your AWX/Tower instance(s)

Throughout the rest of these example documents the above `client` variable will be referred to as a correctly
configured client to an operational AWX/Tower instance.

## Basic method
```go
package main

import (
    "log"
    awx "github.com/denouche/goawx/client"
)

func main() {
    client, err := awx.AWXClient("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", "", nil)
    // ...
}
```

## Token method

```go
package main

import (
    awx "github.com/denouche/goawx/client"
    "log"
)

func main() {
    client, err := awx.AWXClient("http://awx.your_server_host.com", "", "", "my-token", nil)
    // ...
}
```
