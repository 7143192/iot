package defines

// TestBody used for go request test.
type TestBody struct {
	Val1 string
	Val2 string
}

type Auth struct {
	IdentityInfo Identity `json:"identity"`
}

type Identity struct {
	Methods []string `json:"methods"`
	PwdInfo Password `json:"password"`
}

type Password struct {
	UserInfo User `json:"user"`
}

type User struct {
	DomainInfo Domain `json:"domain"`
	Name       string `json:"name"`
	Pwd        string `json:"password"`
}

type Domain struct {
	Name string `json:"name"`
}

type TestEntireBody struct {
	MapInfo   *RoadMap `json:"mapInfo"`
	GraphInfo *Graph   `json:"graphInfo"`
	AuthInfo  Auth     `json:"auth"`
}
