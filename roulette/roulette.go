/*
 * @Author: your name
 * @Date: 2019-11-10 20:48:56
 * @LastEditTime: 2019-11-17 10:50:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/roulette/roulette.go
 */
package roulette

import (
	"math/rand"
	"strconv"
	// "strings"
	"time"
)

type Player struct {
	QQ    int64
	Name  string
	Point int64
	Life  bool
}

func newPlayer(NP *Player, qq int64, name string, point int64, life bool) {
	NP.QQ = qq
	NP.Name = name
	NP.Point = point
	NP.Life = life
}

//地狱绘图的代码,因为没有深入地学习数据结构,后面需要重写

//先手状态,false为主场先手
var Offe bool

var PlayerOne Player
var PlayerTwo Player

//子弹仓
var Bullet [6]bool

//当前回合数
var NowRound int64

//房间号,线上环境会一直存在但在显示模拟情况下不可重复获取,而且对于多房间支持并不好
var Room int

var GameInProgram bool

//恶魔镜头
var Demon bool

//用于异常情况下游戏重置,不过我也不清楚是否有效
func shootInit() {
	GameInProgram = false
	NowRound = 0
	Room = 0
	Offe = false
}

//子弹随机装填
func bulletInit(ptr *[6]bool) {

	for i := 0; i < 6; i++ {
		ptr[i] = false
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(5)
	ptr[num] = true
}

//恶魔镜头开关
func DemonShot() (emm string) {
	Demon = !Demon
	if Demon {
		emm = "恶魔镜头已经启动\n叶轮的部分功能将发生改变\n"
	} else {
		emm = "恶魔镜头已关闭"
	}
	return
}

//游戏发起,qq和name穿入发起者,返回值为一串文字
func GameStart(qq int64, name string) (mess string) {
	shootInit()
	PlayerOne.Life = true
	PlayerTwo.Life = true
	NowRound = 0 //回合数清空
	newPlayer(&PlayerOne, qq, name, 0, true)
	rand.Seed(time.Now().UnixNano())
	for {
		Room = rand.Intn(10000)
		if Room > 1000 {
			break
		}
	}
	bulletInit(&Bullet)
	mess = "确认游戏发起者:" + PlayerOne.Name + "\n" +
		PlayerOne.Name + "拥有积分:" + strconv.FormatInt(PlayerOne.Point, 10) + "\n" +
		"当前游戏模式支持人数:" + "2" + "人\n" +
		// "唯一房间号码已经生成" + strconv.FormatInt(int64(Room), 10) + "\n"
		"请通过输入<叶轮轮盘 -" + strconv.FormatInt(int64(Room), 10) + "> 的形式加入房间\n" +
		"提醒:重复发起游戏将销毁之前创建的游戏房间\n" +
		"发起者请勿重复加入房间,否者将变成直播吃花生米(绝句)"
	return
}

//参与游戏,参数同发起者
func GameJoin(qq int64, name string) (mess string) {
	newPlayer(&PlayerTwo, qq, name, 0, true)
	mess = "玩家加入:" + PlayerTwo.Name + "\n" +
		PlayerTwo.Name + "拥有积分:" + strconv.FormatInt(PlayerOne.Point, 10) + "\n" +
		"左轮装填完成,游戏开始\n"
	seq()
	GameInProgram = true
	if Demon {
		mess += "请对准你的皮炎(绝句)\n提示:恶魔镜头(特殊功能)已经开启,\n请勿在试验场以外的地方继续游戏\n"

	}
	return
}

func Shoot(qq int64) (messShoot string, ispLayer bool, obqq int64) {
	// var onplayer1 *Player
	// var onplayer2 *Player

	// if !Offe {
	// 	onplayer1 = &PlayerOne
	// 	onplayer2 = &PlayerTwo
	// } else {
	// 	onplayer1 = &PlayerTwo
	// 	onplayer2 = &PlayerOne
	// }
	// if qq == onplayer1.QQ {
	// 	messShoot = trigger(onplayer1)
	// 	obqq = onplayer2.QQ
	// 	ispLayer = true
	// 	Offe = !Offe
	// 	return

	if (!Offe) && (qq == PlayerOne.QQ) {
		messShoot = trigger(&PlayerOne)
		obqq = PlayerTwo.QQ
		ispLayer = true
		Offe = !Offe

	} else if (Offe) && (qq == PlayerTwo.QQ) {
		messShoot = trigger(&PlayerTwo)
		obqq = PlayerOne.QQ
		ispLayer = true
		Offe = !Offe
	} else {
		ispLayer = false
	}
	return
}

//扳机函数,不过判断交给了外部函数,实际上存在大量需要改进的地方
func trigger(pl *Player) (mess string) {
	if GameInProgram && (NowRound < 6) {
		mess = "本轮的持者为" + pl.Name + "\n" + "发了,一射\n"
		if Bullet[NowRound] {
			pl.Life = false
			if NowRound == 0 {
				mess += "不幸的," + pl.Name + "第一枪就被打死,我们将永远缅怀他\n"
			}

			if NowRound == 5 {
				mess += "最后一发,毫无疑问的实弹,但规则亦为规则.\n开冲,请(无慈悲).\n"
			}
			// if NowRound == 5 {
			// 	PlayerOne.Life = false
			// 	PlayerTwo.Life = false
			// 	pl.Life = true
			// 	mess += "这是最后一发,你活到了最后一轮"
			// }
			Room = 0
			mess += pl.Name + "一枪打中皮炎,当场暴毙\n可喜可贺,可喜可贺(无慈悲)\n"
			mess += "游戏结束,获胜者为:" + lifeCon() + "\n在第 " + strconv.FormatInt(NowRound+1, 10) + " 轮决出了胜负\n"
			shootInit()
		} else {
			mess += "轮盘为空,对家请继续\n"
		}
		if NowRound < 6 {
			NowRound++
		} else {
			shootInit()
		}
	}
	return
}

func Reset() (mess string) {
	shootInit()
	mess = "游戏已经重置"
	return
}

func SeqDisplay() (mess string) {

	mess = "通过输入<叶轮轮盘开火>触发扳机\n"
	if !Offe {
		mess += "次序已经确定.\n主场(发起者)先行.\n祝你好运.\n"
	} else {

		mess += "次序已经确定.\n客场(加入者)先行.\n祝你好运.\n"
	}
	return
}

//次序决定
func seq() {
	rand.Seed(time.Now().UnixNano())
	aaa := rand.Intn(114)
	if aaa%2 == 0 {
		Offe = false
	} else {
		Offe = true
	}
}

//存活确认,以判断游戏是否需要继续或者执行死刑
func lifeCon() string {
	if PlayerOne.Life {
		return PlayerOne.Name
	} else {
		return PlayerTwo.Name
	}
}

func execute() (kq int64) {
	if !PlayerOne.Life {
		kq = PlayerOne.QQ
	}
	if !PlayerTwo.Life {
		kq = PlayerTwo.QQ
	}
	PlayerOne.QQ = -1
	PlayerTwo.QQ = -1
	return
}

func Help() (text string) {
	text = "欢迎进行俄罗斯轮盘\n" +
		"俄罗斯轮盘是经典赌命游戏,经典规则为在一把6发左轮塞入一枚*,然后双方交替对自己开枪,最后活下来为胜利的游戏(笑)\n" +
		"规则及命令描述:\n " +
		"1.使用命令<叶轮轮盘新游戏>开启一轮新的游戏,并会生成唯一的房间号码\n" +
		"2.使用房间号加入游戏,如命令<叶轮轮盘-1919>.\n" +
		"3.按随机分配的顺序进行开火,命令<叶轮轮盘开火>\n" +
		"4.每轮开火结束将at下一轮的开火者,直到游戏结束\n" +
		"若游戏异常,请输入<叶轮轮盘重置>重置" +
		"\n<恶魔镜头>相关\n" +
		"<恶魔镜头>开启将使叶轮和这个游戏具有真实伤害性(禁言和出群留学)\n" +
		"在叶轮拥有管理员权限和<恶魔镜头>开启状态显现,杀伤取决与<权限者>的设置" +
		"通过<sudo 叶轮状态>可确认权限状态\n"
	return
}
