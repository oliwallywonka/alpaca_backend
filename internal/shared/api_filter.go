package shared

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	/* "github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/sqlite"
	"github.com/go-jet/jet/v2/generator/template"
	sqlite2 "github.com/go-jet/jet/v2/sqlite" */
	//"github.com/go-jet/jet/v2/qrm"
	. "example.com/db/table"
	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Node interface{}
type BinaryExpr struct {
	Left  Node
	Op    string
	Right Node
}
type GroupExpr struct {
	Expr Node
}
type Literal string
type Identifier string

// --------------------
// 🧠 Tokenizer
// --------------------
var tokenRegex = regexp.MustCompile(`\s*(\(|\)|&&|\|\||[=!<>?~!~]+|[A-Za-z_][A-Za-z0-9_.]*|'[^']*'|"[^"]*"|[0-9.]+|true|false|null)\s*`)

func tokenize(input string) []string {
	matches := tokenRegex.FindAllStringSubmatch(input, -1)
	tokens := []string{}
	for _, m := range matches {
		tokens = append(tokens, strings.TrimSpace(m[1]))
	}
	return tokens
}

// --------------------
// 🧩 Parser
// --------------------
type Parser struct {
	tokens []string
	pos    int
}

func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}
func (p *Parser) next() string {
	tok := p.peek()
	p.pos++
	return tok
}

func (p *Parser) parseExpression() (Node, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for {
		op := p.peek()
		if op == "||" {
			p.next()
			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}
			left = &BinaryExpr{Left: left, Op: "OR", Right: right}
		} else {
			break
		}
	}
	return left, nil
}

func (p *Parser) parseTerm() (Node, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for {
		op := p.peek()
		if op == "&&" {
			p.next()
			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			left = &BinaryExpr{Left: left, Op: "AND", Right: right}
		} else {
			break
		}
	}
	return left, nil
}

func (p *Parser) parseFactor() (Node, error) {
	tok := p.peek()
	if tok == "(" {
		p.next()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.next() != ")" {
			return nil, fmt.Errorf("missing closing parenthesis")
		}
		return &GroupExpr{Expr: expr}, nil
	}

	// field op value
	left := p.next()
	op := p.next()
	right := p.next()
	return &BinaryExpr{Left: Identifier(left), Op: op, Right: Literal(right)}, nil
}

//
// ---------------- QUERY BUILDER ----------------
//

func UserRel(table Statement, relation string) {
}

type QueryResult struct {
	SQL   string
	Args  map[string]any
	Where string
}

var tables = struct {
	User *UserTable
	Role *RoleTable
}{
	User: User,
	Role: Role,
}

func GetTableByName(name string) interface{} {
	v := reflect.ValueOf(tables)
	field := v.FieldByName(name)
	if !field.IsValid() {
		return nil
	}
	return field.Type()
}

type Ta struct {
	Table
	AllColumns     ColumnList
	MutableColumns ColumnList
	DefaultColumns ColumnList
}

var selectColumns = map[string]ColumnList{
	"user": User.AllColumns,
	"role": Role.AllColumns,
}

func BuildSelectSQL(rootTable interface{}, tableNames []string) (SelectStatement, error) {
	
	return nil, nil
}

func columnBuilder(tableNames []string) []ColumnList {
	columnsList := make([]ColumnList, 0)
	for _, tableName := range tableNames {
		switch tableName {
		case "user":
			columnsList = append(columnsList, User.AllColumns)
		case "role":
			columnsList = append(columnsList, Role.AllColumns)
		default:
			return nil
		}
	}
	return columnsList
}

func BuildSQLQuery(
	table, filter string,
) (*QueryResult, error) {
	tokens := tokenize(filter)
	parser := &Parser{tokens: tokens}
	ast, err := parser.parseExpression()
	if err != nil {
		return nil, err
	}
	counter := 1
	sql, args, err := buildWhere(ast, &counter)
	if err != nil {
		return nil, err
	}
	final := fmt.Sprintf("SELECT * FROM %s WHERE %s",
		table, sql)
	return &QueryResult{SQL: final, Args: args, Where: sql}, nil
}

func buildWhere(node Node, counter *int) (string, map[string]any, error) {
	switch n := node.(type) {
	case *GroupExpr:
		s, a, e := buildWhere(n.Expr, counter)
		return "(" + s + ")", a, e
	case *BinaryExpr:
		if n.Op == "AND" || n.Op == "OR" {
			lSQL, lArgs, _ := buildWhere(n.Left, counter)
			rSQL, rArgs, _ := buildWhere(n.Right, counter)
			merged := map[string]any{}
			for k, v := range lArgs {
				merged[k] = v
			}
			for k, v := range rArgs {
				merged[k] = v
			}
			return fmt.Sprintf("%s %s %s", lSQL, n.Op, rSQL), merged, nil
		}
		sqlOp := toSQLOperator(n.Op)
		val := parseLiteral(string(n.Right.(Literal)), sqlOp)
		idx := *counter
		(*counter)++
		key := fmt.Sprintf("$%d", idx)
		args := map[string]any{key: val}
		return fmt.Sprintf("%s %s %s", string(n.Left.(Identifier)), sqlOp, key), args, nil
	}
	return "", nil, fmt.Errorf("unexpected node")
}

func parseLiteral(v string, op string) any {
	if strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'") {
		val := strings.Trim(v, "'")
		if op == "LIKE" || op == "NOT LIKE" {
			val = "%" + val + "%"
		}
		return val
	}
	if v == "true" {
		return true
	}
	if v == "false" {
		return false
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	return v
}

func toSQLOperator(op string) string {
	switch op {
	case "=":
		return "="
	case "!=":
		return "<>"
	case ">":
		return ">"
	case ">=":
		return ">="
	case "<":
		return "<"
	case "<=":
		return "<="
	case "~":
		return "LIKE"
	case "!~":
		return "NOT LIKE"
	default:
		return op
	}
}

func main() {
	/* sql, args, err := BuildSQLWhere("(id='abc' && created>'2022-01-01' || name.role = 1) || (status='active' && updated ~ '2023')")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sql)

	// Output: ((id = ? AND created > ?) OR (status = ?))
	fmt.Println(args)
	// Output: [abc 2022-01-01 active] */

	db, err := sql.Open("sqlite3", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/* err := sqlite.GenerateDSN(
		"file:dev.db", "./db",
		template.Default(sqlite2.Dialect).
			UseSchema(func(schema metadata.Schema) template.Schema {
				return template.DefaultSchema(schema).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(column metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(column)
									return defaultTableModelField
								})
						}),
					)
			}),
	)
	if err != nil {
		panic(err)
	} */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filter := r.URL.Query().Get("filter")

		query, err := BuildSQLQuery(
			"user",
			filter,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(query.SQL)
		fmt.Println(query.Args)

		var users []struct {
			ID     string `db:"id" sql:"primary_key"`
			Name   string `db:"name"`
			RoleID string `db:"role_id"`
		}
		total, err := qrm.Query(context.Background(), db, query.SQL, argsMapToSlice(query.Args), &users)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(total)
		fmt.Printf("%+v\n", users)
	})
	http.ListenAndServe(":8090", nil)
}

func argsMapToSlice(m map[string]any) []any {
	if len(m) == 0 {
		return nil
	}
	max := 0
	for k := range m {
		if strings.HasPrefix(k, "$") {
			if n, err := strconv.Atoi(strings.TrimPrefix(k, "$")); err == nil && n > max {
				max = n
			}
		}
	}
	res := make([]any, max)
	for k, v := range m {
		if strings.HasPrefix(k, "$") {
			if n, err := strconv.Atoi(strings.TrimPrefix(k, "$")); err == nil && n >= 1 && n <= max {
				res[n-1] = v
			}
		}
	}
	return res
}
