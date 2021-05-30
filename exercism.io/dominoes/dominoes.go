package dominoes

type queueItem struct {
	pool  []Domino
	seen  []Domino
	chain []Domino
}

func (i *queueItem) Tail() int {
	return i.chain[len(i.chain)-1][1]
}

type Domino = [2]int

func MakeChain(dominoes []Domino) ([]Domino, bool) {
	if len(dominoes) == 0 {
		return []Domino{}, true
	}

	// First queue initialization
	queue := []queueItem{queueItem{dominoes[1:], []Domino{}, []Domino{dominoes[0]}}}

	var item queueItem
	var d Domino
	for len(queue) > 0 {
		// Take first item out of queue
		item, queue = queue[0], queue[1:]

		// Exit if no items in pool of current item, or continue with next item
		if len(item.pool) == 0 {
			if len(item.seen) == 0 && item.chain[0][0] == item.Tail() {
				return item.chain, true
			}
			continue
		}

		d, item.pool = item.pool[0], item.pool[1:]
		// Reverse item if necessary
		if d[0] != item.Tail() {
			d = Domino{d[1], d[0]}
		}
		// If matching, add new item to queue with updated chain and where seen items are put back in pool
		if d[0] == item.Tail() {
			queue = append(queue, queueItem{
				append(item.pool, item.seen...),
				[]Domino{},
				append(item.chain, d)})
		}
		// Lastly, declare the current domino as seen
		queue = append(queue, queueItem{item.pool, append(item.seen, d), item.chain})
	}
	return nil, false
}
