package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

// アプリのバージョン情報
// ビルド時に値をセットする
var Version string

const (
	Second = 1
	Minute = 60
	Hour   = 60 * Minute
)

// options はコマンドラインオプションを保持する。
type options struct {
	Version    func() `short:"v" long:"version" description:"version"`
	HourFlag   bool   `long:"hour" description:"hour flag"`
	MinuteFlag bool   `short:"m" long:"minute" description:"minute flag"`
	SecondFlag bool   `short:"s" long:"second" description:"second flag"`
}

func main() {
	// オプション引数解析
	var opts options
	opts.Version = func() {
		fmt.Println(Version)
		os.Exit(0)
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	// durationを計算するために最低２つは引数が必要
	// 無ければ終了
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Need 2 arguments. See -h. args=%v\n", args)
		os.Exit(1)
	}

	// durationの秒数値を計算
	ds := calcTimeDuration(args[0], args[1])

	// 時として出力
	if opts.HourFlag {
		d := secondToHour(ds)
		fmt.Println(d)
		return
	}

	// 分として出力
	if opts.MinuteFlag {
		d := secondToMinute(ds)
		fmt.Println(d)
		return
	}

	// 秒として出力
	if opts.SecondFlag {
		fmt.Println(ds)
		return
	}

	// オプション指定がなければ時として出力
	d := secondToHour(ds)
	fmt.Println(d)
}

// calcTimeDuration は２つの時刻の差を整数秒で返します。
func calcTimeDuration(st, et string) int {
	ss := timeToSecond(st)
	if ss < 0 {
		return -1
	}
	es := timeToSecond(et)
	if es < 0 {
		return -1
	}
	return es - ss
}

// timeToSecond は時刻文字列HH:MMを秒数で返します。
// 不正な文字列が渡された場合は-1を返します。
func timeToSecond(t string) int {
	ts := strings.Split(t, ":")
	if len(ts) < 2 {
		return -1
	}

	hstr := ts[0]
	mstr := ts[1]

	h, err := strconv.Atoi(hstr)
	if err != nil {
		return -1
	}
	m, err := strconv.Atoi(mstr)
	if err != nil {
		return -1
	}

	hsec := h * Hour
	msec := m * Minute
	return hsec + msec
}

// secondToMinute は秒数を分の少数で返す。
// 0未満の数値を渡したら-1を返す。
func secondToMinute(sec int) float64 {
	if sec < 0 {
		return -1
	}
	return float64(sec) / Minute
}

// secondToHour は秒数を時の少数で返す。
// 0未満の数値を渡したら-1を返す。
func secondToHour(sec int) float64 {
	if sec < 0 {
		return -1
	}
	return float64(sec) / Hour
}
