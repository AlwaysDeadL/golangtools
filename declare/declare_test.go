package declare

import "testing"


/**
一般功能测试: go test .
 */
func TestClare(t *testing.T) {
	DoDeclareSamples()
}

/**
基准测试: 测试方法名需要以 Benchmark开头  func BenchXXX(b *testing.B)
go test -bench .
 */
func BenchmarkDoDeclareSamples(b *testing.B) {
	for i :=0; i < b.N; i++ {
		DoAdd(1, 2)
	}
}

/**
FuzzDoDeclareSamples:
go test -fuzztime 10s -fuzz .
 */
func FuzzDoDeclareSamples(f *testing.F) {
	f.Add(1,2)
	f.Add(100, 100)
	f.Add(10, 10)
	f.Fuzz(func(t *testing.T, a int, b int) {
		result := DoAdd(a, b)
		if a <= 0 && b <= 0 && result > 0 {
			f.Errorf("DoAdd(%d, %d) = %d; not expected result", a, b, result)
		}
	})
}
