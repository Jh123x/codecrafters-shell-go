package trie

import "fmt"

type Node struct {
	next   map[byte]*Node
	hasVal bool
}

func NewNode() *Node {
	return &Node{
		next:   make(map[byte]*Node, 26),
		hasVal: false,
	}
}

func (n *Node) Print() {
	fmt.Print("Node{")
	for letter, node := range n.next {
		if node == nil {
			continue
		}

		fmt.Printf("%s(%v):", string(letter), node.hasVal)
		node.Print()
	}
	fmt.Print("}")
}

func (n *Node) AddWord(letters string) {
	if len(letters) == 0 {
		n.hasVal = true
		return
	}

	start := letters[0]
	if _, ok := n.next[start]; !ok {
		n.next[start] = NewNode()
	}

	n.next[start].AddWord(letters[1:])
}

func (n *Node) HasWord(letters string) bool {
	if len(letters) == 0 {
		return n.hasVal
	}

	start := letters[0]
	if _, ok := n.next[start]; !ok {
		return false
	}

	return n.next[start].HasWord(letters[1:])
}

func (n *Node) GetPrefixWords(prefix string) []string {
	if len(prefix) == 0 {
		return n.GetAllWords()
	}

	start := prefix[0]
	if _, ok := n.next[start]; !ok {
		return []string{}
	}

	return n.next[start].GetPrefixWords(prefix[1:])
}

func (n *Node) GetAllWords() []string {
	acc := make([]string, 0)

	if n.hasVal {
		acc = append(acc, "")
	}

	for letter, nodeVal := range n.next {
		if nodeVal == nil {
			continue
		}

		for _, word := range nodeVal.GetAllWords() {
			acc = append(acc, string(letter)+word)
		}
	}

	return acc
}
