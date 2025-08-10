# Temperature Sensor

A temperature sensor simulation using raw values from a text file representing values from a 12-bit ADC with reference voltage of 3.3V.

## Usage
You can run pre-build binaries from the bin folder, or build your own either using Go syntax or the provided Makefile.

Build to your own host platform:
```
go build .
```

Or

```
make build
```

Build for 64 bit Linux:
```
make build-linux
```

Build for 64 bit Mac:
```
make build-mac
```

Build for 64 bit Windows
```
make build-windows
```