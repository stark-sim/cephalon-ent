package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// WithdrawAccount holds the schema definition for the WithdrawAccount entity.
type WithdrawAccount struct {
	ent.Schema
}

// Fields of the WithdrawAccount.
func (WithdrawAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("外键用户 id"),
		field.String("business_name").Default("").StructTag(`json:"business_name"`).Comment("威付通商户名称，对公时为户名"),
		field.Int64("business_id").Default(0).StructTag(`json:"business_id"`).Comment("商户 id"),
		field.Enum("business_type").GoType(enums.BusinessTypeYun).Default(string(enums.BusinessTypeYun)).StructTag(`json:"business_type"`).Comment("商户类型"),
		field.String("id_card").Default("").StructTag(`json:"id_card"`).Comment("身份证号码"),
		field.String("personal_name").Default("未设置昵称").StructTag(`json:"personal_name"`).Comment("个人商户名称"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("个人商户手机号"),
		field.String("bank_card_number").Default("").StructTag(`json:"bank_card_number"`).Comment("银行卡号"),
		field.String("bank").Default("未知银行").StructTag(`json:"bank"`).Comment("开户支行"),
		field.Enum("way").GoType(enums.WithdrawTypeWithdrawAlipay).Default(string(enums.WithdrawTypeUnknown)).StructTag(`json:"way"`).Comment("提现方式"),
		field.String("alipay_card_no").Default("").StructTag(`json:"alipay_card_no"`).Comment("支付宝账户"),
		field.String("company_account").Default("").StructTag(`json:"company_account"`).Comment("对公账号"),
	}
}

// Edges of the WithdrawAccount.
func (WithdrawAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("withdraw_account").Field("user_id").Unique().Required(),
	}
}

// Indexes of the WithdrawAccount
func (WithdrawAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of WithdrawAccount
func (WithdrawAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of WithdrawAccount
func (WithdrawAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("提现账户，用来提供提现渠道"),
	}
}
