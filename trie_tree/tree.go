package trie_tree

type Node[V any] struct {
    Children map[rune]*Node[V]
    IsEnd    bool
    Value    V
}

func NewNode[V any]() *Node[V] {
    return &Node[V]{Children: make(map[rune]*Node[V])}
}

type Tree[V any] struct {
    Root *Node[V]
}

func NewTree[V any]() *Tree[V] {
    return &Tree[V]{Root: NewNode[V]()}
}

func (t *Tree[V]) Put(word string, value V) {
    node := t.Root
    for _, char := range word {
        if _, exists := node.Children[char]; !exists {
            node.Children[char] = NewNode[V]()
        }
        node = node.Children[char]
    }
    node.IsEnd = true
    node.Value = value
}

func (t *Tree[V]) Search(word string) (V, bool) {
    node := t.Root
    for _, char := range word {
        if _, exists := node.Children[char]; !exists {
            var zero V
            return zero, false
        }
        node = node.Children[char]
    }
    return node.Value, node.IsEnd
}

func (t *Tree[V]) StartsWith(prefix string) bool {
    node := t.Root
    for _, char := range prefix {
        if _, exists := node.Children[char]; !exists {
            return false
        }
        node = node.Children[char]
    }
    return true
}

func (t *Tree[V]) Delete(word string) {
    t.deleteRecursive(t.Root, word, 0)
}

func (t *Tree[V]) deleteRecursive(node *Node[V], word string, depth int) bool {
    if node == nil {
        return false
    }

    if depth == len(word) {
        if node.IsEnd {
            node.IsEnd = false
        }
        return len(node.Children) == 0
    }

    char := rune(word[depth])
    if t.deleteRecursive(node.Children[char], word, depth+1) {
        delete(node.Children, char)
        return len(node.Children) == 0 && !node.IsEnd
    }
    return false
}

type Entry[V any] struct {
    Word  string
    Value V
}

func (t *Tree[V]) GetWordsWithPrefix(prefix string) []Entry[V] {
    node := t.findPrefixNode(prefix)
    if node == nil {
        return nil
    }

    var results []Entry[V]
    t.collectWords(node, prefix, &results)
    return results
}

func (t *Tree[V]) findPrefixNode(prefix string) *Node[V] {
    node := t.Root
    for _, char := range prefix {
        if _, exists := node.Children[char]; !exists {
            return nil
        }
        node = node.Children[char]
    }
    return node
}

func (t *Tree[V]) collectWords(node *Node[V], prefix string, results *[]Entry[V]) {
    if node.IsEnd {
        *results = append(*results, Entry[V]{Word: prefix, Value: node.Value})
    }
    for char, child := range node.Children {
        t.collectWords(child, prefix+string(char), results)
    }
}
