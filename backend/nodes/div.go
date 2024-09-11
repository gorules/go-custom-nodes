package nodes

import "github.com/gorules/zen-go"

type divNode struct {
}

func (a divNode) Handle(request zen.NodeRequest) (zen.NodeResponse, error) {
	left, err := zen.GetNodeField[float64](request, "a")
	if err != nil {
		return zen.NodeResponse{}, err
	}

	right, err := zen.GetNodeField[float64](request, "b")
	if err != nil {
		return zen.NodeResponse{}, err
	}

	key, err := zen.GetNodeFieldRaw[string](request, "key")
	if err != nil {
		return zen.NodeResponse{}, err
	}

	output := make(map[string]any)
	output[key] = left / right

	return zen.NodeResponse{Output: output}, nil
}