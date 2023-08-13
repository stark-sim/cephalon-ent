// Code generated by ent, DO NOT EDIT.

package hmackeypair

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the hmackeypair type in the database.
	Label = "hmac_key_pair"
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
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldCaller holds the string denoting the caller field in the database.
	FieldCaller = "caller"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeMissionProductions holds the string denoting the mission_productions edge name in mutations.
	EdgeMissionProductions = "mission_productions"
	// EdgeCreatedMissions holds the string denoting the created_missions edge name in mutations.
	EdgeCreatedMissions = "created_missions"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the hmackeypair in the database.
	Table = "hmac_key_pairs"
	// MissionProductionsTable is the table that holds the mission_productions relation/edge.
	MissionProductionsTable = "mission_productions"
	// MissionProductionsInverseTable is the table name for the MissionProduction entity.
	// It exists in this package in order to avoid circular dependency with the "missionproduction" package.
	MissionProductionsInverseTable = "mission_productions"
	// MissionProductionsColumn is the table column denoting the mission_productions relation/edge.
	MissionProductionsColumn = "hmac_key_pair_id"
	// CreatedMissionsTable is the table that holds the created_missions relation/edge.
	CreatedMissionsTable = "missions"
	// CreatedMissionsInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	CreatedMissionsInverseTable = "missions"
	// CreatedMissionsColumn is the table column denoting the created_missions relation/edge.
	CreatedMissionsColumn = "hmac_key_pair_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "hmac_key_pairs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_hmac_key_pair"
)

// Columns holds all SQL columns for hmackeypair fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldKey,
	FieldSecret,
	FieldCaller,
	FieldUserID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hmac_key_pairs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_hmac_key_pair",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultCaller holds the default value on creation for the "caller" field.
	DefaultCaller string
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the HmacKeyPair queries.
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

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByCaller orders the results by the caller field.
func ByCaller(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaller, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByMissionProductionsCount orders the results by mission_productions count.
func ByMissionProductionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMissionProductionsStep(), opts...)
	}
}

// ByMissionProductions orders the results by mission_productions terms.
func ByMissionProductions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionProductionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedMissionsCount orders the results by created_missions count.
func ByCreatedMissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedMissionsStep(), opts...)
	}
}

// ByCreatedMissions orders the results by created_missions terms.
func ByCreatedMissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedMissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newMissionProductionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionProductionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MissionProductionsTable, MissionProductionsColumn),
	)
}
func newCreatedMissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedMissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedMissionsTable, CreatedMissionsColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
