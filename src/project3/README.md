This program manage a dictionary database to manage words and affected definitions

It allow you to specify following parameters with CLI:

- list
- add
- define
- remove

To run it, type the following command:

Install package:
````
go get github.com/dgraph-io/badger/...
````

Compile sources:
````
go build -o dict
````

Run the program, (example):
````
go run .\main.go -action list
````

The output for the following file with be:

````
badger 2022/02/01 23:50:30 INFO: All 2 tables opened in 4ms
badger 2022/02/01 23:50:30 INFO: Discard stats nextEmptySlot: 0
badger 2022/02/01 23:50:30 INFO: Set nextTxnTs to 2
badger 2022/02/01 23:50:30 INFO: Deleting empty file: ./badger\000002.vlog
Dictionary Content
Golang          Is a great langage!                               Feb  1 23:49:42
Python          Is an easy langage!                               Feb  1 23:50:17
badger 2022/02/01 23:50:30 INFO: Lifetime L0 stalled for: 0s
badger 2022/02/01 23:50:30 INFO:
````
