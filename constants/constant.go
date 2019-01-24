// Go in action
// @jeffotoni
// 2019-01-16

package main

const (
	PATH        = "/myhome/app"
	KB     int  = 1 << (10 * iota) // 1Kb
	MB     int  = 1 << (10 * iota) // 1MB
	GB     int  = 1 << (10 * iota) // 1G
	MALE   bool = true
	FEMALE      = true
	DOLLAR      = 3.99
)

// EXPRESSIONS PERMITTED
const a = 2 + 3.0        // a == 5.0   (untyped floating-point constant)
const b = 15 / 4         // b == 3     (untyped integer constant)
const c = 15 / 4.0       // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3 / 2  // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3 / 2. // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0       // d == 8     (untyped integer constant)
const e = 1.0 << 3       // e == 8     (untyped integer constant)
//const f = int32(1) << 33 // illegal    (constant 8589934592 overflows int32)
//const g = float64(2) >> 1 // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar" // h == true  (untyped boolean constant)
const j = true          // j == true  (untyped boolean constant)
const k = 'w' + 1       // k == 'x'   (untyped rune constant)
const l = "hi"          // l == "hi"  (untyped string constant)
const m = string(k)     // m == "x"   (type string)

const Σ = 1 - 0.707i     //            (untyped complex constant)
const Δ = Σ + 2.0e-4     //            (untyped complex constant)
const Φ = iota*1i - 1/1i //            (untyped complex constant)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)

const Pi float64 = 3.14159265358979323846
const zero = 0.0 // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1 // untyped integer constant
)
const xa, xb, xc = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
const xu, xv float32 = 0, 3    // u = 0.0, v = 3.0

func main() {

	println("###############")
	println(Sunday)
	println(Monday)
	println(Tuesday)
	println(Wednesday)
	println(Thursday)
	println(Friday)
	println(Partyday)
	println(numberOfDays)
	println(PATH)
	println(KB)
	println(MB)
	println(GB)
	println(MALE)
	println(FEMALE)
	println(DOLLAR)
	println("###############")
	println("")
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(h)
	println(j)
	println(k)
	println(l)
	println(m)
	println(Σ)
	println(Δ)
	println(Φ)
	println("###############")
}
