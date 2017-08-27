package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	substr    = flag.String("str", ":", "指定の文字列")
	tabindent = flag.Bool("t", false, "タブインデントにする(色文字を整形する場合など)")
	last      = flag.Bool("l", false, "インデントを最後から検索する")
	stop      = flag.Bool("s", false, "表示ごとに停止する")
)

var (
	indexMethod = strings.Index
)

func init() {
	flag.Parse()
	if *last {
		indexMethod = strings.LastIndex
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		gp := ""       // グループ化したあとの出力結果
		prevpref := "" // substrでの検索結果を保存する
		for {
			b, _, err := r.ReadLine()
			if err != nil {
				fmt.Print(gp)
				return
			}
			line := string(b)
			// 行をまとめる為
			if i := indexMethod(line, *substr); i != -1 {
				pref := line[:i+1]
				if prevpref == "" {
					if !*stop {
						fmt.Println("\n" + pref)
					} else {
						// fmt.Print("\n" + pref)
						// s := ""
						// fmt.Scanln(&s)
					}
				}
				if prevpref != pref && prevpref != "" {
					prevpref = pref
					break
				}
				prevpref = pref
				indent := strings.Repeat(" ", i+1)
				if *tabindent {
					indent = "\t"
				}
				gp += indent + line[i+1:] + "\n"
			} else {
				gp += line + "\n"
				break
			}
		}
		fmt.Print(gp)
	}
}
