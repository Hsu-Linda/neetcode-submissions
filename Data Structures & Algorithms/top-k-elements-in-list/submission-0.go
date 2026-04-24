
type Node struct {
	num   int
	count int
}

type MinHeap struct {
	array []Node
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	h := &MinHeap{array: make([]Node, 0, k)}

	for num, count := range counts {
		newNode := Node{num, count}
		if len(h.array) < k {
			// 1. 沒滿，直接塞
			h.Insert(newNode)
		} else if count > h.array[0].count {
			// 2. 滿了，且新人比較強，直接替換 Root
			h.Replace(newNode)
		}
		// 如果新人比較弱，就直接無視
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = h.array[i].num
	}
	return res
}

// --- Heap 核心方法 ---

func (h *MinHeap) Insert(item Node) {
	h.array = append(h.array, item) // Go 的 append 是小寫
	h.siftUp(len(h.array) - 1)
}

func (h *MinHeap) Replace(item Node) {
	h.array[0] = item
	h.siftDown(0)
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.array[i].count < h.array[p].count {
			h.array[i], h.array[p] = h.array[p], h.array[i]
			i = p
		} else {
			break
		}
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.array)
	for {
		smallest := i
		l, r := 2*i+1, 2*i+2

		// 檢查左小孩是否存在且比較小
		if l < n && h.array[l].count < h.array[smallest].count {
			smallest = l
		}
		// 檢查右小孩是否存在且比較小
		if r < n && h.array[r].count < h.array[smallest].count {
			smallest = r
		}

		if smallest == i {
			break
		}

		h.array[i], h.array[smallest] = h.array[smallest], h.array[i]
		i = smallest
	}
}

