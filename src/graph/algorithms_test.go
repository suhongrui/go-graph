package graph

import (
	"testing"
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

func ReduceDirectPathsSpec(c gospec.Context) {
	c.Specify("Reduced triangle", func() {
		gr := NewDirectedMap()
		gr.AddArc(1, 2)
		gr.AddArc(2, 3)
		gr.AddArc(1, 3)
		
		rgr := NewDirectedMap()
		ReduceDirectPaths(gr, rgr, nil)
		
		expectedGraph := NewDirectedMap()
		expectedGraph.AddArc(1, 2)
		expectedGraph.AddArc(2, 3)
		c.Expect(rgr, DirectedGraphEquals, expectedGraph)
	})
	
	c.Specify("A bit more complex example", func() {
		gr := NewDirectedMap()
		gr.AddArc(1, 2)
		gr.AddArc(2, 3)
		gr.AddArc(3, 4)
		gr.AddArc(2, 4)
		gr.AddArc(4, 5)
		gr.AddArc(1, 6)
		gr.AddArc(2, 6)
		
		rgr := NewDirectedMap()
		ReduceDirectPaths(gr, rgr, nil)
		
		expectedGraph := NewDirectedMap()
		expectedGraph.AddArc(1, 2)
		expectedGraph.AddArc(2, 3)
		expectedGraph.AddArc(3, 4)
		expectedGraph.AddArc(4, 5)
		expectedGraph.AddArc(2, 6)
		c.Expect(rgr, DirectedGraphEquals, expectedGraph)
	})
}

func TopologicalSortSpec(c gospec.Context) {
	gr := NewDirectedMap()
	c.Specify("Single node graph", func() {
		gr.AddNode(NodeId(1))
		nodes, hasCycle := TopologicalSort(gr)
		c.Expect(hasCycle, IsFalse)
		c.Expect(nodes, ContainsExactly, Values(NodeId(1)))
	})
	
	c.Specify("Simple two nodes graph", func() {
		gr.AddArc(1, 2)
		nodes, hasCycle := TopologicalSort(gr)
		c.Expect(hasCycle, IsFalse)
		c.Expect(nodes, ContainsExactly, Values(NodeId(1), NodeId(2)))
	})
} 

func TestAlgorithms(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(ReduceDirectPathsSpec)
	r.AddSpec(TopologicalSortSpec)
	gospec.MainGoTest(r, t)
}
