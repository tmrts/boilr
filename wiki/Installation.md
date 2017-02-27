# Install Script
Download the latest version
of [install](https://raw.githubusercontent.com/tmrts/boilr/master/install)
script, which is also included in
every [release](https://github.com/tmrts/boilr/releases), and run it to install
the `boilr` binary. The `boilr` binary will be installed to `~/bin/boilr`.

# Binary Release
You can find the latest binary
releases [here](https://github.com/tmrts/boilr/releases). Grab the one the suits
your architecture and operating system and start using it.

# Building from Source
Make sure you have setup a Go >= 1.7 development environment and the `GOPATH`
environment variable is configured
(see [official docs](https://golang.org/doc/code.html#GOPATH) for instructions)
and your `PATH` includes `$GOPATH/bin`.

Then use the following command
```bash
go get github.com/tmrts/boilr
```

The binary will be installed into `$GOPATH/bin`.
