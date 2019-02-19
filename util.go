package main

import (
	"io/ioutil"
)

func readFile(path string) string {
	// ファイルを読み出し用にオープン
	bytes, err := ioutil.ReadFile(path);
	if err != nil {
		panic(err)
	}

	//もしかしたら改行はCRLFだと動かないのかもしれんが知らん

	//newProgram時に渡すstringはゼロ終端じゃないと
	//動いたり動かなかったり
	//しかも動かないときは存在しないsyntaxエラーを返す
	return string(bytes) + "\x00"
}
