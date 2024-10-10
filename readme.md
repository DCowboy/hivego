# HiveGo - A client for the Hive blockchain

At this time, there are only a few functions from the client. More will be added.

*Functions added by the DCowboy fork added below

### Example usage:
create a client:
```
hrpc := hivego.NewHiveRpc("https://api.myHiveBlockchainNode.com")
```

submit a custom json tx:
```
txid, err := hrpc.BroadcastJson([]string{submittingAccount}, []string{}, id, string(jsonPayload), &activeWif)
```

vote a post:
```
txid, err := hrpc.VotePost(voter, author, permlink, weight, &wif)
```

get n blocks starting from block x as the raw response from the rpc (in bytes):
```
responseBytes, err := hrpc.GetBlockRangeFast(startBlock int, count int)
```

### Added by DCowboy:
Validate account:
```
acct, err := hrpc.CheckAccount(account)
// will return first account name that matches (expected exact match)
```

Validate key:
```
valid, err := hrpc.CheckKey(&wifVar, "alice")
//returns false and error for any reason the check failed. Returns true if valid
```

WARNING: It is not recommended to stream blocks from public APIs. They are provided as a service to users and saturating them with block requests may (rightfully) result in your IP getting banned
