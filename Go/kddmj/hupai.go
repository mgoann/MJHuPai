package kddmj

import (
	"fmt"
)

var tableMgr *TableMgr

func init() {
	tableMgr = NewTableMgr()
	tableMgr.Load("E:\\Go\\GOPATH\\src\\github.com\\MDGSF\\MJHuPai\\Go\\kddmj\\genTable\\")
}

/*
CanHuWithLaiZi 赖子胡牌
handCards: 手牌数组，最多14张。
laizi: 赖子数组，保存可能的赖子。
return: true可以胡牌，false不可以胡牌。
return int: 点数。
*/
func CanHuWithLaiZi(handCards []int, laizi []int) (bool, int) {

	if !IsValidHandCards(handCards) {
		return false, 0
	}

	var bHasCommonCard = false

	laiziNum := 0
	var slots [TILEMAX]int
	for _, c := range handCards {
		if isLaizi(c, laizi) {
			laiziNum++
		} else {
			bHasCommonCard = true
			slots[c]++
		}
	}

	if !bHasCommonCard && laiziNum > 0 {
		//手牌全部都是赖子。
		return true, 10
	}

	XuShu, Zi := getArray(slots)

	ret1, dianshu1 := walkThroughTable(laiziNum, XuShu, Zi, tableMgr.TableXuShuWithEye, tableMgr.TableXuShu, tableMgr.TableZi)

	ret2, dianshu2 := walkThroughTable(laiziNum, Zi, XuShu, tableMgr.TableZiWithEye, tableMgr.TableZi, tableMgr.TableXuShu)

	if ret1 && ret2 {
		return true, max(dianshu1, dianshu2)
	} else if !ret1 && ret2 {
		return true, dianshu2
	} else if ret1 && !ret2 {
		return true, dianshu1
	}

	return false, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getArray(slots [TILEMAX]int) (XuShu []int, Zi []int) {
	if getWan(slots) > 0 {
		XuShu = append(XuShu, getWan(slots))
	}
	if getTiao(slots) > 0 {
		XuShu = append(XuShu, getTiao(slots))
	}
	if getTong(slots) > 0 {
		XuShu = append(XuShu, getTong(slots))
	}

	if getFeng(slots) > 0 {
		Zi = append(Zi, getFeng(slots))
	}
	if getJian(slots) > 0 {
		Zi = append(Zi, getJian(slots))
	}

	return
}

func walkThroughTable(laiziNum int, jiangArray []int, commonArray []int,
	jTableWithEye *Table, jTableNoEye *Table, cTableNoEye *Table) (bool, int) {

	//保存最大的点数
	var maxDianShu = 0

	//保存当前这次计算时最大的点数
	var curMaxDianShu = 0

	var bCanHu = false

	for i, iNum := range jiangArray {
		success := true
		hasLaiZiNum := laiziNum

		//这里会把8个赖子全部遍历，但是没有必要，因为也许我一共也只有3个赖子。
		dianShu, jiangNeedLaiZiNum, ok := jTableWithEye.IsInTable(iNum)
		if !ok || jiangNeedLaiZiNum > hasLaiZiNum {
			continue
		}
		hasLaiZiNum -= jiangNeedLaiZiNum
		if jiangNeedLaiZiNum > 0 {
			curMaxDianShu = dianShu
		}

		for j, jNum := range jiangArray {
			if i == j {
				continue
			}
			dianShu, needLaiZiNum, ok := jTableNoEye.IsInTable(jNum)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
			hasLaiZiNum -= needLaiZiNum
			if needLaiZiNum > 0 && dianShu > curMaxDianShu {
				curMaxDianShu = dianShu
			}
		}
		if !success {
			continue
		}

		for _, num := range commonArray {
			dianShu, needLaiZiNum, ok := cTableNoEye.IsInTable(num)
			if !ok || needLaiZiNum > hasLaiZiNum {
				success = false
				break
			}
			hasLaiZiNum -= needLaiZiNum
			if needLaiZiNum > 0 && dianShu > curMaxDianShu {
				curMaxDianShu = dianShu
			}
		}
		if !success {
			continue
		}

		if hasLaiZiNum >= 3 || curMaxDianShu == 10 {
			return true, 10
		}

		if jiangNeedLaiZiNum == 2 {
			_, ok := jTableNoEye.IsInTableMap(iNum, 0)
			if ok {
				return true, 10
			}
		} else if jiangNeedLaiZiNum == 3 {
			_, ok := jTableNoEye.IsInTableMap(iNum, 1)
			if ok {
				return true, 10
			}
		} else if jiangNeedLaiZiNum == 4 {
			_, ok := jTableNoEye.IsInTableMap(iNum, 2)
			if ok {
				return true, 10
			}
		}

		if curMaxDianShu > maxDianShu {
			maxDianShu = curMaxDianShu
		}
		bCanHu = true
	}

	if bCanHu {
		return true, maxDianShu
	}
	return false, 0
}

func isLaizi(c int, LaiZiArr []int) bool {
	for _, laizi := range LaiZiArr {
		if c == laizi {
			return true
		}
	}
	return false
}

/*
CanHu 胡牌
handCards: 手牌数组，最多14张。
return: true可以胡牌，false不可以胡牌。
*/
func CanHu(handCards []int) bool {

	if !IsValidHandCards(handCards) {
		return false
	}

	slots := GenSlots(handCards)

	// fmt.Println("getWan = ", getWan(slots))
	// fmt.Println("getTiao = ", getTiao(slots))
	// fmt.Println("getTong = ", getTong(slots))
	// fmt.Println("getFeng = ", getFeng(slots))
	// fmt.Println("getJian = ", getJian(slots))

	wan := getWan(slots)
	tiao := getTiao(slots)
	tong := getTong(slots)
	feng := getFeng(slots)
	jian := getJian(slots)

	var XuShu []int
	if wan > 0 {
		XuShu = append(XuShu, wan)
	}
	if tiao > 0 {
		XuShu = append(XuShu, tiao)
	}
	if tong > 0 {
		XuShu = append(XuShu, tong)
	}

	var Zi []int
	if feng > 0 {
		Zi = append(Zi, feng)
	}
	if jian > 0 {
		Zi = append(Zi, jian)
	}

	for i, iNum := range XuShu {
		_, ok := tableMgr.TableXuShuWithEye.IsInTableMap(iNum, 0)
		if !ok {
			continue
		}

		for j, jNum := range XuShu {
			if i == j {
				continue
			}

			_, ok := tableMgr.TableXuShu.IsInTableMap(jNum, 0)
			if !ok {
				return false
			}
		}

		for _, num := range Zi {
			_, ok := tableMgr.TableZi.IsInTableMap(num, 0)
			if !ok {
				return false
			}
		}

		return true
	}

	for i, iNum := range Zi {
		_, ok := tableMgr.TableZiWithEye.IsInTableMap(iNum, 0)
		if !ok {
			continue
		}

		for j, jNum := range Zi {
			if i == j {
				continue
			}

			_, ok := tableMgr.TableZi.IsInTableMap(jNum, 0)
			if !ok {
				return false
			}
		}

		for _, num := range XuShu {
			_, ok := tableMgr.TableXuShu.IsInTableMap(num, 0)
			if !ok {
				return false
			}
		}

		return true
	}

	return false
}

/*
IsValidHandCards 判断手牌是否合法
handCards: 手牌数组，最多14张。
return: true手牌合法，false手牌非法。
*/
func IsValidHandCards(handCards []int) bool {
	cardsNum := len(handCards)
	if cardsNum <= 0 || cardsNum%3 != 2 || cardsNum > MaxHandCardNum {
		return false
	}

	for _, c := range handCards {
		if !IsValidCard(c) {
			return false
		}
	}

	slots := GenSlots(handCards)
	for _, num := range slots {
		if num > 4 {
			return false
		}
	}

	return true
}

/*
GenSlots 将手牌转换为对应的表示每张牌数量的数组。
例子：
handCards []int = {0x01, 0x01, 0x11, 0x12, 0x13, 0x33, 0x43, 0x43}
[TILEMAX]int = {
	2, 0, 0, 0, 0, 0, 0, 0, 0, //万
	1, 1, 1, 0, 0, 0, 0, 0, 0, //条
	0, 0, 0, 0, 0, 0, 0, 0, 0, //筒
	0, 0, 1, 0, //风
	0, 0, 2, //箭
}
*/
func GenSlots(handCards []int) [TILEMAX]int {
	var slots [TILEMAX]int

	for _, c := range handCards {
		slots[c]++
	}

	return slots
}

func getWan(slots [TILEMAX]int) int {
	return getNum(slots, CharacterOne, CharacterNine)
}

func getTiao(slots [TILEMAX]int) int {
	return getNum(slots, BambooOne, BambooNine)
}

func getTong(slots [TILEMAX]int) int {
	return getNum(slots, DotOne, DotNine)
}

func getFeng(slots [TILEMAX]int) int {
	return getNum(slots, EastWind, NorthWind)
}

func getJian(slots [TILEMAX]int) int {
	return getNum(slots, WhiteDragon, RedDragon)
}

func getNum(slots [TILEMAX]int, iStart int, iEnd int) int {
	num := 0
	for i := iStart; i <= iEnd; i++ {
		num = num*10 + int(slots[i])
	}
	return num
}

/*
ShowHandCards 打印手牌
*/
func ShowHandCards(handCards []int) {
	for _, c := range handCards {
		//fmt.Printf("0x%02X, ", c)
		fmt.Printf("%d, ", c)
	}
	fmt.Println()
}

/*
核心思想：
1、名词解释：eye(将），字牌（feng、东南西北中发白），花色（万、筒、条、字牌）
2、分而治之：检查手牌是否能胡是依次检查万、筒、条、字牌四种花色是否能组成胡牌的一部分。
3、单一花色要能满足胡牌的部分，则要么是3*n(不带将)，要么是3*n+2（带将）。3*n中的3带表三张牌一样的刻子，或三张连续的牌如1筒2筒3筒。
4、判断是否满足胡牌的单一花色部分，需要根据是否有将，有几个赖子，查询不同的表。表内容表示表里的元素加上对应的赖子数量能组成3*n 或3*n+2。
赖子数是表名最后的数字，带有eye的表表示满足3*n+2，没有的表示满足3*n。
5、查表的key值，是直接根据1-9有几张牌就填几（赖子不算），如1-9万各一张，则key为111111111。如1万3张，9万2张，则key为300000002。
6、组合多种花色，判断是否能胡牌。将赖子分配给不同的花色，有若干种分配方式，只要有一种分配能让所有花色满足单一花色胡牌部分，则手牌能胡。
如：手上有3个赖子，可以分配万、筒、条各一张，也可以万、同、字牌各一张
7、根据是否有将、是否字牌分为4种表，每种表又根据赖子个数0-8分别建表，共36张表，具体如下：

赖子个数     带将表        不带将的表         字牌带将表             字牌不带将表
0        eye_table_0     table_0       feng_eye_table_0        feng_table_0
1        eye_table_1     table_1       feng_eye_table_0        feng_table_1
2        eye_table_2     table_2       feng_eye_table_0        feng_table_2
3        eye_table_3     table_3       feng_eye_table_0        feng_table_3
4        eye_table_4     table_4       feng_eye_table_0        feng_table_4
5        eye_table_5     table_5       feng_eye_table_0        feng_table_5
6        eye_table_6     table_6       feng_eye_table_0        feng_table_6
7        eye_table_7     table_7       feng_eye_table_0        feng_table_7
8        eye_table_8     table_8       feng_eye_table_0        feng_table_8

步骤：
1、统计手牌中鬼牌个数 nGui，将鬼牌从牌数据中去除.
2、不同花色分开处理，分别校验是否能满足 将、顺子、刻子
3、分析东南西北风中发白时，不需要分析顺子的情况，简单很多
4、分析单一花色时，直接根据1-9点对应数字得出一个9位的整数，每位上为0-4代表该点数上有几张牌
比如：1筒2筒3筒3筒3筒3筒6筒7筒8筒2万3万3万3万4万
数字：筒: 1,1,4,0,0,1,1,1,0 得出的数字为114001110
      万 0,1,3,1,0,0,0,0,0 得出的数字为13100000
5、组合多种花色，判断是否能胡牌。将赖子分配给不同的花色，有若干种分配方式，只要有一种分配能让所有花色满足单一花色胡牌部分，则手牌能胡。
如：手上有3个赖子，可以分配万、筒、条各一张，也可以万、同、字牌各一张
每种花色与赖子组合，如果所有花色都能配型成功则可胡牌
检查配型时，每种花色的牌数量必需是3*n 或者 3*n + 2
根据赖子个数、带不带将，查找对应表，看能否满足3*n 或 3*n+2的牌型

非字牌表的产生:
1、穷举万字牌所有满足胡牌胡可能，将对应的牌型记录为数字，根据是否有将、放入eye_table_0或table_0中。
  具体是每次加入一个刻子，顺子或是将（将只能加入一对）,最多加入四组外带将牌。
2、将table_0中牌去掉一张，放入table_1中，表示去掉的牌用1张赖子代替。eye_table_0中牌去掉一张，放入到eye_table_1中。
3、将table_1中牌去掉一张，放入table_2中，表示去掉的牌用1张赖子代替。eye_table_1中牌去掉一张，放入到eye_table_2中。
4、将table_2中牌去掉一张，放入table_3中，表示去掉的牌用1张赖子代替。eye_table_2中牌去掉一张，放入到eye_table_3中。
5、将table_3中牌去掉一张，放入table_4中，表示去掉的牌用1张赖子代替。eye_table_3中牌去掉一张，放入到eye_table_4中。
6、将table_4中牌去掉一张，放入table_5中，表示去掉的牌用1张赖子代替。eye_table_4中牌去掉一张，放入到eye_table_5中。
7、将table_5中牌去掉一张，放入table_6中，表示去掉的牌用1张赖子代替。eye_table_5中牌去掉一张，放入到eye_table_6中。
8、将table_6中牌去掉一张，放入table_7中，表示去掉的牌用1张赖子代替。eye_table_6中牌去掉一张，放入到eye_table_7中。
9、将table_7中牌去掉一张，放入table_8中，表示去掉的牌用1张赖子代替。eye_table_7中牌去掉一张，放入到eye_table_8中。

字牌表的产生：
与非字牌表的产生方法相同，只是第一步中，不能加入顺子（除非麻将玩法字牌是能组成顺子的）

表的大小：总量在2M左右

表生成耗时：2-3S
*/
var (
	DianShuTable = make(map[int]int)
)

func init() {
	initDianShuTable()
}

func initDianShuTable() {
	DianShuTable[0] = 1
	DianShuTable[1] = 2
	DianShuTable[2] = 3
	DianShuTable[3] = 4
	DianShuTable[4] = 5
	DianShuTable[5] = 6
	DianShuTable[6] = 7
	DianShuTable[7] = 8
	DianShuTable[8] = 9

	DianShuTable[9] = 1
	DianShuTable[10] = 2
	DianShuTable[11] = 3
	DianShuTable[12] = 4
	DianShuTable[13] = 5
	DianShuTable[14] = 6
	DianShuTable[15] = 7
	DianShuTable[16] = 8
	DianShuTable[17] = 9

	DianShuTable[18] = 1
	DianShuTable[19] = 2
	DianShuTable[20] = 3
	DianShuTable[21] = 4
	DianShuTable[22] = 5
	DianShuTable[23] = 6
	DianShuTable[24] = 7
	DianShuTable[25] = 8
	DianShuTable[26] = 9

	DianShuTable[27] = 10
	DianShuTable[28] = 10
	DianShuTable[29] = 10
	DianShuTable[30] = 10
	DianShuTable[31] = 10
	DianShuTable[32] = 10
	DianShuTable[33] = 10
}
