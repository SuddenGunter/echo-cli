# echo-cli

Echo-CLI app created as an experiment where main goal was get some skills with Cobra. It allows to save auth token to temp dir, and it it's presented - connect and communicate with websocket server (see echo-server submodule).

Usage:
  `echo-cli [command]`

Available Commands:
  ```
  auth        Authorize local user to echo-server
  help        Help about any command
  user        Manage users
```

# build

Use `make build` to build app

# test

Use `make test` to run tests

# demo

Use `make demo` to run app in demo mode. You need working server on ws://localhost:8080/ws to see demo. Server included as submodule in same repo

