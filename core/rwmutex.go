package core

/* 
   To my future colleagues:
   This code is the very core, the heart and soul of all concurrent GPU-CPU logic.
   Please, take a deep breath before editing. Raptors will kill you if break it.
   -Arne.
*/

import "sync"

// RWMutex protects an array for safe access by
// one writer and many readers. 
// RWMutex makes sure the readers receive all data
// exactly once and in the correct order.
// The functionality is like a Go channel, but
// without copying the data.
type RWMutex struct {
	cond       sync.Cond  // wait condition: read/write is safe
	state      sync.Mutex // protects the internal state, used in cond.
	absA, absB int64      // half-open interval locked for writing
	n          int        // Total number of elements in protected array.
	readers    []*RMutex  // all readers who can access this rwmutex
}

// RWMutex to protect an array of length N.
func NewRWMutex(N int) *RWMutex {
	m := new(RWMutex)
	m.n = N
	m.cond = *(sync.NewCond(&m.state))
	return m
}

// Lock the next delta elements for writing.
func (m *RWMutex) WriteNext(delta int) {
	m.cond.L.Lock()

	delta64 := int64(delta)
	if m.absA != m.absB {
		panic("rwmutex: lock of locked mutex")
	}
	for !m.canWLock(m.absA, m.absA+delta64) {
		m.cond.Wait()
	}
	m.absB = m.absA + delta64

	m.cond.L.Unlock()
	m.cond.Broadcast()
}

// WRange returns the currently write-locked range.
// It is not thread-safe because the RWMutex is only
// supposed to be accessed by one writer thread.
func (m *RWMutex) WRange() (start, stop int) {
	return int(m.absA % int64(m.n)), int((m.absB-1)%int64(m.n)) + 1
}

// Unlock the previous interval locked for writing.
func (m *RWMutex) WriteDone() {
	m.cond.L.Lock()
	if m.absA == m.absB {
		panic("rwmutex: unlock of unlocked mutex")
	}
	m.absA = m.absB

	m.cond.L.Unlock()
	m.cond.Broadcast()
}

// Can m safely lock for writing [start, stop[ ?
func (m *RWMutex) canWLock(a, b int64) (ok bool) {
	for _, r := range m.readers {
		if a < r.absD || b > (r.absC+int64(m.n)) {
			Debug("canWLock", a, b, ": false")
			return false
		}
	}
	Debug("canWLock", a, b, ": true")
	return true
}
