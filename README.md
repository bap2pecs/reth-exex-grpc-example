## Context

This is a PoC to build Reth ExEx that can depend on:
- ANY external systems
- writing in ANY language

# How to run the demo

```bash
make proto-gen
make start-grpc
make build-docker
make demo-run
make demo-stop
```

to restart the demo, run `make demo-restart`

# How to test

```bash
make logs
make ls
# replace <RPC_URL> with the reth RPC URL seen in previous step
make query-block rpc=<RPC_URL>
```