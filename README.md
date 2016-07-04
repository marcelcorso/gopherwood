# gopherwood

![Noahs Ark](https://upload.wikimedia.org/wikipedia/commons/2/23/Noahs_Ark.jpg)

Go maps are not thread safe. 

You can implement maps with binary trees.

Binary trees are pretty. 

Thread safe trees are hard to implement effeciently (because you need to write-lock sections of the tree and other tricks).

...

_Do not communicate by sharing memory; instead, share memory by communicating._

**What if every tree node was handled by a goroutine?**

**Bang!** Gopherwood. (some trees are made of wood)

## Dialogue

_You_: OMG. That doesn't scale.

_Me_: very probably. But goroutines are cheap so this thing should behave ok for small Ns.

_You_: is it fast? 

_Me_: don't know yet. but I hope so because there are no locks.. but maybe a large number of goroutines running may slow things.  



## Cool tricks

### 1 

Its common on textbook binary tree implemantations to define a node like 

```go
type Node struct {
  key string
  left *Node
  right *Node
}
```

Gopherwood does nothing of this!

Data is stored on local variables ["inside" the goroutine](https://play.golang.org/p/uI5pyqPecK).  

By data I mean the "key" and channels. Yep. No pointers to other nodes. The goroutine saves 4 channels: left**To**, left**Fro**, right**To**, right**Fro**. 

These are used to send messages to and fro. 

### 2

When doing a change to a node like adding a key code normally looks like

```go
func (v *Tree) Add(n *Node, key string) {
  if n.key > key {
    if n.right == nil {
      n.right = &Node{key: key}
    } else {
      v.Add(n.right, key)
    }
  } else {
    if n.left == nil {
      n.left = &Node{key: key}
    } else {
      v.Add(n.left, key)
    }
  }
}
```

But that's not so DRY. Pointers to the rescue. Pointers to channels on gopherwood:

```go
        sideTo := &rightTo
        sideFro := &rightFro
        if key > newKey {
          sideTo = &leftTo
          sideFro = &leftFro
        } else if key == newKey {
          continue
        }

        if *sideTo == nil {
          *sideTo, *sideFro = createNode(newKey)
        } else {
          *sideTo <- "add"
          *sideTo <- newKey
        }
```

## Benchmarks

Comparing Gopherwood with a textbook implementation of binary tree. Adding N random keys with one goroutine. 

```
go test -bench=.
PASS
BenchmarkGopherwoodAdd-8       10000        395670 ns/op
BenchmarkGotreeAdd-8           10000        344732 ns/op
ok      github.com/marcelcorso/gopherwood   8.073s
```

## TODO 

 * rm nodes
 * rebalance the tree
 * informal benchmark 


