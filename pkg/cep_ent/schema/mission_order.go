package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionOrder holds the schema definition for the MissionOrder entity.
type MissionOrder struct {
	ent.Schema
}

// Fields of the MissionOrder.
func (MissionOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id,string"`).Comment("任务 id，外键关联任务"),
		field.Enum("status").GoType(enums.MissionOrderStatusWaiting).Default(string(enums.MissionOrderStatusUnknown)).StructTag(`json:"status"`).Comment("任务订单的状态，注意不强关联任务的状态"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("币种 id"),
		field.Int64("consume_user_id").StructTag(`json:"consume_user_id,string"`).Default(0).Comment("外键关联发起任务的用户 id"),
		field.Int64("consume_amount").Default(0).StructTag(`json:"consume_amount"`).Comment("订单的货币消耗量"),
		field.Int64("produce_user_id").StructTag(`json:"produce_user_id,string"`).Default(0).Comment("外键关联完成任务的用户 id"),
		field.Int64("produce_amount").Default(0).StructTag(`json:"produce_amount"`).Comment("订单的货币分润量"),
		field.Int64("gas_amount").Default(0).StructTag(`json:"gas_amount"`).Comment("订单的平台收量，不记录用户 id，因为都是记载到 genesis 用户"),
		field.Enum("mission_type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeUnknown)).StructTag(`json:"mission_type"`).Comment("任务类型，等于任务表的类型字段"),
		field.Enum("mission_billing_type").GoType(enums.MissionBillingTypeTime).Default(string(enums.MissionBillingTypeUnknown)).StructTag(`json:"mission_billing_type"`).Comment("是否为计时类型任务"),
		field.Enum("call_way").GoType(enums.MissionCallWayApi).Default(string(enums.MissionCallWayUnknown)).StructTag(`json:"call_way"`).Comment("调用方式，API 调用或者微信小程序调用"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("订单序列号"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始执行时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务结束执行时刻"),
		field.Int64("buy_duration").Default(0).StructTag(`json:"buy_duration"`).Comment("包时任务订单购买的时长"),
		field.Time("plan_started_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"plan_started_at"`).Comment("任务计划开始时间（包时）"),
		field.Time("plan_finished_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"plan_finished_at"`).Comment("任务计划结束时间（包时）"),
		field.Time("expired_warning_time").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"expired_warning_time"`).Comment("包时任务到期提醒时间（发了通知提醒就更新该时间）"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id,string"`).Comment("任务批次外键"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号，用于方便检索"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("关联的设备 id"),
		field.Int64("total_amount").Default(0).StructTag(`json:"total_amount"`).Comment("订单总金额"),
		field.Int64("settled_amount").Default(0).StructTag(`json:"settled_amount"`).Comment("已结算金额"),
		field.Int64("settled_count").Default(0).StructTag(`json:"settled_count"`).Comment("已结算次数"),
		field.Int64("total_settle_count").Default(0).StructTag(`json:"total_settle_count"`).Comment("总结算次数"),
		field.Time("lately_settled_at").Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}).Default(common.ZeroTime).StructTag(`json:"lately_settled_at"`).SchemaType(map[string]string{dialect.Postgres: "timestamptz"}).Comment("上一次结算时间"),
	}
}

// Edges of the MissionOrder.
func (MissionOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("consume_user", User.Type).Ref("consume_mission_orders").Field("consume_user_id").Unique().Required(),
		edge.From("produce_user", User.Type).Ref("produce_mission_orders").Field("produce_user_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("mission_orders").Field("symbol_id").Unique().Required(),
		edge.To("bills", Bill.Type),
		edge.From("mission_batch", MissionBatch.Type).Ref("mission_orders").Field("mission_batch_id").Unique().Required(),
		edge.From("mission", Mission.Type).Ref("mission_orders").Field("mission_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("mission_orders").Field("device_id").Unique().Required(),
		edge.To("extra_service_orders", ExtraServiceOrder.Type),
	}
}

// Indexes of the MissionOrder
func (MissionOrder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("consume_user_id"),
		index.Fields("produce_user_id"),
		index.Fields("symbol_id"),
		index.Fields("mission_batch_id"),
		index.Fields("mission_id"),
		index.Fields("device_id"),
	}
}

// Mixin of MissionOrder
func (MissionOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务订单，记录由任务产生的一些金额变动情况，取代任务消耗订单和任务生产订单"),
	}
}
