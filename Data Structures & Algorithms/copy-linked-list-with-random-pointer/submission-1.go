/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// [ONode1.Val, ONode1.next (ONode2), ONode1.random (Oaddr4)]  (head)
// 


// [NNode2]

// 1. Create Head NNode1  // [NNode1.Val,nil, nil]
// 2. ONode has next 
// -> 2-1. create a empty NNode(NNode2)
// -> 2-2. let the NNode1.Next is empty's addr
// ONodeCur 2 
// [NNode1.Val, NNode2Addr, nil]
// [NNode2, nil, nil]
// 3. ONode no next


// [node1.Val, node1.next (node2), node1.random (addr4)]  (head)
// -> [node2.Val, node2.next(node3), node2.random]
// -> [node3.Val, node3.next(node4), node3.random]
// -> [node4.Val, node4.next(nil), node4.random]
// -> null

// node.random  use the random addr to find which node (node index)
// ONodeMap
// {
// 	ONode1Addr: 1
// 	ONode2Addr: 2
// 	ONode3Addr: 3
// 	ONode4Addr: 4
//}

// ONode.Random we get the ONode addr, we can use map to find which one 
// We get ONodeX.Random = FSDFJSDOIJ 
// get from ONodeMap it's ONode3's Addr  which means ONodeX.Random = ONode3.Addr

// mapping  to NNodeMap get the NNode3's addr 
// and let NNodeX.Random = NNode3.Addr

// NNodeMap
// {
// 	1: NNodeAddr1
// 	2: NNodeAddr2
// 	3: NNodeAddr3
// 	4: NNodeAddr4
//}







// [node1.Val, node1.next, node1.random (node4) ]  (head)
// [node4.Val, node4.next(nil) node4.random (node3) ]
// [node3.Val, node3.next, node4.random (nil) ]
// we can't creat node4 since we already create it one the "next for loop"



// at first we can check the next/random node is nil or not 
// if is nil, then we linked to null, we finished
// if isn't, then we prepare the node for link to that

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	
	nNode := Node{
		Val: head.Val,
		Next: nil,
		Random: nil,
	}
	nNodeHead := &nNode
	nNodeCur  := &nNode
	oNodeCur  := head

	oNodeMap := make(map[*Node]int)  // ONodeAddr : index
	nNodeMap := make(map[int]*Node)  // index : NNodeAddr
	index := 0
	oNodeMap[oNodeCur] = index
	nNodeMap[index] = nNodeCur

	for oNodeCur.Next != nil {
		oNodeCur = oNodeCur.Next
		index ++
		oNodeMap[oNodeCur] = index
		nNodeCur.Next = &Node{
			Val: oNodeCur.Val,
			Next: nil,
			Random: nil,
		}
		nNodeCur = nNodeCur.Next
		nNodeMap[index] = nNodeCur
	}

	oNodeCur = head
	nNodeCur = nNodeHead
	for oNodeCur != nil {
		if randomIndex, ok := oNodeMap[oNodeCur.Random]; !ok {
			nNodeCur.Random = nil
		} else {
			nNodeRandomAddr := nNodeMap[randomIndex]
			nNodeCur.Random = nNodeRandomAddr
		}
		
		nNodeCur = nNodeCur.Next
		oNodeCur = oNodeCur.Next
	}

	return nNodeHead
}
