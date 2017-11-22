package mj

//Card 麻将子定义
type Card uint8

//MAX_GAME_PLAYER_USER 最大游戏玩家数
const MAX_GAME_PLAYER_USER = 4

//CARD_MAX_COUNT 麻将中牌数
const CARD_MAX_COUNT = 136

//MAX_HANDCARD_NUM 最多的手牌数量为14张
const MAX_HANDCARD_NUM = 14

//MAX_FIXCARD_NUM 最多的固定牌数量
const MAX_FIXCARD_NUM = 4

//MAX_XUSHU_NUM 序数牌的张数: 1~9
const MAX_XUSHU_NUM = 9

//MAX_FENG_NUM 风牌的张数: 东南西北
const MAX_FENG_NUM = 4

//MAX_JIAN_NUM 箭牌的张数: 中发白
const MAX_JIAN_NUM = 3

//MAX_GANG_COUNT 最大杠操作,因为手牌最多14张，一个杠要用4张牌，所以手中最多同时存在3个杠。
const MAX_GANG_COUNT = 3

//INVALID_CARD 一张非真实性的牌，填充假数据之用
const INVALID_CARD = 0xFF

//INVALID_CHAIR 一个非真实性的玩家椅子号
const INVALID_CHAIR = 0x7F

const MJ_TYPE_WAN = 0    //万, 0-8,各4张，共36张
const MJ_TYPE_TIAO = 1   //条, 0-8,各4张，共36张
const MJ_TYPE_BING = 2   //饼, 0-8,各4张，共36张
const MJ_TYPE_FENG = 3   //东南西北各4张，共16张
const MJ_TYPE_JIAN = 4   //中发白  各4张，共12张
const MJ_TYPE_FLOWER = 5 //花

//MJ_TYPE_NUM 牌的种类数量(万，条，饼，风，箭）
const MJ_TYPE_NUM = 5

//MAX_ALL_GANG_COUNT 牌数
const MAX_ALL_GANG_COUNT = 34

//MAX_CARD 最大牌面值
const MAX_CARD = 0x43

//MAX_CARD_ARRAY_SIZE 数组大小，要比最大的牌面值大一。
const MAX_CARD_ARRAY_SIZE = (MAX_CARD + 1)

//g_CardWangData 万
var g_CardWangData []Card = []Card{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x00,
}

//g_CardTiaoData 条
var g_CardTiaoData []Card = []Card{
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x00,
}

//g_CardTongData 筒
var g_CardTongData []Card = []Card{
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x00,
}

//g_CardFengData 风
var g_CardFengData []Card = []Card{
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x31, 0x32, 0x33, 0x34,
	0x00,
}

//g_CardJianData 箭
var g_CardJianData []Card = []Card{
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x41, 0x42, 0x43,
	0x00,
}

//g_CardHuaData 花
var g_CardHuaData []Card = []Card{
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x00,
}

//g_CardGangData 所有牌类型
var g_CardGangData []Card = []Card{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //条
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //筒
	0x31, 0x32, 0x33, 0x34, //风
	0x41, 0x42, 0x43, //箭
}

//CardYiWan 一万
const CardYiWan = 0x01

//CardJiuWan 九万
const CardJiuWan = 0x09

//CardYaoJi 一条，幺鸡
const CardYaoJi = 0x11

//CardJiuTiao 九条
const CardJiuTiao = 0x19

//CardYiTong 一筒
const CardYiTong = 0x21

//CardJiuTong 九筒
const CardJiuTong = 0x29

//CardDONG 东
const CardDONG = 0x31

//CardNAN 南
const CardNAN = 0x32

//CardXI 西
const CardXI = 0x33

//CardBEI 北
const CardBEI = 0x34

//CardZHONG 中
const CardZHONG = 0x41

//CardFA 发
const CardFA = 0x42

//CardBAI 白
const CardBAI = 0x43

//CardNum 获取牌对应的值
func CardNum(c Card) uint8 {
	return uint8((c) & 0x0F)
}

//CardType 获取牌对应的类型
func CardType(c Card) uint8 {
	return uint8(((c) & 0xF0) >> 4)
}

//IsValidCard 判断是否是有效的牌
func IsValidCard(c Card) bool {
	if IsXuShu(c) || IsZi(c) {
		return true
	}
	return false
}

//IsXuShu 判断是否是序数牌
func IsXuShu(c Card) bool {
	if IsWan(c) || IsTiao(c) || IsTong(c) {
		return true
	}
	return false
}

//IsZi 判断是否是字牌
func IsZi(c Card) bool {
	if IsFeng(c) || IsJian(c) {
		return true
	}
	return false
}

//IsWan 判断是否是万牌
func IsWan(c Card) bool {
	return c >= 0x01 && c <= 0x09
}

//IsTiao 判断是否是条牌
func IsTiao(c Card) bool {
	return c >= 0x11 && c <= 0x19
}

//IsTong 判断是否是筒牌
func IsTong(c Card) bool {
	return c >= 0x21 && c <= 0x29
}

//IsFeng 判断是否是风牌
func IsFeng(c Card) bool {
	return c >= 0x31 && c <= 0x34
}

//IsJian 判断是否是箭牌
func IsJian(c Card) bool {
	return c >= 0x41 && c <= 0x43
}
