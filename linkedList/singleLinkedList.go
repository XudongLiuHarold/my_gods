package linkedlist

type node struct {
	value interface{}
	next  *node
}

type List struct {
	head   *node
	tail   *node
	length int
}

func init() {

}

func (l *List) isEmpty() bool {
	return true
}

func (l *List) insert(val ...interface{}) {

}

func (l *List) deleteByIndex(idx int) {

}

func (l *List) deleteByData(val ...interface{}) {

}

func (l *List) addToHead(val ...interface{}) {

}

func (l *List) addToTail(val ...interface{}) {

}

func (l *List) getHead() {

}

func (l *List) getTail() {

}
