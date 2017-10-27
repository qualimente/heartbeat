# Overview #

Heartbeat is a simple program to produce a periodic heartbeat message.

[![Build Status](https://travis-ci.org/qualimente/heartbeat.svg?branch=master)](https://travis-ci.org/qualimente/heartbeat)

# Developing #

1. Clean up build env `make clean`
2. Build the software `make build test`
3. Run the software `bin/darwin/amd64/heartbeat run`
```
Running heartbeat periodically
{"appName":"heartbeat","level":"info","msg":"Running periodically","time":"2017-10-26T20:02:14-07:00"}
{"level":"info","msg":"Every heartbeat bears your name","time":"2017-10-26T20:02:14-07:00","type":"heartbeat"}
{"level":"info","msg":"Every heartbeat bears your name","time":"2017-10-26T20:02:15-07:00","type":"heartbeat"}
{"level":"info","msg":"Every heartbeat bears your name","time":"2017-10-26T20:02:16-07:00","type":"heartbeat"}
```

# Building Releases #

1. Clean up build env `make clean`
2. Build the software `make all`

Releases will be created in the `releases` directory

# Running #

`heartbeat` prints heartbeat messages to standard out with one message printed in json format every second, by default.

options can be printed with `--help`, `heartbeat run`'s are most interesting permitting control over format (text, json) and period:

```
NAME:
   heartbeat run - Run the program

USAGE:
   heartbeat run [command options] [arguments...]

OPTIONS:
   --format value  Specify the output format: text, json
   --period value  Specify the period between heartbeat messages. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' (default: 1s)
```

# Homage #

Yes, the default message is a reference to Amy Grant's classic [Every Heartbeat](https://www.youtube.com/watch?v=kcYIZ-cduvM)
