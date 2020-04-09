/*
 * @Author: your name
 * @Date: 2020-03-30 20:47:53
 * @LastEditTime: 2020-04-01 11:18:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/BiliiliTransform/BiliiliTransform.go
 */
package bv2av

// package main

import (
	"math"
	"strconv"
	"strings"
)

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var s = [6]int64{11, 10, 3, 8, 4, 6}

const xor int64 = 177451812
const add int64 = 8728348608

//规则为 叶轮<BVXXXXXXXXXX>

func deb2a(x string) (bv string) {

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
	bv = bv + "AV" + strconv.FormatInt(av, 10)
	return
}

func B2afunc(onMsg string) (outMsg string, exp bool) {

	getBvMsg := ""

	if strings.Index(onMsg, "BV") == 0 {
		if (len(onMsg) == 12) && (onMsg[1] == 'V') {
			getBvMsg = onMsg
		}
	}

	start := strings.Index(onMsg, "叶轮<")
	end := strings.Index(onMsg, ">")
	if (start != -1) && (end > start) {
		str0 := strings.SplitN(onMsg, "叶轮<", 2)
		str1 := strings.SplitN(str0[1], ">", 2)
		str2 := str1[0]
		if (len(str2) == 12) && (str2[1] == 'V') {

			getBvMsg = str2
		}
	}

	strlenght := len(getBvMsg)

	if strlenght == 0 {
		outMsg = ""
		exp = false
		return
	}

	if (strlenght == 12) && (!strings.ContainsAny(getBvMsg, "l | I | O | 0")) {
		outMsg = deb2a(getBvMsg)
		exp = true
		return
	} else {
		outMsg = "转换错误,BV号不规范,请区分\"i,o,0\"的大小写"
		exp = true
		return
	}
	// } else {
	// 	outMsg = ""
	// 	exp = false
	// 	return
	return
}
