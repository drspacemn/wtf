package main

import (
	"fmt"
	"strings"
)

type AcyclicGraph struct {
	Graph
}

type Graph struct {
	vertices Set
	edges Set
	downEdges map[interface{}]Set
	upEdges map[interface{}]Set
}

// Subgrapher allows a Vertex to be a Graph itself, by returning a Grapher.
type Subgrapher interface {
	Subgraph() Grapher
}

type Grapher interface {
	DirectedGraph() Grapher
}

type Vertex interface{}

type NamedVertex interface{
	Vertex
	Name() string
}

func (g *Graph) DirectedGraph() Grapher {
	return g
}

func (g *Graph) Vertices() []Vertex {
	result := make([]Vertex, 0, len(g.vertices()))
	for _, v := range g.vertices {
		result = append(result, v.(Vertex))
	}

	return result
}

func (g *Graph) Edges() []Edge {
	result := make([]Edge, 0, len(g.edges))
	for _, v := range g.edges {
		result = append(result, v.(Edge))
	}

	return result
}

func (g *Graph) EdgesFrom(v Vertex) (edges []Edge) {
	from := hashcode(v)
	for _, e := range g.Edges() {
		if hashcode(e.Target()) == from {
			edges = append(edges, e)
		}
	}

	return edges
}

func (g *Graph) EdgesTo(v Vertex) (edges []Edge) {
	search := hashcode(v)
	for _, e := range g.Edges() {
		if hashcode(e.Target()) == search {
			edges = append(edges, e)
		}
	}

	return edges
}

func (g *Graph) HasVertex(v Vertex) bool {
	return g.vertices.Include(v)
}

func (g *Graph) HasEdge(e Edge) bool {
	return g.edges.Include(e)
}

func (g *Graph) Add(v Vertex) Vertex {
	g.init()
	g.vertices.Add(v)
	return v
}

func (g *Graph) Remove(v Vertex) {
	g.vertices.Delete(v)

	for _, target := range g.downEdgesNoCopy(v) {
		g.RemoveEdge(BasicEdge(v, target))
	}
	for _, source := range g.upEdgesNoCopy(v) {
		g.RemoveEdge(BasicEdge(source, v))
	}
}

func (g *Graph) Replace(original, replacement Vertex) bool {
	if !g.vertices.Include(original) {
		return false
	}

	if original == replacement {
		return true
	}

	g.Add(replacement)
	for _, target := range g.downEdgesNoCopy(original) {
		g.Connect(BasicEdge(replacement, target))
	}
	for _, source := range g.upEdgesNoCopy(original) {
		g.Connect(basicEdge(source, replacement))
	}

	g.Remove(original)

	return true
}

func (g *Graph) RemoveEdge(edge Edge) {
	g.init()

	g.edges.Delete(edge)

	if s, ok := g.downEdges[hashcode(edge.Source())]; ok {
		s.Delete(edge.Target())
	}
	if s, ok := g.upEdges[hashcode(edge.Target())]; ok {
		s.Delete(edge.Source())
	}
}

func (g *Graph) UpEdges(v Vertex) Set {
	return g.upEdgesNoCopy(v).Copy()
}

func (g *Graph) DownEdges(v Vertex) Set {
	return g.downEdgesNoCopy(v).Copy()
}

func (g *Graph) upEdgesNoCopy(v Vertex) Set {
	g.init()
	return g.upEdges[hashcode(v)]
}

func (g *Graph) downEdgesNoCopy(v Vertex) Set {
	g.init()
	return g.downEdges[hashcode(v)]
}

func (g *Graph) Connect(edge Edge) {
	g.init()

	source := edge.Source()
	target := edge.Target()
	sourceCode := hashcode(source)
	targetCode := hashcode(target)

	if s, ok := g.downEdges[sourceCode]; ok && s.Include(target) {
		return
	}

	g.edges.Add(edge)

	s, ok := g.downEdges[sourceCode]
	if !ok {
		s = make(Set)
		g.downEdges[sourceCode] = s
	}
	s.Add(target)

	s, ok = g.upEdges[targetCode]
	if !ok {
		s = make(Set)
		g.upEdges[targetCode] = s
	}
	s.Add(source)
}

func (g *Graph) StringWithNodeTypes() string {
	var buf bytes.Buffer

	vertices := g.Vertices()
	names := make([]string, 0, len(vertices))
	mapping := make(map[string]Vertex, len(vertices))
	for _, v := range vertices {
		name := VertexName(v)
		names = append(names, name)
		mapping[name] = v
	}
	sort.Strings(names)

	for _, name := range names {
		v := mapping[name]
		targets := g.downEdges[hashcode(v)]

		buf.WriteString(fmt.Sprintf("%s - %T\n", name, v))

		deps := make([]string, 0, targets.Len())
		targetNodes := make(map[string]Vertex)
		for _, target := range targets {
			dep := VertexName(target)
			deps = append(deps, dep)
			targetNodes[dep] = target
		}
		sort.Strings(deps)

		for _, d := range deps {
			buf.WriteString(fmt.Sprintf(" %s - %T\n", d, targetNodes[d]))
		}
	}

	return buf.String()
}

// String outputs some human-friendly output for the graph structure.
func (g *Graph) String() string {
	var buf bytes.Buffer

	// Build the list of node names and a mapping so that we can more
	// easily alphabetize the output to remain deterministic.
	vertices := g.Vertices()
	names := make([]string, 0, len(vertices))
	mapping := make(map[string]Vertex, len(vertices))
	for _, v := range vertices {
		name := VertexName(v)
		names = append(names, name)
		mapping[name] = v
	}
	sort.Strings(names)

	// Write each node in order...
	for _, name := range names {
		v := mapping[name]
		targets := g.downEdges[hashcode(v)]

		buf.WriteString(fmt.Sprintf("%s\n", name))

		// Alphabetize dependencies
		deps := make([]string, 0, targets.Len())
		for _, target := range targets {
			deps = append(deps, VertexName(target))
		}
		sort.Strings(deps)

		// Write dependencies
		for _, d := range deps {
			buf.WriteString(fmt.Sprintf("  %s\n", d))
		}
	}

	return buf.String()
}

func (g *Graph) init() {
	if g.vertices == nil {
		g.vertices = make(Set)
	}
	if g.edges == nil {
		g.edges = make(Set)
	}
	if g.downEdges == nil {
		g.downEdges = make(map[interface{}]Set)
	}
	if g.upEdges == nil {
		g.upEdges = make(map[interface{}]Set)
	}
}

func (g *Graph) Dot(opts *DotOpts) []byte {
	return newMarshalGraph("", g).Dot(opts)
}

func VertexName(raw Vertex) string {
	switch v := raw.(type) {
	case NamedVertex:
		return v.Name()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
