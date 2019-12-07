/*
 * @Author: your name
 * @Date: 2019-10-29 22:09:18
 * @LastEditTime: 2019-11-19 20:35:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/app.go
 */
package main

import (
	"github.com/CanNev/Impeller_bot/roulette"
)
import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

import "strings"

import "strconv"
import "encoding/json"
import "io/ioutil"
import "github.com/Tnze/CoolQ-Golang-SDK/cqp/util"

//go:generate cqcfg -c .
// cqp: 名称: Impeller叶轮
// cqp: 版本: 1.0.0:0
// cqp: 作者: CanNev
// cqp: 简介: 基于coolQ的自律性机器人

//一些环境变量
var szm bool
var emj uint64
var anaNum bool

func main() { /*此处应当留空*/ }

func init() {
	szm = true
	anaNum = true
	emj = 128659
	cqp.AppID = "me.cqp.CanNev.demo" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onMsg
}

//试验场管理者身份确认,算法需要重写
func admid(infoQQ int64) (a bool) {
	switch infoQQ {
	case 1683941741:
		a = true
	case 1550843570:
		a = true
	case 287859992:
		a = true
	case 1174652322:
		a = true
	default:
		a = false
	}
	return
}

// func qjms(inMsg string) (newMsg string) {
// 	em := "[CQ:emoji,id={" + strconv.FormatUint(emj, 10) + "}]"
// 	newMsg = em + em + em + em + em + em + em + em + em + em + em + em + inMsg
// 	return
// }

func onMsg(subType, msgID int32, fromGroup int64, fromQQ int64, fromAnoymous, msg string, font int32) int32 {

	idcode := cqp.GetGroupMemberInfo2(fromGroup, fromQQ, false)
	Inf, _ := util.UnpackGroupMemberInfo(idcode)

	//拳交模块,防止被tx屏蔽
	em := "[CQ:emoji,id=" + strconv.FormatUint(emj, 10) + "]"
	if !szm {
		em = ""
	}

	emx := em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + "\n"
	if strings.Index(msg, "叶轮统计改") != -1 {
		anaNum = !anaNum
		cqp.SendGroupMsg(fromGroup, emx+"用语统计计数状态改变")
		if anaNum {
			cqp.SendGroupMsg(fromGroup, emx+"当前状态:语录统计开启")
		} else {
			cqp.SendGroupMsg(fromGroup, emx+"当前状态:语录统计关闭,统计数据已发送")
		}
	}

	if strings.HasPrefix(msg, "叶轮at测试") {
		var obqq int64
		obqq = 1683941741
		messAt := "[CQ:at,qq=" + strconv.FormatInt(obqq, 10) + "]"

		cqp.SendGroupMsg(fromGroup, messAt)
	}

	//语录统计
	if anaNum {
		handleJ(msg, "Z:/src/github.com/CanNev/Impeller_bot/ana.json", "Z:/src/github.com/CanNev/Impeller_bot/ana.json")
	}

	//俄罗斯轮盘赌

	go roulettePlay(fromGroup, fromQQ, msg, emx, Inf.Name)

	if strings.Index(msg, "叶轮统计信息") != -1 {

		aaa := yltj("Z:/src/github.com/CanNev/Impeller_bot/ana.json", "Z:/src/github.com/CanNev/Impeller_bot/ana.json")
		cqp.SendGroupMsg(fromGroup, emx+"生成JSON文件输出中")
		cqp.SendGroupMsg(fromGroup, aaa)
	}
	if strings.Index(msg, "叶轮拳交") != -1 {
		cqp.SendGroupMsg(fromGroup, emx+"拳交设置启动")
		if (strings.Index(msg, "改变")) != -1 {
			szm = !szm
			if szm {
				cqp.SendGroupMsg(fromGroup, emx+"拳交已经开启")
			} else {
				cqp.SendGroupMsg(fromGroup, emx+"拳交已经关闭")
			}
		}
		if (strings.Index(msg, "黑色高级车")) > 0 {
			emj = 128659
			cqp.SendGroupMsg(fromGroup, emx+"已经恢复黑色高级车")
		}
		if (strings.Index(msg, "UP")) > 0 {
			emj++
			cqp.SendGroupMsg(fromGroup, emx+"已经提升拳交等级")
		}
		if (strings.Index(msg, "DOWN")) > 0 {
			emj--
			cqp.SendGroupMsg(fromGroup, emx+"已经降低拳交等级")
		}
	}
	// if szm {
	// 	msg = "[CQ:emoji,id={128659}]"*14 + msg
	// }
	// if strings.HasSuffix(msg, "叶轮应答") {
	// 	aqw := "成员信息应答反馈\n" + "应答用户: " + Inf.Name + "\n" + "QQ号码: " + strconv.FormatInt(Inf.QQ, 10) + "\n"
	// 	cqp.SendGroupMsg(fromGroup, aqw)
	// }

	if strings.Index(msg, "叶轮状态") != -1 {
		sdu := "自律性应答bot-Impeller叶轮-正常工作中\n提交bug,变态粪亲父QQ:1683941741\n模式选项列表\n1.俄罗斯轮盘赌(new)\n还原了古典俄罗斯轮盘的的基本玩法,输入<叶轮轮盘帮助>获得游戏玩法,欢迎相互枪毙(激寒)\n\n2.企鹅五子棋企鹅扫雷移植中\n\n叶轮最后更新时间2019年11月16日"
		cqp.SendGroupMsg(fromGroup, emx+sdu)
	}
	if strings.Index(msg, "sudo ") != -1 {
		dd := admid(Inf.QQ)
		cqp.SendGroupMsg(fromGroup, emx+"检测到管理授权请求,确认身份中")
		if dd {
			cqp.SendGroupMsg(fromGroup, emx+"身份确认,<权限者>"+Inf.Name)
			if (strings.Index(msg, "-DS")) != -1 {
				aaa := roulette.DemonShot()

				cqp.SendGroupMsg(fromGroup, emx+aaa)
			}

			if (strings.Index(msg, "-a")) != -1 {
				cqp.SendGroupMsg(fromGroup, emx+"接收到权限指令")
				aaa := "功能设置已更新\n权限模式选项列表\n1.启动<恶魔镜头>  -DS\n警告,请确认是否有开启的必要"
				cqp.SendGroupMsg(fromGroup, aaa)
			}

		} else {
			cqp.SendGroupMsg(fromGroup, emx+"非<权限者>-"+Inf.Name+",访问拒绝")
		}

	}

	// if strings.Index(msg, "叶轮") != -1 {
	// 	if strings.Contains(msg, "水母祝寿") {
	// 		switch {
	// 		case strings.HasSuffix(msg, "01"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿筵开，画堂深映花如绣。瑞烟喷兽。帘幕香风透。一点台星，化作人间秀。韶音奏。两行红袖。齐劝长生酒。")
	// 		case strings.HasSuffix(msg, "02"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿2庄周浪说华封祝，汉帝虚传嵩岳声。争似寿宁嘉节日，千门万户愿长生。")
	// 		case strings.HasSuffix(msg, "03"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿3祝寿祝寿。筵开锦绣。拈起香来玉也似手。拈起盏来金也似酒。祝寿祝寿。 命比乾坤久。")
	// 		case strings.HasSuffix(msg, "04"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿4去年会祝夫人寿。今岁也又还依旧。鬓绿与颜朱。神仙想不如。 骨相真难老。疑是居蓬岛。")
	// 		case strings.HasSuffix(msg, "05"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿5九天宫上圣，降世共昭回。万汇须亭毓，群仙送下来。承乾当否极，庶事尽康哉。")
	// 		case strings.HasSuffix(msg, "06"):
	// 			cqp.SendGroupMsg(fromGroup, emx+"祝寿6帝命当敷佑，民生有厥初。千秋唐节日，万国禹朝车。韶美笙镛外，需亨饮食馀。神仙似姑射，梦想即华胥。")
	// 		default:
	// 		}
	// 		cqp.SendGroupMsg(fromGroup, emx+"水母万寿无疆")
	// 	}
	// }
	return 0
}

//关键词探知,通用模块
func handleJ(anaMsg string, jsonFile string, outFile string) error {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	// We have known the outer json object is a map, so we define  result as map.
	// otherwise, result could be defined as slice if outer is an array
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}

	// handle peers
	nodes := result["ana"].([]interface{})
	//anakeep,分析msg是否有关键词
	for _, anakeep := range nodes {
		anaKey := anakeep.(map[string]interface{})
		keyWord, _ := anaKey["name"]
		as := strings.Index(anaMsg, keyWord.(string))
		if as != -1 {
			a, _ := strconv.ParseUint(anaKey["time"].(string), 10, 32)
			a++
			anaKey["time"] = strconv.FormatUint(a, 10)
		}
	}
	// for _, node := range nodes {
	// 	m := node.(map[string]interface{})
	// 	if name, exists := m["name"]; exists {
	// 		if name == anaMsg {
	// 			a, _ := strconv.ParseUint(m["time"].(string), 10, 32)
	// 			a++
	// 			m["time"] = strconv.FormatUint(a, 10)
	// 			// time := strconv.FormatUint(m["time"] , 10)
	// 		}

	// 	}
	// }

	// Convert golang object back to byte
	byteValue, err = json.Marshal(result)
	if err != nil {
		return err
	}

	// Write back to file
	err = ioutil.WriteFile(outFile, byteValue, 0644)
	return err
}

func yltj(jsonFile string, outFile string) string {
	// Read json buffer from jsonFile
	byteValue, _ := ioutil.ReadFile(jsonFile)

	return string(byteValue)
}
func roulettePlay(fromGroup int64, fromQQ int64, msg string, emx string, InfName string) {
	if strings.HasPrefix(msg, "叶轮轮盘") {

		//潜在的bug,外界重置会打破正常游戏
		if strings.HasSuffix(msg, "重置") {

			mess := roulette.Reset()
			// fmt.Println(wer)
			cqp.SendGroupMsg(fromGroup, emx+mess)
		}

		if strings.HasSuffix(msg, "帮助") {
			text1 := roulette.Help()

			cqp.SendGroupMsg(fromGroup, emx+text1)
		}

		if roulette.GameInProgram {
			if strings.HasSuffix(msg, "开火") {

				messShoot, playerTime, obqq := roulette.Shoot(fromQQ)
				messAt := "[CQ:at,qq=" + strconv.FormatInt(obqq, 10) + "]"
				if playerTime {
					if roulette.Demon {
						var subQQ int64
						if (roulette.NowRound > 0) && (!roulette.PlayerOne.Life || !roulette.PlayerTwo.Life) {
							if !roulette.PlayerOne.Life {
								subQQ = roulette.PlayerOne.QQ
							}
							if !roulette.PlayerTwo.Life {
								subQQ = roulette.PlayerTwo.QQ
							}

							messShoot += "<恶魔镜头>已经启动\n执行对败者的处刑\n"
							cqp.SetGroupBan(fromGroup, subQQ, 30)
						}
					}
					cqp.SendGroupMsg(fromGroup, emx+messShoot+messAt)
				} else {
					cqp.SendGroupMsg(fromGroup, "[CQ:at,qq="+strconv.FormatInt(fromQQ, 10)+"]"+"恁并不是当局游戏玩家或不是这个回合,请等待游戏结束,或直接重置游戏干烂当局游戏?")
				}
			}
		} else {
			if (strings.HasSuffix(msg, "新游戏")) && (roulette.Room == 0) {
				mess := roulette.GameStart(fromQQ, InfName)
				// fmt.Println(wer)
				cqp.SendGroupMsg(fromGroup, emx+mess)
			}
			if strings.HasSuffix(msg, "-"+strconv.FormatInt(int64(roulette.Room), 10)) {
				mess := roulette.GameJoin(fromQQ, InfName)
				// fmt.Println(wer)
				cqp.SendGroupMsg(fromGroup, emx+mess)
				mess2 := roulette.SeqDisplay()
				cqp.SendGroupMsg(fromGroup, emx+mess2)
			}
		}
	}

}
