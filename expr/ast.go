package expr

type Node struct {
	Val   Token
	Left  *Node
	Right *Node
}
