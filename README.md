# go-fdkaac

fork from lib-fdkaac(https://github.com/izern/fdk-aac) (https://github.com/winlinvip/go-fdkaac)

## Usage

First, get the source code:

```
go get -d github.com/seastart/go-fdkaac
```

Then, compile the fdk-aac:

```
wget https://github.com/seastart/go-fdkaac/raw/refs/heads/master/fdk-aac-2.0.0.tar.gz
tar -zxvf fdk-aac-2.0.0.tar.gz && cd fdk-aac-2.0.0/ && ./configure --prefix=/usr/local/lib/fdk-aac-2.0.0 && make && make install
```

Done, import and use the package:

* [ExampleAacDecoder_RAW](fdkaac/example_test.go#L29), decode the aac frame to PCM samples.
* [ExampleAacEncoder_LC](fdkaac/example_test.go#L316), encode the PCM samples to aac frame.

There are an example of AAC audio packets in ADTS:

* [avatar aac over ADTS](doc/adts_data.go), user can use this file to decode to PCM then encode.

To run all examples:

```
go test ./...
```

Winlin 2016
