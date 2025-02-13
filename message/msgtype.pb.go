// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: msg/msgtype.proto

package message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MSG_TYPE int32

const (
	MSG_TYPE_E_M_MSG_NULL          MSG_TYPE = 0
	MSG_TYPE_E_MSG_DB_UPDATE       MSG_TYPE = 1
	MSG_TYPE_E_Test_C              MSG_TYPE = 2
	MSG_TYPE_E_Test_S              MSG_TYPE = 3
	MSG_TYPE_E_Gm_C                MSG_TYPE = 998
	MSG_TYPE_E_Gm_S                MSG_TYPE = 999
	MSG_TYPE_E_Auth_C              MSG_TYPE = 1000
	MSG_TYPE_E_Auth_S              MSG_TYPE = 1001
	MSG_TYPE_E_Login_C             MSG_TYPE = 1002
	MSG_TYPE_E_Login_S             MSG_TYPE = 1003
	MSG_TYPE_E_PlayerInfo_S        MSG_TYPE = 1004 //玩家数据
	MSG_TYPE_E_CrossDay_C          MSG_TYPE = 1005 //跨天请求
	MSG_TYPE_E_CrossDay_S          MSG_TYPE = 1006
	MSG_TYPE_E_OpenFuncChange_S    MSG_TYPE = 1007 //功能开启
	MSG_TYPE_E_BagInfo_S           MSG_TYPE = 2000 //背包背包信息
	MSG_TYPE_E_BagItemUpdate_S     MSG_TYPE = 2001 //道具更新
	MSG_TYPE_E_BagLevelUp_C        MSG_TYPE = 2002 //背包升级
	MSG_TYPE_E_BagLevelUp_S        MSG_TYPE = 2003
	MSG_TYPE_E_BagSellItem_C       MSG_TYPE = 2004 //出售背包中道具
	MSG_TYPE_E_BagSellItem_S       MSG_TYPE = 2005
	MSG_TYPE_E_BuyItems_C          MSG_TYPE = 2006 //钻石购买道具
	MSG_TYPE_E_BuyItems_S          MSG_TYPE = 2007
	MSG_TYPE_E_Reward_S            MSG_TYPE = 2008
	MSG_TYPE_E_PlayerMailInfo_C    MSG_TYPE = 2009 //个人邮件
	MSG_TYPE_E_PlayerMailInfo_S    MSG_TYPE = 2010
	MSG_TYPE_E_PlayerMailChange_S  MSG_TYPE = 2011
	MSG_TYPE_E_ReceiveMailItems_C  MSG_TYPE = 2012 //领附件
	MSG_TYPE_E_ReceiveMailItems_S  MSG_TYPE = 2013
	MSG_TYPE_E_ElementBagInfo_S    MSG_TYPE = 2100 //所有物品信息
	MSG_TYPE_E_ElementLevelUp_C    MSG_TYPE = 2101 //地上物升级请求
	MSG_TYPE_E_ElementLevelUp_S    MSG_TYPE = 2102 //地上物升级返回
	MSG_TYPE_E_ElementAutoReward_C MSG_TYPE = 2103 //自动生产物品，领取奖励
	MSG_TYPE_E_ElementAutoReward_S MSG_TYPE = 2104 //自动生产物品，领取奖励
	MSG_TYPE_E_ElementUpdate_S     MSG_TYPE = 2105 //物品更新或新增
	//  E_ElementDelete_S = 2106; //物品删除
	MSG_TYPE_E_ElementBagInfo_C   MSG_TYPE = 2106 //获取背包
	MSG_TYPE_E_MapInfo_S          MSG_TYPE = 2107 //地图信息
	MSG_TYPE_E_MapElementChange_C MSG_TYPE = 2108 //地编地上物变更请求
	MSG_TYPE_E_MapElementChange_S MSG_TYPE = 2109 //地编地上物变更响应
	//  E_ObstacleClear_C = 2110; //障碍物清理请求
	//  E_ObstacleClear_S = 2111;
	MSG_TYPE_E_MapTileGroupUnlock_C         MSG_TYPE = 2112 //地块组解锁请求
	MSG_TYPE_E_MapTileGroupUnlock_S         MSG_TYPE = 2113 //地块组解锁响应
	MSG_TYPE_E_ElementMakeItem_C            MSG_TYPE = 2114 //生产道具请求
	MSG_TYPE_E_ElementMakeItem_S            MSG_TYPE = 2115 //生产道具请求
	MSG_TYPE_E_ElementCancelMakeItem_C      MSG_TYPE = 2116 //删除生产道具请求
	MSG_TYPE_E_ElementCancelMakeItem_S      MSG_TYPE = 2117 //删除生产道具请求
	MSG_TYPE_E_ElementMakeRewards_C         MSG_TYPE = 2118 //批量领取生产道具
	MSG_TYPE_E_ElementMakeRewards_S         MSG_TYPE = 2119 //批量领取生产道具
	MSG_TYPE_E_ElementMakeBuySlot_C         MSG_TYPE = 2120 //生产道具槽位购买请求
	MSG_TYPE_E_ElementMakeBuySlot_S         MSG_TYPE = 2121 //生产道具槽位购买请求
	MSG_TYPE_E_MapTileGroupFinishUnlock_C   MSG_TYPE = 2122 //完成解锁
	MSG_TYPE_E_MapTileGroupFinishUnlock_S   MSG_TYPE = 2123 //完成解锁
	MSG_TYPE_E_ElementLevelUpFinish_C       MSG_TYPE = 2124 //完成升级
	MSG_TYPE_E_ElementLevelUpFinish_S       MSG_TYPE = 2125 //完成升级
	MSG_TYPE_E_MapInfo_C                    MSG_TYPE = 2126 //获取地图
	MSG_TYPE_E_ElementShopInfo_C            MSG_TYPE = 2127 //地上物商店
	MSG_TYPE_E_ElementShopInfo_S            MSG_TYPE = 2128 //地上物商店
	MSG_TYPE_E_ElementShopBuyBuilding_C     MSG_TYPE = 2129 //地上物商店购买建筑
	MSG_TYPE_E_ElementShopBuyBuilding_S     MSG_TYPE = 2130
	MSG_TYPE_E_ElementShopBuyDecorate_C     MSG_TYPE = 2131 //地上物商店购买装饰物
	MSG_TYPE_E_ElementShopBuyDecorate_S     MSG_TYPE = 2132
	MSG_TYPE_E_ElementDecorateGroupReward_C MSG_TYPE = 2133 //领取组奖励
	MSG_TYPE_E_ElementDecorateGroupReward_S MSG_TYPE = 2134
	MSG_TYPE_E_ElementAutoGenRewardInfo_C   MSG_TYPE = 2135 //获取某个地上物自动产出的详情
	MSG_TYPE_E_ElementAutoGenRewardInfo_S   MSG_TYPE = 2136
	MSG_TYPE_E_GetHandbooks_C               MSG_TYPE = 2137 //获取图鉴信息
	MSG_TYPE_E_GetHandbooks_S               MSG_TYPE = 2138
	MSG_TYPE_E_WishTreeInfo_C               MSG_TYPE = 2139 //获取许愿树信息
	MSG_TYPE_E_WishTreeInfo_S               MSG_TYPE = 2140
	MSG_TYPE_E_WishTreeSubmit_C             MSG_TYPE = 2141 //许愿树 提交订单
	MSG_TYPE_E_WishTreeSubmit_S             MSG_TYPE = 2142
	MSG_TYPE_E_WishTreeOpenBox_C            MSG_TYPE = 2143 //许愿树 宝箱领取
	MSG_TYPE_E_WishTreeOpenBox_S            MSG_TYPE = 2144
	MSG_TYPE_E_WishTreeOpen_C               MSG_TYPE = 2145 //许愿树 展开订单
	MSG_TYPE_E_WishTreeOpen_S               MSG_TYPE = 2146
	MSG_TYPE_E_RefreshOrderItem_C           MSG_TYPE = 2147 //许愿树 刷新订单
	MSG_TYPE_E_RefreshOrderItem_S           MSG_TYPE = 2148
	MSG_TYPE_E_TrainInfo_C                  MSG_TYPE = 2149 //获取小火车信息
	MSG_TYPE_E_TrainInfo_S                  MSG_TYPE = 2150
	MSG_TYPE_E_UnlockTrack_C                MSG_TYPE = 2151 //小火车解锁轨道
	MSG_TYPE_E_UnlockTrack_S                MSG_TYPE = 2152
	MSG_TYPE_E_CommitTrack_C                MSG_TYPE = 2153 //小火车 交付
	MSG_TYPE_E_CommitTrack_S                MSG_TYPE = 2154
	MSG_TYPE_E_TrackReward_C                MSG_TYPE = 2155 //小火车 领奖
	MSG_TYPE_E_TrackReward_S                MSG_TYPE = 2156
	MSG_TYPE_E_Accelerate_C                 MSG_TYPE = 2157 //加速
	MSG_TYPE_E_Accelerate_S                 MSG_TYPE = 2158
	MSG_TYPE_E_ModuleNumberInfo_C           MSG_TYPE = 2159 //计数器
	MSG_TYPE_E_ModuleNumberInfo_S           MSG_TYPE = 2160
	MSG_TYPE_E_ModuleNumberChangeInfo_S     MSG_TYPE = 2161
	MSG_TYPE_E_ObstacleClear_C              MSG_TYPE = 2162 //障碍物清理请求
	MSG_TYPE_E_ObstacleClear_S              MSG_TYPE = 2163
	MSG_TYPE_E_ObstacleClearFinish_C        MSG_TYPE = 2164 //障碍物清理请求
	MSG_TYPE_E_ObstacleClearFinish_S        MSG_TYPE = 2165
	MSG_TYPE_E_SkinEdit_C                   MSG_TYPE = 2166
	MSG_TYPE_E_SkinEdit_S                   MSG_TYPE = 2167
	MSG_TYPE_E_SkinBuy_C                    MSG_TYPE = 2168
	MSG_TYPE_E_SkinBuy_S                    MSG_TYPE = 2169
	MSG_TYPE_E_SkinChange_S                 MSG_TYPE = 2170
	MSG_TYPE_E_FunctionInfo_S               MSG_TYPE = 2200 //功能列表
	MSG_TYPE_E_HeroInfo_C                   MSG_TYPE = 2300 //英雄信息
	MSG_TYPE_E_HeroInfo_S                   MSG_TYPE = 2301
	MSG_TYPE_E_HeroInfoChange_S             MSG_TYPE = 2302
	MSG_TYPE_E_HeroUpdateLevel_C            MSG_TYPE = 2303 //升级
	MSG_TYPE_E_HeroUpdateLevel_S            MSG_TYPE = 2304
	MSG_TYPE_E_HeroUpdateSkillLevel_C       MSG_TYPE = 2305 //升级技能
	MSG_TYPE_E_HeroUpdateSkillLevel_S       MSG_TYPE = 2306
	MSG_TYPE_E_HeroUpdateStar_C             MSG_TYPE = 2307 //升星
	MSG_TYPE_E_HeroUpdateStar_S             MSG_TYPE = 2308
	MSG_TYPE_E_HeroWorkingStatusChange_S    MSG_TYPE = 2309 //工作状态修改
)

// Enum value maps for MSG_TYPE.
var (
	MSG_TYPE_name = map[int32]string{
		0:    "E_M_MSG_NULL",
		1:    "E_MSG_DB_UPDATE",
		2:    "E_Test_C",
		3:    "E_Test_S",
		998:  "E_Gm_C",
		999:  "E_Gm_S",
		1000: "E_Auth_C",
		1001: "E_Auth_S",
		1002: "E_Login_C",
		1003: "E_Login_S",
		1004: "E_PlayerInfo_S",
		1005: "E_CrossDay_C",
		1006: "E_CrossDay_S",
		1007: "E_OpenFuncChange_S",
		2000: "E_BagInfo_S",
		2001: "E_BagItemUpdate_S",
		2002: "E_BagLevelUp_C",
		2003: "E_BagLevelUp_S",
		2004: "E_BagSellItem_C",
		2005: "E_BagSellItem_S",
		2006: "E_BuyItems_C",
		2007: "E_BuyItems_S",
		2008: "E_Reward_S",
		2009: "E_PlayerMailInfo_C",
		2010: "E_PlayerMailInfo_S",
		2011: "E_PlayerMailChange_S",
		2012: "E_ReceiveMailItems_C",
		2013: "E_ReceiveMailItems_S",
		2100: "E_ElementBagInfo_S",
		2101: "E_ElementLevelUp_C",
		2102: "E_ElementLevelUp_S",
		2103: "E_ElementAutoReward_C",
		2104: "E_ElementAutoReward_S",
		2105: "E_ElementUpdate_S",
		2106: "E_ElementBagInfo_C",
		2107: "E_MapInfo_S",
		2108: "E_MapElementChange_C",
		2109: "E_MapElementChange_S",
		2112: "E_MapTileGroupUnlock_C",
		2113: "E_MapTileGroupUnlock_S",
		2114: "E_ElementMakeItem_C",
		2115: "E_ElementMakeItem_S",
		2116: "E_ElementCancelMakeItem_C",
		2117: "E_ElementCancelMakeItem_S",
		2118: "E_ElementMakeRewards_C",
		2119: "E_ElementMakeRewards_S",
		2120: "E_ElementMakeBuySlot_C",
		2121: "E_ElementMakeBuySlot_S",
		2122: "E_MapTileGroupFinishUnlock_C",
		2123: "E_MapTileGroupFinishUnlock_S",
		2124: "E_ElementLevelUpFinish_C",
		2125: "E_ElementLevelUpFinish_S",
		2126: "E_MapInfo_C",
		2127: "E_ElementShopInfo_C",
		2128: "E_ElementShopInfo_S",
		2129: "E_ElementShopBuyBuilding_C",
		2130: "E_ElementShopBuyBuilding_S",
		2131: "E_ElementShopBuyDecorate_C",
		2132: "E_ElementShopBuyDecorate_S",
		2133: "E_ElementDecorateGroupReward_C",
		2134: "E_ElementDecorateGroupReward_S",
		2135: "E_ElementAutoGenRewardInfo_C",
		2136: "E_ElementAutoGenRewardInfo_S",
		2137: "E_GetHandbooks_C",
		2138: "E_GetHandbooks_S",
		2139: "E_WishTreeInfo_C",
		2140: "E_WishTreeInfo_S",
		2141: "E_WishTreeSubmit_C",
		2142: "E_WishTreeSubmit_S",
		2143: "E_WishTreeOpenBox_C",
		2144: "E_WishTreeOpenBox_S",
		2145: "E_WishTreeOpen_C",
		2146: "E_WishTreeOpen_S",
		2147: "E_RefreshOrderItem_C",
		2148: "E_RefreshOrderItem_S",
		2149: "E_TrainInfo_C",
		2150: "E_TrainInfo_S",
		2151: "E_UnlockTrack_C",
		2152: "E_UnlockTrack_S",
		2153: "E_CommitTrack_C",
		2154: "E_CommitTrack_S",
		2155: "E_TrackReward_C",
		2156: "E_TrackReward_S",
		2157: "E_Accelerate_C",
		2158: "E_Accelerate_S",
		2159: "E_ModuleNumberInfo_C",
		2160: "E_ModuleNumberInfo_S",
		2161: "E_ModuleNumberChangeInfo_S",
		2162: "E_ObstacleClear_C",
		2163: "E_ObstacleClear_S",
		2164: "E_ObstacleClearFinish_C",
		2165: "E_ObstacleClearFinish_S",
		2166: "E_SkinEdit_C",
		2167: "E_SkinEdit_S",
		2168: "E_SkinBuy_C",
		2169: "E_SkinBuy_S",
		2170: "E_SkinChange_S",
		2200: "E_FunctionInfo_S",
		2300: "E_HeroInfo_C",
		2301: "E_HeroInfo_S",
		2302: "E_HeroInfoChange_S",
		2303: "E_HeroUpdateLevel_C",
		2304: "E_HeroUpdateLevel_S",
		2305: "E_HeroUpdateSkillLevel_C",
		2306: "E_HeroUpdateSkillLevel_S",
		2307: "E_HeroUpdateStar_C",
		2308: "E_HeroUpdateStar_S",
		2309: "E_HeroWorkingStatusChange_S",
	}
	MSG_TYPE_value = map[string]int32{
		"E_M_MSG_NULL":                   0,
		"E_MSG_DB_UPDATE":                1,
		"E_Test_C":                       2,
		"E_Test_S":                       3,
		"E_Gm_C":                         998,
		"E_Gm_S":                         999,
		"E_Auth_C":                       1000,
		"E_Auth_S":                       1001,
		"E_Login_C":                      1002,
		"E_Login_S":                      1003,
		"E_PlayerInfo_S":                 1004,
		"E_CrossDay_C":                   1005,
		"E_CrossDay_S":                   1006,
		"E_OpenFuncChange_S":             1007,
		"E_BagInfo_S":                    2000,
		"E_BagItemUpdate_S":              2001,
		"E_BagLevelUp_C":                 2002,
		"E_BagLevelUp_S":                 2003,
		"E_BagSellItem_C":                2004,
		"E_BagSellItem_S":                2005,
		"E_BuyItems_C":                   2006,
		"E_BuyItems_S":                   2007,
		"E_Reward_S":                     2008,
		"E_PlayerMailInfo_C":             2009,
		"E_PlayerMailInfo_S":             2010,
		"E_PlayerMailChange_S":           2011,
		"E_ReceiveMailItems_C":           2012,
		"E_ReceiveMailItems_S":           2013,
		"E_ElementBagInfo_S":             2100,
		"E_ElementLevelUp_C":             2101,
		"E_ElementLevelUp_S":             2102,
		"E_ElementAutoReward_C":          2103,
		"E_ElementAutoReward_S":          2104,
		"E_ElementUpdate_S":              2105,
		"E_ElementBagInfo_C":             2106,
		"E_MapInfo_S":                    2107,
		"E_MapElementChange_C":           2108,
		"E_MapElementChange_S":           2109,
		"E_MapTileGroupUnlock_C":         2112,
		"E_MapTileGroupUnlock_S":         2113,
		"E_ElementMakeItem_C":            2114,
		"E_ElementMakeItem_S":            2115,
		"E_ElementCancelMakeItem_C":      2116,
		"E_ElementCancelMakeItem_S":      2117,
		"E_ElementMakeRewards_C":         2118,
		"E_ElementMakeRewards_S":         2119,
		"E_ElementMakeBuySlot_C":         2120,
		"E_ElementMakeBuySlot_S":         2121,
		"E_MapTileGroupFinishUnlock_C":   2122,
		"E_MapTileGroupFinishUnlock_S":   2123,
		"E_ElementLevelUpFinish_C":       2124,
		"E_ElementLevelUpFinish_S":       2125,
		"E_MapInfo_C":                    2126,
		"E_ElementShopInfo_C":            2127,
		"E_ElementShopInfo_S":            2128,
		"E_ElementShopBuyBuilding_C":     2129,
		"E_ElementShopBuyBuilding_S":     2130,
		"E_ElementShopBuyDecorate_C":     2131,
		"E_ElementShopBuyDecorate_S":     2132,
		"E_ElementDecorateGroupReward_C": 2133,
		"E_ElementDecorateGroupReward_S": 2134,
		"E_ElementAutoGenRewardInfo_C":   2135,
		"E_ElementAutoGenRewardInfo_S":   2136,
		"E_GetHandbooks_C":               2137,
		"E_GetHandbooks_S":               2138,
		"E_WishTreeInfo_C":               2139,
		"E_WishTreeInfo_S":               2140,
		"E_WishTreeSubmit_C":             2141,
		"E_WishTreeSubmit_S":             2142,
		"E_WishTreeOpenBox_C":            2143,
		"E_WishTreeOpenBox_S":            2144,
		"E_WishTreeOpen_C":               2145,
		"E_WishTreeOpen_S":               2146,
		"E_RefreshOrderItem_C":           2147,
		"E_RefreshOrderItem_S":           2148,
		"E_TrainInfo_C":                  2149,
		"E_TrainInfo_S":                  2150,
		"E_UnlockTrack_C":                2151,
		"E_UnlockTrack_S":                2152,
		"E_CommitTrack_C":                2153,
		"E_CommitTrack_S":                2154,
		"E_TrackReward_C":                2155,
		"E_TrackReward_S":                2156,
		"E_Accelerate_C":                 2157,
		"E_Accelerate_S":                 2158,
		"E_ModuleNumberInfo_C":           2159,
		"E_ModuleNumberInfo_S":           2160,
		"E_ModuleNumberChangeInfo_S":     2161,
		"E_ObstacleClear_C":              2162,
		"E_ObstacleClear_S":              2163,
		"E_ObstacleClearFinish_C":        2164,
		"E_ObstacleClearFinish_S":        2165,
		"E_SkinEdit_C":                   2166,
		"E_SkinEdit_S":                   2167,
		"E_SkinBuy_C":                    2168,
		"E_SkinBuy_S":                    2169,
		"E_SkinChange_S":                 2170,
		"E_FunctionInfo_S":               2200,
		"E_HeroInfo_C":                   2300,
		"E_HeroInfo_S":                   2301,
		"E_HeroInfoChange_S":             2302,
		"E_HeroUpdateLevel_C":            2303,
		"E_HeroUpdateLevel_S":            2304,
		"E_HeroUpdateSkillLevel_C":       2305,
		"E_HeroUpdateSkillLevel_S":       2306,
		"E_HeroUpdateStar_C":             2307,
		"E_HeroUpdateStar_S":             2308,
		"E_HeroWorkingStatusChange_S":    2309,
	}
)

func (x MSG_TYPE) Enum() *MSG_TYPE {
	p := new(MSG_TYPE)
	*p = x
	return p
}

func (x MSG_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_msgtype_proto_enumTypes[0].Descriptor()
}

func (MSG_TYPE) Type() protoreflect.EnumType {
	return &file_msg_msgtype_proto_enumTypes[0]
}

func (x MSG_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_TYPE.Descriptor instead.
func (MSG_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_msg_msgtype_proto_rawDescGZIP(), []int{0}
}

var File_msg_msgtype_proto protoreflect.FileDescriptor

var file_msg_msgtype_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0xe9, 0x14, 0x0a, 0x08, 0x4d, 0x53, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x5f, 0x4d, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x44, 0x42, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x45, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x5f, 0x43, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x5f,
	0x54, 0x65, 0x73, 0x74, 0x5f, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x06, 0x45, 0x5f, 0x47, 0x6d,
	0x5f, 0x43, 0x10, 0xe6, 0x07, 0x12, 0x0b, 0x0a, 0x06, 0x45, 0x5f, 0x47, 0x6d, 0x5f, 0x53, 0x10,
	0xe7, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x45, 0x5f, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x43, 0x10, 0xe8,
	0x07, 0x12, 0x0d, 0x0a, 0x08, 0x45, 0x5f, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x53, 0x10, 0xe9, 0x07,
	0x12, 0x0e, 0x0a, 0x09, 0x45, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x43, 0x10, 0xea, 0x07,
	0x12, 0x0e, 0x0a, 0x09, 0x45, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x53, 0x10, 0xeb, 0x07,
	0x12, 0x13, 0x0a, 0x0e, 0x45, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x5f, 0x53, 0x10, 0xec, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x43, 0x72, 0x6f, 0x73, 0x73,
	0x44, 0x61, 0x79, 0x5f, 0x43, 0x10, 0xed, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x44, 0x61, 0x79, 0x5f, 0x53, 0x10, 0xee, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x45,
	0x5f, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x53, 0x10, 0xef, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x45, 0x5f, 0x42, 0x61, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x5f, 0x53, 0x10, 0xd0, 0x0f, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x5f, 0x42, 0x61, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x53, 0x10, 0xd1, 0x0f, 0x12, 0x13,
	0x0a, 0x0e, 0x45, 0x5f, 0x42, 0x61, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x5f, 0x43,
	0x10, 0xd2, 0x0f, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x5f, 0x42, 0x61, 0x67, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x55, 0x70, 0x5f, 0x53, 0x10, 0xd3, 0x0f, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f, 0x42, 0x61,
	0x67, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x5f, 0x43, 0x10, 0xd4, 0x0f, 0x12, 0x14,
	0x0a, 0x0f, 0x45, 0x5f, 0x42, 0x61, 0x67, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x5f,
	0x53, 0x10, 0xd5, 0x0f, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x5f, 0x43, 0x10, 0xd6, 0x0f, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x42, 0x75, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x53, 0x10, 0xd7, 0x0f, 0x12, 0x0f, 0x0a, 0x0a, 0x45, 0x5f,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x53, 0x10, 0xd8, 0x0f, 0x12, 0x17, 0x0a, 0x12, 0x45,
	0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x5f,
	0x43, 0x10, 0xd9, 0x0f, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0xda, 0x0f, 0x12, 0x19, 0x0a,
	0x14, 0x45, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x53, 0x10, 0xdb, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x43,
	0x10, 0xdc, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x53, 0x10, 0xdd, 0x0f, 0x12, 0x17,
	0x0a, 0x12, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x5f, 0x53, 0x10, 0xb4, 0x10, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x5f, 0x43, 0x10, 0xb5, 0x10,
	0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x55, 0x70, 0x5f, 0x53, 0x10, 0xb6, 0x10, 0x12, 0x1a, 0x0a, 0x15, 0x45, 0x5f, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x43, 0x10, 0xb7, 0x10, 0x12, 0x1a, 0x0a, 0x15, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x53, 0x10, 0xb8,
	0x10, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x53, 0x10, 0xb9, 0x10, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10,
	0xba, 0x10, 0x12, 0x10, 0x0a, 0x0b, 0x45, 0x5f, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x5f,
	0x53, 0x10, 0xbb, 0x10, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x4d, 0x61, 0x70, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x43, 0x10, 0xbc, 0x10, 0x12,
	0x19, 0x0a, 0x14, 0x45, 0x5f, 0x4d, 0x61, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x53, 0x10, 0xbd, 0x10, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x5f,
	0x4d, 0x61, 0x70, 0x54, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x43, 0x10, 0xc0, 0x10, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x5f, 0x4d, 0x61, 0x70,
	0x54, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x53, 0x10, 0xc1, 0x10, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x6b, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x5f, 0x43, 0x10, 0xc2, 0x10, 0x12, 0x18,
	0x0a, 0x13, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x5f, 0x53, 0x10, 0xc3, 0x10, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x5f, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x6b, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x5f, 0x43, 0x10, 0xc4, 0x10, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x5f, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x6b, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x5f, 0x53, 0x10, 0xc5, 0x10, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x5f, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x43, 0x10, 0xc6, 0x10, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x53, 0x10,
	0xc7, 0x10, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x6b, 0x65, 0x42, 0x75, 0x79, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x43, 0x10, 0xc8, 0x10, 0x12,
	0x1b, 0x0a, 0x16, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6b, 0x65,
	0x42, 0x75, 0x79, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x53, 0x10, 0xc9, 0x10, 0x12, 0x21, 0x0a, 0x1c,
	0x45, 0x5f, 0x4d, 0x61, 0x70, 0x54, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x43, 0x10, 0xca, 0x10, 0x12,
	0x21, 0x0a, 0x1c, 0x45, 0x5f, 0x4d, 0x61, 0x70, 0x54, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x53, 0x10,
	0xcb, 0x10, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x43, 0x10, 0xcc,
	0x10, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x55, 0x70, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x53, 0x10, 0xcd, 0x10,
	0x12, 0x10, 0x0a, 0x0b, 0x45, 0x5f, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10,
	0xce, 0x10, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10, 0xcf, 0x10, 0x12, 0x18, 0x0a, 0x13,
	0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x5f, 0x53, 0x10, 0xd0, 0x10, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x43, 0x10, 0xd1, 0x10, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x5f, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x53, 0x10, 0xd2, 0x10, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x5f, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x44, 0x65, 0x63, 0x6f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x10, 0xd3, 0x10, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x5f, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x44, 0x65, 0x63,
	0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x53, 0x10, 0xd4, 0x10, 0x12, 0x23, 0x0a, 0x1e, 0x45, 0x5f,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x43, 0x10, 0xd5, 0x10, 0x12,
	0x23, 0x0a, 0x1e, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x6f,
	0x72, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x53, 0x10, 0xd6, 0x10, 0x12, 0x21, 0x0a, 0x1c, 0x45, 0x5f, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x75, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x5f, 0x43, 0x10, 0xd7, 0x10, 0x12, 0x21, 0x0a, 0x1c, 0x45, 0x5f, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0xd8, 0x10, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f,
	0x47, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x43, 0x10, 0xd9,
	0x10, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x5f, 0x53, 0x10, 0xda, 0x10, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x57, 0x69,
	0x73, 0x68, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10, 0xdb, 0x10, 0x12,
	0x15, 0x0a, 0x10, 0x45, 0x5f, 0x57, 0x69, 0x73, 0x68, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x5f, 0x53, 0x10, 0xdc, 0x10, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x57, 0x69, 0x73, 0x68,
	0x54, 0x72, 0x65, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x43, 0x10, 0xdd, 0x10, 0x12,
	0x17, 0x0a, 0x12, 0x45, 0x5f, 0x57, 0x69, 0x73, 0x68, 0x54, 0x72, 0x65, 0x65, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x5f, 0x53, 0x10, 0xde, 0x10, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x5f, 0x57, 0x69,
	0x73, 0x68, 0x54, 0x72, 0x65, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x6f, 0x78, 0x5f, 0x43, 0x10,
	0xdf, 0x10, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x5f, 0x57, 0x69, 0x73, 0x68, 0x54, 0x72, 0x65, 0x65,
	0x4f, 0x70, 0x65, 0x6e, 0x42, 0x6f, 0x78, 0x5f, 0x53, 0x10, 0xe0, 0x10, 0x12, 0x15, 0x0a, 0x10,
	0x45, 0x5f, 0x57, 0x69, 0x73, 0x68, 0x54, 0x72, 0x65, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x5f, 0x43,
	0x10, 0xe1, 0x10, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x57, 0x69, 0x73, 0x68, 0x54, 0x72, 0x65,
	0x65, 0x4f, 0x70, 0x65, 0x6e, 0x5f, 0x53, 0x10, 0xe2, 0x10, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x5f, 0x43, 0x10, 0xe3, 0x10, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x5f, 0x53, 0x10, 0xe4, 0x10,
	0x12, 0x12, 0x0a, 0x0d, 0x45, 0x5f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x5f,
	0x43, 0x10, 0xe5, 0x10, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x5f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0xe6, 0x10, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x43, 0x10, 0xe7, 0x10, 0x12, 0x14,
	0x0a, 0x0f, 0x45, 0x5f, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x5f,
	0x53, 0x10, 0xe8, 0x10, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x43, 0x10, 0xe9, 0x10, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x53, 0x10, 0xea, 0x10,
	0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x43, 0x10, 0xeb, 0x10, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x5f, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x53, 0x10, 0xec, 0x10, 0x12, 0x13, 0x0a, 0x0e,
	0x45, 0x5f, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x10, 0xed,
	0x10, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x5f, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x53, 0x10, 0xee, 0x10, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10, 0xef,
	0x10, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x5f, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0xf0, 0x10, 0x12, 0x1f, 0x0a, 0x1a,
	0x45, 0x5f, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0xf1, 0x10, 0x12, 0x16, 0x0a,
	0x11, 0x45, 0x5f, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x5f, 0x43, 0x10, 0xf2, 0x10, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x5f, 0x4f, 0x62, 0x73, 0x74, 0x61,
	0x63, 0x6c, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x53, 0x10, 0xf3, 0x10, 0x12, 0x1c, 0x0a,
	0x17, 0x45, 0x5f, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x43, 0x10, 0xf4, 0x10, 0x12, 0x1c, 0x0a, 0x17, 0x45,
	0x5f, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x53, 0x10, 0xf5, 0x10, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x53,
	0x6b, 0x69, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x5f, 0x43, 0x10, 0xf6, 0x10, 0x12, 0x11, 0x0a, 0x0c,
	0x45, 0x5f, 0x53, 0x6b, 0x69, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x5f, 0x53, 0x10, 0xf7, 0x10, 0x12,
	0x10, 0x0a, 0x0b, 0x45, 0x5f, 0x53, 0x6b, 0x69, 0x6e, 0x42, 0x75, 0x79, 0x5f, 0x43, 0x10, 0xf8,
	0x10, 0x12, 0x10, 0x0a, 0x0b, 0x45, 0x5f, 0x53, 0x6b, 0x69, 0x6e, 0x42, 0x75, 0x79, 0x5f, 0x53,
	0x10, 0xf9, 0x10, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x5f, 0x53, 0x6b, 0x69, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x53, 0x10, 0xfa, 0x10, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x53, 0x10, 0x98, 0x11, 0x12,
	0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x43, 0x10,
	0xfc, 0x11, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x6e, 0x66, 0x6f,
	0x5f, 0x53, 0x10, 0xfd, 0x11, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x53, 0x10, 0xfe, 0x11, 0x12, 0x18,
	0x0a, 0x13, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x5f, 0x43, 0x10, 0xff, 0x11, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x53, 0x10,
	0x80, 0x12, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x43, 0x10, 0x81,
	0x12, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x53, 0x10, 0x82, 0x12,
	0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x5f, 0x43, 0x10, 0x83, 0x12, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x5f, 0x48,
	0x65, 0x72, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x5f, 0x53, 0x10,
	0x84, 0x12, 0x12, 0x20, 0x0a, 0x1b, 0x45, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x57, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x53, 0x10, 0x85, 0x12, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_msgtype_proto_rawDescOnce sync.Once
	file_msg_msgtype_proto_rawDescData = file_msg_msgtype_proto_rawDesc
)

func file_msg_msgtype_proto_rawDescGZIP() []byte {
	file_msg_msgtype_proto_rawDescOnce.Do(func() {
		file_msg_msgtype_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_msgtype_proto_rawDescData)
	})
	return file_msg_msgtype_proto_rawDescData
}

var file_msg_msgtype_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_msgtype_proto_goTypes = []interface{}{
	(MSG_TYPE)(0), // 0: msg.MSG_TYPE
}
var file_msg_msgtype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_msgtype_proto_init() }
func file_msg_msgtype_proto_init() {
	if File_msg_msgtype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_msgtype_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_msgtype_proto_goTypes,
		DependencyIndexes: file_msg_msgtype_proto_depIdxs,
		EnumInfos:         file_msg_msgtype_proto_enumTypes,
	}.Build()
	File_msg_msgtype_proto = out.File
	file_msg_msgtype_proto_rawDesc = nil
	file_msg_msgtype_proto_goTypes = nil
	file_msg_msgtype_proto_depIdxs = nil
}
