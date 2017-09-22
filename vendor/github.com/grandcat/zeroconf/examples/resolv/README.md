Browse and Resolve
==================
Compile:
```bash
go build -v
```

Browse for available services in your local network:
```bash
./resolv
```
By default, it shows all working stations in your network running
a mDNS service like Avahi.
The output should look similar to this one:
```
2016/12/04 00:40:23 &{{stefanserver _workstation._tcp local.   } stefan.local. 50051 [] 120 [192.168.42.42] [fd00::86a6:c8ff:fe62:4242]}
2016/12/04 00:40:23 stefanserver
2016/12/04 00:40:28 No more entries.
```
The `-wait` parameter enables to wait for a specific time until
it stops listening for new services.

For a list of all possible options, just have a look at:
```bash
./resolv --help
```