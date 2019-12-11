package trigrs

/*
	Модуль реализует RS-триггер.
	Обратите внимание, что прямой выход -- это выход СБРОСА!!!
	А обратный выход -- это выход УСТАНОВКИ!!!
*/

import (
	мИне "../andnot"
)

//ТРсТриггер -- модель RS-триггер
type ТРсТриггер struct {
	сброс *мИне.Т2Ине
	уст   *мИне.Т2Ине
}

//РсТриггерНов -- возвращает указатель на ТРсТриггер
func РсТриггерНов() (рс *ТРсТриггер) {
	_рс := ТРсТриггер{
		сброс: мИне.Нов2Ине(),
		уст:   мИне.Нов2Ине(),
	}
	_рс.Сброс()
	return &_рс
}

//Уст -- установить RS-триггер
func (сам *ТРсТриггер) Уст() {
	сам.сброс.Уст(0, сам.уст.Получ())
	сам.уст.Уст(сам.сброс.Получ(), 1)
}

//Сброс -- сбрасывает RS-триггер
func (сам *ТРсТриггер) Сброс() {
	сам.сброс.Уст(1, сам.уст.Получ())
	сам.уст.Уст(сам.сброс.Получ(), 0)
}

//Выход -- возвращает прямой выход RS-триггера
func (сам *ТРсТриггер) Выход() int8 {
	return сам.сброс.Получ()
}

//ВыходНе -- возвращает инверсный выход RS-триггера
func (сам *ТРсТриггер) ВыходНе() int8 {
	return сам.уст.Получ()
}
