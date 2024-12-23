package enums

// Weekday 常量定义， 通过 type xx 声明一个类型，并使用常量生成器默认生成值 (默认从 0 开始, 后续自动 + 1)
type Weekday int
type Season int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)
