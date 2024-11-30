package recognizer

type RecognizeResponse map[string][][]string

type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FileResult struct {
	Filename  string     `json:"filename"`
	IsWolf    bool       `json:"is_wolf"`
	Intervals []Interval `json:"intervals"`
}

type RecognizeResult struct {
	Result []FileResult `json:"result"`
}
