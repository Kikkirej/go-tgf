package ast

import (
  "code.google.com/p/go-uuid/uuid"
)

var (
  Nodes = make(map[string]Node)
  Edges = make(map[string]Edge)
)

func MakeNode(id string, label string) Node {
  node := Node{id, label, make([]string, 0), make([]string, 0)}
  Nodes[id] = node
  return node
}

//
// InboundNode ---- [label] ----> OutboundNode
//
func MakeEdge(label string, InboundNodeId string, OutboundNodeId string) Edge {
  id           := uuid.New()
  InboundNode  := Nodes[InboundNodeId]
  OutboundNode := Nodes[OutboundNodeId]

  InboundNode.AppendOutboundEdgeId(id)
  OutboundNode.AppendInboundEdgeId(id)

  edge := Edge{id, label, InboundNodeId, InboundNode, OutboundNodeId, OutboundNode}
  Edges[id] = edge

  return edge
}

func RootNodes() []Node {
  rootNodes := make([]Node, 0)

  for _, node := range Nodes {
    if len(node.InboundEdgeIds) == 0 {
      rootNodes = append(rootNodes, node)
    }
  }

  return rootNodes
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

func (n *Node) AppendOutboundEdgeId(id string) {
  n.OutboundEdgeIds = append(n.OutboundEdgeIds, id)
  Nodes[n.Id] = *n
}

func (n *Node) AppendInboundEdgeId(id string) {
  n.InboundEdgeIds = append(n.InboundEdgeIds, id)
  Nodes[n.Id] = *n
}

func (n *Node) OutboundEdges() []Edge {
  OutboundEdges := make([]Edge, 0)

  for _, edgeId := range n.OutboundEdgeIds {
    OutboundEdges = append(OutboundEdges, Edges[edgeId])
  }

  return OutboundEdges
}

func (n *Node) InboundEdges() []Edge {
  InboundEdges := make([]Edge, 0)

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
  InboundNodeId   string
  InboundNode     Node
  OutboundNodeId  string
  OutboundNode    Node
}