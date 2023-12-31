package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionProduceOrder holds the schema definition for the MissionProduceOrder entity.
type MissionProduceOrder struct {
	ent.Schema
}

// Fields of the MissionProduceOrder.
func (MissionProduceOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键关联用户 id"),
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id"`).Comment("任务 id，关联任务中枢的任务"),
		field.Int64("mission_production_id").Optional().Nillable().StructTag(`json:"mission_production_id"`).Comment("任务执行情况 id"),
		field.Enum("status").GoType(enums.MissionOrderStatusWaiting).Default(string(enums.MissionOrderStatusWaiting)).StructTag(`json:"status"`).Comment("任务订单的状态，注意不强关联任务的状态"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("任务收益的本金 cep 量"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("任务收益的赠送 cep 量"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id""`).Comment("币种 id"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("任务订单分润"),
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeUnknown)).StructTag(`json:"type"`).Comment("任务类型，计时或者次数任务"),
		field.Bool("is_time").Default(false).StructTag(`json:"is_time"`).Comment("是否为计时类型任务"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id"`).Comment("生产者接该任务用的设备 id"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("订单序列号"),
		field.Int64("mission_consume_order_id").Default(0).StructTag(`json:"mission_consume_order_id"`).Comment("外键任务消费订单，一个任务消费订单可能会对应多个任务生产订单"),
	}
}

// Edges of the MissionProduceOrder.
func (MissionProduceOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("mission_produce_orders").Field("user_id").Unique().Required(),
		edge.To("earn_bills", EarnBill.Type),
		edge.From("device", Device.Type).Ref("mission_produce_orders").Field("device_id").Unique().Required(),
		edge.From("mission_consume_order", MissionConsumeOrder.Type).Ref("mission_produce_orders").Field("mission_consume_order_id").Unique().Required(),
		edge.From("mission_production", MissionProduction.Type).Ref("mission_produce_order").Field("mission_production_id").Unique(),
	}
}

// Mixin of MissionProduceOrder
func (MissionProduceOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionProduceOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务生产订单，记录任务生产情况对应的订单业务，比如分润情况，分润比例等信息"),
	}
}
