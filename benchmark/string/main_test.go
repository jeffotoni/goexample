package main_test

import (
    "fmt"
    "strconv"
    "strings"
    "testing"
)

var str, longStr string = "string_jeffotoni", `qwertyuiopqwertyuiopqwertyuio
qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop`

const cStr = "string_jeffotoni"

func BenchmarkPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "string_jeffotoni" + str
    }
}

func BenchmarkLongPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "string_jeffotoni" + longStr
    }
}

func BenchmarkConstPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "string_jeffotoni" + cStr
    }
}

func BenchmarkStr2Plus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "string_jeffotoni" + "jeffotoni" + strconv.Itoa(1919) + "quero somente testar como" + strconv.Itoa(32343) + " se fosse long   "
    }
}

func BenchmarkJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"string_jeffotoni%s", str}, "")
    }
}

func BenchmarkLongJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"string_jeffotoni%s", longStr}, "")
    }
}

func BenchmarkConstJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"string_jeffotoni%s", cStr}, "")
    }
}
func BenchmarkSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("string_jeffotoni%s", str)
    }
}

func BenchmarkLongSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("string_jeffotoni%s", longStr)
    }
}

func BenchmarkConstSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("string_jeffotoni%s", cStr)
    }
}

func BenchmarkBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("string_jeffotoni")
        b.WriteString(str)
    }
}

func BenchmarkLongBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("string_jeffotoni")
        b.WriteString(longStr)
    }
}

func BenchmarkConstBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("string_jeffotoni")
        b.WriteString(cStr)
    }
}

func BenchmarkConcatLongBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        Concat(longStr)
    }
}

func BenchmarkConcatStrBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        Concat(str)
    }
}
func BenchmarkConcatStr2Builder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        Concat("jeffotoni", 1919, "quero somente testar como", 23454, " se fosse long   ")
    }
}
func BenchmarkConcatConstBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        Concat(cStr)
    }
}

func Concat(strs ...interface{}) string {
    var sb strings.Builder
    for _, str := range strs {
        sb.WriteString(buildStr(str))
    }
    return sb.String()
}

func buildStr(str interface{}) string {

    switch str.(type) {
    case nil:
        return "nil"
    //
    case bool:
        return strconv.FormatBool(bool(str.(bool)))
    //
    case int:
        return strconv.Itoa(int(str.(int)))

    case []int:
        concat := ""
        for _, val := range str.([]int) {
            concat = Concat(concat, val)
        }
        return concat
    case uint:
        return strconv.FormatUint(uint64(str.(uint)), 10)
    case []uint:
        concat := ""
        for _, val := range str.([]uint) {
            concat = Concat(concat, val)
        }
        return concat
    case int8:
        return strconv.Itoa(int(str.(int8)))
    //provavelmente funciona para byte, pois é um aliais para uint8
    case uint8:
        return strconv.FormatUint(uint64(str.(uint8)), 10)
    case []int8:
        concat := ""
        for _, val := range str.([]int8) {
            concat = Concat(concat, val)
        }
        return concat

    case []uint8:
        concat := ""
        for _, val := range str.([]uint8) {
            concat = Concat(concat, val)
        }
        return concat

    case int16:
        return strconv.Itoa(int(str.(int16)))
    case uint16:
        return strconv.FormatUint(uint64(str.(uint16)), 10)
    case []int16:
        concat := ""
        for _, val := range str.([]int16) {
            concat = Concat(concat, val)
        }
        return concat

    case []uint16:
        concat := ""
        for _, val := range str.([]uint16) {
            concat = Concat(concat, val)
        }
        return concat

    //probably work for rune too, since rune is a alias for int32
    case int32:
        return strconv.FormatInt(int64(str.(int32)), 10)
    case uint32:
        return strconv.FormatUint(uint64(str.(uint32)), 10)
    case []int32:
        concat := ""
        for _, val := range str.([]int32) {
            concat = Concat(concat, val)
        }
        return concat
    case []uint32:
        concat := ""
        for _, val := range str.([]uint32) {
            concat = Concat(concat, val)
        }
        return concat
    case int64:
        return strconv.FormatInt(int64(str.(int64)), 10)
    case uint64:
        return strconv.FormatUint(uint64(str.(uint64)), 10)
    case []int64:
        concat := ""
        for _, val := range str.([]int64) {
            concat = Concat(concat, val)
        }
        return concat
    case []uint64:
        concat := ""
        for _, val := range str.([]uint64) {
            concat = Concat(concat, val)
        }
        return concat
    case string:
        return string(str.(string))
    case []string:
        concat := ""
        for _, val := range str.([]string) {
            concat = Concat(concat, val)
        }
        return concat
    case float64:
        return strconv.FormatFloat(str.(float64), 'f', 6, 64)
    case []float64:
        concat := ""
        for _, val := range str.([]float64) {
            concat = Concat(concat, val)
        }
        return concat
    case float32:
        return strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
    case []float32:
        concat := ""
        for _, val := range str.([]float32) {
            concat = Concat(concat, val)
        }
        return concat
    // case uintptr:
    //     return string(str.(uintptr))
    // case []uintptr:
    //     concat := ""
    //     for _, val := range str.([]uintptr) {
    //         concat = Concat(concat, val)
    //     }
    //     return concat
    case complex64:
        return "not suport complex 64"
    case complex128:
        return "not suport complex 128"
    default:
        println("type no encontrado...")
        break
    }
    return ""
}
