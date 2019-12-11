package andnot

/*
	Модуль предоставляет элемент 2И-НЕ
	2И + выход НЕ
*/

import (
	мИ "../and"
)

//Т2Ине -- логический элемент 2И с инверсией на выходе
type Т2Ине struct {
	_2и *мИ.Т2И
}

//Нов2Ине -- возвращает указатель на новый Т2Ине
func Нов2Ине() (ине *Т2Ине) {
	_2ине := Т2Ине{
		_2и: мИ.Нов2И(),
	}
	_2ине._2и.Уст(0, 0)
	return &_2ине
}

//Уст -- устанавливает входы Т2Ине
func (сам *Т2Ине) Уст(пВх1, пВх2 int8) {
	сам._2и.Уст(пВх1, пВх2)
}

//Получ -- получение хранимого значения
func (сам *Т2Ине) Получ() int8 {
	if сам._2и.Получ() == 1 {
		return 0
	}
	return 1
}
