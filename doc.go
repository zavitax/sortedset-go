/*
Package sortedset was inspired by SortedSet in Redis: it provides fast access to elements in sorted set by key or score(order).

Introduction

Every node in the set is associated with these properties.

    type SortedSetNode[K constraints.Ordered, SCORE Number, V interface{}] struct {
        key      K       // unique key of this node
        Value    V       // associated data
        score    SCORE   // int64 score to determine the order of this node in the set
    }

Each node in the set is associated with a key. While keys are unique, scores may repeat. Nodes are kept ordered by score and key. Each node in the set can be also accessed by rank: the position of the node in the sorted set.

With sorted sets you can add, remove or update nodes in a very efficient way (O(N) = log(N)). You can also retrieve ranges of nodes by score or by rank in a very efficient way. Accessing the middle of a sorted set is also very efficient, so you can use Sorted Sets as smart lists of non repeating nodes, where you can quickly access everything you need: nodes in order, fast existence test, fast access to nodes in the middle.

Use Case

A typical use case of sorted set is a leaderboard in a massive online game, where you update the player's position every time a new score is submitted. You can easily take the top players using GetRangeByRank() method. You can also, given an user id, return their rank in the listing using FindRank() method. Using FindRank() and GetRangeByRank() together you can show users with a score similar to a given user.

Examples

    // create a new set
    set := sortedset.New[string, int64, string]()

    // fill in new node
    set.AddOrUpdate("a", 89, "Kelly")
    set.AddOrUpdate("b", 100, "Staley")
    set.AddOrUpdate("c", 100, "Jordon")
    set.AddOrUpdate("d", -321, "Park")
    set.AddOrUpdate("e", 101, "Albert")
    set.AddOrUpdate("f", 99, "Lyman")
    set.AddOrUpdate("g", 99, "Singleton")
    set.AddOrUpdate("h", 70, "Audrey")

    // update an existing node by key
    set.AddOrUpdate("e", 99, "ntrnrt")

    // get the node by key
    set.GetByKey("b")

    // remove node by key
    set.Remove("b")

    // get the number of nodes in this set
    set.GetCount()

    // find the rank(postion) in the set.
    set.FindRank("d") // return 1 here

    // get and remove the node with minimum score
    set.PopMin()

    // get the node with maximum score
    set.PeekMax()

    // get the node at rank 1 (the node with minimum score)
    set.GetByRank(1, false)

    // get & remove the node at rank -1 (the node with maximum score)
    set.GetByRank(-1, true)

    // get the node with the 2nd highest maximum score
    set.GetByRank(-2, false)

    // get nodes with in rank range [1, -1],  that is all nodes actually
    set.GetRangeByRank(1, -1, false)

    // get & remove the 2nd/3rd nodes in reserve order
    set.GetRangeByRank(-2, -3, true)

    // get the nodes whose score are within the interval [60,100]
    set.GetRangeByScore(60, 100, nil)

    // get the nodes whose score are within the interval (60,100]
    set.GetRangeByScore(60, 100, &GetRangeByScoreOptions{
        ExcludeStart: true,
    })

    // get the nodes whose score are within the interval [60,100)
    set.GetRangeByScore(60, 100, &GetRangeByScoreOptions{
        ExcludeEnd: true,
    })

    // get the nodes whose score are within the interval [60,100] in reverse order
    set.GetRangeByScore(100, 60, nil)

    // get the top 2 nodes with lowest scores within the interval [60,100]
    set.GetRangeByScore(60, 100, &GetRangeByScoreOptions{
        Limit: 2,
    })

    // get the top 2 nodes with highest scores within the interval [60,100]
    set.GetRangeByScore(100, 60, &GetRangeByScoreOptions{
        Limit: 2,
    })

    // get the top 2 nodes with highest scores within the interval (60,100)
    set.GetRangeByScore(100, 60, &GetRangeByScoreOptions{
        Limit: 2,
        ExcludeStart: true,
        ExcludeEnd: true,
    })
*/
package sortedset
