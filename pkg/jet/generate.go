package jetfx

import (
	"log"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/sqlite"
	"github.com/go-jet/jet/v2/generator/template"
	sqlite2 "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func Generate() {
	err := sqlite.GenerateDSN("file:dev.db", "./db",
		template.Default(sqlite2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseSQLBuilder(template.DefaultSQLBuilder().
						UseTable(func(table metadata.Table) template.TableSQLBuilder {
							if table.Name == "_prisma_migrations" {
								return template.TableSQLBuilder{
									Skip: true,
								}
							}
							return template.DefaultTableSQLBuilder(table)
						}),
					).
					UseModel(template.DefaultModel().
						// IGNORE _prisma_migrations table
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(column metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(column)
									if shared.HasImageField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.ImageField{})
									}
									if shared.HasLangField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.LangField{})
									}
									if shared.HasPermissionField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.PermissionField{})
									}
									if shared.HasContactField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.ContactField{})
									}
									if shared.HasPriceField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.PriceField{})
									}
									if shared.HasLocationField(table.Name, column.Name) {
										defaultTableModelField.Type = template.NewType(shared.LocationField{})
									}
									if column.Name == "created_at" || column.Name == "updated_at" {
										defaultTableModelField.Type = template.NewType(int64(0))
									}
									return defaultTableModelField
								})
						}),
					)
			}),
	)

	if err != nil {
		log.Fatal(err)
	}
}
