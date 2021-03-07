# How to Use

### Server

> For running the grpc server you can run the below command in the main folder:

```bash
make run-stand-alone
```

### Client

> Grpc client supports some arguments, you can see how to use the client with the command bellow :

```bash
make run-client ARGS="-h"
```

> Make sure before running the clinet that the server is running  
> you can pass your desired arguments to the command after wards for example:

```bash
make run-client ARGS="-rent 1700 -rent-inc-rate 2.5 -down-payment -2000 -intrest 2.91 -term 20 -price 250000 -p-tax 2 -t-tax 0 -res 11"
```

> And to see an example you can run:

```bash
make run-client-example
```

### Tests

> For running the tests you can run the below command in the main folder:

```bash
make test
```

> _this command supports ARGS aswell_

# Features

This repository demonstrates a lot of the features that are needed in an enterprise and service oriented ready enviroment

#### For example:

- Containarization
- Advanced error handling and error reporting
- Service oriented architecture
- Compositon
- Advanced config support
- Efficent logging capablities
- etc
