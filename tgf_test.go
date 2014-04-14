package tgf

import (
  "testing"
  . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TgfSuite struct{}

var _ = Suite(&TgfSuite{})

func (s *TgfSuite) TestIntegrationOptimistic(c *C) {
  allNodes, rootNodes, allEdges, _ := ParseFile("example.tgf")

  c.Assert(len(allNodes), Equals, 6)
  c.Assert(len(rootNodes), Equals, 2)
  c.Assert(len(allEdges), Equals, 4)

  c.Assert(rootNodes[0].OutboundEdges()[0].OutboundNodeId, Equals, "2")
  c.Assert(rootNodes[1].OutboundEdges()[0].OutboundNodeId, Equals, "4")
}
