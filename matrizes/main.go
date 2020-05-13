package main

import (
	"errors"
	"runtime"
	"sync"
)

// Errors
var (
	ErrNumElements = errors.New("Error number of elements")
	ErrMatrixSize  = errors.New("Error size of matrix")
)

// Matrix is a 2d array
type Matrix struct {
	N    int
	data [][]float64
	mux  sync.Mutex
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		A = &Matrix{
			N: 8,
			data: [][]float64{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{9, 1, 2, 3, 4, 5, 6, 7},
				{8, 9, 1, 2, 3, 4, 5, 6},
				{7, 8, 9, 1, 2, 3, 4, 5},
				{6, 7, 8, 9, 1, 2, 3, 4},
				{5, 6, 7, 8, 9, 1, 2, 3},
				{4, 5, 6, 7, 8, 9, 1, 2},
				{3, 4, 5, 6, 7, 8, 9, 0},
			},
		}
		B = &Matrix{
			N: 8,
			data: [][]float64{
				{9, 8, 7, 6, 5, 4, 3, 2},
				{1, 9, 8, 7, 6, 5, 4, 3},
				{2, 1, 9, 8, 7, 6, 5, 4},
				{3, 2, 1, 9, 8, 7, 6, 5},
				{4, 3, 2, 1, 9, 8, 7, 6},
				{5, 4, 3, 2, 1, 9, 8, 7},
				{6, 5, 4, 3, 2, 1, 9, 8},
				{7, 6, 5, 4, 3, 2, 1, 0},
			},
		}

		C = &Matrix{
			N: 8,
			data: [][]float64{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		}
	)

	C.MultNaive(A, B)

}

// New a size by size matrix
func New(size int) func(...float64) (*Matrix, error) {
	wg := sync.WaitGroup{}
	d := make([][]float64, size)
	for i := range d {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			d[i] = make([]float64, size)
		}(i)
	}
	wg.Wait()
	m := &Matrix{N: size, data: d}
	return func(es ...float64) (*Matrix, error) {
		if len(es) != size*size {
			return nil, ErrNumElements
		}
		for i := range es {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				m.mux.Lock()
				m.data[i/size][i%size] = es[i]
				m.mux.Unlock()
			}(i)
		}
		wg.Wait()
		return m, nil
	}
}

// At access element (i, j)
func (A *Matrix) At(i, j int) float64 {
	return A.data[i][j]
}

// Set set element (i, j) with val
func (A *Matrix) Set(i, j int, val float64) {
	A.mux.Lock()
	A.data[i][j] = val
	A.mux.Unlock()
}

// MultNaive matrix multiplication O(n^3)
func (A *Matrix) MultNaive(B, C *Matrix) (err error) {
	var (
		i, j, k int
		sum     float64
		N       = A.N
	)

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			sum = 0.0
			for k = 0; k < N; k++ {
				sum += A.At(i, k) * B.At(k, j)
			}
			C.Set(i, j, sum)
		}
	}
	return
}

// ParalMultNaivePerRow matrix multiplication O(n^3) in concurrency per row
func (A *Matrix) ParalMultNaivePerRow(B, C *Matrix) (err error) {
	var N = A.N

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < N; j++ {
				sum := 0.0
				for k := 0; k < N; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Set(i, j, sum)
			}
		}(i)
	}
	wg.Wait()
	return
}

// ParalMultNaivePerElem matrix multiplication O(n^3) in concurrency per element
func (A *Matrix) ParalMultNaivePerElem(B, C *Matrix) (err error) {
	var N = A.N

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0.0
				for k := 0; k < N; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Set(i, j, sum)
			}(i, j)
		}
	}
	wg.Wait()
	return
}
