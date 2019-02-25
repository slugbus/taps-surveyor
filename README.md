# taps-surveyor
Survery the UCSC TAPS API via the command line

[![Go Report Card](https://goreportcard.com/badge/github.com/slugbus/taps-surveyor)](https://goreportcard.com/report/github.com/slugbus/taps-surveyor)


## Installation
```shell
$ go get -u github.com/slugbus/taps-surveyor
```
## Documentation

Usage:
  ```
 $ taps-surveyor [flags]
  ```

Flags:

  `-d`, `--duration` (duration)   how long to ping the TAPS server for (default 30s)

  `-h`, `--help`                help for taps-surveyor

  `-i`, `--interval` (duration)   how often to ping the TAPS server (default 3s)

  `-n`, `--number` (uint)         how many times to ping the TAPS server. If set this flag takes precedence over the duration flag (default 10)

  `-s`, `--server` (string)       specify a custom server to ping (default "http://bts.ucsc.edu:8081/location/get")