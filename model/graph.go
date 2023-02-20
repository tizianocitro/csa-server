package model

type GraphData struct {
	Description GraphDescription `json:"description"`
	Edges       []GraphEdge      `json:"edges"`
	Nodes       []GraphNode      `json:"nodes"`
}

type GraphNodeData struct {
	IsUrlHashed bool   `json:"isUrlHashed"`
	Label       string `json:"label"`
	URL         string `json:"url"`
}

type GraphNodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GraphNode struct {
	Data     GraphNodeData     `json:"data"`
	ID       string            `json:"id"`
	Position GraphNodePosition `json:"position"`
	Type     string            `json:"type"`
}

type GraphEdge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Type   string `json:"type"`
}

type GraphDescription struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
