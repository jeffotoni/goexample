package main

import (
    //"strconv"
    "hash/crc32"
    "sync"
    "testing"
)

func BenchmarkStoreLoad(b *testing.B) {
    for N := 0; N < b.N; N++ {
        var m sync.Map
        for i := 0; i < 64*1024; i++ {
            for k := 0; k < 256; k++ {

                v := ChecksumIEEE(string(k))
                m.Store(v, k)
                var v1 interface{}
                a, ok := m.Load(v)
                if ok {
                    // ok foi gravado
                    v1 = a
                }
                _ = v1
            }
        }
    }
}

func BenchmarkStoreLoadInterface(b *testing.B) {
    for N := 0; N < b.N; N++ {
        var m sync.Map
        for i := 0; i < 64*1024; i++ {
            for k := 0; k < 256; k++ {
                var v1 interface{} = ChecksumIEEE(string(k))
                m.Store(v1, k)
                a, ok := m.Load(v1)
                if ok {
                    // ok foi gravado
                    v1 = a
                }
                _ = v1
            }
        }
    }
}

func BenchmarkLoadOrStore(b *testing.B) {
    for N := 0; N < b.N; N++ {
        var m sync.Map
        for i := 0; i < 64*1024; i++ {
            for k := 0; k < 256; k++ {
                var v1 interface{}
                v := ChecksumIEEE(string(k))
                a, loaded := m.LoadOrStore(v, k)
                if loaded {
                    v1 = a
                }
                _ = v1

            }
        }
    }
}

func BenchmarkLoadOrStoreInterface(b *testing.B) {
    for N := 0; N < b.N; N++ {
        var m sync.Map
        for i := 0; i < 64*1024; i++ {
            for k := 0; k < 256; k++ {
                var v1 interface{} = ChecksumIEEE(string(k))
                a, loaded := m.LoadOrStore(v1, k)
                if loaded {
                    v1 = a
                }
                _ = v1

            }
        }
    }
}

func BenchmarkAllZero(b *testing.B) {
    for N := 0; N < b.N; N++ {
        for i := 0; i < 64*1024; i++ {
            for k := 0; k < 256; k++ {
                var v1 int
                v1 = k
                _ = v1
            }
        }
    }
}

func ChecksumIEEE(str string) uint64 {
    return uint64(crc32.ChecksumIEEE([]byte(str)))
}
