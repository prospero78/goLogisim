package trigrssync

/*
	Модуль предоставляет синхронный RS-триггер

*/

import(
	м2Ине "../andnot"
	мТриг "../trigrs"
)
//ТСинхРсТриг -- синхронный RS-триггер
type ТСинхРсТриг struct{
	устИне *м2Ине.Т2Ине
	сбросИне *м2Ине.Т2Ине
	тригРС *мТриг.ТРсТриггер
}

func СинхРсТригНов()(триг *ТСинхРсТриг){
	
}