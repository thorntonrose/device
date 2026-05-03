# Device

...


## Build

Prequisites:

* Go 1.26+
* Mage 1.15+

To build:

```
mage build
```

To run tests:

```
[PACKAGE=<path>] [TESTS=<pattern>] mage test
```

\<path> = a Go package path, e.g. `./internal/parser`
\<pattern> = a test name pattern, e.g. `TestParse`, `^TestParse$`


## Run

To run:

```
./bin/device [<flags>] <file>
```
