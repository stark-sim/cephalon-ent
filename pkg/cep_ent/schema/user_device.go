package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserDevice struct {
	ent.Schema
}

func (UserDevice) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("外键用户 id"),
		field.Int64("device_id").StructTag(`json:"device_id,string"`).Default(0).Comment("外键设备 id"),
	}
}

// Edges of the ServerOrder.
func (UserDevice) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("user_devices").Field("user_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("user_devices").Field("device_id").Unique().Required(),
	}
}

// Mixin of ServerOrder
func (UserDevice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (UserDevice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户与设备的中间关系，依赖时间字段记录用户与设备的解绑换绑等操作"),
	}
}
