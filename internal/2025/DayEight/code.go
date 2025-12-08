package dayeight

import (
	"cmp"
	"container/heap"
	"os"
	"slices"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D8", ParseInput, PartOne, PartTwo)
}

type JunctionBox struct {
	X, Y, Z int
}

type UnionFind struct {
	parent []int
	size   []int
}

type Connection struct {
	firstJunctionBox, secondJunctionBox int
	squaredDistance                     int
}

type ConnectionHeap []Connection

func (h ConnectionHeap) Len() int           { return len(h) }
func (h ConnectionHeap) Less(i, j int) bool { return h[i].squaredDistance > h[j].squaredDistance }
func (h ConnectionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ConnectionHeap) Push(x any)        { *h = append(*h, x.(Connection)) }
func (h *ConnectionHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewUnionFind(count int) *UnionFind {
	parent := make([]int, count)
	size := make([]int, count)
	for idx := range parent {
		parent[idx] = idx
		size[idx] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(junctionBox int) int {
	if uf.parent[junctionBox] != junctionBox {
		uf.parent[junctionBox] = uf.Find(uf.parent[junctionBox])
	}
	return uf.parent[junctionBox]
}

func (uf *UnionFind) Union(junctionBox1, junctionBox2 int) (bool, int) {
	root1 := uf.Find(junctionBox1)
	root2 := uf.Find(junctionBox2)
	if root1 == root2 {
		return false, root1
	}
	if uf.size[root1] < uf.size[root2] {
		root1, root2 = root2, root1
	}
	uf.parent[root2] = root1
	uf.size[root1] += uf.size[root2]
	return true, root1
}

func parseJunctionBoxes(data []string) []JunctionBox {
	var junctionBoxes []JunctionBox
	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		junctionBoxes = append(junctionBoxes, JunctionBox{X: x, Y: y, Z: z})
	}
	return junctionBoxes
}

func generateAllConnections(junctionBoxes []JunctionBox) []Connection {
	junctionBoxCount := len(junctionBoxes)
	connections := make([]Connection, 0, junctionBoxCount*(junctionBoxCount-1)/2)

	for i := range junctionBoxCount {
		for j := i + 1; j < junctionBoxCount; j++ {
			dx := junctionBoxes[i].X - junctionBoxes[j].X
			dy := junctionBoxes[i].Y - junctionBoxes[j].Y
			dz := junctionBoxes[i].Z - junctionBoxes[j].Z
			squaredDistance := dx*dx + dy*dy + dz*dz

			connections = append(connections, Connection{
				firstJunctionBox:  i,
				secondJunctionBox: j,
				squaredDistance:   squaredDistance,
			})
		}
	}

	slices.SortFunc(connections, func(a, b Connection) int {
		return cmp.Compare(a.squaredDistance, b.squaredDistance)
	})

	return connections
}

func findKSmallestConnections(junctionBoxes []JunctionBox, k int) []Connection {
	junctionBoxCount := len(junctionBoxes)
	connectionHeap := make(ConnectionHeap, 0, k)
	heap.Init(&connectionHeap)

	for i := range junctionBoxCount {
		for j := i + 1; j < junctionBoxCount; j++ {
			dx := junctionBoxes[i].X - junctionBoxes[j].X
			dy := junctionBoxes[i].Y - junctionBoxes[j].Y
			dz := junctionBoxes[i].Z - junctionBoxes[j].Z
			squaredDistance := dx*dx + dy*dy + dz*dz

			connection := Connection{
				firstJunctionBox:  i,
				secondJunctionBox: j,
				squaredDistance:   squaredDistance,
			}

			if len(connectionHeap) < k {
				heap.Push(&connectionHeap, connection)
			} else if connection.squaredDistance < connectionHeap[0].squaredDistance {
				heap.Pop(&connectionHeap)
				heap.Push(&connectionHeap, connection)
			}
		}
	}

	connections := make([]Connection, 0, len(connectionHeap))
	for connectionHeap.Len() > 0 {
		connections = append(connections, heap.Pop(&connectionHeap).(Connection))
	}

	slices.SortFunc(connections, func(a, b Connection) int {
		return cmp.Compare(a.squaredDistance, b.squaredDistance)
	})

	return connections
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func PartOne(data []string) int {
	junctionBoxes := parseJunctionBoxes(data)
	junctionBoxCount := len(junctionBoxes)

	targetPairs := 1000
	if junctionBoxCount <= 20 {
		targetPairs = 10
	}

	connections := findKSmallestConnections(junctionBoxes, targetPairs)

	uf := NewUnionFind(junctionBoxCount)
	for _, connection := range connections {
		uf.Union(connection.firstJunctionBox, connection.secondJunctionBox)
	}

	circuitSizes := make(map[int]int)
	for idx := range junctionBoxes {
		root := uf.Find(idx)
		circuitSizes[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[0] * sizes[1] * sizes[2]
}

func PartTwo(data []string) int {
	junctionBoxes := parseJunctionBoxes(data)
	connections := generateAllConnections(junctionBoxes)

	uf := NewUnionFind(len(junctionBoxes))
	var lastConnection Connection

	for _, connection := range connections {
		merged, root := uf.Union(connection.firstJunctionBox, connection.secondJunctionBox)
		if merged {
			lastConnection = connection
			if uf.size[root] == len(junctionBoxes) {
				break
			}
		}
	}

	return junctionBoxes[lastConnection.firstJunctionBox].X * junctionBoxes[lastConnection.secondJunctionBox].X
}
