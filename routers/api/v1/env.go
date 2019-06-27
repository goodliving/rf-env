package v1

import (

    "github.com/gin-gonic/gin"
    "goodliving/rf-env/pkg/util"
    "fmt"
    "net/http"

)

func GetEnv(c *gin.Context) {
    
    env := c.Param("env")

    data, err := util.ReadSource(env)
    if err != nil {
        fmt.Println(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : "1000",
        "msg" : "处理成功",
        "env" : env,
        "data" : data,
    })

}