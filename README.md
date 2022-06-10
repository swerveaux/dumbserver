This is just a quick and dirty webserver that logs what path was requested and the headers passed in. It is very much nothing special.

You can just build it normally or run it in a docker container if you don't feel like installing go.

```
docker build -t dumbserver .
docker run --rm -p 8080:8080 dumbserver
```

Options are:

* --port 1234 - Change the port the server listens on. By default, it's 8080. Or, if you run it in a docker container, just change the port mapping when you run the container.
* --log-to-stderr - By default Go logs to STDERR, but since this is mostly meant to be fed into something like jq, I changed it to use STDOUT as default. Use this to set it back.
* --show-body - Show the request body in addition to headers and path. This can get super noisy, so it's off by default.