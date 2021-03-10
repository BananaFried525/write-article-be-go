package controllers

import (
	"log"
	"strconv"
	"strings"
	"write-article/src/models"

	"github.com/gin-gonic/gin"
)

func PostCreateArticle(c *gin.Context) {
	var r models.PostCreateArticleRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Println(err.Error())
		ret["error"] = err.Error()
		c.JSON(400, ret)
	}
	r.ArticleTxt = strings.ReplaceAll(r.ArticleTxt, "'", "''")
	_, err := conn.Exec(`
		INSERT INTO articles (
			article_name,
			article_txt
		)
		VALUES (
			'` + r.ArticleName + `',
			'` + r.ArticleTxt + `'
		)
	`)
	if err != nil {
		ret["error"] = err.Error()
		log.Println(ret)
		c.JSON(500, ret)
		return
	}
	ret["data"] = "Success"
	log.Println(ret)
	c.JSON(200, ret)
}

func GetArticleList(c *gin.Context) {
	var r models.GetArticleRequrst
	thread := make(chan bool)
	if err := c.ShouldBindQuery(&r); err != nil {
		ret["error"] = err.Error()
		c.JSON(400, ret)
		return
	}

	var (
		article     models.ArticleModel
		articleList []models.ArticleModel
	)
	go func(sig chan bool) {
		queryString := `
		SELECT 
			article_id AS articleId,
			article_name AS articleName,
			article_txt AS articleTxt,
			create_at AS createAt,
			update_at AS updateAt
		FROM articles
		`

		if r.ArticleId != 0 || r.ArticleName != "" {
			queryString += `
			WHERE article_id = ` + strconv.Itoa(r.ArticleId) + `
				and article_name  like '%` + r.ArticleName + `%'
			`
		}

		rows, err := conn.Query(queryString)
		if err != nil {
			ret["error"] = err
			log.Println(ret)
			c.JSON(500, ret)
			return
		}
		for rows.Next() {
			err = rows.Scan(
				&article.ArticleId,
				&article.ArticleName,
				&article.ArticleTxt,
				&article.CreateAt,
				&article.UpdateAt,
			)
			if err != nil {
				log.Println(err)
			} else {
				articleList = append(articleList, article)
			}
		}
		sig <- true
	}(thread)
	_ = <-thread

	ret["data"] = articleList
	c.JSON(200, ret)
	return
}
