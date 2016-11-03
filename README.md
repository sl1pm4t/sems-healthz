# sems-healthz

A simple application that monitors the health of SEMS SBC node and can provide that info via an HTTP endpoint.

## Usage

* Launch sems-healthz

```bash
$ ./sems-healthz
2016/11/03 11:48:28 Starting server...
2016/11/03 11:48:28 Health service listening on 0.0.0.0:4480
```

Configure Google health check to poll protocol = `HTTP`, port = `4480`, path = `/healthz`

## Why

Google Compute health checks can only poll an HTTP endpoint to verify the health of node.
Upon receiving a `/heathz` request, this app attempts to query the sems-stats RPC interface, and if it gets a positive response will return 200ok to the HTTP request.