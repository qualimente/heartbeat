# Overview #

Heartbeat is a simple program to produce a periodic heartbeat message.

# Developing #

1. Clean up build env `make clean`
2. Build the software `make build test`
3. Run the software `bin/darwin/amd64/heartbeat run`
```
Running heartbeat periodically
2017-08-08T04:26:44Z Every heartbeat bears your name
2017-08-08T04:26:45Z Every heartbeat bears your name
2017-08-08T04:26:46Z Every heartbeat bears your name
```

# Building Releases #

1. Clean up build env `make clean`
2. Build the software `make all`

Releases will be created in the `releases` directory

# Homage #

Yes, the default message is a reference to Amy Grant's classic [Every Heartbeat](https://www.youtube.com/watch?v=kcYIZ-cduvM)
