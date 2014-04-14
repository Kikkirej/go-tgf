**[Trivial Graph Format](http://en.wikipedia.org/wiki/Trivial_Graph_Format)** is a simple text-based file format for describing graphs. It consists of a list of node definitions, which map node IDs to labels, followed by a list of edges, which specify node pairs and an optional edge label. Node IDs can be arbitrary identifiers, whereas labels for both nodes and edges are plain strings.

This library parses .tgf file and creates Node and Edge structs.

### Example
```go
package main

import (
  "fmt"
  "github.com/didip/gotgf"
  "github.com/didip/gotgf/ast"
)

func main() {
  nodes, _, err := tgf.ParseFile("example.tgf")
  if err != nil { panic(err) }

  // Print array of root nodes
  fmt.Println("Root Nodes: ", ast.RootNodes(), "\n")

  for index, node := range nodes {
    // Print node.Id
    fmt.Println("Node ID:", node.Id)

    // Print node.Label
    fmt.Println("Node Label:", node.Label)

    // Print array of edges that go out from a particular node.
    if index == "1" || index == "3" || index == "4" {
      fmt.Println("Node OutboundEdges", node.OutboundEdges(), "\n")
    }

    // Print array of edges that come into a particular node.
    if index == "2" || index == "4" || index == "5" {
      fmt.Println("Node InboundEdges:", node.InboundEdges(), "\n")
    }
  }
}
```