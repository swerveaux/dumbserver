This is just a quick and dirty webserver that logs what path was requested and the headers passed in. It is very much nothing special.

You can just build it normally or run it in a docker container if you don't feel like installing go.

```
docker build -t dumbserver .
docker run --rm -p 8080:8080 dumbserver
```