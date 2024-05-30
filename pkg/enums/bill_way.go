package enums

type BillWay string

const (
	BillWayUnknown                   BillWay = "unknown"
	BillWayRechargeWechat            BillWay = "recharge_wechat"
	BillWayRechargeAlipay            BillWay = "recharge_alipay"
	BillWayMissionTime               BillWay = "mission_time"
	BillWayMissionTimePlanHour       BillWay = "mission_time_plan_hour"
	BillWayMissionTimePlanDay        BillWay = "mission_time_plan_day"
	BillWayMissionTimePlanWeek       BillWay = "mission_time_plan_week"
	BillWayMissionTimePlanMonth      BillWay = "mission_time_plan_month"
	BillWayMissionCount              BillWay = "mission_count"
	BillWayMissionHold               BillWay = "mission_hold"
	BillWayMissionVolume             BillWay = "mission_volume"
	BillWayActiveRegister            BillWay = "active_register"
	BillWayActiveShare               BillWay = "active_share"
	BillWayActiveBind                BillWay = "active_bind"
	BillWayActiveInviteRecharge      BillWay = "active_invite_recharge"
	BillWaySpecialChannelRecharge    BillWay = "special_channel_recharge"
	BillWayActiveRecharge            BillWay = "active_recharge"
	BillWayTransferManual            BillWay = "transfer_manual"
	BillWayTransferWithdraw          BillWay = "transfer_withdraw"
	BillWayFirstInviteRecharge       BillWay = "first_invite_recharge"
	BillWayExtraServiceTimePlanHour  BillWay = "extra_service_time_plan_hour"
	BillWayExtraServiceTimePlanDay   BillWay = "extra_service_time_plan_day"
	BillWayExtraServiceTimePlanWeek  BillWay = "extra_service_time_plan_week"
	BillWayExtraServiceTimePlanMonth BillWay = "extra_service_time_plan_month"
	BillWayExtraServiceHold          BillWay = "extra_service_hold"
	BillWayExtraServiceVolume        BillWay = "extra_service_volume"
	BillWayExtraServiceTime          BillWay = "extra_service_time"
	BillWayWithdrawVX                BillWay = "withdraw_wechat"
	BillWayWithdrawAlipay            BillWay = "withdraw_alipay"
	BillWayBankCard                  BillWay = "withdraw_bank_card"
	BillWayWithdrawRefund            BillWay = "withdraw_refund"
	BillWayCdkExchange               BillWay = "cdk_exchange"
	BillWayLottoPrize                BillWay = "lotto_prize"
	BillWayNodeTrouble               BillWay = "node_trouble"       // 节点故障扣费
	BillWayIncomeReplacement         BillWay = "income_replacement" // 收益补发
	BillWayIncomeDeduct              BillWay = "income_deduct"      // 收益扣除
)

func (BillWay) Values() []string {
	return []string{
		string(BillWayBankCard),
		string(BillWayWithdrawAlipay),
		string(BillWayWithdrawVX),
		string(BillWayWithdrawRefund),
		string(BillWayUnknown),
		string(BillWayRechargeWechat),
		string(BillWayRechargeAlipay),
		string(BillWayMissionTime),
		string(BillWayMissionTimePlanHour),
		string(BillWayMissionTimePlanDay),
		string(BillWayMissionTimePlanWeek),
		string(BillWayMissionTimePlanMonth),
		string(BillWayMissionCount),
		string(BillWayActiveBind),
		string(BillWayMissionHold),
		string(BillWayMissionVolume),
		string(BillWayActiveRegister),
		string(BillWayActiveShare),
		string(BillWayActiveRecharge),
		string(BillWayTransferManual),
		string(BillWayFirstInviteRecharge),
		string(BillWayTransferWithdraw),
		string(BillWaySpecialChannelRecharge),
		string(BillWayExtraServiceTimePlanHour),
		string(BillWayExtraServiceTimePlanDay),
		string(BillWayExtraServiceTimePlanWeek),
		string(BillWayExtraServiceTimePlanMonth),
		string(BillWayExtraServiceHold),
		string(BillWayExtraServiceVolume),
		string(BillWayActiveInviteRecharge),
		string(BillWayExtraServiceTime),
		string(BillWayCdkExchange),
		string(BillWayLottoPrize),
		string(BillWayNodeTrouble),
		string(BillWayIncomeReplacement),
		string(BillWayIncomeDeduct),
	}
}

func (obj BillWay) Ptr() *BillWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
