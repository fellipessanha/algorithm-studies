package leetcodesgo

import "sort"

type Packet struct {
	source, destination, timestamp int
}

type Router struct {
	limit    int
	queue    []Packet
	packages map[int][]int
	id       int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit,
		make([]Packet, 0, memoryLimit),
		make(map[int][]int),
		0,
	}
}

func (this *Router) popTimestamp(destination int) {
	this.packages[destination] = this.packages[destination][1:]
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := Packet{source, destination, timestamp}

	for i := len(this.queue) - 1; i >= 0; i-- {
		pkt := this.queue[i]
		if pkt == packet {
			return false
		}
		if pkt.timestamp < packet.timestamp {
			break
		}
	}
	if len(this.queue) >= this.limit {
		this.popTimestamp(this.queue[0].destination)
		this.queue = this.queue[1:]
	}

	this.packages[packet.destination] = append(
		this.packages[packet.destination],
		packet.timestamp,
	)

	this.queue = append(this.queue, packet)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.queue) < 1 {
		return []int{}
	}
	packet := this.queue[0]
	this.popTimestamp(packet.destination)
	this.queue = this.queue[1:]
	return []int{packet.source, packet.destination, packet.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	counter := this.packages[destination]
	leftIndex := sort.Search(len(counter), func(i int) bool { return counter[i] >= startTime })
	rightIndex := sort.Search(len(counter), func(i int) bool { return counter[i] > endTime })
	return rightIndex - leftIndex
}
