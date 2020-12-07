package main

import (
	"log"
	"os"
	"strings"

	"strconv"

	"github.com/fbegyn/aoc2020/go/helpers"
	"regexp"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	graph := NewGraph()
	childBagsRgx := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
	
	for inp := range input {
		line := inp[:len(inp)-1]
		split := strings.Split(line, " contain ")
		parentRaw, childRaw := split[0], split[1]
		parentKey := strings.ReplaceAll(parentRaw, " bags", "")

		for _, childInp := range strings.Split(childRaw, ", ") {
			if childInp == "no other bags" {
				continue
			}
			childData := childBagsRgx.FindStringSubmatch(childInp)
			if len(childData) == 0 {
				log.Fatalln("did not find child data")
			}

			weight, err := strconv.Atoi(childData[1])
			if err != nil {
				log.Fatalln(err)
			}

			childKey := childData[2]
			graph.AddWeightedChildToParent(parentKey, childKey, weight)
		}
	}

	bag := graph.nodeKeys["shiny gold"]
	part1 := len(graph.GetAllParents(bag))
	part2 := graph.ChildrenWeight(bag)

	log.Printf("solution for part 1: %d", part1)
	log.Printf("solution for part 2: %d", part2)
}

type Graph struct {
	nodeKeys map[string]*Node
}

type Node struct {
	key string
	parentEdge []Edge
	childEdge []Edge
}

type Edge struct {
	parent *Node
	child *Node
	weight int
}

func NewGraph() *Graph {
	return &Graph{
		make(map[string]*Node),
	}
}

func (g *Graph) AddNode(key string) *Node {
	exists, ok := g.nodeKeys[key]
	if !ok {
		node := &Node{key: key}
		g.nodeKeys[key] = node
		return node
	}
	return exists
}

func (g *Graph) AddWeightedChildToParent(parent, child string, weight int) {
	parentNode, ok := g.nodeKeys[parent]
	if !ok {
		parentNode = g.AddNode(parent)
	}

	childNode, ok := g.nodeKeys[child]
	if !ok {
		childNode = g.AddNode(child)
	}

	edge := Edge{
		parent: parentNode,
		child: childNode,
		weight: weight,
	}

	parentNode.childEdge = append(parentNode.childEdge, edge)
	childNode.parentEdge = append(childNode.parentEdge, edge)
}


func (g *Graph) GetAllParents(n *Node) []*Node {
	parentSet := make(map[*Node]struct{})
	for _, parentEdge := range n.parentEdge {
		parentSet[parentEdge.parent] = struct{}{}
		for _, parent := range g.GetAllParents(parentEdge.parent) {
			parentSet[parent] = struct{}{}
		}
	}

	keys := make([]*Node, 0, len(parentSet))
	for k := range parentSet {
		keys = append(keys, k)
	}
	return keys
}

func (g *Graph) ChildrenWeight(n *Node) (weight int) {
	for _, childEdge := range n.childEdge {
		weight += childEdge.weight
		weight += g.ChildrenWeight(childEdge.child) * childEdge.weight
	}
	return
}

type Bag struct {
	typ     string
	color   string
	count   []int
	content []*Bag
}

func (b *Bag) Compare(ba Bag) bool {
	return b.typ == ba.typ && b.color == ba.color
}

func (b *Bag) CountBags() int {
	count := 0
	for ind := range b.content {
		count += b.count[ind] + b.count[ind] * b.content[ind].CountBags()
	}
	return count
}

func (b *Bag) Contains(ba Bag) bool {
	for ind := range b.content {
		if b.content[ind].Compare(ba) {
			return true
		}
		if b.content[ind].Contains(ba) {
			return true
		}
	}
	return false
}

func NewBagFromLine(line string) (b *Bag) {
	bag := Bag{}
	line = strings.TrimRight(line, ".")
	split := strings.SplitAfter(line, "contain")

	bagDesc := strings.Split(split[0], " ")
	bag.typ = bagDesc[0]
	bag.color = bagDesc[1]

	allContents := strings.TrimSpace(split[1])
	contents := strings.Split(allContents, ",")

	for _, cont := range contents {
		cont = strings.TrimSpace(cont)
		b := ParseContents(cont)
		if b != nil {
			bag.content = append(bag.content, b)
			bag.count = append(bag.count, b.count...)
		}
	}

	return &bag
}

func ParseContents(desc string) *Bag {
	content := strings.Split(desc, " ")
	if content[0] == "no" {
		return nil
	}
	count, err := strconv.Atoi(content[0])
	if err != nil {
		log.Fatalln(err)
	}
	return &Bag{
		typ:     content[1],
		color:   content[2],
		count:   []int{count},
		content: nil,
	}
}
