// Demo code for the TextView primitive.
package main

import (
	"github.com/rivo/tview"
)

const corporate = `See: https://ru.wikipedia.org/wiki/Go
	Go (часто также Golang) — компилируемый многопоточный язык программирования, разработанный внутри компании Google[4]. Разработка Go началась в сентябре 2007 года, его непосредственным проектированием занимались Роберт Гризмер, Роб Пайк и Кен Томпсон[5], занимавшиеся до этого проектом разработки операционной системы Inferno. Официально язык был представлен в ноябре 2009 года.
	На данный момент поддержка официального компилятора, разрабатываемого создателями языка, осуществляется для операционных систем FreeBSD, OpenBSD, Linux, macOS, Windows, DragonFly BSD, Plan 9, Solaris, Android, AIX.[6].
	
See: https://golang.org/

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}`

func main() {
	app := tview.NewApplication()
	editBox := tview.NewEditBox(). // TODO: need EditBox
					SetWrap(true).          // TODO: need modify cursor movement modification if false
					SetDynamicColors(true). // TODO: not clear - may be remove
					SetChangedFunc(func() {
			app.Draw()
		}).
		SetText(corporate)
	editBox.SetBorder(true) // TODO: add test with and without border
	if err := app.SetRoot(editBox, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
