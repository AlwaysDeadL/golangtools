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
边界值测试： fuzz 会基于给定的值，重新生成不同的输入组合作为新的参数进行测试
go test -fuzztime 10s -fuzz .

当输入的参数不符合预期时：会生成当前错误的场景参数， 如：
Failing input written to testdata/fuzz/FuzzDoDeclareSamples/301107ff18bbebcb
    To re-run:
    go test -run=FuzzDoDeclareSamples/301107ff18bbebcb
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
