package main

import "fmt"

func main() {

	// 値を標準出力に出力
	// 最後の「\n」は改行
	fmt.Print("東京", "の降水確率は", 90, "%です。\n")

	// 値を標準出力に出力
	// 改行は自動出力されるが、値の間にスペースが出力される
	fmt.Println("東京", "の降水確率は", 90, "%です。")

	// 書式を使用して標準出力に出力
	// 書式は「値を文字列にフォーマットする」を参照
	fmt.Printf("%sの降水確率は%d%%です。\n", "東京", 90)
}