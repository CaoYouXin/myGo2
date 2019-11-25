package router

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func somejson(engine *gin.Engine) {
	engine.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.JSON(http.StatusOK, data)
	})
	engine.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
	engine.GET("/pureJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.PureJSON(http.StatusOK, data)
	})
	engine.GET("/secureJSON", func(c *gin.Context) {
		// data := map[string]interface{}{
		// 	"lang": "GO语言",
		// 	"tag":  "<br/>",
		// }

		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})
}

func xml(engine *gin.Engine) {
	engine.GET("/xml", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.XML(http.StatusOK, gin.H(data))
	})
}

func yaml(engine *gin.Engine) {
	engine.GET("/yaml", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		}

		c.YAML(http.StatusOK, data)
	})
}

type protobufexample map[string]interface{}

func (data *protobufexample) Marshal() ([]byte, error) {
	return json.Marshal(data)
}
func (data *protobufexample) Reset() {}
func (data *protobufexample) String() string {
	bytes, err := data.Marshal()
	if err != nil {
		log.Fatalf("protobuf: %v\n", err)
		return ""
	}
	return string(bytes)
}
func (data *protobufexample) ProtoMessage() {}

func protobuf(engine *gin.Engine) {
	engine.GET("/protobuf", func(c *gin.Context) {
		data := protobufexample(map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br/>",
		})

		c.ProtoBuf(http.StatusOK, &data)
	})
}

func formpost(engine *gin.Engine) {
	engine.POST("/query_form_post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"id":      id,
			"page":    page,
			"message": message,
			"nick":    nick,
		})
	})
}

func upload(engine *gin.Engine) {
	engine.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		f, _ := file.Open()
		input := bufio.NewScanner(f)

		output := make([]string, 0, 3)
		for input.Scan() {
			output = append(output, input.Text())
		}

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "%s", strings.Join(output, "<br/>"))
	})

	engine.POST("/upload/2/seperated", func(c *gin.Context) {
		file1, _ := c.FormFile("file1")
		file2, _ := c.FormFile("file2")

		c.String(http.StatusOK, file1.Filename+" "+file2.Filename)
	})

	engine.POST("/upload/2/in/array", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		outputs := make([]string, 0, 3)
		for _, f := range files {
			outputs = append(outputs, f.Filename)
		}

		c.String(http.StatusOK, strings.Join(outputs, " "))
	})
}

func download(engine *gin.Engine) {
	engine.GET("gopher.png", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment: "gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
}

func checkCookie(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		act, err := c.Cookie("act")
		if err != nil || act != key {
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

func auth(engine *gin.Engine) {
	doc := engine.Group("/doc")
	doc.Use(checkCookie("doc"))
	doc.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "doctors ping OK!")
	})

	fac := engine.Group("/fac")
	fac.Use(checkCookie("fac"))
	fac.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "factory ping OK!")
	})
}
