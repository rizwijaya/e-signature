package html

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func Render(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/dashboard/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/dashboard/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	layouts2, err := filepath.Glob(templatesDir + "/landing/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes2, err := filepath.Glob(templatesDir + "/landing/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include2 := range includes2 {
		layoutCopy2 := make([]string, len(layouts2))
		copy(layoutCopy2, layouts2)
		files2 := append(layoutCopy2, include2)
		r.AddFromFiles(filepath.Base(include2), files2...)
	}

	r.AddFromFiles("error_404.html", templatesDir+"/common/error_404.html")
	r.AddFromFiles("error.html", templatesDir+"/common/error.html")
	r.AddFromFiles("login.html", templatesDir+"/common/login.html")
	r.AddFromFiles("register.html", templatesDir+"/common/register.html")
	return r
}

func ManualRender(tmpDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	//r.AddFromFiles("landing", tmpDir+"landing/landing_index.html", tmpDir+"landing/layouts/base.html")
	return r
}
