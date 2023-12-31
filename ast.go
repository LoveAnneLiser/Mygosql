package gosql

type AstKind uint

const (
	SelectKind AstKind = iota
	CreateTableKind
	InsertKind
)

type expressionKind uint

const (
	literalKind expressionKind = iota
)

type expression struct {
	literal *Token
}

// 插入语句具有表名和要插入的值列表：
type InsertStatement struct {
	table  Token
	values *[]expression
}

// 创建语句具有表名以及列名和类型的列表：
type columnDefinition struct {
	name     Token
	datatype Token
}

type CreateTableStatement struct {
	name Token
	cols *[]*columnDefinition
}

// select语句有一个表名和一个列名列表
type SelectStatement struct {
	item []*expression
	from Token
}

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	InsertStatement      *InsertStatement
	Kind                 AstKind
}

type Ast struct {
	Statements []*Statement
}
