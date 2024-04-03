package token

type TokenType string

type Token struct {
    // 词法单元类型
    Type    TokenType
    // 字面量
    Literal string
}
// token/token.go

const (
    ILLEGAL = "ILLEGAL"
    // End of file
    EOF     = "EOF"

    // 标识符+字面量
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1343456

    // 运算符
    ASSIGN = "="
    PLUS   = "+"

    // 分隔符
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // 关键字
    FUNCTION = "FUNCTION"
    LET      = "LET"
)
