package ast

import (
  "code.google.com/p/go-uuid/uuid"
)

var (
  Nodes = make(map[string]Node)
  Edges = make(map[string]Edge)
)

func NewNode(id string, label string) Node {
  node := Node{id, label, make([]string, 0), make([]string, 0)}
  Nodes[id] = node
  return node
}

func NewEdge(label string, fromNodeId string, toNodeId string) Edge {
  id       := uuid.New()
  fromNode := Nodes[fromNodeId]
  toNode   := Nodes[toNodeId]

  fromNode.AppendToEdgeId(id)
  fromNode.AppendFromEdgeId(id)

  edge := Edge{id, label, fromNodeId, fromNode, toNodeId, toNode}
  Edges[id] = edge

  return edge
}


// ----------------------------
// Node
//
type Node struct {
  Id          string
  Label       string
  FromEdgeIds []string
  ToEdgeIds   []string
}

func (n Node) AppendToEdgeId(id string) {
  n.ToEdgeIds = append(n.ToEdgeIds, id)
  Nodes[n.Id] = n
}

func (n Node) AppendFromEdgeId(id string) {
  n.FromEdgeIds = append(n.FromEdgeIds, id)
  Nodes[n.Id] = n
}

func (n Node) ToEdges() []Edge {
  toEdges := make([]Edge, 0)

  // for _, edgeId := range n.ToEdgeIds {
  //   toEdges = append(toEdges, Nodes[edgeId])
  // }

  return toEdges
}


// ----------------------------
// Edge
//
type Edge struct {
  Id          string
  Label       string
  FromNodeId  string
  FromNode    Node
  ToNodeId    string
  ToNode      Node
}