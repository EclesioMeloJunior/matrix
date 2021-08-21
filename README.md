## matrix

This repository contains the operations with matrixes. The current matrix is a `type Matrix [][]int64`
### Current features

- [x] Sum
- [] Substraction
- [x] Multiply
- [x] MultiplyByScalar
- [x] Transpose
- [x] IsSquared
- [x] Determinant
- [] LU Decomposition (work in progress)
- [x] Zero (creates a matrix R x C and all entries are 0)
- [x] RandInt (creates a matrix R x C where all entries are choices from 0 - 255)

### Other methods

- [x] Get (retrieve a value given a row and column)
- [x] Equals
- [x] Cols
- [x] Rows


### Tests

To run unit tests

```
make test or go test ./...
```

To run benchmark tests

```
make bench or go test ./... -bench=.
```