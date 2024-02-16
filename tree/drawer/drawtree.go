package drawer

import (
	"algorithm/tree/nodeinterface"
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type drawTreeNode struct {
	Value any

	X int
	Y int

	Child []*drawTreeNode
}

func makeDrawTree(node nodeinterface.Node, level int, order *int) *drawTreeNode {
	if node == nil {
		return nil
	}

	drawNode := &drawTreeNode{
		Value: node.GetValue(),
		Y:     level,
	}

	// in-order
	childs := node.GetChilds()
	half := len(childs) / 2
	for i := 0; i < half; i++ {
		child := childs[i]
		drawNode.Child = append(drawNode.Child, makeDrawTree(child, level-1, order))
	}

	// setting x position
	drawNode.X = *order
	(*order)++

	// right-side
	for i := half; i < len(childs); i++ {
		child := childs[i]
		drawNode.Child = append(drawNode.Child, makeDrawTree(child, level-1, order))
	}

	return drawNode
}

func SaveTreeGraph(t nodeinterface.Node, filepath string) error {
	var order int
	drawTree := makeDrawTree(t, 0, &order)
	if drawTree == nil {
		return fmt.Errorf("empty tree")
	}

	// plot
	p := plot.New()

	// drawing cycle
	var xys plotter.XYs
	drawTree.getLocations(&xys)

	points, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{G: 255, A: 255}
	points.Radius = vg.Points(20)

	// draw lines
	err = drawLines(drawTree, p)
	if err != nil {
		return err
	}
	p.Add(points)

	// add labels
	err = addLabels(drawTree, p)
	if err != nil {
		return err
	}

	return p.Save(1000, 600, filepath)
}

func (d *drawTreeNode) getLocations(xys *plotter.XYs) {
	*xys = append(*xys, plotter.XY{
		X: float64(d.X),
		Y: float64(d.Y),
	})

	for _, c := range d.Child {
		if c == nil {
			continue
		}

		c.getLocations(xys)
	}
}

func drawLines(node *drawTreeNode, p *plot.Plot) error {
	for _, c := range node.Child {
		if c == nil {
			continue
		}

		pts := plotter.XYs{
			{X: float64(node.X), Y: float64(node.Y)},
			{X: float64(c.X), Y: float64(c.Y)},
		}

		line, err := plotter.NewLine(pts)
		if err != nil {
			return err
		}

		sctter, err := plotter.NewScatter(pts)
		if err != nil {
			return err
		}

		p.Add(line, sctter)

		err = drawLines(c, p)
		if err != nil {
			return err
		}
	}

	return nil
}

func addLabels(node *drawTreeNode, p *plot.Plot) error {
	label, err := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{X: float64(node.X), Y: float64(node.Y)},
		},
		Labels: []string{fmt.Sprint(node.Value)},
	})
	if err != nil {
		return err
	}

	p.Add(label)
	for _, c := range node.Child {
		if c == nil {
			continue
		}

		err = addLabels(c, p)
		if err != nil {
			return err
		}
	}

	return nil
}
