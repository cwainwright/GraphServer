package graph

import (
	linkedlist "GraphServer/linkedlist"
)

// --- Vertex ---

type Vertex struct {
	ID    int                          // Identifier for Vertex
	Edges *linkedlist.LinkedList[Edge] // Edges Connecting Vertex
}

func NewVertex(id int) Vertex {
	return Vertex{id, &linkedlist.LinkedList[Edge]{}}
}

// --- Edge ---

type Edge struct {
	Weight float64 // Cost of Traversal
	Vertex *Vertex // Connecting Vertex
}

func NewEdge(weight float64) Edge {
	return Edge{weight, nil}
}

// --- Graph ---

type Graph struct{ linkedlist.LinkedList[Vertex] } // List of Graph Vertices

func (graph *Graph) FindNode(id int) *linkedlist.Node[Vertex] {
	for node := graph.Head; node != nil; node = node.Next {
		if node.Value.ID == id {
			return node
		}
	}
	return nil
}

func (graph *Graph) FindVertex(id int) *Vertex {
	for node := graph.Head; node != nil; node = node.Next {
		if node.Value.ID == id {
			return &node.Value
		}
	}
	return nil
}

func (graph *Graph) AddVertex(id int) *Vertex {
	if graph.FindVertex(id) == nil {
		vertex := NewVertex(id)
		graph.AppendNode(vertex)
		return &vertex
	}
	return nil
}

func (graph *Graph) RemoveVertex(id int) {
	graph.RemoveNode(func(a Vertex, b Vertex) bool { return a.ID == b.ID }, NewVertex(id))
}
