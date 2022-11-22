package a

import "fmt"

//在上面程序的第 8 行，创建了一个名为 VowelsFinder 的接口，该接口有一个 FindVowels() []rune 的方法。
//在接下来的一行，我们创建了一个 MyString 类型。
//在第 15 行，我们给接受者类型（Receiver Type） MyString 添加了方法 FindVowels() []rune。现在，我们称 MyString 实现了 VowelsFinder 接口。
// 这就和其他语言（如 Java）很不同，其他一些语言要求一个类使用 implement 关键字，来显式地声明该类实现了接口。而在 Go 中，并不需要这样。
// 如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口。
//在第 28 行，v 的类型为 VowelsFinder，name 的类型为 MyString，我们把 name 赋值给了 v。由于 MyString 实现了 VowelFinder，因此这是合法的。
// 在下一行，v.FindVowels() 调用了 MyString 类型的 FindVowels 方法，打印字符串 Sam Anderson 里所有的元音。该程序输出 Vowels are [a e o]。

// 定义接口
type VowelsFinder interface {
	FindVowels() []rune
}

// 定义string别名
type MyString string

// 给接受者类型（Receiver Type） MyString 添加了方法 FindVowels() []rune
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	//把 name 赋值给了 v。由于 MyString 实现了 VowelFinder，因此这是合法的。
	name := MyString("Sam Anderson") // 实现了接口的 string
	var v VowelsFinder               //可以直接先声明接口
	v = name                         // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

}
