package Controllers

import (
    models "first/Model"
    "first/Repository"
    utils "first/Utils"

    "github.com/gin-gonic/gin"
    "net/http"
)

func GetArticles(c *gin.Context) {

    pagination := utils.GeneratePaginationFromRequest(c)
    var article models.Article
    articleList, err := Repository.GetAllArticles(&article, &pagination)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err,
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "data": articleList,
    })

}

func GetArticleById(c *gin.Context) {
    var article models.Article

    articleList, err := Repository.FindArticleById(&article, c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "GetArticleById error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "ID":      articleList.ID,
        "Title":   articleList.Title,
        "Content": articleList.Content,
    })
    return
}
