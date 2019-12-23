module project1

require (
	pkg/math v0.0.0
	pkg/util v0.0.0
)

replace (
	pkg/math => ./pkg/math
	pkg/util => ./pkg/util
)

go 1.12
