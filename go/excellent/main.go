package main // main パッケージを宣言

// EvenOrOdd 関数は、引数として受け取った整数が偶数か奇数かを判定します。
func EvenOrOdd(number int) string {
	// 数の剰余（余り）が 0 かどうかをチェックします。
	// number % 2 == 0 は、「number を 2 で割った余りが 0」であることを意味します。
	if number%2 == 0 {
		return "even" // 余りが 0 の場合（偶数）、"even" を返します
	} else {
		return "odd" // それ以外の場合（奇数）、"odd" を返します
	}
}

// 💡 実行例:
// func main() {
//     fmt.Println(EvenOrOdd(4)) // 出力: even
//     fmt.Println(EvenOrOdd(7)) // 出力: odd
// }