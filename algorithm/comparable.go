package algorithm

type Comparable interface {
	Compare(other interface{}) int
}