package ast

import (
  "testing"
  . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type AstSuite struct{}

var _ = Suite(&AstSuite{})

func (s *AstSuite) TestMakeNode(c *C) {
  node := MakeNode("1", "This is a Node with ID: 1")

  c.Assert(node.Id, Equals, "1")
  c.Assert(node.Label, Equals, "This is a Node with ID: 1")
}

func (s *AstSuite) TestMakeEdge(c *C) {
  MakeNode("1", "This is a Node with ID: 1")
  MakeNode("2", "This is a Node with ID: 2")
  edge  := MakeEdge("This is edge connecting 1->2", "1", "2")

  c.Assert(edge.Label, Equals, "This is edge connecting 1->2")
  c.Assert(edge.OutboundNodeId, Equals, "2")
  c.Assert(edge.InboundNodeId, Equals, "1")
}

func (s *AstSuite) TestRootNodes(c *C) {
  MakeNode("1", "This is a Node with ID: 1")
  MakeNode("2", "This is a Node with ID: 2")

  c.Assert(len(RootNodes()), Equals, 2)
}
