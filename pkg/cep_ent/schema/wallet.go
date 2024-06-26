package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("外键用户 id"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("外键币种 id"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("货币余额"),
		field.Int64("total_amount").Default(0).StructTag(`json:"total_amount"`).Comment("货币总量，当货币是收益货币时，代表总收益，当货币是充值货币时，代表总充值金额"),
		field.Int64("withdraw_amount").Default(0).StructTag(`json:"withdraw_amount"`).Comment("已提现金额，目前只有一种货币可以提现"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("wallets").Field("user_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("wallets").Field("symbol_id").Unique().Required(),
	}
}

// Mixin of Wallet
func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Indexes of Wallet
func (Wallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "symbol_id", "deleted_at").Unique(),
		index.Fields("user_id"),
		index.Fields("symbol_id"),
	}
}

// Annotations of Wallet
func (Wallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("钱包，作为用户和各币种的中间关系，记录各余额"),
	}
}
