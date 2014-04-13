package tgf

type Node struct {
  Id          string
  Label       string
  FromEdgeIds []string
  ToEdgeIds   []string
}

type Edge struct {
  Id          string
  Label       string
  FromNodeId  string
  FromNode    Node
  ToNodeId    string
  ToNode      Node
}