package main

import "fmt"

type Node struct {
	children [26]*Node
	wordEnd  bool
}

type Trie struct {
	root *Node
}

func (trie *Trie) insert(word string) {
	currentNode := trie.root
	for _, letter := range word {
		charIndex := letter - 'a' //Find the index to which the letter should be inserted. "letter" gives ascii value of the alphabet
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.wordEnd = true
}

func (trie *Trie) search(word string) {
	currentNode := trie.root
	for _, letter := range word {
		charIndex := letter - 'a'
		if currentNode.children[charIndex] == nil {
			fmt.Println(word, "is not in trie")
			return
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.wordEnd {
		fmt.Println(word, "found in trie")
	} else {
		fmt.Println(word, "is not in trie")
	}
}

func (trie *Trie) remove(word string) {
	currentNode := trie.root
	for _, letter := range word {
		charIndex := letter - 'a'
		if currentNode.children[charIndex] == nil {
			fmt.Println(word, "is not in trie")
			return
		}
		currentNode = currentNode.children[charIndex]
	}
	if len(currentNode.children) == 0 {
		currentNode = nil
		return
	}
	currentNode.wordEnd = false
}

func main() {
	t := &Trie{root: &Node{}}
	t.insert("hello")
	t.insert("hi")
	t.search("hello")
	t.remove("hello")
	t.search("hi")
	t.remove("hi")
}
