package linkedlist

type node struct {
	value interface{}
	next  *node
}

type List struct {
	head *node
	tail *node
	size int
}

func init() {

}

func (l *List) isEmpty() bool {
	return true
}

func (l *List) length() int {
	return 0
}

func (l *List) get(idx int) {

}

func (l *List) insert(idx int, val ...interface{}) {

}

func (l *List) deleteByIndex(idx int) {

}

func (l *List) deleteByData(val ...interface{}) {

}

func (l *List) display() {

}
