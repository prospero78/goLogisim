package and

/*
	Модуль обеспечивает элемент 2И
*/

import (
	мТранс "../trans"
)

//Т2И -- элемент 2И
type Т2И struct {
	транс1 *мТранс.ТТранс
	транс2 *мТранс.ТТранс
}

//Нов2И -- возвращает указатель на новый Т2И
func Нов2И() (л2и *Т2И) {
	_л2и := Т2И{
		транс1: мТранс.ТрансНов(0),
		транс2: мТранс.ТрансНов(0),
	}
	_л2и.Уст(0, 0)
	return &_л2и
}

//Уст -- устанавливает входы элемента 2И
func (сам *Т2И) Уст(пВх1, пВх2 int8) {
	сам.транс1.БазаУст(пВх1)
	сам.транс2.БазаУст(пВх2)
}

//Получ -- возвращает хранимое значение 2ИЛИ
func (сам *Т2И) Получ() int8 {
	if сам.транс1.Эм() == 1 && сам.транс2.Эм() == 1 {
		return 1
	}
	return 0
}
