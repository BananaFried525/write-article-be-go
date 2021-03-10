package models

type ArticleModel struct {
	ArticleId   int    `json:"articleId"`
	ArticleName string `json:"articleName"`
	ArticleTxt  string `json:"articleTxt"`
	CreateAt    string `json:"createAt"`
	UpdateAt    string `json:"updateAt"`
}

type PostCreateArticleRequest struct {
	ArticleName string `json:"articleName" binding:"required"`
	ArticleTxt  string `json:"articleTxt"`
}

type GetArticleRequrst struct {
	ArticleId   int    `json:"articleId" form:"id"`
	ArticleName string `json:"articleName" form:"name"`
}

type DeleteArticleRequest struct {
	Id int `json:"id" form:"id"`
}
