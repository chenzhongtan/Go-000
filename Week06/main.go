package main

import (
	"container/list"
	"sync"
	"time"
)

type census struct {
	Success int
	Fail    int
}

type slidingWindow struct {
	mu         sync.Mutex
	List       *list.List
	TimeSecond int64
	Size       int
	Census     map[int64]*census
}

func (w *slidingWindow) Add(now int64) {
	w.List.PushFront(now)
}


func (w *slidingWindow) Remove() {
	w.List.Remove(w.List.Back())
}

func (w *slidingWindow) IsSuccess() bool {
	w.mu.Lock()
	now := time.Now().Unix()
	res := false
	if w.List.Len() >= w.Size {
		if (now - w.List.Back().Value.(int64)) >= w.TimeSecond {
			w.Add(now)
			w.Remove()
			res = true
		}
	} else {
		res = true
		w.Add(now)
	}

	v1 := w.Census[now]
	if v1 == nil {
		v1 = new(census)
	}
	if res {
		v1.Success++
	} else {
		v1.Fail++
	}
	w.Census[now] = v1
	w.mu.Unlock()
	return res
}