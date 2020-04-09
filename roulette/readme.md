<!--
 * @Author: your name
 * @Date: 2020-04-01 11:08:48
 * @LastEditTime: 2020-04-01 11:12:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/roulette/readme.md
 -->

# TDN 轮盘赌(单线程)

## 介于只能完成单线程工作,需要重新组织,已经将部分代码移到此文件下

## 业务代码

```golang
	// 俄罗斯轮盘赌

        roulettePlay(fromGroup, fromQQ, msg, emx, Inf.Name)


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

```
