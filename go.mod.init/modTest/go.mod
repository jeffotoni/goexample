module modTest

require mymath v0.0.0

require mymath2 v0.0.0

replace mymath => ./mymath

replace mymath2 => ./mymath2

go 1.16
