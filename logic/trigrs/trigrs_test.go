package trigrs

/*
	Модуль предоставляет тест для RS-триггера
*/

import (
	мТест "testing"
	мВрем "time"
)

var (
	триг *ТРсТриггер
)

func TestРсТРиггер(тест *мТест.T) {
	_Проверить := func(пВх int8) {
		if пВх == 1 {
			рез := триг.Выход()
			if рез != 1 {
				тест.Errorf("1.1 _Проверить(): ОШИБКА триг должен быть установлен\n")
			}
			рез = триг.ВыходНе()
			if рез != 0 {
				тест.Errorf("1.2 _Проверить(): ОШИБКА триг должен быть установлен\n")
			}
		} else {
			рез := триг.Выход()
			if рез != 0 {
				тест.Errorf("1.3 _Проверить(): ОШИБКА триг должен быть сброшен\n")
			}
			рез = триг.ВыходНе()
			if рез != 1 {
				тест.Errorf("1.4 _Проверить(): ОШИБКА триг должен быть сброшен\n")
			}
		}
	}
	_Создать := func() {
		тест.Logf("_Создать()\n")
		if триг = РсТриггерНов(); триг == nil {
			тест.Errorf("_Создать(): ОШИБКА триг не может быть nil\n")
		}
		мВрем.Sleep(мВрем.Microsecond * 5)
		_Проверить(0)
	}
	_Уст := func() {
		тест.Logf("_Уст()\n")
		триг.Уст()
		мВрем.Sleep(мВрем.Microsecond * 5)
		_Проверить(1)
	}
	_Создать()
	_Уст()
}