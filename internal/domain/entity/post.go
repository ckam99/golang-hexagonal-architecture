package entity

type Post struct {
	ID      int
	Title   string
	Content string
	UserId  int
}

type PostFilter struct {
	Title   string
	Content string
	UserId  int
}
