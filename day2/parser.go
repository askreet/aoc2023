package day2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

const (
	TGame = iota
	TNumber
	TColor
	TComma
	TColon
	TSemicolon
	TEOF
	TNewline
)

var _TyName = map[int]string{
	TGame:      "TGame",
	TNumber:    "TNumber",
	TColor:     "TColor",
	TComma:     "TComma",
	TColon:     "TColon",
	TSemicolon: "TSemicolon",
	TEOF:       "TEOF",
	TNewline:   "TNewline",
}

type Token struct {
	Ty     int
	IntVal int
	StrVal string
	Line   int
	Col    int
}

func (t Token) TyName() string {
	return _TyName[t.Ty]
}

type Tokens []Token

type Reveal struct {
	Color  string
	Number int
}

func (t *Tokens) AddGame() *Tokens {
	*t = append(*t, Token{Ty: TGame})
	return t
}

func (t *Tokens) AddColor(name string) *Tokens {
	*t = append(*t, Token{Ty: TColor, StrVal: name})
	return t
}

func (t *Tokens) AddNumber(val int) *Tokens {
	*t = append(*t, Token{Ty: TNumber, IntVal: val})
	return t
}

func (t *Tokens) AddColon() *Tokens {
	*t = append(*t, Token{Ty: TColon})
	return t
}

func (t *Tokens) AddComma() *Tokens {
	*t = append(*t, Token{Ty: TComma})
	return t
}

func (t *Tokens) AddSemicolon() *Tokens {
	*t = append(*t, Token{Ty: TSemicolon})
	return t
}

func (t *Tokens) AddNewline() *Tokens {
	*t = append(*t, Token{Ty: TNewline})
	return t
}

func (t *Tokens) AddEOF() *Tokens {
	*t = append(*t, Token{Ty: TEOF})
	return t
}

func (t *Tokens) At(line, col int) {
	lastToken := &(*t)[len(*t)-1]
	lastToken.Line = line
	lastToken.Col = col
}

func Lex(in io.Reader) (Tokens, error) {
	line := 1
	col := 1

	word := bytes.NewBuffer([]byte{})
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanRunes)
	var tokens Tokens

	for scanner.Scan() {
		text := scanner.Bytes()
		switch text[0] {
		case ':':
			if word.Len() > 0 {
				if intVal, err := strconv.Atoi(word.String()); err == nil {
					tokens.AddNumber(intVal).At(line, col-word.Len())
				} else {
					return nil, fmt.Errorf("unexpected TColon at line %d, col %d", line, col)
				}
			} else {
				return nil, fmt.Errorf("unexpected TColon at line %d, col %d", line, col)
			}
			tokens.AddColon().At(line, col)
			word.Truncate(0)
		case ';':
			if word.Len() > 0 {
				tokens.AddColor(word.String()).At(line, col-word.Len())
			} else {
				return nil, fmt.Errorf("unexpected TSemicolon at line %d, col %d", line, col)
			}
			tokens.AddSemicolon().At(line, col)
			word.Truncate(0)
		case ',':
			if word.Len() > 0 {
				tokens.AddColor(word.String()).At(line, col-word.Len())
			} else {
				return nil, fmt.Errorf("unexpected TComma at line %d, col %d", line, col)
			}
			tokens.AddComma().At(line, col)
			word.Truncate(0)
		case ' ':
			fallthrough
		case '\n':
			if word.Len() > 0 {
				if intVal, err := strconv.Atoi(word.String()); err == nil {
					tokens.AddNumber(intVal).At(line, col-word.Len())
				} else if word.String() == "Game" {
					tokens.AddGame().At(line, col-len("Game"))
				} else {
					tokens.AddColor(word.String()).At(line, col-word.Len())
				}
			}

			if text[0] == '\n' {
				tokens.AddNewline().At(line, col)
				line += 1
				col = 1
			}

			word.Truncate(0)
		default:
			word.Write(text)
		}

		col += 1
	}

	if word.Len() > 0 {
		tokens.AddColor(word.String()).At(line, col-word.Len())
	}

	tokens.AddEOF().At(line, col)

	return tokens, nil
}

type TokenStream struct {
	tokens Tokens
	idx    int
}

func (ts *TokenStream) Peek() *Token {
	if len(ts.tokens) < ts.idx+1 {
		return &ts.tokens[ts.idx+1]
	} else {
		return nil
	}
}

func (ts *TokenStream) Take() *Token {
	if len(ts.tokens) > ts.idx {
		token := &ts.tokens[ts.idx]
		ts.idx++
		return token
	} else {
		panic("attempt to take tokens past end of stream")
	}
}

func (ts *TokenStream) HasNext() bool {
	return len(ts.tokens) > ts.idx+1
}

func Parse(tokens Tokens) ([]Game, error) {
	stream := TokenStream{tokens, 0}
	var games []Game
	var token *Token

	for stream.HasNext() {
		token = stream.Take()
		switch token.Ty {
		case TGame:
			break
		case TEOF:
			goto Done
		default:
			return nil, UnexpectedErr(token)
		}

		if token = stream.Take(); token.Ty != TNumber {
			return nil, ExpectedErr(TNumber, token)
		}
		game := Game{Id: token.IntVal}

		if token = stream.Take(); token.Ty != TColon {
			return nil, ExpectedErr(TColon, token)
		}

		var set []Reveal
		for {
			if token = stream.Take(); token.Ty != TNumber {
				return nil, ExpectedErr(TNumber, token)
			}
			reveal := Reveal{Number: token.IntVal}

			if token = stream.Take(); token.Ty != TColor {
				return nil, ExpectedErr(TColor, token)
			}
			reveal.Color = token.StrVal
			set = append(set, reveal)

			token = stream.Take()
			switch token.Ty {
			case TComma:
				continue

			case TSemicolon:
				game.Sets = append(game.Sets, set)
				set = nil
				break

			case TNewline:
				fallthrough
			case TEOF:
				game.Sets = append(game.Sets, set)
				games = append(games, game)
				goto NextGame

			default:
				return nil, UnexpectedErr(token)
			}
		}

	NextGame:
	}

Done:
	return games, nil
}

func ExpectedErr(want int, token *Token) error {
	return fmt.Errorf("expected %v, found %v at line %d, col %d",
		(*Token).TyName(&Token{Ty: want}), token.TyName(), token.Line, token.Col)
}

func UnexpectedErr(token *Token) error {
	return fmt.Errorf("unexpected %v at line %d, col %d",
		token.TyName(), token.Line, token.Col)
}
