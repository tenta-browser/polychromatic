Polychromatic
=============

[![Go Report Card](https://goreportcard.com/badge/github.com/tenta-browser/polychromatic)](https://goreportcard.com/report/github.com/tenta-browser/polychromatic)
[![GoDoc](https://godoc.org/github.com/tenta-browser/polychromatic?status.svg)](https://godoc.org/github.com/tenta-browser/polychromatic)

Polychromatic provides an opinionated wrapper around [logrus](https://github.com/sirupsen/logrus) including the [prefixed](https://github.com/x-cray/logrus-prefixed-formatter)
formatter and [colored](https://github.com/mattn/go-colorable) terminal output. It also includes an `EventualLogger` which may be used when
the desire to log a particular set of entries is not known immediately.

Contact: developer@tenta.io

Installation
------------

`go get github.com/tenta-browser/polychromatic`

Usage
-----

In general, just call `GetLogger("prefix")` wherever you want a prefixed logger. The function `SetLogLevel(level)` sets
the _global_ log level for all loggers which have been issued from polychromatic.

For eventual logging, call `NewEventualLogger()` to get an eventual logger which has `Queue` and `Queuef` (Printf style)
methods to enqueue pending log entries. If you decide you want those log items, then call `Flush()`. If you don't want
them, then drop it like it's hot.

Performance
-----------

The main log component just provides some convenience functions and setup wrappers around logrus, so performance wise
it's identical. The EventualLogger does take locks to deal with threading issues, and is primarily designed for developemnt
(which is a nice way of saying we haven't benchmarked it).

License
-------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

For any questions, please contact developer@tenta.io

Contributing
------------

We welcome contributions, feedback and plain old complaining. Feel free to open
an issue or shoot us a message to developer@tenta.io. If you'd like to contribute,
please open a pull request and send us an email to sign a contributor agreement.

About Tenta
-----------

This logging library is brought to you by Team Tenta. Tenta is your [private, encrypted browser](https://tenta.com) that protects your data instead of selling. We're building a next-generation browser that combines all the privacy tools you need, including built-in OpenVPN. Everything is encrypted by default. That means your bookmarks, saved tabs, web history, web traffic, downloaded files, IP address and DNS. A truly incognito browser that's fast and easy.
