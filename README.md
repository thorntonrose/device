# Device

Device is an emulator for a fictional I/O control device.


## Reference Manual

See [docs](docs/index.md).


## Build

Prequisites:

* Go 1.26+
* Mage 1.15+

To build:

```
mage build
```

Binaries will be created in `./bin/device`.


## Test

To run tests:

```
[PACKAGE=<path>] [TESTS=<pattern>] mage test
```

\<path> = a Go package path, e.g. `./internal/parser`
\<pattern> = a test name pattern, e.g. `TestParse`, `^TestParse$`
