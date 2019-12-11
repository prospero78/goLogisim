package and

/*
	Модуль предоставляет тест элемента 2И
*/

import (
	мТест "testing"
)

var (
	_2и *Т2И
)

func Test2Или(тест *мТест.T) {
	_Создать := func() {
		if _2и = Нов2И(); _2и == nil {
			тест.Errorf("_Создать(): ОШИБКА 2и не может быть nil\n")
		}
	}
	_И := func(п1, п2 int8) {
		_2и.Уст(п1, п2)
		if п1+п2 == 2 {
			if _2и.Получ() != 1 {
				тест.Errorf("_И(): ОШИБКА 2и не равно 1\n")
			}
		} else {
			if _2и.Получ() != 0 {
				тест.Errorf("_И(): ОШИБКА 2и не равно 9\n")
			}
		}
	}
	_Создать()
	_И(0, 0)
	_И(1, 0)
	_И(0, 1)
	_И(1, 1)
}