package http

type Attach struct {
	SourceId uint64 `json:"source_id"`
	TargetId uint64 `json:"target_id"`
}

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Delete struct {
	Id uint64 `json:"target_id"`
}

type ResponseList struct {
	Items []*User `json:"items"`
}
