package ast

import (
  "fmt"
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

//
// OutboundNode ---- [label] ----> InboundNode
//
func NewEdge(label string, OutboundNodeId string, InboundNodeId string) Edge {
  id           := uuid.New()
  OutboundNode := Nodes[OutboundNodeId]
  InboundNode  := Nodes[InboundNodeId]

  OutboundNode.AppendInboundEdgeId(id)
  InboundNode.AppendOutboundEdgeId(id)

  edge := Edge{id, label, InboundNodeId, InboundNode, OutboundNodeId, OutboundNode}
  Edges[id] = edge

  return edge
}


// ----------------------------
// Node
//
type Node struct {
  Id              string
  Label           string
  InboundEdgeIds  []string
  OutboundEdgeIds []string
}

func (n Node) AppendOutboundEdgeId(id string) {
  n.OutboundEdgeIds = append(n.OutboundEdgeIds, id)
  Nodes[n.Id] = n
}

func (n Node) AppendInboundEdgeId(id string) {
  n.InboundEdgeIds = append(n.InboundEdgeIds, id)
  Nodes[n.Id] = n
}

func (n Node) OutboundEdges() []Edge {
  OutboundEdges := make([]Edge, len(n.OutboundEdgeIds))

  for _, edgeId := range n.OutboundEdgeIds {
    OutboundEdges = append(OutboundEdges, Edges[edgeId])
  }

  return OutboundEdges
}

func (n Node) InboundEdges() []Edge {
  InboundEdges := make([]Edge, len(n.InboundEdgeIds))

  for _, edgeId := range n.InboundEdgeIds {
    InboundEdges = append(InboundEdges, Edges[edgeId])
  }

  return InboundEdges
}


// ----------------------------
// Edge
//
type Edge struct {
  Id              string
  Label           string
  OutboundNodeId  string
  OutboundNode    Node
  InboundNodeId   string
  InboundNode     Node
}