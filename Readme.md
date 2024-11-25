### Pre requists
* librocksdb
* libsnappy
* libz
* liblz4
* libzstd
* libbz2 (optional)

### Commands
* GET <key>  get value of <key>
* SET <key> <value> set <key> <value>
* PUT <key> <value> set <key> <value>. same with set
* ITER <count> [starting-key] return list of key and values as array by starting from [starting-key] and [count] items
