package entities

type Post struct{
	Id  int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	LikeCount int `json:"like_count"`
	DisLikeCount int `json:"dislike_count"`
	IdUser int `json:"iduser"`
}

type PostResponse struct{
	Id  int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	LikeCount int `json:"like_count"`
	DisLikeCount int `json:"dislike_count"`
	Date string `json:"created_at"`
	IdUser       int    `json:"iduser"`
	UserName     string `json:"user_name"`
	UserCareer   string `json:"user_career"`
}