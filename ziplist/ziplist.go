package ziplist

import "unsafe"

type ZipList struct {

}

type Head struct {
	zlbytes uint32 //整个占用的内存字节数，对 ziplist 进行内存重分配，或者计算末端时使用。
	zltail uint32 //到达表尾节点的偏移量。 通过这个偏移量，可以在不遍历整个 ziplist 的前提下，弹出表尾节点。
	ziplist uint32 //中节点的数量。 当这个值小于 UINT16_MAX （65535）时，这个值就是 ziplist 中节点的数量； 当这个值等于 UINT16_MAX 时，节点的数量需要遍历整个 ziplist 才能计算得出。
}

type EntryData struct {
	PreEntryLength uint32
	Encoding uint32
	Content Content
}

type Content interface {
	Len() uint32
	Point() unsafe.Pointer
}

func (list *ZipList) Append() {

}

func (list *ZipList) RemoveEnd() {

}

func (list *ZipList) Len() uint32 {

}

func (list *ZipList) Size() uint32 {

}
