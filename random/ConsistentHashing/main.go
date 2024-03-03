package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

var crc32q = crc32.MakeTable(0xD5828281)

func main() {
	ring := NewRing()
	fmt.Println(ring)
	ring.AddNode("00") // virtual node 0
	ring.AddNode("01") // virtual node 0
	ring.AddNode("02") // virtual node 0
	ring.AddNode("03") // virtual node 0
	ring.AddNode("04") // virtual node 0
	fmt.Println(ring)
	ring.AddNode("10") // virtual node 1
	ring.AddNode("11") // virtual node 1
	ring.AddNode("12") // virtual node 1
	ring.AddNode("13") // virtual node 1
	ring.AddNode("14") // virtual node 1

	fmt.Printf("Node %v\n", ring.Get("00"))
	fmt.Printf("Node %v\n", ring.Get("01"))
	fmt.Printf("Node %v\n", ring.Get("02"))
	fmt.Printf("Node %v\n", ring.Get("03"))
	fmt.Printf("Node %v\n", ring.Get("04"))
	fmt.Printf("Node %v\n", ring.Get("11"))
	fmt.Printf("Node %v\n", ring.Get("14"))

	arr := []Node{ring.Get("00"), ring.Get("01"), ring.Get("02"), ring.Get("03"), ring.Get("04"), ring.Get("10"), ring.Get("11"), ring.Get("12")}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].HashId < arr[j].HashId
	})
	fmt.Println(arr)
}

func NewRing() Ring {
	return []Node{}
}

type Ring []Node

func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	*r = append(*r, node)
	sort.Sort(r)
}

func (r *Ring) Get(key string) Node {
	searchfn := func(i int) bool {
		return (*r)[i].HashId >= crc32.Checksum([]byte(key), crc32q)
	}

	i := sort.Search(r.Len(), searchfn)
	if i >= r.Len() {
		i = 0
	}
	return (*r)[i]
}

func (r *Ring) Len() int {
	return len(*r)
}

func (r *Ring) Less(i, j int) bool {
	return (*r)[i].HashId < (*r)[j].HashId
}

func (r *Ring) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

type Node struct {
	Id     string
	HashId uint32
}

func NewNode(id string) Node {
	return Node{
		Id:     id,
		HashId: crc32.Checksum([]byte(id), crc32q),
	}
}
