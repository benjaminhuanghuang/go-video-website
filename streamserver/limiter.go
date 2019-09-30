package main

import (
	"log"
)

type ConnectionLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnectionLimiter {
	return &ConnectionLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnectionLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnectionLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
