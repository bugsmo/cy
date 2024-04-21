package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*Time)(nil)

type Time struct{ mixin.Schema }

func (Time) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreateTime{}.Fields()...)
	fields = append(fields, UpdateTime{}.Fields()...)
	fields = append(fields, DeleteTime{}.Fields()...)
	return fields
}

var _ ent.Mixin = (*CreateTime)(nil)

type CreateTime struct{ mixin.Schema }

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*UpdateTime)(nil)

type UpdateTime struct{ mixin.Schema }

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Comment("更新时间").
			Immutable().
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*DeleteTime)(nil)

type DeleteTime struct{ mixin.Schema }

func (DeleteTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Comment("删除时间").
			Immutable().
			Optional().
			Nillable(),
	}
}
