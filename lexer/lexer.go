// 词法分析器, 将源代码作为输入，并输出对应的词法单元
// 在生产环境中，应该将文件名和行号附加到词法单元中，以便更好地跟踪可能出现的词法分析错误和语法分析错误。

package lexer

type Lexer struct {
    input        string
    position     int  // 所输入字符串中的当前位置（指向当前字符）
    readPosition int  // 所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
    ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    return l
}
