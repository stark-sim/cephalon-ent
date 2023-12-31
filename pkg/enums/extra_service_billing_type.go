package enums

type ExtraServiceBillingType string

const (
	ExtraServiceBillingTypeUnknown ExtraServiceBillingType = "unknown"
	// ExtraServiceBillingTypeTimePlanHour 附加服务包时 - 按小时
	ExtraServiceBillingTypeTimePlanHour ExtraServiceBillingType = "time_plan_hour"
	// ExtraServiceBillingTypeTimePlanDay  附加服务包时 - 按天
	ExtraServiceBillingTypeTimePlanDay ExtraServiceBillingType = "time_plan_day"
	// ExtraServiceBillingTypeTimePlanWeek 附加服务包时 - 按周
	ExtraServiceBillingTypeTimePlanWeek ExtraServiceBillingType = "time_plan_week"
	// ExtraServiceBillingTypeTimePlanMonth 附加服务包时 - 按月
	ExtraServiceBillingTypeTimePlanMonth ExtraServiceBillingType = "time_plan_month"
	// ExtraServiceBillingTypeVolume 附加服务按量，单次任务，内部计时 * 时间单价
	ExtraServiceBillingTypeVolume ExtraServiceBillingType = "time_plan_volume"
	// ExtraServiceBillingTypeHold 按位，特点是开启期间具备某些功能，结束后失去功能
	ExtraServiceBillingTypeHold ExtraServiceBillingType = "hold"
	// ExtraServiceBillingTypeTime 按时，时间 * 分钟单价，特点是有供应中状态 supplying
	ExtraServiceBillingTypeTime ExtraServiceBillingType = "time"
)

func (obj ExtraServiceBillingType) Values() []string {
	return []string{
		string(ExtraServiceBillingTypeUnknown),
		string(ExtraServiceBillingTypeTimePlanHour),
		string(ExtraServiceBillingTypeTimePlanDay),
		string(ExtraServiceBillingTypeTimePlanWeek),
		string(ExtraServiceBillingTypeTimePlanMonth),
		string(ExtraServiceBillingTypeVolume),
		string(ExtraServiceBillingTypeHold),
		string(ExtraServiceBillingTypeTime),
	}
}
