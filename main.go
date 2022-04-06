package ips

const (
	topN int = 100
)

var (
	// cached stores all request IPs.
	cached = make(map[string]*node)
	// topSorted is a binary tree of sorted pointers to nodes to grab the top N quickly
	topSorted *node
)

type (
	node struct {
		right, left *node
		ip          string
		count       int
	}
)

func request_handled(ip string) {
	n, ok := cached[ip]
	if !ok {
		n = &node{nil, nil, ip, 0}
	}
	n.count++
	cached[ip] = n
	if topSorted == nil {
		topSorted = n
		return
	}
	insert(topSorted, n)
}

// top100 returns the top 100. This is done by maintaining a binary tree and then performing a depth first search for the 100 largest values.
func top100() []*node {
	var t []*node
	return dft(topSorted, t)
}

func dft(root *node, s []*node) []*node {
	if len(s) == topN {
		return s
	}
	if root.right == nil {
		return append(s, root)
	}
	if root.left == nil {
		return append(s, root)
	}
	return append(dft(root.right, s), dft(root.left, s)...)
}

func clear() {
	cached = make(map[string]*node)
}

func insert(root *node, n *node) {
	if root == nil {
		return
	}
	if root.count > n.count {
		if root.left == nil {
			root.left = n
			return
		}
	} else {
		if root.right == nil {
			root.right = n
			return
		}
		insert(root.left, n)
		insert(root.right, n)
	}
}
