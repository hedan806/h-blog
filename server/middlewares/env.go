package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	// 文件视图跟目录.
	viewPathPattern = "server/views/*/*.html"

	// 当前域
	Host = ""

	// 静态域
	staticCluster []string

	// 时差
	timeZoneOffset = int64(8)

	adminMenus = []map[string]string{
		{"url": "/om/setting", "name": "网站设置"},
		{"url": "/om/user", "name": "用户管理"},
		{"url": "/om/cms/category", "name": "文章管理"},
		{"url": "/om/picture", "name": "图片管理"},
		{"url": "/om/tag", "name": "Tag管理"},
		{"url": "/om/keyword", "name": "关键字管理"},
	}

	adminSubMenus = map[string][]map[string]string{
		"/om/setting": {
			{"url": "/om/setting", "name": "网站设置"},
		},
		"/om/user": {
			{"url": "/om/user", "name": "用户管理"},
		},
		"/om/cms/category": {
			{"url": "/om/cms/category", "name": "文章管理"},
		},
		"/om/picture": {
			{"url": "/om/picture", "name": "图片管理"},
		},
		"/om/tag": {
			{"url": "/om/tag", "name": "Tag管理"},
		},
		"/om/keyword": {
			{"url": "/om/keyword", "name": "关键字管理"},
		},
	}
)

func SetEnv(app *gin.Engine) {

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	loadHTMLGlob(app)
}

func loadHTMLGlob(app *gin.Engine) {
	if gin.IsDebugging() {
		templ := getTemplateObj()
		debugPrintLoadTemplate(templ)
		app.HTMLRender = render.HTMLDebug{Glob: viewPathPattern}
	} else {
		templ := getTemplateObj()
		app.SetHTMLTemplate(templ)
	}
}

// Reconstruction the gin same name function in gin.go for Custiom tmplate functions.
func debugPrintLoadTemplate(tmpl *template.Template) {
	if gin.IsDebugging() {
		var buf bytes.Buffer
		for _, tmpl := range tmpl.Templates() {
			buf.WriteString("\t- ")
			buf.WriteString(tmpl.Name())
			buf.WriteString("\n")
		}
		debugPrint("Loaded HTML Templates (%d): \n%s\n", len(tmpl.Templates()), buf.String())
	}
}

// Reconstruction the gin same name function in gin.go for Custiom tmplate functions.
func debugPrint(format string, values ...interface{}) {
	if gin.IsDebugging() {
		log.Printf("[EasyCMS-debug] "+format, values...)
	}
}

func getTemplateObj() *template.Template {
	funcMap := template.FuncMap{
		"staticUrl":      getStaticUrl,
		"staticDomain":   getStaticDomain,
		"host":           getHost,
		"showTitle":      showTitle,
		"showAdminTitle": showAdminTitle,
		//"adminSession":    func() sessions.Session { return CurAdminSession },
		//"isGuestForAdmin": IsGuestForAdmin,
		"showAdminMenus": showAdminMenus,
		"getKeywords":    getKeywords,
		"getDescription": getDescription,
		"getCurTimestamp": func() int64 {
			return time.Now().Unix()
		},
		"html": func(text string) template.HTML {
			return template.HTML(text)
		},
		"loadtimes": func(startTime time.Time) string {
			// 加载时间
			return fmt.Sprintf("%dms", time.Now().Sub(startTime)/1000000)
		},
		"url": func(url string) string {
			// 没有http://或https://开头的增加http://
			if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
				return url
			}

			return "http://" + url
		},
		"add": func(a, b int) int {
			// 加法运算
			return a + b
		},
		//"renderInput":  wtforms.RenderInput,
		//"renderInputH": wtforms.RenderInputH,
		"formatdate": func(t time.Time) string {
			// 格式化日期
			return t.Format(time.RFC822)
		},
		"formattime": func(t time.Time) string {
			// 格式化时间
			now := time.Now()
			duration := now.Sub(t)
			if duration.Seconds() < 60 {
				return fmt.Sprintf("刚刚")
			} else if duration.Minutes() < 60 {
				return fmt.Sprintf("%.0f 分钟前", duration.Minutes())
			} else if duration.Hours() < 24 {
				return fmt.Sprintf("%.0f 小时前", duration.Hours())
			}
			t = t.Add(time.Hour * time.Duration(timeZoneOffset))
			return t.Format("2006-01-02 15:04")
		},
		"formatdatetime": func(t time.Time) string {
			// 格式化时间成 2006-01-02 15:04:05
			return t.Add(time.Hour * time.Duration(timeZoneOffset)).Format("2006-01-02 15:04:05")
		},
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(text, "\n", "<br>", -1))
		},
	}

	return template.Must(template.New("").Funcs(funcMap).ParseGlob(viewPathPattern))
}

func getKeywords() string {
	return "keyword"
}

func getDescription() string {
	return "description"
}

func getThumb(file, way string, width, height int) string {
	return getStaticUrl(fmt.Sprintf("thumb/%s/%d/%d/%s", way, width, height, file))
}

func getStaticUrl(file string) string {
	staticHostLen := len(staticCluster)
	domain := ""
	if staticHostLen == 1 {
		domain = staticCluster[0] + "/" + file
	} else if staticHostLen > 1 {
		// get rand static host url source
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		domain = staticCluster[r.Intn(staticHostLen)] + "/" + file
	}

	if domain != "" {
		domain = "http://" + domain
	}

	return domain
}

func getStaticDomain() string {
	staticUrl := staticCluster[0]
	domains := strings.Split(staticUrl, ".")
	domainLen := len(domains)

	return fmt.Sprintf("%s.%s", domains[domainLen-2], domains[domainLen-1])
}

func getHost() string {
	return Host
}

func showTitle() string {
	return "Honlyc"
}

func showAdminTitle(name string) string {
	return "" + name
}

// to render the admin menu list.
func showAdminMenus() template.HTML {
	menuStr := ""

	for _, one := range adminMenus {
		menuStr += `<dl class="menu expand">
    <dt class="menu_title clickable"><a href="` + one["url"] + `"><label>` + one["name"] + `</label></a></dt>
`
		for _, child := range adminSubMenus[ one["url"] ] {
			menuStr += `<dd class="menu_item"><a href="` + child["url"] + `"><span>` + child["name"] + `</span></a></dd>
`
		}

		menuStr += "</dl>"
	}

	return template.HTML(menuStr)
}
