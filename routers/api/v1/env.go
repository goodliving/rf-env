package v1

import (
	"github.com/gin-gonic/gin"
    "os"
    "path/filepath"
    "fmt"
    "github.com/beevik/etree"
)

func GetEnv(c *gin.Context) {
    
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    env := c.Param("env")

    filename := "env_info.xml"

    abspath := filepath.Join(pwd, filename)

    doc := etree.NewDocument()
    if err := doc.ReadFromFile(abspath); err != nil {
        fmt.Printf("文件读取失败，请检查文件%s是否存在", abspath)
    }

    //读取文件内容
    root := doc.SelectElement("environment")
    fmt.Println("ROOT element:", root.Tag)

    for _, info := range root.SelectElements(env) {
        fmt.Printf("查询环境为%s", env)
        fmt.Println(info)
    }

}