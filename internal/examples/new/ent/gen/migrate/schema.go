// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DefaultModelsColumns holds the columns for the "default_models" table.
	DefaultModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
	}
	// DefaultModelsTable holds the schema information for the "default_models" table.
	DefaultModelsTable = &schema.Table{
		Name:       "default_models",
		Columns:    DefaultModelsColumns,
		PrimaryKey: []*schema.Column{DefaultModelsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DefaultModelsTable,
	}
)

func init() {
}