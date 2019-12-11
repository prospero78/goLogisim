package trans

/*
	Модуль предоставляет тест модели транзистора
*/

import (
	мТест "testing"
	мВрем "time"
)

const (
	подтяжПит   = 1
	подтяжЗемля = 0
	подтяжПлав  = -1
)

var (
	транс *ТТранс
)

func TestТранс(тест *мТест.T) {
	_Проверить := func(пЗнач int8) {
		if пЗнач == 1 || пЗнач == 0 {
			if транс.Эм() != пЗнач {
				тест.Errorf("_Проверить(): ОШИБКА эммитер(%v)!=%v\n", транс.Эм(), пЗнач)
			}
			if транс.Кол() == пЗнач {
				тест.Errorf("_Проверить(): ОШИБКА коллектор(%v)==%v\n", транс.Кол(), пЗнач)
			}
		}

	}
	_Создание := func() {
		тест.Logf("_Создание()\n")
		if транс = ТрансНов(подтяжПлав); транс == nil {
			тест.Errorf("_Создание(): ОШИБКА транс не может быть nil\n")
		}
		if транс.База() != -1 {
			тест.Errorf("_Проверить(): ОШИБКА база(%v)!=-1\n", транс.База())
		}
		_Проверить(подтяжПлав)
	}
	_БазаУст := func() {
		тест.Logf("_БазаУст()\n")
		транс.БазаУст(подтяжПит)
		мВрем.Sleep(мВрем.Millisecond * 2)
		if транс.База() != подтяжПит {
			тест.Errorf("_БазаУст(): ОШИБКА база(%v)!=%v\n", транс.База(), подтяжПит)
		}
		_Проверить(подтяжПит)
	}
	_ПлавУст := func() {
		for транс.База() != -1 {
			транс.БазаУст(подтяжПлав)
		}
		знач := транс.База()
		_Проверить(знач)
		for транс.Подтяжка() != -1 {
			транс.ПодтяжкаУст(подтяжПлав)
		}
		знач = транс.База()
		_Проверить(знач)
	}
	_БазаПлохУст := func(пЗнач int8) {
		defer func() {
			if паника := recover(); паника == nil {
				тест.Errorf("_БазаПлохУст(): ОШИБКА в нулевой панике\n")
			}
		}()
		тест.Logf("_БазаПлохУст()\n")
		транс.БазаУст(пЗнач)
		//мВрем.Sleep(мВрем.Millisecond*2)
	}
	_ПодтяжПлохУст := func(пЗнач int8) {
		defer func() {
			if паника := recover(); паника == nil {
				тест.Errorf("_ПодтяжПлохУст(): ОШИБКА в нулевой панике\n")
			}
		}()
		тест.Logf("_ПодтяжПлохУст()\n")
		транс.ПодтяжкаУст(пЗнач)
		//мВрем.Sleep(мВрем.Millisecond*2)
	}
	_БазаПодтяжПит := func() {
		тест.Logf("_БазаПодтяж()\n")
		транс.БазаУст(подтяжПлав)
		транс.ПодтяжкаУст(подтяжПит)
		мВрем.Sleep(мВрем.Millisecond * 2)
		if транс.База() != подтяжПит {
			тест.Errorf("_БазаПодтяжПит(): ОШИБКА база(%v)!=%v\n", транс.База(), подтяжПит)
		}
		_Проверить(подтяжПит)
	}
	_БазаПодтяжЗемля := func() {
		тест.Logf("_БазаПодтяжЗемля()\n")
		транс.БазаУст(подтяжПлав)
		транс.ПодтяжкаУст(подтяжЗемля)
		мВрем.Sleep(мВрем.Millisecond * 2)
		if транс.База() != подтяжЗемля {
			тест.Errorf("_БазаПодтяжЗемля(): ОШИБКА база(%v)!=%v\n", транс.База(), подтяжЗемля)
		}
		_Проверить(подтяжЗемля)
	}
	_БазаПодтяжПлав := func() {
		тест.Logf("_БазаПодтяжПлав()\n")
		транс.БазаУст(подтяжПлав)
		транс.ПодтяжкаУст(подтяжПлав)
		мВрем.Sleep(мВрем.Millisecond * 2)
		if транс.База() != подтяжПлав {
			тест.Errorf("_БазаПодтяжПлав(): ОШИБКА база(%v)!=%v\n", транс.База(), подтяжПлав)
		}
		//_Проверить(подтяжПлав)
	}
	_Создание()
	_БазаУст()
	_ПлавУст()
	_БазаПлохУст(-2)
	_ПодтяжПлохУст(-2)
	_БазаПодтяжПит()
	_БазаПодтяжЗемля()
	_БазаПодтяжПлав()
	_БазаПодтяжПлав()
	_БазаПодтяжПлав()
}
