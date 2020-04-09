/*
 * @Author: your name
 * @Date: 2019-10-29 22:09:18
 * @LastEditTime: 2020-04-08 01:39:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/app.go
 */
package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"impeller_bot/bv2av"
	"time"

	// import "reflect"
	"strings"

	"encoding/json"
	"io/ioutil"
	"strconv"
)

// import "github.com/CanNev/Impeller_bot_pkg/bv2av"

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
	// anaNum = true
	emj = 128659
	cqp.AppID = "me.cqp.CanNev.demo" // TODO: 修改为这个插件的ID

	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onMsg
	cqp.GroupRequest = onGroupRequest
	cqp.FriendRequest = onFriendRequest
	cqp.Enable = onEnable
	// cqp.Exit = onExit

}

func onEnable() int32 {
	// groupList := cqp.GetGroupList()
	// lenList := len(groupList)
	// i := 0
	emg := "欧芒果(启动音)\n"
	// for i = 0; i < lenList; i++ {
	// 	time.Sleep(time.Second * 1)
	// 	cqp.SendGroupMsg(groupList[i].ID, emg)
	// }
	cqp.SendPrivateMsg(1683941741, emg)
	return 0
}

// func onExit() int32 {
// 	groupList := cqp.GetGroupList()
// 	lenList := len(groupList)
// 	i := 0
// 	emg := "奇酷比感激酱(关机音)"
// 	for i = 0; i < lenList; i++ {
// 		time.Sleep(time.Second * 1)
// 		cqp.SendGroupMsg(groupList[i].ID, emg)
// 	}
// 	cqp.SendPrivateMsg(1683941741, emg)
// 	return 0
// }

//试验场管理者身份确认,算法需要重写
func admid(infoQQ int64) (a bool) {
	switch infoQQ {
	case 1683941741: //应天
		a = true
	case 1550843570: //水母
		a = true
	case 287859992: //荧曈
		a = true
	case 1174652322: //拉菲
		a = true

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

func onFriendRequest(subType int32, sendTime int32, fromQQ int64, msg string, responseFlag string) int32 {
	return cqp.SetFriendAddRequest(responseFlag, 1, "")
}

func onGroupRequest(subType int32, sendTime int32, fromGroup int64, fromQQ int64, msg string, responseFlag string) int32 {
	// cqp.SendPrivateMsg(1683941741, strconv.FormatInt(int64(subType), 10))
	// cqp.SendPrivateMsg(1683941741, strconv.FormatInt(int64(sendTime), 10))
	// cqp.SendPrivateMsg(1683941741, strconv.FormatInt(int64(fromGroup), 10))
	// cqp.SendPrivateMsg(1683941741, strconv.FormatInt(int64(fromQQ), 10))
	// cqp.SendPrivateMsg(1683941741, msg)
	// cqp.SendPrivateMsg(1683941741, responseFlag)
	// cqp.SendPrivateMsg(1683941741, "返回值:"+strconv.FormatInt(int64(qa), 10))
	return cqp.SetGroupAddRequest(responseFlag, 2, 1, "叶轮加入完成,使用方法请输入\"叶轮帮助\"")
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {

	reMsg := ""
	getMsg, exp := bv2av.B2afunc(msg)
	reMsg += getMsg
	if exp {
		cqp.SendPrivateMsg(fromQQ, reMsg)
	}

	return 0
}

func onMsg(subType, msgID int32, fromGroup int64, fromQQ int64, fromAnoymous, msg string, font int32) int32 {

	//*******************************************//
	if fromQQ == 80000000 {

		return -1
	}
	//********************************************//

	//BV号转换功能
	reMsg := ""
	getMsg, exp := bv2av.B2afunc(msg)
	reMsg += getMsg
	if exp {

		cqp.SendGroupMsg(fromGroup, reMsg)
	}

	Inf := cqp.GetGroupMemberInfo(fromGroup, fromQQ, false)

	//拳交模块,防止被tx屏蔽
	// em := "[CQ:emoji,id=" + strconv.FormatUint(emj, 10) + "]"
	// if !szm {
	// 	em = ""
	// }

	// emx := em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + em + "\n"
	emx := "  "
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

	// //语录统计
	// if anaNum {
	// 	handleJ(msg, "Z:/src/github.com/CanNev/Impeller_bot/ana.json", "Z:/src/github.com/CanNev/Impeller_bot/ana.json")
	// }

	if strings.Index(msg, "叶轮") != -1 {

		//愚人节的小玩笑,群主有权利一键爆破群聊(迫真)
		jokerOwner := cqp.GetGroupMemberInfo(fromGroup, fromQQ, false)
		jokerOwnerLevel := jokerOwner.Auth
		jokerOwnerName := jokerOwner.Name
		if jokerOwnerLevel == 3 {
			if strings.Index(msg, "权限确认") != -1 {
				cqp.SendGroupMsg(fromGroup, "	检测到权限请求,身份确认中...")
				cqp.SendGroupMsg(fromGroup, "	身份确认,<群主>"+jokerOwnerName)
			}
			if strings.Index(msg, "群聊爆破") != -1 {
				cqp.SendGroupMsg(fromGroup, "	已确认来自<群主>的爆破请求,现在将开始倒计时,请私信取消指令以解除爆破请求")
				var i int64
				for i = 10; i > 0; i-- {
					time.Sleep(time.Second * 2)
					cqp.SendGroupMsg(fromGroup, "	"+strconv.FormatInt(i, 10)+"~~~")
				}
				cqp.SendGroupMsg(fromGroup, "  爆破请求已确认,群聊将于<24小时>内封禁,我们下个群聊再见")
			}
		}
		// if strings.Index(msg, "叶轮统计信息") != -1 {

		// 	aaa := yltj("Z:/src/github.com/CanNev/Impeller_bot/ana.json", "Z:/src/github.com/CanNev/Impeller_bot/ana.json")
		// 	cqp.SendGroupMsg(fromGroup, emx+"生成JSON文件输出中")
		// 	cqp.SendGroupMsg(fromGroup, aaa)
		// }
		// if strings.Index(msg, "叶轮拳交") != -1 {
		// 	cqp.SendGroupMsg(fromGroup, emx+"拳交设置启动")
		// 	if (strings.Index(msg, "改变")) != -1 {
		// 		szm = !szm
		// 		if szm {
		// 			cqp.SendGroupMsg(fromGroup, emx+"拳交已经开启")
		// 		} else {
		// 			cqp.SendGroupMsg(fromGroup, emx+"拳交已经关闭")
		// 		}
		// 	}
		// 	if (strings.Index(msg, "黑色高级车")) > 0 {
		// 		emj = 128659
		// 		cqp.SendGroupMsg(fromGroup, emx+"已经恢复黑色高级车")
		// 	}
		// 	if (strings.Index(msg, "UP")) > 0 {
		// 		emj++
		// 		cqp.SendGroupMsg(fromGroup, emx+"已经提升拳交等级")
		// 	}
		// 	if (strings.Index(msg, "DOWN")) > 0 {
		// 		emj--
		// 		cqp.SendGroupMsg(fromGroup, emx+"已经降低拳交等级")
		// 	}
		// }

		if strings.Index(msg, "叶轮状态") != -1 {
			sdu := "bot-Impeller叶轮(重建版）\n\n输入\"叶轮帮助\"获取功能列表\n\n提交bug,变态粪亲父QQ:1683941741\n\n最后更新时间2020年4月1日"
			cqp.SendGroupMsg(fromGroup, sdu)
		}
		if strings.Index(msg, "叶轮帮助") != -1 {
			sdu := "反馈BUG,变态粪亲父QQ:1683941741\n  功能列表\n1.叶轮群员信息备份\n  >群管理员可激活\n  >名称即指令\n\n2.群聊重建工具\n  >开发中\n\n3.BV号转AV号工具\n  >指令\n  \"叶轮<BVXXXXXXXXXX>\"\n或\n  BVXXX(直接粘贴BV号)\n  >请粘贴通过长按稿件获取的BV号\n  根据屑站规则,BV号为\"包含BV在内的12个字符的字符串\",且不存在l | I | O | 0字符,长度不符合规范的BV将不做应答\n4.叶轮群聊爆破  >字面意思\n  >愚人节玩笑,群主可触发"
			cqp.SendGroupMsg(fromGroup, sdu)
		}

		/************************************************************/
		//需要封装
		if strings.Index(msg, "叶轮群员信息备份") != -1 {

			power := false
			QQTeamRecover := cqp.GetGroupMemberList(fromGroup)
			msgarg := cqp.GetGroupMemberInfo(fromGroup, fromQQ, false)

			power = admid(fromQQ)
			if power {
			} else {
				if msgarg.Auth > 1 {
					cqp.SendGroupMsg(fromGroup, "确认管理员身份，推送开始")
					power = true
				}
			}

			if power {
				cqp.SendPrivateMsg(fromQQ, "群"+strconv.FormatInt(fromGroup, 10)+"群员统计中")

				changdu := len(QQTeamRecover) //迫真汉语拼音，因为叫length的变量太多了
				cqp.SendPrivateMsg(fromQQ, "统计完成，该群群人数"+strconv.FormatInt(int64(changdu), 10)+"人")

				covermsg0 := "当前群员列表 \n{\n"
				covermsg1 := "(续上）\n{\n"
				covermsg2 := "(续上）\n{\n"
				covermsg3 := "(续上）\n{\n"
				i := 0

				if changdu < 50 {

					for i = 0; i < changdu; i++ {
						covermsg0 = covermsg0 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}
					cqp.SendPrivateMsg(fromQQ, covermsg0)
				} else if changdu < 100 {

					for i = 0; i < 50; i++ {
						covermsg0 = covermsg0 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}
					for i = 50; i < changdu; i++ {
						covermsg1 = covermsg1 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}
					cqp.SendPrivateMsg(fromQQ, covermsg0)
					cqp.SendPrivateMsg(fromQQ, covermsg1)
				} else if changdu < 150 {

					for i = 0; i < 50; i++ {
						covermsg0 = covermsg0 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					for i = 50; i < 100; i++ {
						covermsg1 = covermsg1 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					for i = 100; i < changdu; i++ {
						covermsg2 = covermsg2 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}
					cqp.SendPrivateMsg(fromQQ, covermsg0)
					cqp.SendPrivateMsg(fromQQ, covermsg1)
					cqp.SendPrivateMsg(fromQQ, covermsg2)
				} else if changdu < 300 {
					for i = 0; i < 50; i++ {
						covermsg0 = covermsg0 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					for i = 50; i < 100; i++ {
						covermsg1 = covermsg1 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					for i = 100; i < 150; i++ {
						covermsg2 = covermsg2 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					for i = 150; i < changdu; i++ {
						covermsg3 = covermsg3 + "\"yj" + strconv.FormatInt(int64(i), 10) + "\": [\n" + "    \"Name\": \"" + QQTeamRecover[i].Name + "\",\n    \"QQ\": \"" + strconv.FormatInt(QQTeamRecover[i].QQ, 10) + "\"\n],\n"
					}

					cqp.SendPrivateMsg(fromQQ, covermsg0)
					cqp.SendPrivateMsg(fromQQ, covermsg1)
					cqp.SendPrivateMsg(fromQQ, covermsg2)
					cqp.SendPrivateMsg(fromQQ, covermsg3)

				} else {
					covermsg0 = covermsg0 + "\n群员超过处理上线，当前最大支持300人"
					cqp.SendPrivateMsg(fromQQ, covermsg0)
				}
			}
		}
		/************************************************************/

		if strings.Index(msg, "叶轮权限确认") != -1 {
			dd := admid(fromQQ)
			cqp.SendGroupMsg(fromGroup, emx+"	检测到权限请求,身份确认中...")
			if dd {
				cqp.SendGroupMsg(fromGroup, emx+"身份确认,<权限者>"+Inf.Name)
				// if (strings.Index(msg, "-DS")) != -1 {
				// 	aaa := roulette.DemonShot()

				// 	cqp.SendGroupMsg(fromGroup, emx+aaa)
				// }

				if (strings.Index(msg, "-a")) != -1 {
					cqp.SendGroupMsg(fromGroup, emx+"接收到权限指令")
					aaa := "功能设置已更新\n权限模式选项列表\n1.启动<恶魔镜头>  -DS\n警告,请确认是否有开启的必要"
					cqp.SendGroupMsg(fromGroup, aaa)
				}

			} else {
				cqp.SendGroupMsg(fromGroup, emx+"非<权限者>-"+Inf.Name+",访问拒绝")
			}

		}
	}
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
