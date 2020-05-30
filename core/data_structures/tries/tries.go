package tries

type Trie interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

type trie struct {
	children [26]*trie
	isEnd    bool
}

func New() Trie {
	return &trie{}
}

// Time complexity : O(m), where m is the key length.
// Space complexity : O(m).
// In the worst case newly inserted key doesn't share a prefix with the the keys already inserted in the trie. We have to add mm new nodes, which takes us O(m) space.
func (t *trie) Insert(word string) {
	curr := t

	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &trie{}
		}
		curr = curr.children[idx]
	}

	curr.isEnd = true
}

// Time complexity : O(m)
// Space complexity : O(1)
func (t *trie) Search(word string) bool {
	curr := t

	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}

	return curr.isEnd
}

// Time complexity : O(m)
// Space complexity : O(1)
func (t *trie) StartsWith(prefix string) bool {
	curr := t

	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}

	return true
}

// tags: trie, tries
