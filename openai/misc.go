package openai

import (
	"github.com/Alp4ka/classifier"
)

type classes []classifier.Class

func (c classes) Classes() []classifier.ClassStruct {
	ret := make([]classifier.ClassStruct, 0, len(c))
	for i := range c {
		ret = append(ret, c[i].Class())
	}
	return ret
}

var _allowedSymbols = []uint8("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNMйцукенгшщзхъфывапролдячсмитьбюёЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮЁ\n\r\t().,-=+?!$#@%&*\\/1234567890`'\":;>< ")
var _allowedSymbolsSet = make(map[uint8]struct{})

func init() {
	for _, sym := range _allowedSymbols {
		_allowedSymbolsSet[sym] = struct{}{}
	}
}
