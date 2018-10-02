# tdur (Time Duration)

[![Build Status](https://travis-ci.org/jiro4989/tdur.svg?branch=master)](https://travis-ci.org/jiro4989/tdur)

時刻と時刻の差を計算するコマンド。

## 使い方

引数を２つ取って、その出力を秒、分、時で出力するだけ。  
デフォルトは時。

```bash
tdur 10:00 19:00
# -> 9

tdur -m 10:00 19:00
# -> 540

tdur -s 10:00 19:00
# -> 32400
```

## インストール

`go get github.com/jiro4989/tdur`

## ヘルプ

`tdur -h`

    Usage:
      tdur [OPTIONS]

    Application Options:
          --hour    hour flag
      -m, --minute  minute flag
      -s, --second  second flag

    Help Options:
      -h, --help    Show this help message


## 開発方法

```
make deps
make
```
