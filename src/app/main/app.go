package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/handler"
)

// reference our echo instance and create it early
var e = createMux()

// ------
// types
// ------

// TemplateRegistry ...
type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// -----------
// initialize
// -----------

func init() {
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	// Static assets
	e.Static("/static", "asset")

	// Templates
	templates := make(map[string]*template.Template)
	templates["home"] = registerTemplate("layout", "pages/static/home")
	templates["help"] = registerTemplate("layout", "pages/static/help")
	templates["about"] = registerTemplate("layout", "pages/static/about")
	templates["contact"] = registerTemplate("layout", "pages/static/contact")
	templates["howdy"] = registerTemplate("layout", "howdy")

	// Renerer
	e.Renderer = &TemplateRegistry{Templates: templates}

	// Route => handler
	e.GET("/", handler.HomeHandler)
	e.GET("/help", handler.HelpHandler)
	e.GET("/about", handler.AboutHandler)
	e.GET("/contact", handler.ContactHandler)
	e.GET("/howdy", handler.HowdyHandler)
	e.GET("/:message", handler.ParrotHandler)
}

// -------
// custom
// -------

// Render ...
func (tr *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, ok := tr.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	// Add global methods if data is a map
	if value, ok := data.(map[string]interface{}); ok {
		value["reverse"] = c.Echo().Reverse
	}

	return t.ExecuteTemplate(w, "layout", data)
}

// ------
// utils
// ------

func registerTemplate(filenames ...string) *template.Template {
	var files []string
	for _, value := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", value))
	}

	// Templateで使用する関数を登録する
	funcMap := template.FuncMap{
		"debugPrint": DebugPrint,
	}

	t := template.New(files[0]).Funcs(funcMap)

	return template.Must(t.ParseFiles(files...))
}

// DebugPrint ...
func DebugPrint(data interface{}) (res string) {
	// relectの基礎型を取得する
	rv := reflect.ValueOf(data)
	rvk := rv.Kind()

	// ポインタ型やインタフェース型の場合は中身の型を取得する
	switch rvk {
	case reflect.Ptr, reflect.Interface:
		rv = rv.Elem()
		rvk = rv.Kind()
	}

	// 型で場合分けをして出力文字列を生成する
	var output string
	switch rvk {
	case reflect.Struct:
		output = DebugPrintReflectStruct(rv)
	case reflect.Array, reflect.Slice:
		output = DebugPrintReflectArrayOrSlice(rv)
	default:
		output = fmt.Sprint(rv.Interface())
	}
	res = fmt.Sprintf("T: %v, V: %v", rv.Type(), output)
	return
}

// DebugPrintReflectStruct ...
func DebugPrintReflectStruct(rv reflect.Value) string {
	var fields []string

	// reflect.Type型を取得する
	rvt := rv.Type()

	// structのフィールド数でループ処理をする
	for i := 0; i < rv.NumField(); i++ {
		fk := rvt.Field(i) // reflect.StructField型
		fv := rv.Field(i)  // reflect.Value型

		// structのフィールド名を取得する
		fkn := fk.Name

		// structが外部公開されている場合はフィールドの値を取得する
		var fvi interface{}
		if fk.PkgPath == "" {
			fvi = fv.Interface()
		}

		// スライスに要素を詰める
		fields = append(fields, fmt.Sprintf("%v: %v", fkn, fvi))
	}
	// スライスの中身をセパレーターを用いて結合する
	str := strings.Join(fields, ", ")
	return "{" + str + "}"
}

// DebugPrintReflectArrayOrSlice ...
func DebugPrintReflectArrayOrSlice(rv reflect.Value) (str string) {
	// 配列・スライスの中身（0番目の要素）の型がStructの場合は専用のロジックで文字列を生成する
	if rv.Index(0).Kind() == reflect.Struct {
		var elements []string
		for i := 0; i < rv.Len(); i++ {
			// i番目の要素のreflect.Valueを取得する
			rv := rv.Index(i)

			// スライスに要素を詰める
			elements = append(elements, fmt.Sprintf("%d: %s", i, DebugPrintReflectStruct(rv)))
		}
		// スライスの中身をセパレーターを用いて結合する
		str = strings.Join(elements, ", ")

		str = "[" + str + "]"
	} else {
		str = fmt.Sprint(rv.Interface())
	}
	return
}
