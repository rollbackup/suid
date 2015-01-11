package suid

import (
	"fmt"
	"sync"
	"time"
)

const epoch = 1420000000000000000

type SharedID struct {
	seq       int64
	currentMs int64
	sync.Mutex
}

func (s *SharedID) Generate(shardId int) (int64, error) {
	ms := time.Now().Unix() * 1000
	id := ms - epoch
	id = id << 23
	id = id | (int64(shardId) << 10)

	seq, err := s.nextSeq(ms)
	if err != nil {
		return int64(0), err
	}

	fmt.Println(seq)

	id |= seq % 1024
	return id, nil
}

func (s *SharedID) nextSeq(ms int64) (int64, error) {
	s.Lock()
	defer s.Unlock()

	if s.currentMs > ms {
		return int64(0), fmt.Errorf("time goes backward in this machine")
	}

	if s.currentMs < ms {
		s.currentMs = ms
	} else {
		s.seq++
	}

	return s.seq, nil
}

var defaultGenerator = &SharedID{}

func Generate(shardId int) (int64, error) {
	return defaultGenerator.Generate(shardId)
}
