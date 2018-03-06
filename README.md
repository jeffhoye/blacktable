# blacktable

## run
```
blacktable examples/pingself.csv
exit
```

## installation
```
go install github.com/jeffhoye/blacktable/blacktable
```

## Limitations: 
* UDP only
* The receive period doesn't do anything, but it does set up a udp listen socket immediately.
* The send system works with UDP.  So you can schedule a beacon now.
* You can either pipe it the csv or else specify one or more files as arguments. 
* You can comment out commands in the csv with "#"
* Column 1 is the task (send or receive)
* Column 2 is the name of the task, every task must have a unique name (this will come in handy down the road)
* Column 3 is how many seconds until it should start, 
* Column 4 is the period 
* Column 5 is how many times to execute, 0 means that it should only be run in response to something (not yet implemented), negative means run forever, positive is obvious
* Column 6 is the protocol (currently only udp is supported)
* Column 7 is the ip/port
* Column 8 is the message

There are more magic for the receive task that is not yet implemented, but it will like use a regexp on the from/message to then execute a different task, which is very powerful, but not yet needed.
