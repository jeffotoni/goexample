// The Computer Language Benchmarks Game
// http://benchmarksgame.alioth.debian.org/
//
// Go adaptation of binary-trees program with arena allocator
//
// Forked from https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-9.html and
// added arena allocator to have parity with Rust typed_arena version -Dinko Korunic
//
// Final results compared to Rust #2 (https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-rust-2.html)
//
// Benchmark #1: ./go_orig.sh
//   Time (mean ± σ):     12.708 s ±  0.040 s    [User: 80.741 s, System: 1.275 s]
//   Range (min … max):   12.654 s … 12.762 s    10 runs
//
// Benchmark #2: ./go_arena.sh
//   Time (mean ± σ):      1.256 s ±  0.009 s    [User: 7.054 s, System: 0.008 s]
//   Range (min … max):    1.242 s …  1.270 s    10 runs
//
// Benchmark #3: ./rust.sh
//   Time (mean ± σ):     957.1 ms ±  25.1 ms    [User: 6.582 s, System: 0.073 s]
//   Range (min … max):   937.1 ms … 1023.1 ms    10 runs

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"runtime"
	"sort"
	"strconv"
	"time"
	"golang.org/x/sync/semaphore"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

type Message struct {
	Pos  uint32
	Text string
}

type ByPos []Message

func (m ByPos) Len() int           { return len(m) }
func (m ByPos) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByPos) Less(i, j int) bool { return m[i].Pos < m[j].Pos }

type TreeArena []Tree

func (a *TreeArena) Get() *Tree {
	last := len(*a) - 1
	node := &(*a)[last]
	*a = (*a)[:last]

	return node
}

func ArenaNew(depth uint32) TreeArena {
	i := 1 << (depth - 1)
	a := make([]Tree, i)
	return a
}

func itemCheck(tree *Tree) uint32 {
	if tree.Left != nil && tree.Right != nil {
		return uint32(1) + itemCheck(tree.Right) + itemCheck(tree.Left)
	}

	return 1
}

func bottomUpTree(a TreeArena, depth uint32) *Tree {
	tree := a.Get()
	if depth > uint32(0) {
		tree.Right = bottomUpTree(a, depth - 1)
		tree.Left = bottomUpTree(a, depth - 1)
	}
	return tree
}

func inner(a TreeArena, depth, iterations uint32) string {
	chk := uint32(0)
	for i := uint32(0); i < iterations; i++ {
		a := bottomUpTree(a, depth)
		chk += itemCheck(a)
	}
	return fmt.Sprintf("%d\t trees of depth %d\t check: %d",
		iterations, depth, chk)
}

const minDepth = uint32(4)

func main() {
	t1 := time.Now()
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	run(uint32(n))
	t2 := time.Now().Sub(t1)
	fmt.Println("t2:", t2)
}

func run(n uint32) {
	cpuCount := runtime.NumCPU()
	sem := semaphore.NewWeighted(int64(cpuCount))

	maxDepth := n
	if minDepth+2 > n {
		maxDepth = minDepth + 2
	}

	depth := maxDepth + 1

	messages := make(chan Message, cpuCount)
	expected := uint32(2) // initialize with the 2 summary messages

	go func() {
		// do stretch tree and longLivedTree
		if err := sem.Acquire(context.TODO(), 1); err == nil {
			go func() {
				defer sem.Release(1)
				a := ArenaNew(depth)
				tree := bottomUpTree(a, depth)
				messages <- Message{0,
					fmt.Sprintf("stretch tree of depth %d\t check: %d",
						depth, itemCheck(tree))}
			}()
		} else {
			panic(err)
		}

		if err := sem.Acquire(context.TODO(), 1); err == nil {
			go func() {
				defer sem.Release(1)
				a := ArenaNew(maxDepth)
				longLivedTree := bottomUpTree(a, maxDepth)
				messages <- Message{math.MaxUint32,
					fmt.Sprintf("long lived tree of depth %d\t check: %d",
						maxDepth, itemCheck(longLivedTree))}
			}()
		} else {
			panic(err)
		}

		for halfDepth := minDepth / 2; halfDepth < maxDepth/2+1; halfDepth++ {
			depth := halfDepth * 2
			iterations := uint32(1 << (maxDepth - depth + minDepth))
			expected++

			func(d, i uint32) {
				if err := sem.Acquire(context.TODO(), 1); err == nil {
					go func() {
						defer sem.Release(1)
						a := ArenaNew(depth)
						messages <- Message{d, inner(a, d, i)}
					}()
				} else {
					panic(err)
				}
			}(depth, iterations)
		}
	}()

	var sortedMsg []Message
	for m := range messages {
		sortedMsg = append(sortedMsg, m)
		expected--
		if expected == 0 {
			close(messages)
		}
	}

	sort.Sort(ByPos(sortedMsg))
	for _, m := range sortedMsg {
		fmt.Println(m.Text)
	}
}
