package calculator

import (
	"strconv"
)

type token interface {
	String() string
	Evaluate(tokens *[]token) (int, bool)
	Priority() int
}

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func (n number) Evaluate(tokens *[]token) (int, bool) {
	return int(n), true
}

func (n number) Priority() int {
	return 0
}

type plus struct{}

func (p plus) String() string {
	return "+"
}

func (p plus) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}

	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh + rh, true
}

func (p plus) Priority() int {
	return 1
}

type minus struct{}

func (m minus) String() string {
	return "-"
}

func (m minus) Priority() int {
	return 1
}

func (m minus) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}

	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh - rh, true
}

type multiple struct{}

func (m multiple) String() string {

	return "*"
}

func (m multiple) Priority() int {
	return 2
}

func (m multiple) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}

	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh * rh, true
}

type divide struct{}

func (d divide) String() string {
	return "/"
}

func (s divide) Priority() int {
	return 2
}

func (d divide) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}

	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh / rh, true
}

type parser struct {
	eval        []rune
	idx         int
	parsedToken token
}

func (p *parser) parse() bool {
	// ignore spaces
	for {
		if p.idx >= len(p.eval) {
			return false
		}

		if p.eval[p.idx] != ' ' {
			break
		}
		p.idx++
	}

	if p.eval[p.idx] >= '0' && p.eval[p.idx] <= '9' {
		var value int
		for p.idx < len(p.eval) {
			if p.eval[p.idx] >= '0' && p.eval[p.idx] <= '9' {
				value *= 10
				value += int(p.eval[p.idx] - '0')
				p.idx++
			} else {
				break
			}
		}

		p.parsedToken = number(value)
		return true
	} else if p.eval[p.idx] == '+' {
		p.parsedToken = plus{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '-' {
		p.parsedToken = minus{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '*' {
		p.parsedToken = multiple{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '/' {
		p.parsedToken = divide{}
		p.idx++
		return true
	} else {
		return false
	}
}

func tokenize(eval string) []token {
	tokens := []token{}
	p := &parser{
		eval: []rune(eval),
	}

	for p.parse() {
		tokens = append(tokens, p.parsedToken)

	}

	return tokens
}

func postfix(eval string) []token {
	tokens := tokenize(eval)
	if len(tokens) == 0 {
		return tokens
	}

	postfix := make([]token, 0, len(tokens))
	ops := []token{}
	for i := range tokens {
		if no, ok := tokens[i].(number); ok {
			postfix = append(postfix, no)
		} else {
			for len(ops) > 0 && ops[len(ops)-1].Priority() >= tokens[i].Priority() {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}

			ops = append(ops, tokens[i])
		}
	}

	for len(ops) > 0 {
		postfix = append(postfix, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	return postfix
}

func Evaluate(eval string) (rst int, success bool) {
	tokens := postfix(eval)

	top, tokens := tokens[len(tokens)-1], tokens[:len(tokens)-1]
	rst, success = top.Evaluate(&tokens)
	return rst, success
}
