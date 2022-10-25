package main_test

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
    "testing"
)

type Book struct {
    Title    string   `json:"title"`
    Author   string   `json:"author"`
    Pages    int      `json:"num_pages"`
    Chapters []string `json:"chapters"`
}

type BookDef struct {
    Title  string `msg:"title"`
    Author string `msg:"author"`
    Pages  int    `msg:"num_pages"`
}

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

func BenchmarkJSONMarshal(b *testing.B) {
    obj := generateObject()
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        _, err := json.Marshal(obj)
        if err != nil {
            panic(err)
        }
    }
}

func BenchmarkJSONConcat(b *testing.B) {
    //b.ResetTimer()
    for n := 0; n < b.N; n++ {
        // str := "Computação quântica V.5 vamos ver a quantidade de caracter muito mais muitooooooooooooooooooooooo grade aqui heeee..." +
        //     "Jefferson Otoni Lima" + strconv.Itoa(1650) +
        //     strings.Join([]string{"Escala atômica,", "Arithmetic das partículas subatômicas", "vamos testar mais posicoes hereeeeeeeeeeee."}, ",")
        //fmt.Sprintf("%s", str)
        strL := generateString()
        if len(strL) <= 0 {
            panic(errors.New("Error generateString"))
        }
    }
}

func BenchmarkJSONConcatNaMao(b *testing.B) {
    //b.ResetTimer()
    for n := 0; n < b.N; n++ {
        str := "Computação quântica V.5 vamos ver a quantidade de caracter muito mais muitooooooooooooooooooooooo grade aqui heeee..." +
            "Jefferson Otoni Lima" + strconv.Itoa(1650) +
            strings.Join([]string{"Escala atômica,", "Arithmetic das partículas subatômicas", "vamos testar mais posicoes hereeeeeeeeeeee."}, ",")
        fmt.Sprintf("%s", str)
    }
}

func BenchmarkIntToString1(b *testing.B) {
    for n := 0; n < b.N; n++ {
        IntToString1([]int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8, 8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77, 8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7})
    }
}

func BenchmarkIntToString2(b *testing.B) {
    for n := 0; n < b.N; n++ {
        IntToString2([]int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8, 8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77, 8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7})
    }
}

func BenchmarkIntToString3(b *testing.B) {
    for n := 0; n < b.N; n++ {
        IntToString3([]int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8, 8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77, 8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7})
    }
}

func IntToString3(a []int) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func IntToString1(a []int) string {
    b := ""
    for _, v := range a {
        if len(b) > 0 {
            b += ","
        }
        b += strconv.Itoa(v)
    }

    return b
}

func IntToString2(a []int) string {
    b := make([]string, len(a))
    for i, v := range a {
        b[i] = strconv.Itoa(v)
    }

    return strings.Join(b, ",")
}

func generateString() string {
    return Concat("Computação quântica V.5 vamos ver a quantidade de caracter muito mais muitooooooooooooooooooooooo grade aqui heeee...",
        "Jefferson Otoni Lima",
        1650,
        []string{"Escala atômica,", "Arithmetic das partículas subatômicas", "vamos testar mais posicoes hereeeeeeeeeeee."},
    )
}

func generateObject() *Book {
    return &Book{
        Title:    "Computação quântica V.5 vamos ver a quantidade de caracter muito mais muitooooooooooooooooooooooo grade aqui heeee...",
        Author:   "Jefferson Otoni Lima",
        Pages:    1650,
        Chapters: []string{"Escala atômica,", "Arithmetic das partículas subatômicas", "vamos testar mais posicoes hereeeeeeeeeeee."},
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
        // concat := ""
        // for _, val := range str.([]int) {
        //     concat = Concat(concat, val)
        // }
        return IntToString2(str.([]int))
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
        return strings.Join(str.([]string), " ")
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
