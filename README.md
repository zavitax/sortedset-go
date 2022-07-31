# Sorted Set in Golang with Generics

Sorted Set is a data-struct inspired by the one from Redis. It allows fast access by key or score.

| Property | Type | Description |
|---|---|---|
| `key` | constraints.Ordered | The identifier of the node. It must be unique within the set. |
| `value` | any | value associated with this node |
| `score` | constraints.Ordered | score determines the position of the item within the sorted set. |

Each node in the set is associated with a `key`. While `key`s are unique, `score`s may be repeated. 
Nodes are __taken in order instead of ordered afterwards__, from low score to high score. If scores are the same, the node is ordered by its key in lexicographic order. Each node in the set is associated with __rank__, which represents the position of the node in the sorted set. The __rank__ is 1-based, i.e. rank 1 is the item with the lowest score.

Sorted Set is implemented based on skip list and hash map internally. With sorted sets you can add, remove, or update nodes very efficiently with log(N) complexity. You can also get ranges by score or by rank (position) in a very fast way. Accessing the middle of a sorted set is also very fast, so you can use Sorted Sets as a smart list of non repeating nodes where you have quick access to everything you need: nodes in order, existence test, access to nodes in the middle.

A typical use case of sorted set is a leaderboard in a massive online game, where every time a new score is submitted you update it using `AddOrUpdate()`. You can easily take the top users using `GetRangeByRank()`, you can also, given an user id, return its rank in the list using `FindRank()`. Using `FindRank()` and `GetRangeByRank()` together you can show users with a score similar to a given user.

## Documentation

[https://godoc.org/github.com/zavitax/sortedset-go](https://godoc.org/github.com/zavitax/sortedset-go)
