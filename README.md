# Device

Device is an emulator for a fictional I/O control device.


## Manual

See [manual.md](manual.md).


## Install

To install:

```
GOSUMDB=off go install github.com/thorntonrose/device@latest
```


## Build

Prequisites:

  - Go 1.25+
  - Mage 1.15+

To build:

```
mage build
```

Output is `./device`.


## Test

To run tests:

```
[PACKAGE=<path>] [TESTS=<pattern>] mage test
```

\<path> = a Go package path, e.g. `./internal/parser`

\<pattern> = a test name pattern, e.g. `TestParse`, `^TestParse$`
