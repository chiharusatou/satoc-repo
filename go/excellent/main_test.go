package main

import "testing" // 標準のテストライブラリをインポート

// TestEvenOrOdd 関数は、EvenOrOdd 関数のテストケースを定義します。
// Goのテスト関数は、ファイル名が *_test.go であり、関数名が TestXxx の形式でなければなりません。
func TestEvenOrOdd(t *testing.T) {
	// 1. 関数の実行（テスト対象）
	result := EvenOrOdd(10)

	// 2. 結果の検証（アサーション）
	// 期待される結果が "even" であるかを確認
	if result != "even" {
		// 期待した結果と異なっていた場合、エラーを報告
		t.Errorf("expected: even, actual: %s", result)
	}
}