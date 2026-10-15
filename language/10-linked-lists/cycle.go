/*
HasCycle works on raw nodes rather than a LinkedList, because a well-behaved
LinkedList never has a cycle. Hand-built node chains can.
*/
package listkata

func HasCycle[T any](head *Node[T]) bool {
	return false
}
