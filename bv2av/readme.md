<!--
 * @Author: your name
 * @Date: 2020-03-30 22:02:54
 * @LastEditTime: 2020-04-01 16:12:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/bv2av/readme.md
 -->

# [杂物]Bilibili 视频 BV 号与 AV 号的转换

## 概述

- 某种意义上 AV 号时代的结束,
- 一个简易的 BV 转 AV 的程序,
- 由 golang 编写,
- 随便用,用完点个收藏,没问题的

## 核心代码

```golong

package bv2av

import (
	"math"
	"strconv"
)

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var s = [6]int64{11, 10, 3, 8, 4, 6}

const xor int64 = 177451812
const add int64 = 8728348608

func Deb2a(x string) string {

	var r int64
	var av int64
	tr := make(map[int64]int64)
	i := 0
	for i = 0; i < 58; i++ {
		tr[int64(table[i])] = int64(i)
	}

	for i = 0; i < 6; i++ {
		r += tr[int64(x[s[i]])] * int64(math.Pow(float64(58), float64(i)))
	}
	av = (r - add) ^ xor
	return "AV" + strconv.FormatInt(av, 10)
    }

```

## 只言片语

- 现在已经加载到个人的 QQbot"Impeller 叶轮"上
- 支持两种查询指令,私信和群聊皆可以使用
- 苦于没有服务器,佛系上线


