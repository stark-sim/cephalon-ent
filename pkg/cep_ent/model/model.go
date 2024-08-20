// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the model type in the database.
	Label = "model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldModelType holds the string denoting the model_type field in the database.
	FieldModelType = "model_type"
	// FieldModelStatus holds the string denoting the model_status field in the database.
	FieldModelStatus = "model_status"
	// FieldIsOfficial holds the string denoting the is_official field in the database.
	FieldIsOfficial = "is_official"
	// FieldTotalUsage holds the string denoting the total_usage field in the database.
	FieldTotalUsage = "total_usage"
	// EdgeModelPrices holds the string denoting the model_prices edge name in mutations.
	EdgeModelPrices = "model_prices"
	// EdgeUserModels holds the string denoting the user_models edge name in mutations.
	EdgeUserModels = "user_models"
	// EdgeInvokeModelOrders holds the string denoting the invoke_model_orders edge name in mutations.
	EdgeInvokeModelOrders = "invoke_model_orders"
	// Table holds the table name of the model in the database.
	Table = "models"
	// ModelPricesTable is the table that holds the model_prices relation/edge.
	ModelPricesTable = "model_prices"
	// ModelPricesInverseTable is the table name for the ModelPrice entity.
	// It exists in this package in order to avoid circular dependency with the "modelprice" package.
	ModelPricesInverseTable = "model_prices"
	// ModelPricesColumn is the table column denoting the model_prices relation/edge.
	ModelPricesColumn = "model_id"
	// UserModelsTable is the table that holds the user_models relation/edge.
	UserModelsTable = "user_models"
	// UserModelsInverseTable is the table name for the UserModel entity.
	// It exists in this package in order to avoid circular dependency with the "usermodel" package.
	UserModelsInverseTable = "user_models"
	// UserModelsColumn is the table column denoting the user_models relation/edge.
	UserModelsColumn = "model_id"
	// InvokeModelOrdersTable is the table that holds the invoke_model_orders relation/edge.
	InvokeModelOrdersTable = "invoke_model_orders"
	// InvokeModelOrdersInverseTable is the table name for the InvokeModelOrder entity.
	// It exists in this package in order to avoid circular dependency with the "invokemodelorder" package.
	InvokeModelOrdersInverseTable = "invoke_model_orders"
	// InvokeModelOrdersColumn is the table column denoting the invoke_model_orders relation/edge.
	InvokeModelOrdersColumn = "model_id"
)

// Columns holds all SQL columns for model fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldAuthor,
	FieldDescription,
	FieldModelType,
	FieldModelStatus,
	FieldIsOfficial,
	FieldTotalUsage,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultAuthor holds the default value on creation for the "author" field.
	DefaultAuthor string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultIsOfficial holds the default value on creation for the "is_official" field.
	DefaultIsOfficial bool
	// DefaultTotalUsage holds the default value on creation for the "total_usage" field.
	DefaultTotalUsage int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultModelType enums.Model = "unknown"

// ModelTypeValidator is a validator for the "model_type" field enum values. It is called by the builders before save.
func ModelTypeValidator(mt enums.Model) error {
	switch mt {
	case "unknown", "language":
		return nil
	default:
		return fmt.Errorf("model: invalid enum value for model_type field: %q", mt)
	}
}

const DefaultModelStatus enums.ModelStatus = "init"

// ModelStatusValidator is a validator for the "model_status" field enum values. It is called by the builders before save.
func ModelStatusValidator(ms enums.ModelStatus) error {
	switch ms {
	case "unknown", "init":
		return nil
	default:
		return fmt.Errorf("model: invalid enum value for model_status field: %q", ms)
	}
}

// OrderOption defines the ordering options for the Model queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByModelType orders the results by the model_type field.
func ByModelType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelType, opts...).ToFunc()
}

// ByModelStatus orders the results by the model_status field.
func ByModelStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelStatus, opts...).ToFunc()
}

// ByIsOfficial orders the results by the is_official field.
func ByIsOfficial(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOfficial, opts...).ToFunc()
}

// ByTotalUsage orders the results by the total_usage field.
func ByTotalUsage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalUsage, opts...).ToFunc()
}

// ByModelPricesCount orders the results by model_prices count.
func ByModelPricesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModelPricesStep(), opts...)
	}
}

// ByModelPrices orders the results by model_prices terms.
func ByModelPrices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModelPricesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserModelsCount orders the results by user_models count.
func ByUserModelsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserModelsStep(), opts...)
	}
}

// ByUserModels orders the results by user_models terms.
func ByUserModels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserModelsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvokeModelOrdersCount orders the results by invoke_model_orders count.
func ByInvokeModelOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvokeModelOrdersStep(), opts...)
	}
}

// ByInvokeModelOrders orders the results by invoke_model_orders terms.
func ByInvokeModelOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvokeModelOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newModelPricesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModelPricesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ModelPricesTable, ModelPricesColumn),
	)
}
func newUserModelsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserModelsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserModelsTable, UserModelsColumn),
	)
}
func newInvokeModelOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvokeModelOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvokeModelOrdersTable, InvokeModelOrdersColumn),
	)
}