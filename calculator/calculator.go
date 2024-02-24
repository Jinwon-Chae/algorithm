package calculator

import (
	"algorithm/tree/binarytree"
	"strconv"
)

type token interface {
	String() string
	Evaluate(tokens *[]token) (int, bool)
	Priority() int
	MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool)
	EvaluateTree(left, right *binarytree.TreeNode) int
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

func (n number) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return &binarytree.TreeNode{Value: n}, true
}

func (n number) EvaluateTree(left, right *binarytree.TreeNode) int {
	return int(n)
}

func makeOpTreeNode(opToken token, tokens *[]token) (*binarytree.TreeNode, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return nil, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	right, success := top.MakeTreeNode(&newtokens)
	if !success || len(newtokens) == 0 {
		return nil, false
	}

	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	left, success := top.MakeTreeNode(&newtokens)
	if !success {
		return nil, false
	}

	*tokens = newtokens
	return &binarytree.TreeNode{
		Value: opToken,
		Left:  left,
		Right: right,
	}, true
}

type plus struct{}

func (p plus) String() string {
	return "+"
}

func (p plus) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(p, tokens)
}

func (p plus) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("left token shoul be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("right token shoul be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh + rh
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

func (m minus) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(m, tokens)
}

func (m minus) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("left token shoul be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("right token shoul be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh - rh
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

func (m multiple) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(m, tokens)
}

func (m multiple) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("left token shoul be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("right token shoul be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh * rh
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

func (d divide) Priority() int {
	return 2
}

func (d divide) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(d, tokens)
}

func (d divide) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("left token shoul be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("right token shoul be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh / rh
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

func MakeExpressionTree(eval string) (*binarytree.TreeNode, bool) {
	tokens := postfix(eval)

	top, tokens := tokens[len(tokens)-1], tokens[:len(tokens)-1]
	return top.MakeTreeNode(&tokens)
}

func EvaluateExpressionTree(root *binarytree.TreeNode) int {
	t := root.Value.(token)
	if t == nil {
		return 0
	}

	return t.EvaluateTree(root.Left, root.Right)
}
