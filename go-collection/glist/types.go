package glist

type Element interface {
	Equals(e Element) bool

	Compare(v1 Element, v2 Element) bool
}

type Interface interface {
	Add(e Element)

	AddSlice(es []Element)

	Contains(e Element) bool

	Get(index int) Element

	IndexOf(e Element) int

	RemoveIndex(index int)

	Remove(e Element)

	Sort()

	IsEmpty() bool

	Len() int

	Clear()
}
