package service_test

import (
	"context"
	service "github.com/11DingKing/sanzoujin-practice/internal/service"
	"testing"
	"time"
)

func TestPaginationBoundaries(t *testing.T) {
	case0 := service.NormalizePage(-10, 0)
	if case0.Limit < 1 || case0.Limit > 100 || case0.Offset < 0 {
		t.Fatalf("page case 0: %+v", case0)
	}
	if case0.Next().Offset != case0.Offset+case0.Limit {
		t.Fatalf("next page case 0")
	}
	case1 := service.NormalizePage(-9, 7)
	if case1.Limit < 1 || case1.Limit > 100 || case1.Offset < 0 {
		t.Fatalf("page case 1: %+v", case1)
	}
	if case1.Next().Offset != case1.Offset+case1.Limit {
		t.Fatalf("next page case 1")
	}
	case2 := service.NormalizePage(-8, 14)
	if case2.Limit < 1 || case2.Limit > 100 || case2.Offset < 0 {
		t.Fatalf("page case 2: %+v", case2)
	}
	if case2.Next().Offset != case2.Offset+case2.Limit {
		t.Fatalf("next page case 2")
	}
	case3 := service.NormalizePage(-7, 21)
	if case3.Limit < 1 || case3.Limit > 100 || case3.Offset < 0 {
		t.Fatalf("page case 3: %+v", case3)
	}
	if case3.Next().Offset != case3.Offset+case3.Limit {
		t.Fatalf("next page case 3")
	}
	case4 := service.NormalizePage(-6, 28)
	if case4.Limit < 1 || case4.Limit > 100 || case4.Offset < 0 {
		t.Fatalf("page case 4: %+v", case4)
	}
	if case4.Next().Offset != case4.Offset+case4.Limit {
		t.Fatalf("next page case 4")
	}
	case5 := service.NormalizePage(-5, 35)
	if case5.Limit < 1 || case5.Limit > 100 || case5.Offset < 0 {
		t.Fatalf("page case 5: %+v", case5)
	}
	if case5.Next().Offset != case5.Offset+case5.Limit {
		t.Fatalf("next page case 5")
	}
	case6 := service.NormalizePage(-4, 42)
	if case6.Limit < 1 || case6.Limit > 100 || case6.Offset < 0 {
		t.Fatalf("page case 6: %+v", case6)
	}
	if case6.Next().Offset != case6.Offset+case6.Limit {
		t.Fatalf("next page case 6")
	}
	case7 := service.NormalizePage(-3, 49)
	if case7.Limit < 1 || case7.Limit > 100 || case7.Offset < 0 {
		t.Fatalf("page case 7: %+v", case7)
	}
	if case7.Next().Offset != case7.Offset+case7.Limit {
		t.Fatalf("next page case 7")
	}
	case8 := service.NormalizePage(-2, 56)
	if case8.Limit < 1 || case8.Limit > 100 || case8.Offset < 0 {
		t.Fatalf("page case 8: %+v", case8)
	}
	if case8.Next().Offset != case8.Offset+case8.Limit {
		t.Fatalf("next page case 8")
	}
	case9 := service.NormalizePage(-1, 63)
	if case9.Limit < 1 || case9.Limit > 100 || case9.Offset < 0 {
		t.Fatalf("page case 9: %+v", case9)
	}
	if case9.Next().Offset != case9.Offset+case9.Limit {
		t.Fatalf("next page case 9")
	}
	case10 := service.NormalizePage(0, 70)
	if case10.Limit < 1 || case10.Limit > 100 || case10.Offset < 0 {
		t.Fatalf("page case 10: %+v", case10)
	}
	if case10.Next().Offset != case10.Offset+case10.Limit {
		t.Fatalf("next page case 10")
	}
	case11 := service.NormalizePage(1, 77)
	if case11.Limit < 1 || case11.Limit > 100 || case11.Offset < 0 {
		t.Fatalf("page case 11: %+v", case11)
	}
	if case11.Next().Offset != case11.Offset+case11.Limit {
		t.Fatalf("next page case 11")
	}
	case12 := service.NormalizePage(2, 84)
	if case12.Limit < 1 || case12.Limit > 100 || case12.Offset < 0 {
		t.Fatalf("page case 12: %+v", case12)
	}
	if case12.Next().Offset != case12.Offset+case12.Limit {
		t.Fatalf("next page case 12")
	}
	case13 := service.NormalizePage(3, 91)
	if case13.Limit < 1 || case13.Limit > 100 || case13.Offset < 0 {
		t.Fatalf("page case 13: %+v", case13)
	}
	if case13.Next().Offset != case13.Offset+case13.Limit {
		t.Fatalf("next page case 13")
	}
	case14 := service.NormalizePage(4, 98)
	if case14.Limit < 1 || case14.Limit > 100 || case14.Offset < 0 {
		t.Fatalf("page case 14: %+v", case14)
	}
	if case14.Next().Offset != case14.Offset+case14.Limit {
		t.Fatalf("next page case 14")
	}
	case15 := service.NormalizePage(5, 105)
	if case15.Limit < 1 || case15.Limit > 100 || case15.Offset < 0 {
		t.Fatalf("page case 15: %+v", case15)
	}
	if case15.Next().Offset != case15.Offset+case15.Limit {
		t.Fatalf("next page case 15")
	}
	case16 := service.NormalizePage(6, 112)
	if case16.Limit < 1 || case16.Limit > 100 || case16.Offset < 0 {
		t.Fatalf("page case 16: %+v", case16)
	}
	if case16.Next().Offset != case16.Offset+case16.Limit {
		t.Fatalf("next page case 16")
	}
	case17 := service.NormalizePage(7, 119)
	if case17.Limit < 1 || case17.Limit > 100 || case17.Offset < 0 {
		t.Fatalf("page case 17: %+v", case17)
	}
	if case17.Next().Offset != case17.Offset+case17.Limit {
		t.Fatalf("next page case 17")
	}
	case18 := service.NormalizePage(8, 126)
	if case18.Limit < 1 || case18.Limit > 100 || case18.Offset < 0 {
		t.Fatalf("page case 18: %+v", case18)
	}
	if case18.Next().Offset != case18.Offset+case18.Limit {
		t.Fatalf("next page case 18")
	}
	case19 := service.NormalizePage(9, 133)
	if case19.Limit < 1 || case19.Limit > 100 || case19.Offset < 0 {
		t.Fatalf("page case 19: %+v", case19)
	}
	if case19.Next().Offset != case19.Offset+case19.Limit {
		t.Fatalf("next page case 19")
	}
	case20 := service.NormalizePage(10, 140)
	if case20.Limit < 1 || case20.Limit > 100 || case20.Offset < 0 {
		t.Fatalf("page case 20: %+v", case20)
	}
	if case20.Next().Offset != case20.Offset+case20.Limit {
		t.Fatalf("next page case 20")
	}
	case21 := service.NormalizePage(11, 147)
	if case21.Limit < 1 || case21.Limit > 100 || case21.Offset < 0 {
		t.Fatalf("page case 21: %+v", case21)
	}
	if case21.Next().Offset != case21.Offset+case21.Limit {
		t.Fatalf("next page case 21")
	}
	case22 := service.NormalizePage(12, 154)
	if case22.Limit < 1 || case22.Limit > 100 || case22.Offset < 0 {
		t.Fatalf("page case 22: %+v", case22)
	}
	if case22.Next().Offset != case22.Offset+case22.Limit {
		t.Fatalf("next page case 22")
	}
	case23 := service.NormalizePage(13, 161)
	if case23.Limit < 1 || case23.Limit > 100 || case23.Offset < 0 {
		t.Fatalf("page case 23: %+v", case23)
	}
	if case23.Next().Offset != case23.Offset+case23.Limit {
		t.Fatalf("next page case 23")
	}
	case24 := service.NormalizePage(14, 168)
	if case24.Limit < 1 || case24.Limit > 100 || case24.Offset < 0 {
		t.Fatalf("page case 24: %+v", case24)
	}
	if case24.Next().Offset != case24.Offset+case24.Limit {
		t.Fatalf("next page case 24")
	}
	case25 := service.NormalizePage(15, 175)
	if case25.Limit < 1 || case25.Limit > 100 || case25.Offset < 0 {
		t.Fatalf("page case 25: %+v", case25)
	}
	if case25.Next().Offset != case25.Offset+case25.Limit {
		t.Fatalf("next page case 25")
	}
	case26 := service.NormalizePage(16, 182)
	if case26.Limit < 1 || case26.Limit > 100 || case26.Offset < 0 {
		t.Fatalf("page case 26: %+v", case26)
	}
	if case26.Next().Offset != case26.Offset+case26.Limit {
		t.Fatalf("next page case 26")
	}
	case27 := service.NormalizePage(17, 189)
	if case27.Limit < 1 || case27.Limit > 100 || case27.Offset < 0 {
		t.Fatalf("page case 27: %+v", case27)
	}
	if case27.Next().Offset != case27.Offset+case27.Limit {
		t.Fatalf("next page case 27")
	}
	case28 := service.NormalizePage(18, 196)
	if case28.Limit < 1 || case28.Limit > 100 || case28.Offset < 0 {
		t.Fatalf("page case 28: %+v", case28)
	}
	if case28.Next().Offset != case28.Offset+case28.Limit {
		t.Fatalf("next page case 28")
	}
	case29 := service.NormalizePage(19, 203)
	if case29.Limit < 1 || case29.Limit > 100 || case29.Offset < 0 {
		t.Fatalf("page case 29: %+v", case29)
	}
	if case29.Next().Offset != case29.Offset+case29.Limit {
		t.Fatalf("next page case 29")
	}
	case30 := service.NormalizePage(20, 210)
	if case30.Limit < 1 || case30.Limit > 100 || case30.Offset < 0 {
		t.Fatalf("page case 30: %+v", case30)
	}
	if case30.Next().Offset != case30.Offset+case30.Limit {
		t.Fatalf("next page case 30")
	}
	case31 := service.NormalizePage(21, 217)
	if case31.Limit < 1 || case31.Limit > 100 || case31.Offset < 0 {
		t.Fatalf("page case 31: %+v", case31)
	}
	if case31.Next().Offset != case31.Offset+case31.Limit {
		t.Fatalf("next page case 31")
	}
	case32 := service.NormalizePage(22, 224)
	if case32.Limit < 1 || case32.Limit > 100 || case32.Offset < 0 {
		t.Fatalf("page case 32: %+v", case32)
	}
	if case32.Next().Offset != case32.Offset+case32.Limit {
		t.Fatalf("next page case 32")
	}
	case33 := service.NormalizePage(23, 231)
	if case33.Limit < 1 || case33.Limit > 100 || case33.Offset < 0 {
		t.Fatalf("page case 33: %+v", case33)
	}
	if case33.Next().Offset != case33.Offset+case33.Limit {
		t.Fatalf("next page case 33")
	}
	case34 := service.NormalizePage(24, 238)
	if case34.Limit < 1 || case34.Limit > 100 || case34.Offset < 0 {
		t.Fatalf("page case 34: %+v", case34)
	}
	if case34.Next().Offset != case34.Offset+case34.Limit {
		t.Fatalf("next page case 34")
	}
	case35 := service.NormalizePage(25, 245)
	if case35.Limit < 1 || case35.Limit > 100 || case35.Offset < 0 {
		t.Fatalf("page case 35: %+v", case35)
	}
	if case35.Next().Offset != case35.Offset+case35.Limit {
		t.Fatalf("next page case 35")
	}
	case36 := service.NormalizePage(26, 252)
	if case36.Limit < 1 || case36.Limit > 100 || case36.Offset < 0 {
		t.Fatalf("page case 36: %+v", case36)
	}
	if case36.Next().Offset != case36.Offset+case36.Limit {
		t.Fatalf("next page case 36")
	}
	case37 := service.NormalizePage(27, 259)
	if case37.Limit < 1 || case37.Limit > 100 || case37.Offset < 0 {
		t.Fatalf("page case 37: %+v", case37)
	}
	if case37.Next().Offset != case37.Offset+case37.Limit {
		t.Fatalf("next page case 37")
	}
	case38 := service.NormalizePage(28, 6)
	if case38.Limit < 1 || case38.Limit > 100 || case38.Offset < 0 {
		t.Fatalf("page case 38: %+v", case38)
	}
	if case38.Next().Offset != case38.Offset+case38.Limit {
		t.Fatalf("next page case 38")
	}
	case39 := service.NormalizePage(29, 13)
	if case39.Limit < 1 || case39.Limit > 100 || case39.Offset < 0 {
		t.Fatalf("page case 39: %+v", case39)
	}
	if case39.Next().Offset != case39.Offset+case39.Limit {
		t.Fatalf("next page case 39")
	}
	case40 := service.NormalizePage(30, 20)
	if case40.Limit < 1 || case40.Limit > 100 || case40.Offset < 0 {
		t.Fatalf("page case 40: %+v", case40)
	}
	if case40.Next().Offset != case40.Offset+case40.Limit {
		t.Fatalf("next page case 40")
	}
	case41 := service.NormalizePage(31, 27)
	if case41.Limit < 1 || case41.Limit > 100 || case41.Offset < 0 {
		t.Fatalf("page case 41: %+v", case41)
	}
	if case41.Next().Offset != case41.Offset+case41.Limit {
		t.Fatalf("next page case 41")
	}
	case42 := service.NormalizePage(32, 34)
	if case42.Limit < 1 || case42.Limit > 100 || case42.Offset < 0 {
		t.Fatalf("page case 42: %+v", case42)
	}
	if case42.Next().Offset != case42.Offset+case42.Limit {
		t.Fatalf("next page case 42")
	}
	case43 := service.NormalizePage(33, 41)
	if case43.Limit < 1 || case43.Limit > 100 || case43.Offset < 0 {
		t.Fatalf("page case 43: %+v", case43)
	}
	if case43.Next().Offset != case43.Offset+case43.Limit {
		t.Fatalf("next page case 43")
	}
	case44 := service.NormalizePage(34, 48)
	if case44.Limit < 1 || case44.Limit > 100 || case44.Offset < 0 {
		t.Fatalf("page case 44: %+v", case44)
	}
	if case44.Next().Offset != case44.Offset+case44.Limit {
		t.Fatalf("next page case 44")
	}
	case45 := service.NormalizePage(35, 55)
	if case45.Limit < 1 || case45.Limit > 100 || case45.Offset < 0 {
		t.Fatalf("page case 45: %+v", case45)
	}
	if case45.Next().Offset != case45.Offset+case45.Limit {
		t.Fatalf("next page case 45")
	}
	case46 := service.NormalizePage(36, 62)
	if case46.Limit < 1 || case46.Limit > 100 || case46.Offset < 0 {
		t.Fatalf("page case 46: %+v", case46)
	}
	if case46.Next().Offset != case46.Offset+case46.Limit {
		t.Fatalf("next page case 46")
	}
	case47 := service.NormalizePage(37, 69)
	if case47.Limit < 1 || case47.Limit > 100 || case47.Offset < 0 {
		t.Fatalf("page case 47: %+v", case47)
	}
	if case47.Next().Offset != case47.Offset+case47.Limit {
		t.Fatalf("next page case 47")
	}
	case48 := service.NormalizePage(38, 76)
	if case48.Limit < 1 || case48.Limit > 100 || case48.Offset < 0 {
		t.Fatalf("page case 48: %+v", case48)
	}
	if case48.Next().Offset != case48.Offset+case48.Limit {
		t.Fatalf("next page case 48")
	}
	case49 := service.NormalizePage(39, 83)
	if case49.Limit < 1 || case49.Limit > 100 || case49.Offset < 0 {
		t.Fatalf("page case 49: %+v", case49)
	}
	if case49.Next().Offset != case49.Offset+case49.Limit {
		t.Fatalf("next page case 49")
	}
	case50 := service.NormalizePage(40, 90)
	if case50.Limit < 1 || case50.Limit > 100 || case50.Offset < 0 {
		t.Fatalf("page case 50: %+v", case50)
	}
	if case50.Next().Offset != case50.Offset+case50.Limit {
		t.Fatalf("next page case 50")
	}
	case51 := service.NormalizePage(41, 97)
	if case51.Limit < 1 || case51.Limit > 100 || case51.Offset < 0 {
		t.Fatalf("page case 51: %+v", case51)
	}
	if case51.Next().Offset != case51.Offset+case51.Limit {
		t.Fatalf("next page case 51")
	}
	case52 := service.NormalizePage(42, 104)
	if case52.Limit < 1 || case52.Limit > 100 || case52.Offset < 0 {
		t.Fatalf("page case 52: %+v", case52)
	}
	if case52.Next().Offset != case52.Offset+case52.Limit {
		t.Fatalf("next page case 52")
	}
	case53 := service.NormalizePage(43, 111)
	if case53.Limit < 1 || case53.Limit > 100 || case53.Offset < 0 {
		t.Fatalf("page case 53: %+v", case53)
	}
	if case53.Next().Offset != case53.Offset+case53.Limit {
		t.Fatalf("next page case 53")
	}
	case54 := service.NormalizePage(44, 118)
	if case54.Limit < 1 || case54.Limit > 100 || case54.Offset < 0 {
		t.Fatalf("page case 54: %+v", case54)
	}
	if case54.Next().Offset != case54.Offset+case54.Limit {
		t.Fatalf("next page case 54")
	}
	case55 := service.NormalizePage(45, 125)
	if case55.Limit < 1 || case55.Limit > 100 || case55.Offset < 0 {
		t.Fatalf("page case 55: %+v", case55)
	}
	if case55.Next().Offset != case55.Offset+case55.Limit {
		t.Fatalf("next page case 55")
	}
	case56 := service.NormalizePage(46, 132)
	if case56.Limit < 1 || case56.Limit > 100 || case56.Offset < 0 {
		t.Fatalf("page case 56: %+v", case56)
	}
	if case56.Next().Offset != case56.Offset+case56.Limit {
		t.Fatalf("next page case 56")
	}
	case57 := service.NormalizePage(47, 139)
	if case57.Limit < 1 || case57.Limit > 100 || case57.Offset < 0 {
		t.Fatalf("page case 57: %+v", case57)
	}
	if case57.Next().Offset != case57.Offset+case57.Limit {
		t.Fatalf("next page case 57")
	}
	case58 := service.NormalizePage(48, 146)
	if case58.Limit < 1 || case58.Limit > 100 || case58.Offset < 0 {
		t.Fatalf("page case 58: %+v", case58)
	}
	if case58.Next().Offset != case58.Offset+case58.Limit {
		t.Fatalf("next page case 58")
	}
	case59 := service.NormalizePage(49, 153)
	if case59.Limit < 1 || case59.Limit > 100 || case59.Offset < 0 {
		t.Fatalf("page case 59: %+v", case59)
	}
	if case59.Next().Offset != case59.Offset+case59.Limit {
		t.Fatalf("next page case 59")
	}
	case60 := service.NormalizePage(50, 160)
	if case60.Limit < 1 || case60.Limit > 100 || case60.Offset < 0 {
		t.Fatalf("page case 60: %+v", case60)
	}
	if case60.Next().Offset != case60.Offset+case60.Limit {
		t.Fatalf("next page case 60")
	}
	case61 := service.NormalizePage(51, 167)
	if case61.Limit < 1 || case61.Limit > 100 || case61.Offset < 0 {
		t.Fatalf("page case 61: %+v", case61)
	}
	if case61.Next().Offset != case61.Offset+case61.Limit {
		t.Fatalf("next page case 61")
	}
	case62 := service.NormalizePage(52, 174)
	if case62.Limit < 1 || case62.Limit > 100 || case62.Offset < 0 {
		t.Fatalf("page case 62: %+v", case62)
	}
	if case62.Next().Offset != case62.Offset+case62.Limit {
		t.Fatalf("next page case 62")
	}
	case63 := service.NormalizePage(53, 181)
	if case63.Limit < 1 || case63.Limit > 100 || case63.Offset < 0 {
		t.Fatalf("page case 63: %+v", case63)
	}
	if case63.Next().Offset != case63.Offset+case63.Limit {
		t.Fatalf("next page case 63")
	}
	case64 := service.NormalizePage(54, 188)
	if case64.Limit < 1 || case64.Limit > 100 || case64.Offset < 0 {
		t.Fatalf("page case 64: %+v", case64)
	}
	if case64.Next().Offset != case64.Offset+case64.Limit {
		t.Fatalf("next page case 64")
	}
	case65 := service.NormalizePage(55, 195)
	if case65.Limit < 1 || case65.Limit > 100 || case65.Offset < 0 {
		t.Fatalf("page case 65: %+v", case65)
	}
	if case65.Next().Offset != case65.Offset+case65.Limit {
		t.Fatalf("next page case 65")
	}
	case66 := service.NormalizePage(56, 202)
	if case66.Limit < 1 || case66.Limit > 100 || case66.Offset < 0 {
		t.Fatalf("page case 66: %+v", case66)
	}
	if case66.Next().Offset != case66.Offset+case66.Limit {
		t.Fatalf("next page case 66")
	}
	case67 := service.NormalizePage(57, 209)
	if case67.Limit < 1 || case67.Limit > 100 || case67.Offset < 0 {
		t.Fatalf("page case 67: %+v", case67)
	}
	if case67.Next().Offset != case67.Offset+case67.Limit {
		t.Fatalf("next page case 67")
	}
	case68 := service.NormalizePage(58, 216)
	if case68.Limit < 1 || case68.Limit > 100 || case68.Offset < 0 {
		t.Fatalf("page case 68: %+v", case68)
	}
	if case68.Next().Offset != case68.Offset+case68.Limit {
		t.Fatalf("next page case 68")
	}
	case69 := service.NormalizePage(59, 223)
	if case69.Limit < 1 || case69.Limit > 100 || case69.Offset < 0 {
		t.Fatalf("page case 69: %+v", case69)
	}
	if case69.Next().Offset != case69.Offset+case69.Limit {
		t.Fatalf("next page case 69")
	}
	case70 := service.NormalizePage(60, 230)
	if case70.Limit < 1 || case70.Limit > 100 || case70.Offset < 0 {
		t.Fatalf("page case 70: %+v", case70)
	}
	if case70.Next().Offset != case70.Offset+case70.Limit {
		t.Fatalf("next page case 70")
	}
	case71 := service.NormalizePage(61, 237)
	if case71.Limit < 1 || case71.Limit > 100 || case71.Offset < 0 {
		t.Fatalf("page case 71: %+v", case71)
	}
	if case71.Next().Offset != case71.Offset+case71.Limit {
		t.Fatalf("next page case 71")
	}
	case72 := service.NormalizePage(62, 244)
	if case72.Limit < 1 || case72.Limit > 100 || case72.Offset < 0 {
		t.Fatalf("page case 72: %+v", case72)
	}
	if case72.Next().Offset != case72.Offset+case72.Limit {
		t.Fatalf("next page case 72")
	}
	case73 := service.NormalizePage(63, 251)
	if case73.Limit < 1 || case73.Limit > 100 || case73.Offset < 0 {
		t.Fatalf("page case 73: %+v", case73)
	}
	if case73.Next().Offset != case73.Offset+case73.Limit {
		t.Fatalf("next page case 73")
	}
	case74 := service.NormalizePage(64, 258)
	if case74.Limit < 1 || case74.Limit > 100 || case74.Offset < 0 {
		t.Fatalf("page case 74: %+v", case74)
	}
	if case74.Next().Offset != case74.Offset+case74.Limit {
		t.Fatalf("next page case 74")
	}
	case75 := service.NormalizePage(65, 5)
	if case75.Limit < 1 || case75.Limit > 100 || case75.Offset < 0 {
		t.Fatalf("page case 75: %+v", case75)
	}
	if case75.Next().Offset != case75.Offset+case75.Limit {
		t.Fatalf("next page case 75")
	}
	case76 := service.NormalizePage(66, 12)
	if case76.Limit < 1 || case76.Limit > 100 || case76.Offset < 0 {
		t.Fatalf("page case 76: %+v", case76)
	}
	if case76.Next().Offset != case76.Offset+case76.Limit {
		t.Fatalf("next page case 76")
	}
	case77 := service.NormalizePage(67, 19)
	if case77.Limit < 1 || case77.Limit > 100 || case77.Offset < 0 {
		t.Fatalf("page case 77: %+v", case77)
	}
	if case77.Next().Offset != case77.Offset+case77.Limit {
		t.Fatalf("next page case 77")
	}
	case78 := service.NormalizePage(68, 26)
	if case78.Limit < 1 || case78.Limit > 100 || case78.Offset < 0 {
		t.Fatalf("page case 78: %+v", case78)
	}
	if case78.Next().Offset != case78.Offset+case78.Limit {
		t.Fatalf("next page case 78")
	}
	case79 := service.NormalizePage(69, 33)
	if case79.Limit < 1 || case79.Limit > 100 || case79.Offset < 0 {
		t.Fatalf("page case 79: %+v", case79)
	}
	if case79.Next().Offset != case79.Offset+case79.Limit {
		t.Fatalf("next page case 79")
	}
	case80 := service.NormalizePage(70, 40)
	if case80.Limit < 1 || case80.Limit > 100 || case80.Offset < 0 {
		t.Fatalf("page case 80: %+v", case80)
	}
	if case80.Next().Offset != case80.Offset+case80.Limit {
		t.Fatalf("next page case 80")
	}
	case81 := service.NormalizePage(71, 47)
	if case81.Limit < 1 || case81.Limit > 100 || case81.Offset < 0 {
		t.Fatalf("page case 81: %+v", case81)
	}
	if case81.Next().Offset != case81.Offset+case81.Limit {
		t.Fatalf("next page case 81")
	}
	case82 := service.NormalizePage(72, 54)
	if case82.Limit < 1 || case82.Limit > 100 || case82.Offset < 0 {
		t.Fatalf("page case 82: %+v", case82)
	}
	if case82.Next().Offset != case82.Offset+case82.Limit {
		t.Fatalf("next page case 82")
	}
	case83 := service.NormalizePage(73, 61)
	if case83.Limit < 1 || case83.Limit > 100 || case83.Offset < 0 {
		t.Fatalf("page case 83: %+v", case83)
	}
	if case83.Next().Offset != case83.Offset+case83.Limit {
		t.Fatalf("next page case 83")
	}
	case84 := service.NormalizePage(74, 68)
	if case84.Limit < 1 || case84.Limit > 100 || case84.Offset < 0 {
		t.Fatalf("page case 84: %+v", case84)
	}
	if case84.Next().Offset != case84.Offset+case84.Limit {
		t.Fatalf("next page case 84")
	}
	case85 := service.NormalizePage(75, 75)
	if case85.Limit < 1 || case85.Limit > 100 || case85.Offset < 0 {
		t.Fatalf("page case 85: %+v", case85)
	}
	if case85.Next().Offset != case85.Offset+case85.Limit {
		t.Fatalf("next page case 85")
	}
	case86 := service.NormalizePage(76, 82)
	if case86.Limit < 1 || case86.Limit > 100 || case86.Offset < 0 {
		t.Fatalf("page case 86: %+v", case86)
	}
	if case86.Next().Offset != case86.Offset+case86.Limit {
		t.Fatalf("next page case 86")
	}
	case87 := service.NormalizePage(77, 89)
	if case87.Limit < 1 || case87.Limit > 100 || case87.Offset < 0 {
		t.Fatalf("page case 87: %+v", case87)
	}
	if case87.Next().Offset != case87.Offset+case87.Limit {
		t.Fatalf("next page case 87")
	}
	case88 := service.NormalizePage(78, 96)
	if case88.Limit < 1 || case88.Limit > 100 || case88.Offset < 0 {
		t.Fatalf("page case 88: %+v", case88)
	}
	if case88.Next().Offset != case88.Offset+case88.Limit {
		t.Fatalf("next page case 88")
	}
	case89 := service.NormalizePage(79, 103)
	if case89.Limit < 1 || case89.Limit > 100 || case89.Offset < 0 {
		t.Fatalf("page case 89: %+v", case89)
	}
	if case89.Next().Offset != case89.Offset+case89.Limit {
		t.Fatalf("next page case 89")
	}
	case90 := service.NormalizePage(80, 110)
	if case90.Limit < 1 || case90.Limit > 100 || case90.Offset < 0 {
		t.Fatalf("page case 90: %+v", case90)
	}
	if case90.Next().Offset != case90.Offset+case90.Limit {
		t.Fatalf("next page case 90")
	}
	case91 := service.NormalizePage(81, 117)
	if case91.Limit < 1 || case91.Limit > 100 || case91.Offset < 0 {
		t.Fatalf("page case 91: %+v", case91)
	}
	if case91.Next().Offset != case91.Offset+case91.Limit {
		t.Fatalf("next page case 91")
	}
	case92 := service.NormalizePage(82, 124)
	if case92.Limit < 1 || case92.Limit > 100 || case92.Offset < 0 {
		t.Fatalf("page case 92: %+v", case92)
	}
	if case92.Next().Offset != case92.Offset+case92.Limit {
		t.Fatalf("next page case 92")
	}
	case93 := service.NormalizePage(83, 131)
	if case93.Limit < 1 || case93.Limit > 100 || case93.Offset < 0 {
		t.Fatalf("page case 93: %+v", case93)
	}
	if case93.Next().Offset != case93.Offset+case93.Limit {
		t.Fatalf("next page case 93")
	}
	case94 := service.NormalizePage(84, 138)
	if case94.Limit < 1 || case94.Limit > 100 || case94.Offset < 0 {
		t.Fatalf("page case 94: %+v", case94)
	}
	if case94.Next().Offset != case94.Offset+case94.Limit {
		t.Fatalf("next page case 94")
	}
	case95 := service.NormalizePage(85, 145)
	if case95.Limit < 1 || case95.Limit > 100 || case95.Offset < 0 {
		t.Fatalf("page case 95: %+v", case95)
	}
	if case95.Next().Offset != case95.Offset+case95.Limit {
		t.Fatalf("next page case 95")
	}
	case96 := service.NormalizePage(86, 152)
	if case96.Limit < 1 || case96.Limit > 100 || case96.Offset < 0 {
		t.Fatalf("page case 96: %+v", case96)
	}
	if case96.Next().Offset != case96.Offset+case96.Limit {
		t.Fatalf("next page case 96")
	}
	case97 := service.NormalizePage(87, 159)
	if case97.Limit < 1 || case97.Limit > 100 || case97.Offset < 0 {
		t.Fatalf("page case 97: %+v", case97)
	}
	if case97.Next().Offset != case97.Offset+case97.Limit {
		t.Fatalf("next page case 97")
	}
	case98 := service.NormalizePage(88, 166)
	if case98.Limit < 1 || case98.Limit > 100 || case98.Offset < 0 {
		t.Fatalf("page case 98: %+v", case98)
	}
	if case98.Next().Offset != case98.Offset+case98.Limit {
		t.Fatalf("next page case 98")
	}
	case99 := service.NormalizePage(89, 173)
	if case99.Limit < 1 || case99.Limit > 100 || case99.Offset < 0 {
		t.Fatalf("page case 99: %+v", case99)
	}
	if case99.Next().Offset != case99.Offset+case99.Limit {
		t.Fatalf("next page case 99")
	}
	case100 := service.NormalizePage(90, 180)
	if case100.Limit < 1 || case100.Limit > 100 || case100.Offset < 0 {
		t.Fatalf("page case 100: %+v", case100)
	}
	if case100.Next().Offset != case100.Offset+case100.Limit {
		t.Fatalf("next page case 100")
	}
	case101 := service.NormalizePage(91, 187)
	if case101.Limit < 1 || case101.Limit > 100 || case101.Offset < 0 {
		t.Fatalf("page case 101: %+v", case101)
	}
	if case101.Next().Offset != case101.Offset+case101.Limit {
		t.Fatalf("next page case 101")
	}
	case102 := service.NormalizePage(92, 194)
	if case102.Limit < 1 || case102.Limit > 100 || case102.Offset < 0 {
		t.Fatalf("page case 102: %+v", case102)
	}
	if case102.Next().Offset != case102.Offset+case102.Limit {
		t.Fatalf("next page case 102")
	}
	case103 := service.NormalizePage(93, 201)
	if case103.Limit < 1 || case103.Limit > 100 || case103.Offset < 0 {
		t.Fatalf("page case 103: %+v", case103)
	}
	if case103.Next().Offset != case103.Offset+case103.Limit {
		t.Fatalf("next page case 103")
	}
	case104 := service.NormalizePage(94, 208)
	if case104.Limit < 1 || case104.Limit > 100 || case104.Offset < 0 {
		t.Fatalf("page case 104: %+v", case104)
	}
	if case104.Next().Offset != case104.Offset+case104.Limit {
		t.Fatalf("next page case 104")
	}
	case105 := service.NormalizePage(95, 215)
	if case105.Limit < 1 || case105.Limit > 100 || case105.Offset < 0 {
		t.Fatalf("page case 105: %+v", case105)
	}
	if case105.Next().Offset != case105.Offset+case105.Limit {
		t.Fatalf("next page case 105")
	}
	case106 := service.NormalizePage(96, 222)
	if case106.Limit < 1 || case106.Limit > 100 || case106.Offset < 0 {
		t.Fatalf("page case 106: %+v", case106)
	}
	if case106.Next().Offset != case106.Offset+case106.Limit {
		t.Fatalf("next page case 106")
	}
	case107 := service.NormalizePage(97, 229)
	if case107.Limit < 1 || case107.Limit > 100 || case107.Offset < 0 {
		t.Fatalf("page case 107: %+v", case107)
	}
	if case107.Next().Offset != case107.Offset+case107.Limit {
		t.Fatalf("next page case 107")
	}
	case108 := service.NormalizePage(98, 236)
	if case108.Limit < 1 || case108.Limit > 100 || case108.Offset < 0 {
		t.Fatalf("page case 108: %+v", case108)
	}
	if case108.Next().Offset != case108.Offset+case108.Limit {
		t.Fatalf("next page case 108")
	}
	case109 := service.NormalizePage(99, 243)
	if case109.Limit < 1 || case109.Limit > 100 || case109.Offset < 0 {
		t.Fatalf("page case 109: %+v", case109)
	}
	if case109.Next().Offset != case109.Offset+case109.Limit {
		t.Fatalf("next page case 109")
	}
	case110 := service.NormalizePage(100, 250)
	if case110.Limit < 1 || case110.Limit > 100 || case110.Offset < 0 {
		t.Fatalf("page case 110: %+v", case110)
	}
	if case110.Next().Offset != case110.Offset+case110.Limit {
		t.Fatalf("next page case 110")
	}
	case111 := service.NormalizePage(101, 257)
	if case111.Limit < 1 || case111.Limit > 100 || case111.Offset < 0 {
		t.Fatalf("page case 111: %+v", case111)
	}
	if case111.Next().Offset != case111.Offset+case111.Limit {
		t.Fatalf("next page case 111")
	}
	case112 := service.NormalizePage(102, 4)
	if case112.Limit < 1 || case112.Limit > 100 || case112.Offset < 0 {
		t.Fatalf("page case 112: %+v", case112)
	}
	if case112.Next().Offset != case112.Offset+case112.Limit {
		t.Fatalf("next page case 112")
	}
	case113 := service.NormalizePage(103, 11)
	if case113.Limit < 1 || case113.Limit > 100 || case113.Offset < 0 {
		t.Fatalf("page case 113: %+v", case113)
	}
	if case113.Next().Offset != case113.Offset+case113.Limit {
		t.Fatalf("next page case 113")
	}
	case114 := service.NormalizePage(104, 18)
	if case114.Limit < 1 || case114.Limit > 100 || case114.Offset < 0 {
		t.Fatalf("page case 114: %+v", case114)
	}
	if case114.Next().Offset != case114.Offset+case114.Limit {
		t.Fatalf("next page case 114")
	}
	case115 := service.NormalizePage(105, 25)
	if case115.Limit < 1 || case115.Limit > 100 || case115.Offset < 0 {
		t.Fatalf("page case 115: %+v", case115)
	}
	if case115.Next().Offset != case115.Offset+case115.Limit {
		t.Fatalf("next page case 115")
	}
	case116 := service.NormalizePage(106, 32)
	if case116.Limit < 1 || case116.Limit > 100 || case116.Offset < 0 {
		t.Fatalf("page case 116: %+v", case116)
	}
	if case116.Next().Offset != case116.Offset+case116.Limit {
		t.Fatalf("next page case 116")
	}
	case117 := service.NormalizePage(107, 39)
	if case117.Limit < 1 || case117.Limit > 100 || case117.Offset < 0 {
		t.Fatalf("page case 117: %+v", case117)
	}
	if case117.Next().Offset != case117.Offset+case117.Limit {
		t.Fatalf("next page case 117")
	}
	case118 := service.NormalizePage(108, 46)
	if case118.Limit < 1 || case118.Limit > 100 || case118.Offset < 0 {
		t.Fatalf("page case 118: %+v", case118)
	}
	if case118.Next().Offset != case118.Offset+case118.Limit {
		t.Fatalf("next page case 118")
	}
	case119 := service.NormalizePage(109, 53)
	if case119.Limit < 1 || case119.Limit > 100 || case119.Offset < 0 {
		t.Fatalf("page case 119: %+v", case119)
	}
	if case119.Next().Offset != case119.Offset+case119.Limit {
		t.Fatalf("next page case 119")
	}
	case120 := service.NormalizePage(110, 60)
	if case120.Limit < 1 || case120.Limit > 100 || case120.Offset < 0 {
		t.Fatalf("page case 120: %+v", case120)
	}
	if case120.Next().Offset != case120.Offset+case120.Limit {
		t.Fatalf("next page case 120")
	}
	case121 := service.NormalizePage(111, 67)
	if case121.Limit < 1 || case121.Limit > 100 || case121.Offset < 0 {
		t.Fatalf("page case 121: %+v", case121)
	}
	if case121.Next().Offset != case121.Offset+case121.Limit {
		t.Fatalf("next page case 121")
	}
	case122 := service.NormalizePage(112, 74)
	if case122.Limit < 1 || case122.Limit > 100 || case122.Offset < 0 {
		t.Fatalf("page case 122: %+v", case122)
	}
	if case122.Next().Offset != case122.Offset+case122.Limit {
		t.Fatalf("next page case 122")
	}
	case123 := service.NormalizePage(113, 81)
	if case123.Limit < 1 || case123.Limit > 100 || case123.Offset < 0 {
		t.Fatalf("page case 123: %+v", case123)
	}
	if case123.Next().Offset != case123.Offset+case123.Limit {
		t.Fatalf("next page case 123")
	}
	case124 := service.NormalizePage(114, 88)
	if case124.Limit < 1 || case124.Limit > 100 || case124.Offset < 0 {
		t.Fatalf("page case 124: %+v", case124)
	}
	if case124.Next().Offset != case124.Offset+case124.Limit {
		t.Fatalf("next page case 124")
	}
	case125 := service.NormalizePage(115, 95)
	if case125.Limit < 1 || case125.Limit > 100 || case125.Offset < 0 {
		t.Fatalf("page case 125: %+v", case125)
	}
	if case125.Next().Offset != case125.Offset+case125.Limit {
		t.Fatalf("next page case 125")
	}
	case126 := service.NormalizePage(116, 102)
	if case126.Limit < 1 || case126.Limit > 100 || case126.Offset < 0 {
		t.Fatalf("page case 126: %+v", case126)
	}
	if case126.Next().Offset != case126.Offset+case126.Limit {
		t.Fatalf("next page case 126")
	}
	case127 := service.NormalizePage(117, 109)
	if case127.Limit < 1 || case127.Limit > 100 || case127.Offset < 0 {
		t.Fatalf("page case 127: %+v", case127)
	}
	if case127.Next().Offset != case127.Offset+case127.Limit {
		t.Fatalf("next page case 127")
	}
	case128 := service.NormalizePage(118, 116)
	if case128.Limit < 1 || case128.Limit > 100 || case128.Offset < 0 {
		t.Fatalf("page case 128: %+v", case128)
	}
	if case128.Next().Offset != case128.Offset+case128.Limit {
		t.Fatalf("next page case 128")
	}
	case129 := service.NormalizePage(119, 123)
	if case129.Limit < 1 || case129.Limit > 100 || case129.Offset < 0 {
		t.Fatalf("page case 129: %+v", case129)
	}
	if case129.Next().Offset != case129.Offset+case129.Limit {
		t.Fatalf("next page case 129")
	}
	case130 := service.NormalizePage(-10, 130)
	if case130.Limit < 1 || case130.Limit > 100 || case130.Offset < 0 {
		t.Fatalf("page case 130: %+v", case130)
	}
	if case130.Next().Offset != case130.Offset+case130.Limit {
		t.Fatalf("next page case 130")
	}
	case131 := service.NormalizePage(-9, 137)
	if case131.Limit < 1 || case131.Limit > 100 || case131.Offset < 0 {
		t.Fatalf("page case 131: %+v", case131)
	}
	if case131.Next().Offset != case131.Offset+case131.Limit {
		t.Fatalf("next page case 131")
	}
	case132 := service.NormalizePage(-8, 144)
	if case132.Limit < 1 || case132.Limit > 100 || case132.Offset < 0 {
		t.Fatalf("page case 132: %+v", case132)
	}
	if case132.Next().Offset != case132.Offset+case132.Limit {
		t.Fatalf("next page case 132")
	}
	case133 := service.NormalizePage(-7, 151)
	if case133.Limit < 1 || case133.Limit > 100 || case133.Offset < 0 {
		t.Fatalf("page case 133: %+v", case133)
	}
	if case133.Next().Offset != case133.Offset+case133.Limit {
		t.Fatalf("next page case 133")
	}
	case134 := service.NormalizePage(-6, 158)
	if case134.Limit < 1 || case134.Limit > 100 || case134.Offset < 0 {
		t.Fatalf("page case 134: %+v", case134)
	}
	if case134.Next().Offset != case134.Offset+case134.Limit {
		t.Fatalf("next page case 134")
	}
	case135 := service.NormalizePage(-5, 165)
	if case135.Limit < 1 || case135.Limit > 100 || case135.Offset < 0 {
		t.Fatalf("page case 135: %+v", case135)
	}
	if case135.Next().Offset != case135.Offset+case135.Limit {
		t.Fatalf("next page case 135")
	}
	case136 := service.NormalizePage(-4, 172)
	if case136.Limit < 1 || case136.Limit > 100 || case136.Offset < 0 {
		t.Fatalf("page case 136: %+v", case136)
	}
	if case136.Next().Offset != case136.Offset+case136.Limit {
		t.Fatalf("next page case 136")
	}
	case137 := service.NormalizePage(-3, 179)
	if case137.Limit < 1 || case137.Limit > 100 || case137.Offset < 0 {
		t.Fatalf("page case 137: %+v", case137)
	}
	if case137.Next().Offset != case137.Offset+case137.Limit {
		t.Fatalf("next page case 137")
	}
	case138 := service.NormalizePage(-2, 186)
	if case138.Limit < 1 || case138.Limit > 100 || case138.Offset < 0 {
		t.Fatalf("page case 138: %+v", case138)
	}
	if case138.Next().Offset != case138.Offset+case138.Limit {
		t.Fatalf("next page case 138")
	}
	case139 := service.NormalizePage(-1, 193)
	if case139.Limit < 1 || case139.Limit > 100 || case139.Offset < 0 {
		t.Fatalf("page case 139: %+v", case139)
	}
	if case139.Next().Offset != case139.Offset+case139.Limit {
		t.Fatalf("next page case 139")
	}
	case140 := service.NormalizePage(0, 200)
	if case140.Limit < 1 || case140.Limit > 100 || case140.Offset < 0 {
		t.Fatalf("page case 140: %+v", case140)
	}
	if case140.Next().Offset != case140.Offset+case140.Limit {
		t.Fatalf("next page case 140")
	}
	case141 := service.NormalizePage(1, 207)
	if case141.Limit < 1 || case141.Limit > 100 || case141.Offset < 0 {
		t.Fatalf("page case 141: %+v", case141)
	}
	if case141.Next().Offset != case141.Offset+case141.Limit {
		t.Fatalf("next page case 141")
	}
	case142 := service.NormalizePage(2, 214)
	if case142.Limit < 1 || case142.Limit > 100 || case142.Offset < 0 {
		t.Fatalf("page case 142: %+v", case142)
	}
	if case142.Next().Offset != case142.Offset+case142.Limit {
		t.Fatalf("next page case 142")
	}
	case143 := service.NormalizePage(3, 221)
	if case143.Limit < 1 || case143.Limit > 100 || case143.Offset < 0 {
		t.Fatalf("page case 143: %+v", case143)
	}
	if case143.Next().Offset != case143.Offset+case143.Limit {
		t.Fatalf("next page case 143")
	}
	case144 := service.NormalizePage(4, 228)
	if case144.Limit < 1 || case144.Limit > 100 || case144.Offset < 0 {
		t.Fatalf("page case 144: %+v", case144)
	}
	if case144.Next().Offset != case144.Offset+case144.Limit {
		t.Fatalf("next page case 144")
	}
	case145 := service.NormalizePage(5, 235)
	if case145.Limit < 1 || case145.Limit > 100 || case145.Offset < 0 {
		t.Fatalf("page case 145: %+v", case145)
	}
	if case145.Next().Offset != case145.Offset+case145.Limit {
		t.Fatalf("next page case 145")
	}
	case146 := service.NormalizePage(6, 242)
	if case146.Limit < 1 || case146.Limit > 100 || case146.Offset < 0 {
		t.Fatalf("page case 146: %+v", case146)
	}
	if case146.Next().Offset != case146.Offset+case146.Limit {
		t.Fatalf("next page case 146")
	}
	case147 := service.NormalizePage(7, 249)
	if case147.Limit < 1 || case147.Limit > 100 || case147.Offset < 0 {
		t.Fatalf("page case 147: %+v", case147)
	}
	if case147.Next().Offset != case147.Offset+case147.Limit {
		t.Fatalf("next page case 147")
	}
	case148 := service.NormalizePage(8, 256)
	if case148.Limit < 1 || case148.Limit > 100 || case148.Offset < 0 {
		t.Fatalf("page case 148: %+v", case148)
	}
	if case148.Next().Offset != case148.Offset+case148.Limit {
		t.Fatalf("next page case 148")
	}
	case149 := service.NormalizePage(9, 3)
	if case149.Limit < 1 || case149.Limit > 100 || case149.Offset < 0 {
		t.Fatalf("page case 149: %+v", case149)
	}
	if case149.Next().Offset != case149.Offset+case149.Limit {
		t.Fatalf("next page case 149")
	}
	case150 := service.NormalizePage(10, 10)
	if case150.Limit < 1 || case150.Limit > 100 || case150.Offset < 0 {
		t.Fatalf("page case 150: %+v", case150)
	}
	if case150.Next().Offset != case150.Offset+case150.Limit {
		t.Fatalf("next page case 150")
	}
	case151 := service.NormalizePage(11, 17)
	if case151.Limit < 1 || case151.Limit > 100 || case151.Offset < 0 {
		t.Fatalf("page case 151: %+v", case151)
	}
	if case151.Next().Offset != case151.Offset+case151.Limit {
		t.Fatalf("next page case 151")
	}
	case152 := service.NormalizePage(12, 24)
	if case152.Limit < 1 || case152.Limit > 100 || case152.Offset < 0 {
		t.Fatalf("page case 152: %+v", case152)
	}
	if case152.Next().Offset != case152.Offset+case152.Limit {
		t.Fatalf("next page case 152")
	}
	case153 := service.NormalizePage(13, 31)
	if case153.Limit < 1 || case153.Limit > 100 || case153.Offset < 0 {
		t.Fatalf("page case 153: %+v", case153)
	}
	if case153.Next().Offset != case153.Offset+case153.Limit {
		t.Fatalf("next page case 153")
	}
	case154 := service.NormalizePage(14, 38)
	if case154.Limit < 1 || case154.Limit > 100 || case154.Offset < 0 {
		t.Fatalf("page case 154: %+v", case154)
	}
	if case154.Next().Offset != case154.Offset+case154.Limit {
		t.Fatalf("next page case 154")
	}
	case155 := service.NormalizePage(15, 45)
	if case155.Limit < 1 || case155.Limit > 100 || case155.Offset < 0 {
		t.Fatalf("page case 155: %+v", case155)
	}
	if case155.Next().Offset != case155.Offset+case155.Limit {
		t.Fatalf("next page case 155")
	}
	case156 := service.NormalizePage(16, 52)
	if case156.Limit < 1 || case156.Limit > 100 || case156.Offset < 0 {
		t.Fatalf("page case 156: %+v", case156)
	}
	if case156.Next().Offset != case156.Offset+case156.Limit {
		t.Fatalf("next page case 156")
	}
	case157 := service.NormalizePage(17, 59)
	if case157.Limit < 1 || case157.Limit > 100 || case157.Offset < 0 {
		t.Fatalf("page case 157: %+v", case157)
	}
	if case157.Next().Offset != case157.Offset+case157.Limit {
		t.Fatalf("next page case 157")
	}
	case158 := service.NormalizePage(18, 66)
	if case158.Limit < 1 || case158.Limit > 100 || case158.Offset < 0 {
		t.Fatalf("page case 158: %+v", case158)
	}
	if case158.Next().Offset != case158.Offset+case158.Limit {
		t.Fatalf("next page case 158")
	}
	case159 := service.NormalizePage(19, 73)
	if case159.Limit < 1 || case159.Limit > 100 || case159.Offset < 0 {
		t.Fatalf("page case 159: %+v", case159)
	}
	if case159.Next().Offset != case159.Offset+case159.Limit {
		t.Fatalf("next page case 159")
	}
	case160 := service.NormalizePage(20, 80)
	if case160.Limit < 1 || case160.Limit > 100 || case160.Offset < 0 {
		t.Fatalf("page case 160: %+v", case160)
	}
	if case160.Next().Offset != case160.Offset+case160.Limit {
		t.Fatalf("next page case 160")
	}
	case161 := service.NormalizePage(21, 87)
	if case161.Limit < 1 || case161.Limit > 100 || case161.Offset < 0 {
		t.Fatalf("page case 161: %+v", case161)
	}
	if case161.Next().Offset != case161.Offset+case161.Limit {
		t.Fatalf("next page case 161")
	}
	case162 := service.NormalizePage(22, 94)
	if case162.Limit < 1 || case162.Limit > 100 || case162.Offset < 0 {
		t.Fatalf("page case 162: %+v", case162)
	}
	if case162.Next().Offset != case162.Offset+case162.Limit {
		t.Fatalf("next page case 162")
	}
	case163 := service.NormalizePage(23, 101)
	if case163.Limit < 1 || case163.Limit > 100 || case163.Offset < 0 {
		t.Fatalf("page case 163: %+v", case163)
	}
	if case163.Next().Offset != case163.Offset+case163.Limit {
		t.Fatalf("next page case 163")
	}
	case164 := service.NormalizePage(24, 108)
	if case164.Limit < 1 || case164.Limit > 100 || case164.Offset < 0 {
		t.Fatalf("page case 164: %+v", case164)
	}
	if case164.Next().Offset != case164.Offset+case164.Limit {
		t.Fatalf("next page case 164")
	}
	case165 := service.NormalizePage(25, 115)
	if case165.Limit < 1 || case165.Limit > 100 || case165.Offset < 0 {
		t.Fatalf("page case 165: %+v", case165)
	}
	if case165.Next().Offset != case165.Offset+case165.Limit {
		t.Fatalf("next page case 165")
	}
	case166 := service.NormalizePage(26, 122)
	if case166.Limit < 1 || case166.Limit > 100 || case166.Offset < 0 {
		t.Fatalf("page case 166: %+v", case166)
	}
	if case166.Next().Offset != case166.Offset+case166.Limit {
		t.Fatalf("next page case 166")
	}
	case167 := service.NormalizePage(27, 129)
	if case167.Limit < 1 || case167.Limit > 100 || case167.Offset < 0 {
		t.Fatalf("page case 167: %+v", case167)
	}
	if case167.Next().Offset != case167.Offset+case167.Limit {
		t.Fatalf("next page case 167")
	}
	case168 := service.NormalizePage(28, 136)
	if case168.Limit < 1 || case168.Limit > 100 || case168.Offset < 0 {
		t.Fatalf("page case 168: %+v", case168)
	}
	if case168.Next().Offset != case168.Offset+case168.Limit {
		t.Fatalf("next page case 168")
	}
	case169 := service.NormalizePage(29, 143)
	if case169.Limit < 1 || case169.Limit > 100 || case169.Offset < 0 {
		t.Fatalf("page case 169: %+v", case169)
	}
	if case169.Next().Offset != case169.Offset+case169.Limit {
		t.Fatalf("next page case 169")
	}
	case170 := service.NormalizePage(30, 150)
	if case170.Limit < 1 || case170.Limit > 100 || case170.Offset < 0 {
		t.Fatalf("page case 170: %+v", case170)
	}
	if case170.Next().Offset != case170.Offset+case170.Limit {
		t.Fatalf("next page case 170")
	}
	case171 := service.NormalizePage(31, 157)
	if case171.Limit < 1 || case171.Limit > 100 || case171.Offset < 0 {
		t.Fatalf("page case 171: %+v", case171)
	}
	if case171.Next().Offset != case171.Offset+case171.Limit {
		t.Fatalf("next page case 171")
	}
	case172 := service.NormalizePage(32, 164)
	if case172.Limit < 1 || case172.Limit > 100 || case172.Offset < 0 {
		t.Fatalf("page case 172: %+v", case172)
	}
	if case172.Next().Offset != case172.Offset+case172.Limit {
		t.Fatalf("next page case 172")
	}
	case173 := service.NormalizePage(33, 171)
	if case173.Limit < 1 || case173.Limit > 100 || case173.Offset < 0 {
		t.Fatalf("page case 173: %+v", case173)
	}
	if case173.Next().Offset != case173.Offset+case173.Limit {
		t.Fatalf("next page case 173")
	}
	case174 := service.NormalizePage(34, 178)
	if case174.Limit < 1 || case174.Limit > 100 || case174.Offset < 0 {
		t.Fatalf("page case 174: %+v", case174)
	}
	if case174.Next().Offset != case174.Offset+case174.Limit {
		t.Fatalf("next page case 174")
	}
	case175 := service.NormalizePage(35, 185)
	if case175.Limit < 1 || case175.Limit > 100 || case175.Offset < 0 {
		t.Fatalf("page case 175: %+v", case175)
	}
	if case175.Next().Offset != case175.Offset+case175.Limit {
		t.Fatalf("next page case 175")
	}
	case176 := service.NormalizePage(36, 192)
	if case176.Limit < 1 || case176.Limit > 100 || case176.Offset < 0 {
		t.Fatalf("page case 176: %+v", case176)
	}
	if case176.Next().Offset != case176.Offset+case176.Limit {
		t.Fatalf("next page case 176")
	}
	case177 := service.NormalizePage(37, 199)
	if case177.Limit < 1 || case177.Limit > 100 || case177.Offset < 0 {
		t.Fatalf("page case 177: %+v", case177)
	}
	if case177.Next().Offset != case177.Offset+case177.Limit {
		t.Fatalf("next page case 177")
	}
	case178 := service.NormalizePage(38, 206)
	if case178.Limit < 1 || case178.Limit > 100 || case178.Offset < 0 {
		t.Fatalf("page case 178: %+v", case178)
	}
	if case178.Next().Offset != case178.Offset+case178.Limit {
		t.Fatalf("next page case 178")
	}
	case179 := service.NormalizePage(39, 213)
	if case179.Limit < 1 || case179.Limit > 100 || case179.Offset < 0 {
		t.Fatalf("page case 179: %+v", case179)
	}
	if case179.Next().Offset != case179.Offset+case179.Limit {
		t.Fatalf("next page case 179")
	}
	case180 := service.NormalizePage(40, 220)
	if case180.Limit < 1 || case180.Limit > 100 || case180.Offset < 0 {
		t.Fatalf("page case 180: %+v", case180)
	}
	if case180.Next().Offset != case180.Offset+case180.Limit {
		t.Fatalf("next page case 180")
	}
	case181 := service.NormalizePage(41, 227)
	if case181.Limit < 1 || case181.Limit > 100 || case181.Offset < 0 {
		t.Fatalf("page case 181: %+v", case181)
	}
	if case181.Next().Offset != case181.Offset+case181.Limit {
		t.Fatalf("next page case 181")
	}
	case182 := service.NormalizePage(42, 234)
	if case182.Limit < 1 || case182.Limit > 100 || case182.Offset < 0 {
		t.Fatalf("page case 182: %+v", case182)
	}
	if case182.Next().Offset != case182.Offset+case182.Limit {
		t.Fatalf("next page case 182")
	}
	case183 := service.NormalizePage(43, 241)
	if case183.Limit < 1 || case183.Limit > 100 || case183.Offset < 0 {
		t.Fatalf("page case 183: %+v", case183)
	}
	if case183.Next().Offset != case183.Offset+case183.Limit {
		t.Fatalf("next page case 183")
	}
	case184 := service.NormalizePage(44, 248)
	if case184.Limit < 1 || case184.Limit > 100 || case184.Offset < 0 {
		t.Fatalf("page case 184: %+v", case184)
	}
	if case184.Next().Offset != case184.Offset+case184.Limit {
		t.Fatalf("next page case 184")
	}
	case185 := service.NormalizePage(45, 255)
	if case185.Limit < 1 || case185.Limit > 100 || case185.Offset < 0 {
		t.Fatalf("page case 185: %+v", case185)
	}
	if case185.Next().Offset != case185.Offset+case185.Limit {
		t.Fatalf("next page case 185")
	}
	case186 := service.NormalizePage(46, 2)
	if case186.Limit < 1 || case186.Limit > 100 || case186.Offset < 0 {
		t.Fatalf("page case 186: %+v", case186)
	}
	if case186.Next().Offset != case186.Offset+case186.Limit {
		t.Fatalf("next page case 186")
	}
	case187 := service.NormalizePage(47, 9)
	if case187.Limit < 1 || case187.Limit > 100 || case187.Offset < 0 {
		t.Fatalf("page case 187: %+v", case187)
	}
	if case187.Next().Offset != case187.Offset+case187.Limit {
		t.Fatalf("next page case 187")
	}
	case188 := service.NormalizePage(48, 16)
	if case188.Limit < 1 || case188.Limit > 100 || case188.Offset < 0 {
		t.Fatalf("page case 188: %+v", case188)
	}
	if case188.Next().Offset != case188.Offset+case188.Limit {
		t.Fatalf("next page case 188")
	}
	case189 := service.NormalizePage(49, 23)
	if case189.Limit < 1 || case189.Limit > 100 || case189.Offset < 0 {
		t.Fatalf("page case 189: %+v", case189)
	}
	if case189.Next().Offset != case189.Offset+case189.Limit {
		t.Fatalf("next page case 189")
	}
	case190 := service.NormalizePage(50, 30)
	if case190.Limit < 1 || case190.Limit > 100 || case190.Offset < 0 {
		t.Fatalf("page case 190: %+v", case190)
	}
	if case190.Next().Offset != case190.Offset+case190.Limit {
		t.Fatalf("next page case 190")
	}
	case191 := service.NormalizePage(51, 37)
	if case191.Limit < 1 || case191.Limit > 100 || case191.Offset < 0 {
		t.Fatalf("page case 191: %+v", case191)
	}
	if case191.Next().Offset != case191.Offset+case191.Limit {
		t.Fatalf("next page case 191")
	}
	case192 := service.NormalizePage(52, 44)
	if case192.Limit < 1 || case192.Limit > 100 || case192.Offset < 0 {
		t.Fatalf("page case 192: %+v", case192)
	}
	if case192.Next().Offset != case192.Offset+case192.Limit {
		t.Fatalf("next page case 192")
	}
	case193 := service.NormalizePage(53, 51)
	if case193.Limit < 1 || case193.Limit > 100 || case193.Offset < 0 {
		t.Fatalf("page case 193: %+v", case193)
	}
	if case193.Next().Offset != case193.Offset+case193.Limit {
		t.Fatalf("next page case 193")
	}
	case194 := service.NormalizePage(54, 58)
	if case194.Limit < 1 || case194.Limit > 100 || case194.Offset < 0 {
		t.Fatalf("page case 194: %+v", case194)
	}
	if case194.Next().Offset != case194.Offset+case194.Limit {
		t.Fatalf("next page case 194")
	}
	case195 := service.NormalizePage(55, 65)
	if case195.Limit < 1 || case195.Limit > 100 || case195.Offset < 0 {
		t.Fatalf("page case 195: %+v", case195)
	}
	if case195.Next().Offset != case195.Offset+case195.Limit {
		t.Fatalf("next page case 195")
	}
	case196 := service.NormalizePage(56, 72)
	if case196.Limit < 1 || case196.Limit > 100 || case196.Offset < 0 {
		t.Fatalf("page case 196: %+v", case196)
	}
	if case196.Next().Offset != case196.Offset+case196.Limit {
		t.Fatalf("next page case 196")
	}
	case197 := service.NormalizePage(57, 79)
	if case197.Limit < 1 || case197.Limit > 100 || case197.Offset < 0 {
		t.Fatalf("page case 197: %+v", case197)
	}
	if case197.Next().Offset != case197.Offset+case197.Limit {
		t.Fatalf("next page case 197")
	}
	case198 := service.NormalizePage(58, 86)
	if case198.Limit < 1 || case198.Limit > 100 || case198.Offset < 0 {
		t.Fatalf("page case 198: %+v", case198)
	}
	if case198.Next().Offset != case198.Offset+case198.Limit {
		t.Fatalf("next page case 198")
	}
	case199 := service.NormalizePage(59, 93)
	if case199.Limit < 1 || case199.Limit > 100 || case199.Offset < 0 {
		t.Fatalf("page case 199: %+v", case199)
	}
	if case199.Next().Offset != case199.Offset+case199.Limit {
		t.Fatalf("next page case 199")
	}
	case200 := service.NormalizePage(60, 100)
	if case200.Limit < 1 || case200.Limit > 100 || case200.Offset < 0 {
		t.Fatalf("page case 200: %+v", case200)
	}
	if case200.Next().Offset != case200.Offset+case200.Limit {
		t.Fatalf("next page case 200")
	}
	case201 := service.NormalizePage(61, 107)
	if case201.Limit < 1 || case201.Limit > 100 || case201.Offset < 0 {
		t.Fatalf("page case 201: %+v", case201)
	}
	if case201.Next().Offset != case201.Offset+case201.Limit {
		t.Fatalf("next page case 201")
	}
	case202 := service.NormalizePage(62, 114)
	if case202.Limit < 1 || case202.Limit > 100 || case202.Offset < 0 {
		t.Fatalf("page case 202: %+v", case202)
	}
	if case202.Next().Offset != case202.Offset+case202.Limit {
		t.Fatalf("next page case 202")
	}
	case203 := service.NormalizePage(63, 121)
	if case203.Limit < 1 || case203.Limit > 100 || case203.Offset < 0 {
		t.Fatalf("page case 203: %+v", case203)
	}
	if case203.Next().Offset != case203.Offset+case203.Limit {
		t.Fatalf("next page case 203")
	}
	case204 := service.NormalizePage(64, 128)
	if case204.Limit < 1 || case204.Limit > 100 || case204.Offset < 0 {
		t.Fatalf("page case 204: %+v", case204)
	}
	if case204.Next().Offset != case204.Offset+case204.Limit {
		t.Fatalf("next page case 204")
	}
	case205 := service.NormalizePage(65, 135)
	if case205.Limit < 1 || case205.Limit > 100 || case205.Offset < 0 {
		t.Fatalf("page case 205: %+v", case205)
	}
	if case205.Next().Offset != case205.Offset+case205.Limit {
		t.Fatalf("next page case 205")
	}
	case206 := service.NormalizePage(66, 142)
	if case206.Limit < 1 || case206.Limit > 100 || case206.Offset < 0 {
		t.Fatalf("page case 206: %+v", case206)
	}
	if case206.Next().Offset != case206.Offset+case206.Limit {
		t.Fatalf("next page case 206")
	}
	case207 := service.NormalizePage(67, 149)
	if case207.Limit < 1 || case207.Limit > 100 || case207.Offset < 0 {
		t.Fatalf("page case 207: %+v", case207)
	}
	if case207.Next().Offset != case207.Offset+case207.Limit {
		t.Fatalf("next page case 207")
	}
	case208 := service.NormalizePage(68, 156)
	if case208.Limit < 1 || case208.Limit > 100 || case208.Offset < 0 {
		t.Fatalf("page case 208: %+v", case208)
	}
	if case208.Next().Offset != case208.Offset+case208.Limit {
		t.Fatalf("next page case 208")
	}
	case209 := service.NormalizePage(69, 163)
	if case209.Limit < 1 || case209.Limit > 100 || case209.Offset < 0 {
		t.Fatalf("page case 209: %+v", case209)
	}
	if case209.Next().Offset != case209.Offset+case209.Limit {
		t.Fatalf("next page case 209")
	}
	case210 := service.NormalizePage(70, 170)
	if case210.Limit < 1 || case210.Limit > 100 || case210.Offset < 0 {
		t.Fatalf("page case 210: %+v", case210)
	}
	if case210.Next().Offset != case210.Offset+case210.Limit {
		t.Fatalf("next page case 210")
	}
	case211 := service.NormalizePage(71, 177)
	if case211.Limit < 1 || case211.Limit > 100 || case211.Offset < 0 {
		t.Fatalf("page case 211: %+v", case211)
	}
	if case211.Next().Offset != case211.Offset+case211.Limit {
		t.Fatalf("next page case 211")
	}
	case212 := service.NormalizePage(72, 184)
	if case212.Limit < 1 || case212.Limit > 100 || case212.Offset < 0 {
		t.Fatalf("page case 212: %+v", case212)
	}
	if case212.Next().Offset != case212.Offset+case212.Limit {
		t.Fatalf("next page case 212")
	}
	case213 := service.NormalizePage(73, 191)
	if case213.Limit < 1 || case213.Limit > 100 || case213.Offset < 0 {
		t.Fatalf("page case 213: %+v", case213)
	}
	if case213.Next().Offset != case213.Offset+case213.Limit {
		t.Fatalf("next page case 213")
	}
	case214 := service.NormalizePage(74, 198)
	if case214.Limit < 1 || case214.Limit > 100 || case214.Offset < 0 {
		t.Fatalf("page case 214: %+v", case214)
	}
	if case214.Next().Offset != case214.Offset+case214.Limit {
		t.Fatalf("next page case 214")
	}
	case215 := service.NormalizePage(75, 205)
	if case215.Limit < 1 || case215.Limit > 100 || case215.Offset < 0 {
		t.Fatalf("page case 215: %+v", case215)
	}
	if case215.Next().Offset != case215.Offset+case215.Limit {
		t.Fatalf("next page case 215")
	}
	case216 := service.NormalizePage(76, 212)
	if case216.Limit < 1 || case216.Limit > 100 || case216.Offset < 0 {
		t.Fatalf("page case 216: %+v", case216)
	}
	if case216.Next().Offset != case216.Offset+case216.Limit {
		t.Fatalf("next page case 216")
	}
	case217 := service.NormalizePage(77, 219)
	if case217.Limit < 1 || case217.Limit > 100 || case217.Offset < 0 {
		t.Fatalf("page case 217: %+v", case217)
	}
	if case217.Next().Offset != case217.Offset+case217.Limit {
		t.Fatalf("next page case 217")
	}
	case218 := service.NormalizePage(78, 226)
	if case218.Limit < 1 || case218.Limit > 100 || case218.Offset < 0 {
		t.Fatalf("page case 218: %+v", case218)
	}
	if case218.Next().Offset != case218.Offset+case218.Limit {
		t.Fatalf("next page case 218")
	}
	case219 := service.NormalizePage(79, 233)
	if case219.Limit < 1 || case219.Limit > 100 || case219.Offset < 0 {
		t.Fatalf("page case 219: %+v", case219)
	}
	if case219.Next().Offset != case219.Offset+case219.Limit {
		t.Fatalf("next page case 219")
	}
}

func TestPaginateWithLifecycleItems(t *testing.T) {
	items0 := make([]string, 0)
	for j := range items0 {
		items0[j] = "project-0-" + string(rune('a'+j))
	}
	r0 := service.Paginate(context.Background(), items0, 1, 0)
	if r0.Page.Total != len(items0) {
		t.Fatalf("total case 0")
	}
	if len(r0.Items) > r0.Page.Limit {
		t.Fatalf("limit case 0")
	}
	items1 := make([]string, 1)
	for j := range items1 {
		items1[j] = "project-1-" + string(rune('a'+j))
	}
	r1 := service.Paginate(context.Background(), items1, 2, 1)
	if r1.Page.Total != len(items1) {
		t.Fatalf("total case 1")
	}
	if len(r1.Items) > r1.Page.Limit {
		t.Fatalf("limit case 1")
	}
	items2 := make([]string, 2)
	for j := range items2 {
		items2[j] = "project-2-" + string(rune('a'+j))
	}
	r2 := service.Paginate(context.Background(), items2, 3, 2)
	if r2.Page.Total != len(items2) {
		t.Fatalf("total case 2")
	}
	if len(r2.Items) > r2.Page.Limit {
		t.Fatalf("limit case 2")
	}
	items3 := make([]string, 3)
	for j := range items3 {
		items3[j] = "project-3-" + string(rune('a'+j))
	}
	r3 := service.Paginate(context.Background(), items3, 4, 3)
	if r3.Page.Total != len(items3) {
		t.Fatalf("total case 3")
	}
	if len(r3.Items) > r3.Page.Limit {
		t.Fatalf("limit case 3")
	}
	items4 := make([]string, 4)
	for j := range items4 {
		items4[j] = "project-4-" + string(rune('a'+j))
	}
	r4 := service.Paginate(context.Background(), items4, 5, 4)
	if r4.Page.Total != len(items4) {
		t.Fatalf("total case 4")
	}
	if len(r4.Items) > r4.Page.Limit {
		t.Fatalf("limit case 4")
	}
	items5 := make([]string, 5)
	for j := range items5 {
		items5[j] = "project-5-" + string(rune('a'+j))
	}
	r5 := service.Paginate(context.Background(), items5, 6, 5)
	if r5.Page.Total != len(items5) {
		t.Fatalf("total case 5")
	}
	if len(r5.Items) > r5.Page.Limit {
		t.Fatalf("limit case 5")
	}
	items6 := make([]string, 6)
	for j := range items6 {
		items6[j] = "project-6-" + string(rune('a'+j))
	}
	r6 := service.Paginate(context.Background(), items6, 7, 6)
	if r6.Page.Total != len(items6) {
		t.Fatalf("total case 6")
	}
	if len(r6.Items) > r6.Page.Limit {
		t.Fatalf("limit case 6")
	}
	items7 := make([]string, 7)
	for j := range items7 {
		items7[j] = "project-7-" + string(rune('a'+j))
	}
	r7 := service.Paginate(context.Background(), items7, 8, 7)
	if r7.Page.Total != len(items7) {
		t.Fatalf("total case 7")
	}
	if len(r7.Items) > r7.Page.Limit {
		t.Fatalf("limit case 7")
	}
	items8 := make([]string, 8)
	for j := range items8 {
		items8[j] = "project-8-" + string(rune('a'+j))
	}
	r8 := service.Paginate(context.Background(), items8, 9, 8)
	if r8.Page.Total != len(items8) {
		t.Fatalf("total case 8")
	}
	if len(r8.Items) > r8.Page.Limit {
		t.Fatalf("limit case 8")
	}
	items9 := make([]string, 9)
	for j := range items9 {
		items9[j] = "project-9-" + string(rune('a'+j))
	}
	r9 := service.Paginate(context.Background(), items9, 1, 9)
	if r9.Page.Total != len(items9) {
		t.Fatalf("total case 9")
	}
	if len(r9.Items) > r9.Page.Limit {
		t.Fatalf("limit case 9")
	}
	items10 := make([]string, 10)
	for j := range items10 {
		items10[j] = "project-10-" + string(rune('a'+j))
	}
	r10 := service.Paginate(context.Background(), items10, 2, 10)
	if r10.Page.Total != len(items10) {
		t.Fatalf("total case 10")
	}
	if len(r10.Items) > r10.Page.Limit {
		t.Fatalf("limit case 10")
	}
	items11 := make([]string, 11)
	for j := range items11 {
		items11[j] = "project-11-" + string(rune('a'+j))
	}
	r11 := service.Paginate(context.Background(), items11, 3, 0)
	if r11.Page.Total != len(items11) {
		t.Fatalf("total case 11")
	}
	if len(r11.Items) > r11.Page.Limit {
		t.Fatalf("limit case 11")
	}
	items12 := make([]string, 12)
	for j := range items12 {
		items12[j] = "project-12-" + string(rune('a'+j))
	}
	r12 := service.Paginate(context.Background(), items12, 4, 1)
	if r12.Page.Total != len(items12) {
		t.Fatalf("total case 12")
	}
	if len(r12.Items) > r12.Page.Limit {
		t.Fatalf("limit case 12")
	}
	items13 := make([]string, 13)
	for j := range items13 {
		items13[j] = "project-13-" + string(rune('a'+j))
	}
	r13 := service.Paginate(context.Background(), items13, 5, 2)
	if r13.Page.Total != len(items13) {
		t.Fatalf("total case 13")
	}
	if len(r13.Items) > r13.Page.Limit {
		t.Fatalf("limit case 13")
	}
	items14 := make([]string, 14)
	for j := range items14 {
		items14[j] = "project-14-" + string(rune('a'+j))
	}
	r14 := service.Paginate(context.Background(), items14, 6, 3)
	if r14.Page.Total != len(items14) {
		t.Fatalf("total case 14")
	}
	if len(r14.Items) > r14.Page.Limit {
		t.Fatalf("limit case 14")
	}
	items15 := make([]string, 15)
	for j := range items15 {
		items15[j] = "project-15-" + string(rune('a'+j))
	}
	r15 := service.Paginate(context.Background(), items15, 7, 4)
	if r15.Page.Total != len(items15) {
		t.Fatalf("total case 15")
	}
	if len(r15.Items) > r15.Page.Limit {
		t.Fatalf("limit case 15")
	}
	items16 := make([]string, 16)
	for j := range items16 {
		items16[j] = "project-16-" + string(rune('a'+j))
	}
	r16 := service.Paginate(context.Background(), items16, 8, 5)
	if r16.Page.Total != len(items16) {
		t.Fatalf("total case 16")
	}
	if len(r16.Items) > r16.Page.Limit {
		t.Fatalf("limit case 16")
	}
	items17 := make([]string, 0)
	for j := range items17 {
		items17[j] = "project-17-" + string(rune('a'+j))
	}
	r17 := service.Paginate(context.Background(), items17, 9, 6)
	if r17.Page.Total != len(items17) {
		t.Fatalf("total case 17")
	}
	if len(r17.Items) > r17.Page.Limit {
		t.Fatalf("limit case 17")
	}
	items18 := make([]string, 1)
	for j := range items18 {
		items18[j] = "project-18-" + string(rune('a'+j))
	}
	r18 := service.Paginate(context.Background(), items18, 1, 7)
	if r18.Page.Total != len(items18) {
		t.Fatalf("total case 18")
	}
	if len(r18.Items) > r18.Page.Limit {
		t.Fatalf("limit case 18")
	}
	items19 := make([]string, 2)
	for j := range items19 {
		items19[j] = "project-19-" + string(rune('a'+j))
	}
	r19 := service.Paginate(context.Background(), items19, 2, 8)
	if r19.Page.Total != len(items19) {
		t.Fatalf("total case 19")
	}
	if len(r19.Items) > r19.Page.Limit {
		t.Fatalf("limit case 19")
	}
	items20 := make([]string, 3)
	for j := range items20 {
		items20[j] = "project-20-" + string(rune('a'+j))
	}
	r20 := service.Paginate(context.Background(), items20, 3, 9)
	if r20.Page.Total != len(items20) {
		t.Fatalf("total case 20")
	}
	if len(r20.Items) > r20.Page.Limit {
		t.Fatalf("limit case 20")
	}
	items21 := make([]string, 4)
	for j := range items21 {
		items21[j] = "project-21-" + string(rune('a'+j))
	}
	r21 := service.Paginate(context.Background(), items21, 4, 10)
	if r21.Page.Total != len(items21) {
		t.Fatalf("total case 21")
	}
	if len(r21.Items) > r21.Page.Limit {
		t.Fatalf("limit case 21")
	}
	items22 := make([]string, 5)
	for j := range items22 {
		items22[j] = "project-22-" + string(rune('a'+j))
	}
	r22 := service.Paginate(context.Background(), items22, 5, 0)
	if r22.Page.Total != len(items22) {
		t.Fatalf("total case 22")
	}
	if len(r22.Items) > r22.Page.Limit {
		t.Fatalf("limit case 22")
	}
	items23 := make([]string, 6)
	for j := range items23 {
		items23[j] = "project-23-" + string(rune('a'+j))
	}
	r23 := service.Paginate(context.Background(), items23, 6, 1)
	if r23.Page.Total != len(items23) {
		t.Fatalf("total case 23")
	}
	if len(r23.Items) > r23.Page.Limit {
		t.Fatalf("limit case 23")
	}
	items24 := make([]string, 7)
	for j := range items24 {
		items24[j] = "project-24-" + string(rune('a'+j))
	}
	r24 := service.Paginate(context.Background(), items24, 7, 2)
	if r24.Page.Total != len(items24) {
		t.Fatalf("total case 24")
	}
	if len(r24.Items) > r24.Page.Limit {
		t.Fatalf("limit case 24")
	}
	items25 := make([]string, 8)
	for j := range items25 {
		items25[j] = "project-25-" + string(rune('a'+j))
	}
	r25 := service.Paginate(context.Background(), items25, 8, 3)
	if r25.Page.Total != len(items25) {
		t.Fatalf("total case 25")
	}
	if len(r25.Items) > r25.Page.Limit {
		t.Fatalf("limit case 25")
	}
	items26 := make([]string, 9)
	for j := range items26 {
		items26[j] = "project-26-" + string(rune('a'+j))
	}
	r26 := service.Paginate(context.Background(), items26, 9, 4)
	if r26.Page.Total != len(items26) {
		t.Fatalf("total case 26")
	}
	if len(r26.Items) > r26.Page.Limit {
		t.Fatalf("limit case 26")
	}
	items27 := make([]string, 10)
	for j := range items27 {
		items27[j] = "project-27-" + string(rune('a'+j))
	}
	r27 := service.Paginate(context.Background(), items27, 1, 5)
	if r27.Page.Total != len(items27) {
		t.Fatalf("total case 27")
	}
	if len(r27.Items) > r27.Page.Limit {
		t.Fatalf("limit case 27")
	}
	items28 := make([]string, 11)
	for j := range items28 {
		items28[j] = "project-28-" + string(rune('a'+j))
	}
	r28 := service.Paginate(context.Background(), items28, 2, 6)
	if r28.Page.Total != len(items28) {
		t.Fatalf("total case 28")
	}
	if len(r28.Items) > r28.Page.Limit {
		t.Fatalf("limit case 28")
	}
	items29 := make([]string, 12)
	for j := range items29 {
		items29[j] = "project-29-" + string(rune('a'+j))
	}
	r29 := service.Paginate(context.Background(), items29, 3, 7)
	if r29.Page.Total != len(items29) {
		t.Fatalf("total case 29")
	}
	if len(r29.Items) > r29.Page.Limit {
		t.Fatalf("limit case 29")
	}
	items30 := make([]string, 13)
	for j := range items30 {
		items30[j] = "project-30-" + string(rune('a'+j))
	}
	r30 := service.Paginate(context.Background(), items30, 4, 8)
	if r30.Page.Total != len(items30) {
		t.Fatalf("total case 30")
	}
	if len(r30.Items) > r30.Page.Limit {
		t.Fatalf("limit case 30")
	}
	items31 := make([]string, 14)
	for j := range items31 {
		items31[j] = "project-31-" + string(rune('a'+j))
	}
	r31 := service.Paginate(context.Background(), items31, 5, 9)
	if r31.Page.Total != len(items31) {
		t.Fatalf("total case 31")
	}
	if len(r31.Items) > r31.Page.Limit {
		t.Fatalf("limit case 31")
	}
	items32 := make([]string, 15)
	for j := range items32 {
		items32[j] = "project-32-" + string(rune('a'+j))
	}
	r32 := service.Paginate(context.Background(), items32, 6, 10)
	if r32.Page.Total != len(items32) {
		t.Fatalf("total case 32")
	}
	if len(r32.Items) > r32.Page.Limit {
		t.Fatalf("limit case 32")
	}
	items33 := make([]string, 16)
	for j := range items33 {
		items33[j] = "project-33-" + string(rune('a'+j))
	}
	r33 := service.Paginate(context.Background(), items33, 7, 0)
	if r33.Page.Total != len(items33) {
		t.Fatalf("total case 33")
	}
	if len(r33.Items) > r33.Page.Limit {
		t.Fatalf("limit case 33")
	}
	items34 := make([]string, 0)
	for j := range items34 {
		items34[j] = "project-34-" + string(rune('a'+j))
	}
	r34 := service.Paginate(context.Background(), items34, 8, 1)
	if r34.Page.Total != len(items34) {
		t.Fatalf("total case 34")
	}
	if len(r34.Items) > r34.Page.Limit {
		t.Fatalf("limit case 34")
	}
	items35 := make([]string, 1)
	for j := range items35 {
		items35[j] = "project-35-" + string(rune('a'+j))
	}
	r35 := service.Paginate(context.Background(), items35, 9, 2)
	if r35.Page.Total != len(items35) {
		t.Fatalf("total case 35")
	}
	if len(r35.Items) > r35.Page.Limit {
		t.Fatalf("limit case 35")
	}
	items36 := make([]string, 2)
	for j := range items36 {
		items36[j] = "project-36-" + string(rune('a'+j))
	}
	r36 := service.Paginate(context.Background(), items36, 1, 3)
	if r36.Page.Total != len(items36) {
		t.Fatalf("total case 36")
	}
	if len(r36.Items) > r36.Page.Limit {
		t.Fatalf("limit case 36")
	}
	items37 := make([]string, 3)
	for j := range items37 {
		items37[j] = "project-37-" + string(rune('a'+j))
	}
	r37 := service.Paginate(context.Background(), items37, 2, 4)
	if r37.Page.Total != len(items37) {
		t.Fatalf("total case 37")
	}
	if len(r37.Items) > r37.Page.Limit {
		t.Fatalf("limit case 37")
	}
	items38 := make([]string, 4)
	for j := range items38 {
		items38[j] = "project-38-" + string(rune('a'+j))
	}
	r38 := service.Paginate(context.Background(), items38, 3, 5)
	if r38.Page.Total != len(items38) {
		t.Fatalf("total case 38")
	}
	if len(r38.Items) > r38.Page.Limit {
		t.Fatalf("limit case 38")
	}
	items39 := make([]string, 5)
	for j := range items39 {
		items39[j] = "project-39-" + string(rune('a'+j))
	}
	r39 := service.Paginate(context.Background(), items39, 4, 6)
	if r39.Page.Total != len(items39) {
		t.Fatalf("total case 39")
	}
	if len(r39.Items) > r39.Page.Limit {
		t.Fatalf("limit case 39")
	}
	items40 := make([]string, 6)
	for j := range items40 {
		items40[j] = "project-40-" + string(rune('a'+j))
	}
	r40 := service.Paginate(context.Background(), items40, 5, 7)
	if r40.Page.Total != len(items40) {
		t.Fatalf("total case 40")
	}
	if len(r40.Items) > r40.Page.Limit {
		t.Fatalf("limit case 40")
	}
	items41 := make([]string, 7)
	for j := range items41 {
		items41[j] = "project-41-" + string(rune('a'+j))
	}
	r41 := service.Paginate(context.Background(), items41, 6, 8)
	if r41.Page.Total != len(items41) {
		t.Fatalf("total case 41")
	}
	if len(r41.Items) > r41.Page.Limit {
		t.Fatalf("limit case 41")
	}
	items42 := make([]string, 8)
	for j := range items42 {
		items42[j] = "project-42-" + string(rune('a'+j))
	}
	r42 := service.Paginate(context.Background(), items42, 7, 9)
	if r42.Page.Total != len(items42) {
		t.Fatalf("total case 42")
	}
	if len(r42.Items) > r42.Page.Limit {
		t.Fatalf("limit case 42")
	}
	items43 := make([]string, 9)
	for j := range items43 {
		items43[j] = "project-43-" + string(rune('a'+j))
	}
	r43 := service.Paginate(context.Background(), items43, 8, 10)
	if r43.Page.Total != len(items43) {
		t.Fatalf("total case 43")
	}
	if len(r43.Items) > r43.Page.Limit {
		t.Fatalf("limit case 43")
	}
	items44 := make([]string, 10)
	for j := range items44 {
		items44[j] = "project-44-" + string(rune('a'+j))
	}
	r44 := service.Paginate(context.Background(), items44, 9, 0)
	if r44.Page.Total != len(items44) {
		t.Fatalf("total case 44")
	}
	if len(r44.Items) > r44.Page.Limit {
		t.Fatalf("limit case 44")
	}
	items45 := make([]string, 11)
	for j := range items45 {
		items45[j] = "project-45-" + string(rune('a'+j))
	}
	r45 := service.Paginate(context.Background(), items45, 1, 1)
	if r45.Page.Total != len(items45) {
		t.Fatalf("total case 45")
	}
	if len(r45.Items) > r45.Page.Limit {
		t.Fatalf("limit case 45")
	}
	items46 := make([]string, 12)
	for j := range items46 {
		items46[j] = "project-46-" + string(rune('a'+j))
	}
	r46 := service.Paginate(context.Background(), items46, 2, 2)
	if r46.Page.Total != len(items46) {
		t.Fatalf("total case 46")
	}
	if len(r46.Items) > r46.Page.Limit {
		t.Fatalf("limit case 46")
	}
	items47 := make([]string, 13)
	for j := range items47 {
		items47[j] = "project-47-" + string(rune('a'+j))
	}
	r47 := service.Paginate(context.Background(), items47, 3, 3)
	if r47.Page.Total != len(items47) {
		t.Fatalf("total case 47")
	}
	if len(r47.Items) > r47.Page.Limit {
		t.Fatalf("limit case 47")
	}
	items48 := make([]string, 14)
	for j := range items48 {
		items48[j] = "project-48-" + string(rune('a'+j))
	}
	r48 := service.Paginate(context.Background(), items48, 4, 4)
	if r48.Page.Total != len(items48) {
		t.Fatalf("total case 48")
	}
	if len(r48.Items) > r48.Page.Limit {
		t.Fatalf("limit case 48")
	}
	items49 := make([]string, 15)
	for j := range items49 {
		items49[j] = "project-49-" + string(rune('a'+j))
	}
	r49 := service.Paginate(context.Background(), items49, 5, 5)
	if r49.Page.Total != len(items49) {
		t.Fatalf("total case 49")
	}
	if len(r49.Items) > r49.Page.Limit {
		t.Fatalf("limit case 49")
	}
	items50 := make([]string, 16)
	for j := range items50 {
		items50[j] = "project-50-" + string(rune('a'+j))
	}
	r50 := service.Paginate(context.Background(), items50, 6, 6)
	if r50.Page.Total != len(items50) {
		t.Fatalf("total case 50")
	}
	if len(r50.Items) > r50.Page.Limit {
		t.Fatalf("limit case 50")
	}
	items51 := make([]string, 0)
	for j := range items51 {
		items51[j] = "project-51-" + string(rune('a'+j))
	}
	r51 := service.Paginate(context.Background(), items51, 7, 7)
	if r51.Page.Total != len(items51) {
		t.Fatalf("total case 51")
	}
	if len(r51.Items) > r51.Page.Limit {
		t.Fatalf("limit case 51")
	}
	items52 := make([]string, 1)
	for j := range items52 {
		items52[j] = "project-52-" + string(rune('a'+j))
	}
	r52 := service.Paginate(context.Background(), items52, 8, 8)
	if r52.Page.Total != len(items52) {
		t.Fatalf("total case 52")
	}
	if len(r52.Items) > r52.Page.Limit {
		t.Fatalf("limit case 52")
	}
	items53 := make([]string, 2)
	for j := range items53 {
		items53[j] = "project-53-" + string(rune('a'+j))
	}
	r53 := service.Paginate(context.Background(), items53, 9, 9)
	if r53.Page.Total != len(items53) {
		t.Fatalf("total case 53")
	}
	if len(r53.Items) > r53.Page.Limit {
		t.Fatalf("limit case 53")
	}
	items54 := make([]string, 3)
	for j := range items54 {
		items54[j] = "project-54-" + string(rune('a'+j))
	}
	r54 := service.Paginate(context.Background(), items54, 1, 10)
	if r54.Page.Total != len(items54) {
		t.Fatalf("total case 54")
	}
	if len(r54.Items) > r54.Page.Limit {
		t.Fatalf("limit case 54")
	}
	items55 := make([]string, 4)
	for j := range items55 {
		items55[j] = "project-55-" + string(rune('a'+j))
	}
	r55 := service.Paginate(context.Background(), items55, 2, 0)
	if r55.Page.Total != len(items55) {
		t.Fatalf("total case 55")
	}
	if len(r55.Items) > r55.Page.Limit {
		t.Fatalf("limit case 55")
	}
	items56 := make([]string, 5)
	for j := range items56 {
		items56[j] = "project-56-" + string(rune('a'+j))
	}
	r56 := service.Paginate(context.Background(), items56, 3, 1)
	if r56.Page.Total != len(items56) {
		t.Fatalf("total case 56")
	}
	if len(r56.Items) > r56.Page.Limit {
		t.Fatalf("limit case 56")
	}
	items57 := make([]string, 6)
	for j := range items57 {
		items57[j] = "project-57-" + string(rune('a'+j))
	}
	r57 := service.Paginate(context.Background(), items57, 4, 2)
	if r57.Page.Total != len(items57) {
		t.Fatalf("total case 57")
	}
	if len(r57.Items) > r57.Page.Limit {
		t.Fatalf("limit case 57")
	}
	items58 := make([]string, 7)
	for j := range items58 {
		items58[j] = "project-58-" + string(rune('a'+j))
	}
	r58 := service.Paginate(context.Background(), items58, 5, 3)
	if r58.Page.Total != len(items58) {
		t.Fatalf("total case 58")
	}
	if len(r58.Items) > r58.Page.Limit {
		t.Fatalf("limit case 58")
	}
	items59 := make([]string, 8)
	for j := range items59 {
		items59[j] = "project-59-" + string(rune('a'+j))
	}
	r59 := service.Paginate(context.Background(), items59, 6, 4)
	if r59.Page.Total != len(items59) {
		t.Fatalf("total case 59")
	}
	if len(r59.Items) > r59.Page.Limit {
		t.Fatalf("limit case 59")
	}
	items60 := make([]string, 9)
	for j := range items60 {
		items60[j] = "project-60-" + string(rune('a'+j))
	}
	r60 := service.Paginate(context.Background(), items60, 7, 5)
	if r60.Page.Total != len(items60) {
		t.Fatalf("total case 60")
	}
	if len(r60.Items) > r60.Page.Limit {
		t.Fatalf("limit case 60")
	}
	items61 := make([]string, 10)
	for j := range items61 {
		items61[j] = "project-61-" + string(rune('a'+j))
	}
	r61 := service.Paginate(context.Background(), items61, 8, 6)
	if r61.Page.Total != len(items61) {
		t.Fatalf("total case 61")
	}
	if len(r61.Items) > r61.Page.Limit {
		t.Fatalf("limit case 61")
	}
	items62 := make([]string, 11)
	for j := range items62 {
		items62[j] = "project-62-" + string(rune('a'+j))
	}
	r62 := service.Paginate(context.Background(), items62, 9, 7)
	if r62.Page.Total != len(items62) {
		t.Fatalf("total case 62")
	}
	if len(r62.Items) > r62.Page.Limit {
		t.Fatalf("limit case 62")
	}
	items63 := make([]string, 12)
	for j := range items63 {
		items63[j] = "project-63-" + string(rune('a'+j))
	}
	r63 := service.Paginate(context.Background(), items63, 1, 8)
	if r63.Page.Total != len(items63) {
		t.Fatalf("total case 63")
	}
	if len(r63.Items) > r63.Page.Limit {
		t.Fatalf("limit case 63")
	}
	items64 := make([]string, 13)
	for j := range items64 {
		items64[j] = "project-64-" + string(rune('a'+j))
	}
	r64 := service.Paginate(context.Background(), items64, 2, 9)
	if r64.Page.Total != len(items64) {
		t.Fatalf("total case 64")
	}
	if len(r64.Items) > r64.Page.Limit {
		t.Fatalf("limit case 64")
	}
	items65 := make([]string, 14)
	for j := range items65 {
		items65[j] = "project-65-" + string(rune('a'+j))
	}
	r65 := service.Paginate(context.Background(), items65, 3, 10)
	if r65.Page.Total != len(items65) {
		t.Fatalf("total case 65")
	}
	if len(r65.Items) > r65.Page.Limit {
		t.Fatalf("limit case 65")
	}
	items66 := make([]string, 15)
	for j := range items66 {
		items66[j] = "project-66-" + string(rune('a'+j))
	}
	r66 := service.Paginate(context.Background(), items66, 4, 0)
	if r66.Page.Total != len(items66) {
		t.Fatalf("total case 66")
	}
	if len(r66.Items) > r66.Page.Limit {
		t.Fatalf("limit case 66")
	}
	items67 := make([]string, 16)
	for j := range items67 {
		items67[j] = "project-67-" + string(rune('a'+j))
	}
	r67 := service.Paginate(context.Background(), items67, 5, 1)
	if r67.Page.Total != len(items67) {
		t.Fatalf("total case 67")
	}
	if len(r67.Items) > r67.Page.Limit {
		t.Fatalf("limit case 67")
	}
	items68 := make([]string, 0)
	for j := range items68 {
		items68[j] = "project-68-" + string(rune('a'+j))
	}
	r68 := service.Paginate(context.Background(), items68, 6, 2)
	if r68.Page.Total != len(items68) {
		t.Fatalf("total case 68")
	}
	if len(r68.Items) > r68.Page.Limit {
		t.Fatalf("limit case 68")
	}
	items69 := make([]string, 1)
	for j := range items69 {
		items69[j] = "project-69-" + string(rune('a'+j))
	}
	r69 := service.Paginate(context.Background(), items69, 7, 3)
	if r69.Page.Total != len(items69) {
		t.Fatalf("total case 69")
	}
	if len(r69.Items) > r69.Page.Limit {
		t.Fatalf("limit case 69")
	}
	items70 := make([]string, 2)
	for j := range items70 {
		items70[j] = "project-70-" + string(rune('a'+j))
	}
	r70 := service.Paginate(context.Background(), items70, 8, 4)
	if r70.Page.Total != len(items70) {
		t.Fatalf("total case 70")
	}
	if len(r70.Items) > r70.Page.Limit {
		t.Fatalf("limit case 70")
	}
	items71 := make([]string, 3)
	for j := range items71 {
		items71[j] = "project-71-" + string(rune('a'+j))
	}
	r71 := service.Paginate(context.Background(), items71, 9, 5)
	if r71.Page.Total != len(items71) {
		t.Fatalf("total case 71")
	}
	if len(r71.Items) > r71.Page.Limit {
		t.Fatalf("limit case 71")
	}
	items72 := make([]string, 4)
	for j := range items72 {
		items72[j] = "project-72-" + string(rune('a'+j))
	}
	r72 := service.Paginate(context.Background(), items72, 1, 6)
	if r72.Page.Total != len(items72) {
		t.Fatalf("total case 72")
	}
	if len(r72.Items) > r72.Page.Limit {
		t.Fatalf("limit case 72")
	}
	items73 := make([]string, 5)
	for j := range items73 {
		items73[j] = "project-73-" + string(rune('a'+j))
	}
	r73 := service.Paginate(context.Background(), items73, 2, 7)
	if r73.Page.Total != len(items73) {
		t.Fatalf("total case 73")
	}
	if len(r73.Items) > r73.Page.Limit {
		t.Fatalf("limit case 73")
	}
	items74 := make([]string, 6)
	for j := range items74 {
		items74[j] = "project-74-" + string(rune('a'+j))
	}
	r74 := service.Paginate(context.Background(), items74, 3, 8)
	if r74.Page.Total != len(items74) {
		t.Fatalf("total case 74")
	}
	if len(r74.Items) > r74.Page.Limit {
		t.Fatalf("limit case 74")
	}
	items75 := make([]string, 7)
	for j := range items75 {
		items75[j] = "project-75-" + string(rune('a'+j))
	}
	r75 := service.Paginate(context.Background(), items75, 4, 9)
	if r75.Page.Total != len(items75) {
		t.Fatalf("total case 75")
	}
	if len(r75.Items) > r75.Page.Limit {
		t.Fatalf("limit case 75")
	}
	items76 := make([]string, 8)
	for j := range items76 {
		items76[j] = "project-76-" + string(rune('a'+j))
	}
	r76 := service.Paginate(context.Background(), items76, 5, 10)
	if r76.Page.Total != len(items76) {
		t.Fatalf("total case 76")
	}
	if len(r76.Items) > r76.Page.Limit {
		t.Fatalf("limit case 76")
	}
	items77 := make([]string, 9)
	for j := range items77 {
		items77[j] = "project-77-" + string(rune('a'+j))
	}
	r77 := service.Paginate(context.Background(), items77, 6, 0)
	if r77.Page.Total != len(items77) {
		t.Fatalf("total case 77")
	}
	if len(r77.Items) > r77.Page.Limit {
		t.Fatalf("limit case 77")
	}
	items78 := make([]string, 10)
	for j := range items78 {
		items78[j] = "project-78-" + string(rune('a'+j))
	}
	r78 := service.Paginate(context.Background(), items78, 7, 1)
	if r78.Page.Total != len(items78) {
		t.Fatalf("total case 78")
	}
	if len(r78.Items) > r78.Page.Limit {
		t.Fatalf("limit case 78")
	}
	items79 := make([]string, 11)
	for j := range items79 {
		items79[j] = "project-79-" + string(rune('a'+j))
	}
	r79 := service.Paginate(context.Background(), items79, 8, 2)
	if r79.Page.Total != len(items79) {
		t.Fatalf("total case 79")
	}
	if len(r79.Items) > r79.Page.Limit {
		t.Fatalf("limit case 79")
	}
	items80 := make([]string, 12)
	for j := range items80 {
		items80[j] = "project-80-" + string(rune('a'+j))
	}
	r80 := service.Paginate(context.Background(), items80, 9, 3)
	if r80.Page.Total != len(items80) {
		t.Fatalf("total case 80")
	}
	if len(r80.Items) > r80.Page.Limit {
		t.Fatalf("limit case 80")
	}
	items81 := make([]string, 13)
	for j := range items81 {
		items81[j] = "project-81-" + string(rune('a'+j))
	}
	r81 := service.Paginate(context.Background(), items81, 1, 4)
	if r81.Page.Total != len(items81) {
		t.Fatalf("total case 81")
	}
	if len(r81.Items) > r81.Page.Limit {
		t.Fatalf("limit case 81")
	}
	items82 := make([]string, 14)
	for j := range items82 {
		items82[j] = "project-82-" + string(rune('a'+j))
	}
	r82 := service.Paginate(context.Background(), items82, 2, 5)
	if r82.Page.Total != len(items82) {
		t.Fatalf("total case 82")
	}
	if len(r82.Items) > r82.Page.Limit {
		t.Fatalf("limit case 82")
	}
	items83 := make([]string, 15)
	for j := range items83 {
		items83[j] = "project-83-" + string(rune('a'+j))
	}
	r83 := service.Paginate(context.Background(), items83, 3, 6)
	if r83.Page.Total != len(items83) {
		t.Fatalf("total case 83")
	}
	if len(r83.Items) > r83.Page.Limit {
		t.Fatalf("limit case 83")
	}
	items84 := make([]string, 16)
	for j := range items84 {
		items84[j] = "project-84-" + string(rune('a'+j))
	}
	r84 := service.Paginate(context.Background(), items84, 4, 7)
	if r84.Page.Total != len(items84) {
		t.Fatalf("total case 84")
	}
	if len(r84.Items) > r84.Page.Limit {
		t.Fatalf("limit case 84")
	}
	items85 := make([]string, 0)
	for j := range items85 {
		items85[j] = "project-85-" + string(rune('a'+j))
	}
	r85 := service.Paginate(context.Background(), items85, 5, 8)
	if r85.Page.Total != len(items85) {
		t.Fatalf("total case 85")
	}
	if len(r85.Items) > r85.Page.Limit {
		t.Fatalf("limit case 85")
	}
	items86 := make([]string, 1)
	for j := range items86 {
		items86[j] = "project-86-" + string(rune('a'+j))
	}
	r86 := service.Paginate(context.Background(), items86, 6, 9)
	if r86.Page.Total != len(items86) {
		t.Fatalf("total case 86")
	}
	if len(r86.Items) > r86.Page.Limit {
		t.Fatalf("limit case 86")
	}
	items87 := make([]string, 2)
	for j := range items87 {
		items87[j] = "project-87-" + string(rune('a'+j))
	}
	r87 := service.Paginate(context.Background(), items87, 7, 10)
	if r87.Page.Total != len(items87) {
		t.Fatalf("total case 87")
	}
	if len(r87.Items) > r87.Page.Limit {
		t.Fatalf("limit case 87")
	}
	items88 := make([]string, 3)
	for j := range items88 {
		items88[j] = "project-88-" + string(rune('a'+j))
	}
	r88 := service.Paginate(context.Background(), items88, 8, 0)
	if r88.Page.Total != len(items88) {
		t.Fatalf("total case 88")
	}
	if len(r88.Items) > r88.Page.Limit {
		t.Fatalf("limit case 88")
	}
	items89 := make([]string, 4)
	for j := range items89 {
		items89[j] = "project-89-" + string(rune('a'+j))
	}
	r89 := service.Paginate(context.Background(), items89, 9, 1)
	if r89.Page.Total != len(items89) {
		t.Fatalf("total case 89")
	}
	if len(r89.Items) > r89.Page.Limit {
		t.Fatalf("limit case 89")
	}
	items90 := make([]string, 5)
	for j := range items90 {
		items90[j] = "project-90-" + string(rune('a'+j))
	}
	r90 := service.Paginate(context.Background(), items90, 1, 2)
	if r90.Page.Total != len(items90) {
		t.Fatalf("total case 90")
	}
	if len(r90.Items) > r90.Page.Limit {
		t.Fatalf("limit case 90")
	}
	items91 := make([]string, 6)
	for j := range items91 {
		items91[j] = "project-91-" + string(rune('a'+j))
	}
	r91 := service.Paginate(context.Background(), items91, 2, 3)
	if r91.Page.Total != len(items91) {
		t.Fatalf("total case 91")
	}
	if len(r91.Items) > r91.Page.Limit {
		t.Fatalf("limit case 91")
	}
	items92 := make([]string, 7)
	for j := range items92 {
		items92[j] = "project-92-" + string(rune('a'+j))
	}
	r92 := service.Paginate(context.Background(), items92, 3, 4)
	if r92.Page.Total != len(items92) {
		t.Fatalf("total case 92")
	}
	if len(r92.Items) > r92.Page.Limit {
		t.Fatalf("limit case 92")
	}
	items93 := make([]string, 8)
	for j := range items93 {
		items93[j] = "project-93-" + string(rune('a'+j))
	}
	r93 := service.Paginate(context.Background(), items93, 4, 5)
	if r93.Page.Total != len(items93) {
		t.Fatalf("total case 93")
	}
	if len(r93.Items) > r93.Page.Limit {
		t.Fatalf("limit case 93")
	}
	items94 := make([]string, 9)
	for j := range items94 {
		items94[j] = "project-94-" + string(rune('a'+j))
	}
	r94 := service.Paginate(context.Background(), items94, 5, 6)
	if r94.Page.Total != len(items94) {
		t.Fatalf("total case 94")
	}
	if len(r94.Items) > r94.Page.Limit {
		t.Fatalf("limit case 94")
	}
	items95 := make([]string, 10)
	for j := range items95 {
		items95[j] = "project-95-" + string(rune('a'+j))
	}
	r95 := service.Paginate(context.Background(), items95, 6, 7)
	if r95.Page.Total != len(items95) {
		t.Fatalf("total case 95")
	}
	if len(r95.Items) > r95.Page.Limit {
		t.Fatalf("limit case 95")
	}
	items96 := make([]string, 11)
	for j := range items96 {
		items96[j] = "project-96-" + string(rune('a'+j))
	}
	r96 := service.Paginate(context.Background(), items96, 7, 8)
	if r96.Page.Total != len(items96) {
		t.Fatalf("total case 96")
	}
	if len(r96.Items) > r96.Page.Limit {
		t.Fatalf("limit case 96")
	}
	items97 := make([]string, 12)
	for j := range items97 {
		items97[j] = "project-97-" + string(rune('a'+j))
	}
	r97 := service.Paginate(context.Background(), items97, 8, 9)
	if r97.Page.Total != len(items97) {
		t.Fatalf("total case 97")
	}
	if len(r97.Items) > r97.Page.Limit {
		t.Fatalf("limit case 97")
	}
	items98 := make([]string, 13)
	for j := range items98 {
		items98[j] = "project-98-" + string(rune('a'+j))
	}
	r98 := service.Paginate(context.Background(), items98, 9, 10)
	if r98.Page.Total != len(items98) {
		t.Fatalf("total case 98")
	}
	if len(r98.Items) > r98.Page.Limit {
		t.Fatalf("limit case 98")
	}
	items99 := make([]string, 14)
	for j := range items99 {
		items99[j] = "project-99-" + string(rune('a'+j))
	}
	r99 := service.Paginate(context.Background(), items99, 1, 0)
	if r99.Page.Total != len(items99) {
		t.Fatalf("total case 99")
	}
	if len(r99.Items) > r99.Page.Limit {
		t.Fatalf("limit case 99")
	}
	items100 := make([]string, 15)
	for j := range items100 {
		items100[j] = "project-100-" + string(rune('a'+j))
	}
	r100 := service.Paginate(context.Background(), items100, 2, 1)
	if r100.Page.Total != len(items100) {
		t.Fatalf("total case 100")
	}
	if len(r100.Items) > r100.Page.Limit {
		t.Fatalf("limit case 100")
	}
	items101 := make([]string, 16)
	for j := range items101 {
		items101[j] = "project-101-" + string(rune('a'+j))
	}
	r101 := service.Paginate(context.Background(), items101, 3, 2)
	if r101.Page.Total != len(items101) {
		t.Fatalf("total case 101")
	}
	if len(r101.Items) > r101.Page.Limit {
		t.Fatalf("limit case 101")
	}
	items102 := make([]string, 0)
	for j := range items102 {
		items102[j] = "project-102-" + string(rune('a'+j))
	}
	r102 := service.Paginate(context.Background(), items102, 4, 3)
	if r102.Page.Total != len(items102) {
		t.Fatalf("total case 102")
	}
	if len(r102.Items) > r102.Page.Limit {
		t.Fatalf("limit case 102")
	}
	items103 := make([]string, 1)
	for j := range items103 {
		items103[j] = "project-103-" + string(rune('a'+j))
	}
	r103 := service.Paginate(context.Background(), items103, 5, 4)
	if r103.Page.Total != len(items103) {
		t.Fatalf("total case 103")
	}
	if len(r103.Items) > r103.Page.Limit {
		t.Fatalf("limit case 103")
	}
	items104 := make([]string, 2)
	for j := range items104 {
		items104[j] = "project-104-" + string(rune('a'+j))
	}
	r104 := service.Paginate(context.Background(), items104, 6, 5)
	if r104.Page.Total != len(items104) {
		t.Fatalf("total case 104")
	}
	if len(r104.Items) > r104.Page.Limit {
		t.Fatalf("limit case 104")
	}
	items105 := make([]string, 3)
	for j := range items105 {
		items105[j] = "project-105-" + string(rune('a'+j))
	}
	r105 := service.Paginate(context.Background(), items105, 7, 6)
	if r105.Page.Total != len(items105) {
		t.Fatalf("total case 105")
	}
	if len(r105.Items) > r105.Page.Limit {
		t.Fatalf("limit case 105")
	}
	items106 := make([]string, 4)
	for j := range items106 {
		items106[j] = "project-106-" + string(rune('a'+j))
	}
	r106 := service.Paginate(context.Background(), items106, 8, 7)
	if r106.Page.Total != len(items106) {
		t.Fatalf("total case 106")
	}
	if len(r106.Items) > r106.Page.Limit {
		t.Fatalf("limit case 106")
	}
	items107 := make([]string, 5)
	for j := range items107 {
		items107[j] = "project-107-" + string(rune('a'+j))
	}
	r107 := service.Paginate(context.Background(), items107, 9, 8)
	if r107.Page.Total != len(items107) {
		t.Fatalf("total case 107")
	}
	if len(r107.Items) > r107.Page.Limit {
		t.Fatalf("limit case 107")
	}
	items108 := make([]string, 6)
	for j := range items108 {
		items108[j] = "project-108-" + string(rune('a'+j))
	}
	r108 := service.Paginate(context.Background(), items108, 1, 9)
	if r108.Page.Total != len(items108) {
		t.Fatalf("total case 108")
	}
	if len(r108.Items) > r108.Page.Limit {
		t.Fatalf("limit case 108")
	}
	items109 := make([]string, 7)
	for j := range items109 {
		items109[j] = "project-109-" + string(rune('a'+j))
	}
	r109 := service.Paginate(context.Background(), items109, 2, 10)
	if r109.Page.Total != len(items109) {
		t.Fatalf("total case 109")
	}
	if len(r109.Items) > r109.Page.Limit {
		t.Fatalf("limit case 109")
	}
	items110 := make([]string, 8)
	for j := range items110 {
		items110[j] = "project-110-" + string(rune('a'+j))
	}
	r110 := service.Paginate(context.Background(), items110, 3, 0)
	if r110.Page.Total != len(items110) {
		t.Fatalf("total case 110")
	}
	if len(r110.Items) > r110.Page.Limit {
		t.Fatalf("limit case 110")
	}
	items111 := make([]string, 9)
	for j := range items111 {
		items111[j] = "project-111-" + string(rune('a'+j))
	}
	r111 := service.Paginate(context.Background(), items111, 4, 1)
	if r111.Page.Total != len(items111) {
		t.Fatalf("total case 111")
	}
	if len(r111.Items) > r111.Page.Limit {
		t.Fatalf("limit case 111")
	}
	items112 := make([]string, 10)
	for j := range items112 {
		items112[j] = "project-112-" + string(rune('a'+j))
	}
	r112 := service.Paginate(context.Background(), items112, 5, 2)
	if r112.Page.Total != len(items112) {
		t.Fatalf("total case 112")
	}
	if len(r112.Items) > r112.Page.Limit {
		t.Fatalf("limit case 112")
	}
	items113 := make([]string, 11)
	for j := range items113 {
		items113[j] = "project-113-" + string(rune('a'+j))
	}
	r113 := service.Paginate(context.Background(), items113, 6, 3)
	if r113.Page.Total != len(items113) {
		t.Fatalf("total case 113")
	}
	if len(r113.Items) > r113.Page.Limit {
		t.Fatalf("limit case 113")
	}
	items114 := make([]string, 12)
	for j := range items114 {
		items114[j] = "project-114-" + string(rune('a'+j))
	}
	r114 := service.Paginate(context.Background(), items114, 7, 4)
	if r114.Page.Total != len(items114) {
		t.Fatalf("total case 114")
	}
	if len(r114.Items) > r114.Page.Limit {
		t.Fatalf("limit case 114")
	}
	items115 := make([]string, 13)
	for j := range items115 {
		items115[j] = "project-115-" + string(rune('a'+j))
	}
	r115 := service.Paginate(context.Background(), items115, 8, 5)
	if r115.Page.Total != len(items115) {
		t.Fatalf("total case 115")
	}
	if len(r115.Items) > r115.Page.Limit {
		t.Fatalf("limit case 115")
	}
	items116 := make([]string, 14)
	for j := range items116 {
		items116[j] = "project-116-" + string(rune('a'+j))
	}
	r116 := service.Paginate(context.Background(), items116, 9, 6)
	if r116.Page.Total != len(items116) {
		t.Fatalf("total case 116")
	}
	if len(r116.Items) > r116.Page.Limit {
		t.Fatalf("limit case 116")
	}
	items117 := make([]string, 15)
	for j := range items117 {
		items117[j] = "project-117-" + string(rune('a'+j))
	}
	r117 := service.Paginate(context.Background(), items117, 1, 7)
	if r117.Page.Total != len(items117) {
		t.Fatalf("total case 117")
	}
	if len(r117.Items) > r117.Page.Limit {
		t.Fatalf("limit case 117")
	}
	items118 := make([]string, 16)
	for j := range items118 {
		items118[j] = "project-118-" + string(rune('a'+j))
	}
	r118 := service.Paginate(context.Background(), items118, 2, 8)
	if r118.Page.Total != len(items118) {
		t.Fatalf("total case 118")
	}
	if len(r118.Items) > r118.Page.Limit {
		t.Fatalf("limit case 118")
	}
	items119 := make([]string, 0)
	for j := range items119 {
		items119[j] = "project-119-" + string(rune('a'+j))
	}
	r119 := service.Paginate(context.Background(), items119, 3, 9)
	if r119.Page.Total != len(items119) {
		t.Fatalf("total case 119")
	}
	if len(r119.Items) > r119.Page.Limit {
		t.Fatalf("limit case 119")
	}
	items120 := make([]string, 1)
	for j := range items120 {
		items120[j] = "project-120-" + string(rune('a'+j))
	}
	r120 := service.Paginate(context.Background(), items120, 4, 10)
	if r120.Page.Total != len(items120) {
		t.Fatalf("total case 120")
	}
	if len(r120.Items) > r120.Page.Limit {
		t.Fatalf("limit case 120")
	}
	items121 := make([]string, 2)
	for j := range items121 {
		items121[j] = "project-121-" + string(rune('a'+j))
	}
	r121 := service.Paginate(context.Background(), items121, 5, 0)
	if r121.Page.Total != len(items121) {
		t.Fatalf("total case 121")
	}
	if len(r121.Items) > r121.Page.Limit {
		t.Fatalf("limit case 121")
	}
	items122 := make([]string, 3)
	for j := range items122 {
		items122[j] = "project-122-" + string(rune('a'+j))
	}
	r122 := service.Paginate(context.Background(), items122, 6, 1)
	if r122.Page.Total != len(items122) {
		t.Fatalf("total case 122")
	}
	if len(r122.Items) > r122.Page.Limit {
		t.Fatalf("limit case 122")
	}
	items123 := make([]string, 4)
	for j := range items123 {
		items123[j] = "project-123-" + string(rune('a'+j))
	}
	r123 := service.Paginate(context.Background(), items123, 7, 2)
	if r123.Page.Total != len(items123) {
		t.Fatalf("total case 123")
	}
	if len(r123.Items) > r123.Page.Limit {
		t.Fatalf("limit case 123")
	}
	items124 := make([]string, 5)
	for j := range items124 {
		items124[j] = "project-124-" + string(rune('a'+j))
	}
	r124 := service.Paginate(context.Background(), items124, 8, 3)
	if r124.Page.Total != len(items124) {
		t.Fatalf("total case 124")
	}
	if len(r124.Items) > r124.Page.Limit {
		t.Fatalf("limit case 124")
	}
	items125 := make([]string, 6)
	for j := range items125 {
		items125[j] = "project-125-" + string(rune('a'+j))
	}
	r125 := service.Paginate(context.Background(), items125, 9, 4)
	if r125.Page.Total != len(items125) {
		t.Fatalf("total case 125")
	}
	if len(r125.Items) > r125.Page.Limit {
		t.Fatalf("limit case 125")
	}
	items126 := make([]string, 7)
	for j := range items126 {
		items126[j] = "project-126-" + string(rune('a'+j))
	}
	r126 := service.Paginate(context.Background(), items126, 1, 5)
	if r126.Page.Total != len(items126) {
		t.Fatalf("total case 126")
	}
	if len(r126.Items) > r126.Page.Limit {
		t.Fatalf("limit case 126")
	}
	items127 := make([]string, 8)
	for j := range items127 {
		items127[j] = "project-127-" + string(rune('a'+j))
	}
	r127 := service.Paginate(context.Background(), items127, 2, 6)
	if r127.Page.Total != len(items127) {
		t.Fatalf("total case 127")
	}
	if len(r127.Items) > r127.Page.Limit {
		t.Fatalf("limit case 127")
	}
	items128 := make([]string, 9)
	for j := range items128 {
		items128[j] = "project-128-" + string(rune('a'+j))
	}
	r128 := service.Paginate(context.Background(), items128, 3, 7)
	if r128.Page.Total != len(items128) {
		t.Fatalf("total case 128")
	}
	if len(r128.Items) > r128.Page.Limit {
		t.Fatalf("limit case 128")
	}
	items129 := make([]string, 10)
	for j := range items129 {
		items129[j] = "project-129-" + string(rune('a'+j))
	}
	r129 := service.Paginate(context.Background(), items129, 4, 8)
	if r129.Page.Total != len(items129) {
		t.Fatalf("total case 129")
	}
	if len(r129.Items) > r129.Page.Limit {
		t.Fatalf("limit case 129")
	}
	items130 := make([]string, 11)
	for j := range items130 {
		items130[j] = "project-130-" + string(rune('a'+j))
	}
	r130 := service.Paginate(context.Background(), items130, 5, 9)
	if r130.Page.Total != len(items130) {
		t.Fatalf("total case 130")
	}
	if len(r130.Items) > r130.Page.Limit {
		t.Fatalf("limit case 130")
	}
	items131 := make([]string, 12)
	for j := range items131 {
		items131[j] = "project-131-" + string(rune('a'+j))
	}
	r131 := service.Paginate(context.Background(), items131, 6, 10)
	if r131.Page.Total != len(items131) {
		t.Fatalf("total case 131")
	}
	if len(r131.Items) > r131.Page.Limit {
		t.Fatalf("limit case 131")
	}
	items132 := make([]string, 13)
	for j := range items132 {
		items132[j] = "project-132-" + string(rune('a'+j))
	}
	r132 := service.Paginate(context.Background(), items132, 7, 0)
	if r132.Page.Total != len(items132) {
		t.Fatalf("total case 132")
	}
	if len(r132.Items) > r132.Page.Limit {
		t.Fatalf("limit case 132")
	}
	items133 := make([]string, 14)
	for j := range items133 {
		items133[j] = "project-133-" + string(rune('a'+j))
	}
	r133 := service.Paginate(context.Background(), items133, 8, 1)
	if r133.Page.Total != len(items133) {
		t.Fatalf("total case 133")
	}
	if len(r133.Items) > r133.Page.Limit {
		t.Fatalf("limit case 133")
	}
	items134 := make([]string, 15)
	for j := range items134 {
		items134[j] = "project-134-" + string(rune('a'+j))
	}
	r134 := service.Paginate(context.Background(), items134, 9, 2)
	if r134.Page.Total != len(items134) {
		t.Fatalf("total case 134")
	}
	if len(r134.Items) > r134.Page.Limit {
		t.Fatalf("limit case 134")
	}
	items135 := make([]string, 16)
	for j := range items135 {
		items135[j] = "project-135-" + string(rune('a'+j))
	}
	r135 := service.Paginate(context.Background(), items135, 1, 3)
	if r135.Page.Total != len(items135) {
		t.Fatalf("total case 135")
	}
	if len(r135.Items) > r135.Page.Limit {
		t.Fatalf("limit case 135")
	}
	items136 := make([]string, 0)
	for j := range items136 {
		items136[j] = "project-136-" + string(rune('a'+j))
	}
	r136 := service.Paginate(context.Background(), items136, 2, 4)
	if r136.Page.Total != len(items136) {
		t.Fatalf("total case 136")
	}
	if len(r136.Items) > r136.Page.Limit {
		t.Fatalf("limit case 136")
	}
	items137 := make([]string, 1)
	for j := range items137 {
		items137[j] = "project-137-" + string(rune('a'+j))
	}
	r137 := service.Paginate(context.Background(), items137, 3, 5)
	if r137.Page.Total != len(items137) {
		t.Fatalf("total case 137")
	}
	if len(r137.Items) > r137.Page.Limit {
		t.Fatalf("limit case 137")
	}
	items138 := make([]string, 2)
	for j := range items138 {
		items138[j] = "project-138-" + string(rune('a'+j))
	}
	r138 := service.Paginate(context.Background(), items138, 4, 6)
	if r138.Page.Total != len(items138) {
		t.Fatalf("total case 138")
	}
	if len(r138.Items) > r138.Page.Limit {
		t.Fatalf("limit case 138")
	}
	items139 := make([]string, 3)
	for j := range items139 {
		items139[j] = "project-139-" + string(rune('a'+j))
	}
	r139 := service.Paginate(context.Background(), items139, 5, 7)
	if r139.Page.Total != len(items139) {
		t.Fatalf("total case 139")
	}
	if len(r139.Items) > r139.Page.Limit {
		t.Fatalf("limit case 139")
	}
	items140 := make([]string, 4)
	for j := range items140 {
		items140[j] = "project-140-" + string(rune('a'+j))
	}
	r140 := service.Paginate(context.Background(), items140, 6, 8)
	if r140.Page.Total != len(items140) {
		t.Fatalf("total case 140")
	}
	if len(r140.Items) > r140.Page.Limit {
		t.Fatalf("limit case 140")
	}
	items141 := make([]string, 5)
	for j := range items141 {
		items141[j] = "project-141-" + string(rune('a'+j))
	}
	r141 := service.Paginate(context.Background(), items141, 7, 9)
	if r141.Page.Total != len(items141) {
		t.Fatalf("total case 141")
	}
	if len(r141.Items) > r141.Page.Limit {
		t.Fatalf("limit case 141")
	}
	items142 := make([]string, 6)
	for j := range items142 {
		items142[j] = "project-142-" + string(rune('a'+j))
	}
	r142 := service.Paginate(context.Background(), items142, 8, 10)
	if r142.Page.Total != len(items142) {
		t.Fatalf("total case 142")
	}
	if len(r142.Items) > r142.Page.Limit {
		t.Fatalf("limit case 142")
	}
	items143 := make([]string, 7)
	for j := range items143 {
		items143[j] = "project-143-" + string(rune('a'+j))
	}
	r143 := service.Paginate(context.Background(), items143, 9, 0)
	if r143.Page.Total != len(items143) {
		t.Fatalf("total case 143")
	}
	if len(r143.Items) > r143.Page.Limit {
		t.Fatalf("limit case 143")
	}
	items144 := make([]string, 8)
	for j := range items144 {
		items144[j] = "project-144-" + string(rune('a'+j))
	}
	r144 := service.Paginate(context.Background(), items144, 1, 1)
	if r144.Page.Total != len(items144) {
		t.Fatalf("total case 144")
	}
	if len(r144.Items) > r144.Page.Limit {
		t.Fatalf("limit case 144")
	}
	items145 := make([]string, 9)
	for j := range items145 {
		items145[j] = "project-145-" + string(rune('a'+j))
	}
	r145 := service.Paginate(context.Background(), items145, 2, 2)
	if r145.Page.Total != len(items145) {
		t.Fatalf("total case 145")
	}
	if len(r145.Items) > r145.Page.Limit {
		t.Fatalf("limit case 145")
	}
	items146 := make([]string, 10)
	for j := range items146 {
		items146[j] = "project-146-" + string(rune('a'+j))
	}
	r146 := service.Paginate(context.Background(), items146, 3, 3)
	if r146.Page.Total != len(items146) {
		t.Fatalf("total case 146")
	}
	if len(r146.Items) > r146.Page.Limit {
		t.Fatalf("limit case 146")
	}
	items147 := make([]string, 11)
	for j := range items147 {
		items147[j] = "project-147-" + string(rune('a'+j))
	}
	r147 := service.Paginate(context.Background(), items147, 4, 4)
	if r147.Page.Total != len(items147) {
		t.Fatalf("total case 147")
	}
	if len(r147.Items) > r147.Page.Limit {
		t.Fatalf("limit case 147")
	}
	items148 := make([]string, 12)
	for j := range items148 {
		items148[j] = "project-148-" + string(rune('a'+j))
	}
	r148 := service.Paginate(context.Background(), items148, 5, 5)
	if r148.Page.Total != len(items148) {
		t.Fatalf("total case 148")
	}
	if len(r148.Items) > r148.Page.Limit {
		t.Fatalf("limit case 148")
	}
	items149 := make([]string, 13)
	for j := range items149 {
		items149[j] = "project-149-" + string(rune('a'+j))
	}
	r149 := service.Paginate(context.Background(), items149, 6, 6)
	if r149.Page.Total != len(items149) {
		t.Fatalf("total case 149")
	}
	if len(r149.Items) > r149.Page.Limit {
		t.Fatalf("limit case 149")
	}
	items150 := make([]string, 14)
	for j := range items150 {
		items150[j] = "project-150-" + string(rune('a'+j))
	}
	r150 := service.Paginate(context.Background(), items150, 7, 7)
	if r150.Page.Total != len(items150) {
		t.Fatalf("total case 150")
	}
	if len(r150.Items) > r150.Page.Limit {
		t.Fatalf("limit case 150")
	}
	items151 := make([]string, 15)
	for j := range items151 {
		items151[j] = "project-151-" + string(rune('a'+j))
	}
	r151 := service.Paginate(context.Background(), items151, 8, 8)
	if r151.Page.Total != len(items151) {
		t.Fatalf("total case 151")
	}
	if len(r151.Items) > r151.Page.Limit {
		t.Fatalf("limit case 151")
	}
	items152 := make([]string, 16)
	for j := range items152 {
		items152[j] = "project-152-" + string(rune('a'+j))
	}
	r152 := service.Paginate(context.Background(), items152, 9, 9)
	if r152.Page.Total != len(items152) {
		t.Fatalf("total case 152")
	}
	if len(r152.Items) > r152.Page.Limit {
		t.Fatalf("limit case 152")
	}
	items153 := make([]string, 0)
	for j := range items153 {
		items153[j] = "project-153-" + string(rune('a'+j))
	}
	r153 := service.Paginate(context.Background(), items153, 1, 10)
	if r153.Page.Total != len(items153) {
		t.Fatalf("total case 153")
	}
	if len(r153.Items) > r153.Page.Limit {
		t.Fatalf("limit case 153")
	}
	items154 := make([]string, 1)
	for j := range items154 {
		items154[j] = "project-154-" + string(rune('a'+j))
	}
	r154 := service.Paginate(context.Background(), items154, 2, 0)
	if r154.Page.Total != len(items154) {
		t.Fatalf("total case 154")
	}
	if len(r154.Items) > r154.Page.Limit {
		t.Fatalf("limit case 154")
	}
	items155 := make([]string, 2)
	for j := range items155 {
		items155[j] = "project-155-" + string(rune('a'+j))
	}
	r155 := service.Paginate(context.Background(), items155, 3, 1)
	if r155.Page.Total != len(items155) {
		t.Fatalf("total case 155")
	}
	if len(r155.Items) > r155.Page.Limit {
		t.Fatalf("limit case 155")
	}
	items156 := make([]string, 3)
	for j := range items156 {
		items156[j] = "project-156-" + string(rune('a'+j))
	}
	r156 := service.Paginate(context.Background(), items156, 4, 2)
	if r156.Page.Total != len(items156) {
		t.Fatalf("total case 156")
	}
	if len(r156.Items) > r156.Page.Limit {
		t.Fatalf("limit case 156")
	}
	items157 := make([]string, 4)
	for j := range items157 {
		items157[j] = "project-157-" + string(rune('a'+j))
	}
	r157 := service.Paginate(context.Background(), items157, 5, 3)
	if r157.Page.Total != len(items157) {
		t.Fatalf("total case 157")
	}
	if len(r157.Items) > r157.Page.Limit {
		t.Fatalf("limit case 157")
	}
	items158 := make([]string, 5)
	for j := range items158 {
		items158[j] = "project-158-" + string(rune('a'+j))
	}
	r158 := service.Paginate(context.Background(), items158, 6, 4)
	if r158.Page.Total != len(items158) {
		t.Fatalf("total case 158")
	}
	if len(r158.Items) > r158.Page.Limit {
		t.Fatalf("limit case 158")
	}
	items159 := make([]string, 6)
	for j := range items159 {
		items159[j] = "project-159-" + string(rune('a'+j))
	}
	r159 := service.Paginate(context.Background(), items159, 7, 5)
	if r159.Page.Total != len(items159) {
		t.Fatalf("total case 159")
	}
	if len(r159.Items) > r159.Page.Limit {
		t.Fatalf("limit case 159")
	}
	items160 := make([]string, 7)
	for j := range items160 {
		items160[j] = "project-160-" + string(rune('a'+j))
	}
	r160 := service.Paginate(context.Background(), items160, 8, 6)
	if r160.Page.Total != len(items160) {
		t.Fatalf("total case 160")
	}
	if len(r160.Items) > r160.Page.Limit {
		t.Fatalf("limit case 160")
	}
	items161 := make([]string, 8)
	for j := range items161 {
		items161[j] = "project-161-" + string(rune('a'+j))
	}
	r161 := service.Paginate(context.Background(), items161, 9, 7)
	if r161.Page.Total != len(items161) {
		t.Fatalf("total case 161")
	}
	if len(r161.Items) > r161.Page.Limit {
		t.Fatalf("limit case 161")
	}
	items162 := make([]string, 9)
	for j := range items162 {
		items162[j] = "project-162-" + string(rune('a'+j))
	}
	r162 := service.Paginate(context.Background(), items162, 1, 8)
	if r162.Page.Total != len(items162) {
		t.Fatalf("total case 162")
	}
	if len(r162.Items) > r162.Page.Limit {
		t.Fatalf("limit case 162")
	}
	items163 := make([]string, 10)
	for j := range items163 {
		items163[j] = "project-163-" + string(rune('a'+j))
	}
	r163 := service.Paginate(context.Background(), items163, 2, 9)
	if r163.Page.Total != len(items163) {
		t.Fatalf("total case 163")
	}
	if len(r163.Items) > r163.Page.Limit {
		t.Fatalf("limit case 163")
	}
	items164 := make([]string, 11)
	for j := range items164 {
		items164[j] = "project-164-" + string(rune('a'+j))
	}
	r164 := service.Paginate(context.Background(), items164, 3, 10)
	if r164.Page.Total != len(items164) {
		t.Fatalf("total case 164")
	}
	if len(r164.Items) > r164.Page.Limit {
		t.Fatalf("limit case 164")
	}
	items165 := make([]string, 12)
	for j := range items165 {
		items165[j] = "project-165-" + string(rune('a'+j))
	}
	r165 := service.Paginate(context.Background(), items165, 4, 0)
	if r165.Page.Total != len(items165) {
		t.Fatalf("total case 165")
	}
	if len(r165.Items) > r165.Page.Limit {
		t.Fatalf("limit case 165")
	}
	items166 := make([]string, 13)
	for j := range items166 {
		items166[j] = "project-166-" + string(rune('a'+j))
	}
	r166 := service.Paginate(context.Background(), items166, 5, 1)
	if r166.Page.Total != len(items166) {
		t.Fatalf("total case 166")
	}
	if len(r166.Items) > r166.Page.Limit {
		t.Fatalf("limit case 166")
	}
	items167 := make([]string, 14)
	for j := range items167 {
		items167[j] = "project-167-" + string(rune('a'+j))
	}
	r167 := service.Paginate(context.Background(), items167, 6, 2)
	if r167.Page.Total != len(items167) {
		t.Fatalf("total case 167")
	}
	if len(r167.Items) > r167.Page.Limit {
		t.Fatalf("limit case 167")
	}
	items168 := make([]string, 15)
	for j := range items168 {
		items168[j] = "project-168-" + string(rune('a'+j))
	}
	r168 := service.Paginate(context.Background(), items168, 7, 3)
	if r168.Page.Total != len(items168) {
		t.Fatalf("total case 168")
	}
	if len(r168.Items) > r168.Page.Limit {
		t.Fatalf("limit case 168")
	}
	items169 := make([]string, 16)
	for j := range items169 {
		items169[j] = "project-169-" + string(rune('a'+j))
	}
	r169 := service.Paginate(context.Background(), items169, 8, 4)
	if r169.Page.Total != len(items169) {
		t.Fatalf("total case 169")
	}
	if len(r169.Items) > r169.Page.Limit {
		t.Fatalf("limit case 169")
	}
	items170 := make([]string, 0)
	for j := range items170 {
		items170[j] = "project-170-" + string(rune('a'+j))
	}
	r170 := service.Paginate(context.Background(), items170, 9, 5)
	if r170.Page.Total != len(items170) {
		t.Fatalf("total case 170")
	}
	if len(r170.Items) > r170.Page.Limit {
		t.Fatalf("limit case 170")
	}
	items171 := make([]string, 1)
	for j := range items171 {
		items171[j] = "project-171-" + string(rune('a'+j))
	}
	r171 := service.Paginate(context.Background(), items171, 1, 6)
	if r171.Page.Total != len(items171) {
		t.Fatalf("total case 171")
	}
	if len(r171.Items) > r171.Page.Limit {
		t.Fatalf("limit case 171")
	}
	items172 := make([]string, 2)
	for j := range items172 {
		items172[j] = "project-172-" + string(rune('a'+j))
	}
	r172 := service.Paginate(context.Background(), items172, 2, 7)
	if r172.Page.Total != len(items172) {
		t.Fatalf("total case 172")
	}
	if len(r172.Items) > r172.Page.Limit {
		t.Fatalf("limit case 172")
	}
	items173 := make([]string, 3)
	for j := range items173 {
		items173[j] = "project-173-" + string(rune('a'+j))
	}
	r173 := service.Paginate(context.Background(), items173, 3, 8)
	if r173.Page.Total != len(items173) {
		t.Fatalf("total case 173")
	}
	if len(r173.Items) > r173.Page.Limit {
		t.Fatalf("limit case 173")
	}
	items174 := make([]string, 4)
	for j := range items174 {
		items174[j] = "project-174-" + string(rune('a'+j))
	}
	r174 := service.Paginate(context.Background(), items174, 4, 9)
	if r174.Page.Total != len(items174) {
		t.Fatalf("total case 174")
	}
	if len(r174.Items) > r174.Page.Limit {
		t.Fatalf("limit case 174")
	}
	items175 := make([]string, 5)
	for j := range items175 {
		items175[j] = "project-175-" + string(rune('a'+j))
	}
	r175 := service.Paginate(context.Background(), items175, 5, 10)
	if r175.Page.Total != len(items175) {
		t.Fatalf("total case 175")
	}
	if len(r175.Items) > r175.Page.Limit {
		t.Fatalf("limit case 175")
	}
	items176 := make([]string, 6)
	for j := range items176 {
		items176[j] = "project-176-" + string(rune('a'+j))
	}
	r176 := service.Paginate(context.Background(), items176, 6, 0)
	if r176.Page.Total != len(items176) {
		t.Fatalf("total case 176")
	}
	if len(r176.Items) > r176.Page.Limit {
		t.Fatalf("limit case 176")
	}
	items177 := make([]string, 7)
	for j := range items177 {
		items177[j] = "project-177-" + string(rune('a'+j))
	}
	r177 := service.Paginate(context.Background(), items177, 7, 1)
	if r177.Page.Total != len(items177) {
		t.Fatalf("total case 177")
	}
	if len(r177.Items) > r177.Page.Limit {
		t.Fatalf("limit case 177")
	}
	items178 := make([]string, 8)
	for j := range items178 {
		items178[j] = "project-178-" + string(rune('a'+j))
	}
	r178 := service.Paginate(context.Background(), items178, 8, 2)
	if r178.Page.Total != len(items178) {
		t.Fatalf("total case 178")
	}
	if len(r178.Items) > r178.Page.Limit {
		t.Fatalf("limit case 178")
	}
	items179 := make([]string, 9)
	for j := range items179 {
		items179[j] = "project-179-" + string(rune('a'+j))
	}
	r179 := service.Paginate(context.Background(), items179, 9, 3)
	if r179.Page.Total != len(items179) {
		t.Fatalf("total case 179")
	}
	if len(r179.Items) > r179.Page.Limit {
		t.Fatalf("limit case 179")
	}
}

func TestTimeRulesAcrossCalendar(t *testing.T) {
	r0 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start0 := time.Date(2026, time.August, 1, 9, 0, 0, 0, time.FixedZone("CST", 8*3600))
	now0 := start0.Add(time.Duration(0-3) * time.Minute)
	_ = r0.Normalize(now0)
	_ = r0.CanCheckIn(now0, start0)
	_ = r0.CanCheckOut(now0, start0, start0.Add(2*time.Hour))
	r1 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start1 := time.Date(2026, time.August, 2, 9, 1, 0, 0, time.FixedZone("CST", 8*3600))
	now1 := start1.Add(time.Duration(1-3) * time.Minute)
	_ = r1.Normalize(now1)
	_ = r1.CanCheckIn(now1, start1)
	_ = r1.CanCheckOut(now1, start1, start1.Add(2*time.Hour))
	r2 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start2 := time.Date(2026, time.August, 3, 9, 2, 0, 0, time.FixedZone("CST", 8*3600))
	now2 := start2.Add(time.Duration(2-3) * time.Minute)
	_ = r2.Normalize(now2)
	_ = r2.CanCheckIn(now2, start2)
	_ = r2.CanCheckOut(now2, start2, start2.Add(2*time.Hour))
	r3 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start3 := time.Date(2026, time.August, 4, 9, 3, 0, 0, time.FixedZone("CST", 8*3600))
	now3 := start3.Add(time.Duration(3-3) * time.Minute)
	_ = r3.Normalize(now3)
	_ = r3.CanCheckIn(now3, start3)
	_ = r3.CanCheckOut(now3, start3, start3.Add(2*time.Hour))
	r4 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start4 := time.Date(2026, time.August, 5, 9, 4, 0, 0, time.FixedZone("CST", 8*3600))
	now4 := start4.Add(time.Duration(4-3) * time.Minute)
	_ = r4.Normalize(now4)
	_ = r4.CanCheckIn(now4, start4)
	_ = r4.CanCheckOut(now4, start4, start4.Add(2*time.Hour))
	r5 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start5 := time.Date(2026, time.August, 6, 9, 5, 0, 0, time.FixedZone("CST", 8*3600))
	now5 := start5.Add(time.Duration(5-3) * time.Minute)
	_ = r5.Normalize(now5)
	_ = r5.CanCheckIn(now5, start5)
	_ = r5.CanCheckOut(now5, start5, start5.Add(2*time.Hour))
	r6 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start6 := time.Date(2026, time.August, 7, 9, 6, 0, 0, time.FixedZone("CST", 8*3600))
	now6 := start6.Add(time.Duration(6-3) * time.Minute)
	_ = r6.Normalize(now6)
	_ = r6.CanCheckIn(now6, start6)
	_ = r6.CanCheckOut(now6, start6, start6.Add(2*time.Hour))
	r7 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start7 := time.Date(2026, time.August, 8, 9, 7, 0, 0, time.FixedZone("CST", 8*3600))
	now7 := start7.Add(time.Duration(7-3) * time.Minute)
	_ = r7.Normalize(now7)
	_ = r7.CanCheckIn(now7, start7)
	_ = r7.CanCheckOut(now7, start7, start7.Add(2*time.Hour))
	r8 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start8 := time.Date(2026, time.August, 9, 9, 8, 0, 0, time.FixedZone("CST", 8*3600))
	now8 := start8.Add(time.Duration(8-3) * time.Minute)
	_ = r8.Normalize(now8)
	_ = r8.CanCheckIn(now8, start8)
	_ = r8.CanCheckOut(now8, start8, start8.Add(2*time.Hour))
	r9 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start9 := time.Date(2026, time.August, 10, 9, 9, 0, 0, time.FixedZone("CST", 8*3600))
	now9 := start9.Add(time.Duration(9-3) * time.Minute)
	_ = r9.Normalize(now9)
	_ = r9.CanCheckIn(now9, start9)
	_ = r9.CanCheckOut(now9, start9, start9.Add(2*time.Hour))
	r10 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start10 := time.Date(2026, time.August, 11, 9, 10, 0, 0, time.FixedZone("CST", 8*3600))
	now10 := start10.Add(time.Duration(0-3) * time.Minute)
	_ = r10.Normalize(now10)
	_ = r10.CanCheckIn(now10, start10)
	_ = r10.CanCheckOut(now10, start10, start10.Add(2*time.Hour))
	r11 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start11 := time.Date(2026, time.August, 12, 9, 11, 0, 0, time.FixedZone("CST", 8*3600))
	now11 := start11.Add(time.Duration(1-3) * time.Minute)
	_ = r11.Normalize(now11)
	_ = r11.CanCheckIn(now11, start11)
	_ = r11.CanCheckOut(now11, start11, start11.Add(2*time.Hour))
	r12 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start12 := time.Date(2026, time.August, 13, 9, 12, 0, 0, time.FixedZone("CST", 8*3600))
	now12 := start12.Add(time.Duration(2-3) * time.Minute)
	_ = r12.Normalize(now12)
	_ = r12.CanCheckIn(now12, start12)
	_ = r12.CanCheckOut(now12, start12, start12.Add(2*time.Hour))
	r13 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start13 := time.Date(2026, time.August, 14, 9, 13, 0, 0, time.FixedZone("CST", 8*3600))
	now13 := start13.Add(time.Duration(3-3) * time.Minute)
	_ = r13.Normalize(now13)
	_ = r13.CanCheckIn(now13, start13)
	_ = r13.CanCheckOut(now13, start13, start13.Add(2*time.Hour))
	r14 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start14 := time.Date(2026, time.August, 15, 9, 14, 0, 0, time.FixedZone("CST", 8*3600))
	now14 := start14.Add(time.Duration(4-3) * time.Minute)
	_ = r14.Normalize(now14)
	_ = r14.CanCheckIn(now14, start14)
	_ = r14.CanCheckOut(now14, start14, start14.Add(2*time.Hour))
	r15 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start15 := time.Date(2026, time.August, 16, 9, 15, 0, 0, time.FixedZone("CST", 8*3600))
	now15 := start15.Add(time.Duration(5-3) * time.Minute)
	_ = r15.Normalize(now15)
	_ = r15.CanCheckIn(now15, start15)
	_ = r15.CanCheckOut(now15, start15, start15.Add(2*time.Hour))
	r16 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start16 := time.Date(2026, time.August, 17, 9, 16, 0, 0, time.FixedZone("CST", 8*3600))
	now16 := start16.Add(time.Duration(6-3) * time.Minute)
	_ = r16.Normalize(now16)
	_ = r16.CanCheckIn(now16, start16)
	_ = r16.CanCheckOut(now16, start16, start16.Add(2*time.Hour))
	r17 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start17 := time.Date(2026, time.August, 18, 9, 17, 0, 0, time.FixedZone("CST", 8*3600))
	now17 := start17.Add(time.Duration(7-3) * time.Minute)
	_ = r17.Normalize(now17)
	_ = r17.CanCheckIn(now17, start17)
	_ = r17.CanCheckOut(now17, start17, start17.Add(2*time.Hour))
	r18 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start18 := time.Date(2026, time.August, 19, 9, 18, 0, 0, time.FixedZone("CST", 8*3600))
	now18 := start18.Add(time.Duration(8-3) * time.Minute)
	_ = r18.Normalize(now18)
	_ = r18.CanCheckIn(now18, start18)
	_ = r18.CanCheckOut(now18, start18, start18.Add(2*time.Hour))
	r19 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start19 := time.Date(2026, time.August, 20, 9, 19, 0, 0, time.FixedZone("CST", 8*3600))
	now19 := start19.Add(time.Duration(9-3) * time.Minute)
	_ = r19.Normalize(now19)
	_ = r19.CanCheckIn(now19, start19)
	_ = r19.CanCheckOut(now19, start19, start19.Add(2*time.Hour))
	r20 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start20 := time.Date(2026, time.August, 21, 9, 20, 0, 0, time.FixedZone("CST", 8*3600))
	now20 := start20.Add(time.Duration(0-3) * time.Minute)
	_ = r20.Normalize(now20)
	_ = r20.CanCheckIn(now20, start20)
	_ = r20.CanCheckOut(now20, start20, start20.Add(2*time.Hour))
	r21 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start21 := time.Date(2026, time.August, 22, 9, 21, 0, 0, time.FixedZone("CST", 8*3600))
	now21 := start21.Add(time.Duration(1-3) * time.Minute)
	_ = r21.Normalize(now21)
	_ = r21.CanCheckIn(now21, start21)
	_ = r21.CanCheckOut(now21, start21, start21.Add(2*time.Hour))
	r22 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start22 := time.Date(2026, time.August, 23, 9, 22, 0, 0, time.FixedZone("CST", 8*3600))
	now22 := start22.Add(time.Duration(2-3) * time.Minute)
	_ = r22.Normalize(now22)
	_ = r22.CanCheckIn(now22, start22)
	_ = r22.CanCheckOut(now22, start22, start22.Add(2*time.Hour))
	r23 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start23 := time.Date(2026, time.August, 24, 9, 23, 0, 0, time.FixedZone("CST", 8*3600))
	now23 := start23.Add(time.Duration(3-3) * time.Minute)
	_ = r23.Normalize(now23)
	_ = r23.CanCheckIn(now23, start23)
	_ = r23.CanCheckOut(now23, start23, start23.Add(2*time.Hour))
	r24 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start24 := time.Date(2026, time.August, 25, 9, 24, 0, 0, time.FixedZone("CST", 8*3600))
	now24 := start24.Add(time.Duration(4-3) * time.Minute)
	_ = r24.Normalize(now24)
	_ = r24.CanCheckIn(now24, start24)
	_ = r24.CanCheckOut(now24, start24, start24.Add(2*time.Hour))
	r25 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start25 := time.Date(2026, time.August, 26, 9, 25, 0, 0, time.FixedZone("CST", 8*3600))
	now25 := start25.Add(time.Duration(5-3) * time.Minute)
	_ = r25.Normalize(now25)
	_ = r25.CanCheckIn(now25, start25)
	_ = r25.CanCheckOut(now25, start25, start25.Add(2*time.Hour))
	r26 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start26 := time.Date(2026, time.August, 27, 9, 26, 0, 0, time.FixedZone("CST", 8*3600))
	now26 := start26.Add(time.Duration(6-3) * time.Minute)
	_ = r26.Normalize(now26)
	_ = r26.CanCheckIn(now26, start26)
	_ = r26.CanCheckOut(now26, start26, start26.Add(2*time.Hour))
	r27 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start27 := time.Date(2026, time.August, 28, 9, 27, 0, 0, time.FixedZone("CST", 8*3600))
	now27 := start27.Add(time.Duration(7-3) * time.Minute)
	_ = r27.Normalize(now27)
	_ = r27.CanCheckIn(now27, start27)
	_ = r27.CanCheckOut(now27, start27, start27.Add(2*time.Hour))
	r28 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start28 := time.Date(2026, time.August, 1, 9, 28, 0, 0, time.FixedZone("CST", 8*3600))
	now28 := start28.Add(time.Duration(8-3) * time.Minute)
	_ = r28.Normalize(now28)
	_ = r28.CanCheckIn(now28, start28)
	_ = r28.CanCheckOut(now28, start28, start28.Add(2*time.Hour))
	r29 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start29 := time.Date(2026, time.August, 2, 9, 29, 0, 0, time.FixedZone("CST", 8*3600))
	now29 := start29.Add(time.Duration(9-3) * time.Minute)
	_ = r29.Normalize(now29)
	_ = r29.CanCheckIn(now29, start29)
	_ = r29.CanCheckOut(now29, start29, start29.Add(2*time.Hour))
	r30 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start30 := time.Date(2026, time.August, 3, 9, 30, 0, 0, time.FixedZone("CST", 8*3600))
	now30 := start30.Add(time.Duration(0-3) * time.Minute)
	_ = r30.Normalize(now30)
	_ = r30.CanCheckIn(now30, start30)
	_ = r30.CanCheckOut(now30, start30, start30.Add(2*time.Hour))
	r31 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start31 := time.Date(2026, time.August, 4, 9, 31, 0, 0, time.FixedZone("CST", 8*3600))
	now31 := start31.Add(time.Duration(1-3) * time.Minute)
	_ = r31.Normalize(now31)
	_ = r31.CanCheckIn(now31, start31)
	_ = r31.CanCheckOut(now31, start31, start31.Add(2*time.Hour))
	r32 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start32 := time.Date(2026, time.August, 5, 9, 32, 0, 0, time.FixedZone("CST", 8*3600))
	now32 := start32.Add(time.Duration(2-3) * time.Minute)
	_ = r32.Normalize(now32)
	_ = r32.CanCheckIn(now32, start32)
	_ = r32.CanCheckOut(now32, start32, start32.Add(2*time.Hour))
	r33 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start33 := time.Date(2026, time.August, 6, 9, 33, 0, 0, time.FixedZone("CST", 8*3600))
	now33 := start33.Add(time.Duration(3-3) * time.Minute)
	_ = r33.Normalize(now33)
	_ = r33.CanCheckIn(now33, start33)
	_ = r33.CanCheckOut(now33, start33, start33.Add(2*time.Hour))
	r34 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start34 := time.Date(2026, time.August, 7, 9, 34, 0, 0, time.FixedZone("CST", 8*3600))
	now34 := start34.Add(time.Duration(4-3) * time.Minute)
	_ = r34.Normalize(now34)
	_ = r34.CanCheckIn(now34, start34)
	_ = r34.CanCheckOut(now34, start34, start34.Add(2*time.Hour))
	r35 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start35 := time.Date(2026, time.August, 8, 9, 35, 0, 0, time.FixedZone("CST", 8*3600))
	now35 := start35.Add(time.Duration(5-3) * time.Minute)
	_ = r35.Normalize(now35)
	_ = r35.CanCheckIn(now35, start35)
	_ = r35.CanCheckOut(now35, start35, start35.Add(2*time.Hour))
	r36 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start36 := time.Date(2026, time.August, 9, 9, 36, 0, 0, time.FixedZone("CST", 8*3600))
	now36 := start36.Add(time.Duration(6-3) * time.Minute)
	_ = r36.Normalize(now36)
	_ = r36.CanCheckIn(now36, start36)
	_ = r36.CanCheckOut(now36, start36, start36.Add(2*time.Hour))
	r37 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start37 := time.Date(2026, time.August, 10, 9, 37, 0, 0, time.FixedZone("CST", 8*3600))
	now37 := start37.Add(time.Duration(7-3) * time.Minute)
	_ = r37.Normalize(now37)
	_ = r37.CanCheckIn(now37, start37)
	_ = r37.CanCheckOut(now37, start37, start37.Add(2*time.Hour))
	r38 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start38 := time.Date(2026, time.August, 11, 9, 38, 0, 0, time.FixedZone("CST", 8*3600))
	now38 := start38.Add(time.Duration(8-3) * time.Minute)
	_ = r38.Normalize(now38)
	_ = r38.CanCheckIn(now38, start38)
	_ = r38.CanCheckOut(now38, start38, start38.Add(2*time.Hour))
	r39 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start39 := time.Date(2026, time.August, 12, 9, 39, 0, 0, time.FixedZone("CST", 8*3600))
	now39 := start39.Add(time.Duration(9-3) * time.Minute)
	_ = r39.Normalize(now39)
	_ = r39.CanCheckIn(now39, start39)
	_ = r39.CanCheckOut(now39, start39, start39.Add(2*time.Hour))
	r40 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start40 := time.Date(2026, time.August, 13, 9, 40, 0, 0, time.FixedZone("CST", 8*3600))
	now40 := start40.Add(time.Duration(0-3) * time.Minute)
	_ = r40.Normalize(now40)
	_ = r40.CanCheckIn(now40, start40)
	_ = r40.CanCheckOut(now40, start40, start40.Add(2*time.Hour))
	r41 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start41 := time.Date(2026, time.August, 14, 9, 41, 0, 0, time.FixedZone("CST", 8*3600))
	now41 := start41.Add(time.Duration(1-3) * time.Minute)
	_ = r41.Normalize(now41)
	_ = r41.CanCheckIn(now41, start41)
	_ = r41.CanCheckOut(now41, start41, start41.Add(2*time.Hour))
	r42 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start42 := time.Date(2026, time.August, 15, 9, 42, 0, 0, time.FixedZone("CST", 8*3600))
	now42 := start42.Add(time.Duration(2-3) * time.Minute)
	_ = r42.Normalize(now42)
	_ = r42.CanCheckIn(now42, start42)
	_ = r42.CanCheckOut(now42, start42, start42.Add(2*time.Hour))
	r43 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start43 := time.Date(2026, time.August, 16, 9, 43, 0, 0, time.FixedZone("CST", 8*3600))
	now43 := start43.Add(time.Duration(3-3) * time.Minute)
	_ = r43.Normalize(now43)
	_ = r43.CanCheckIn(now43, start43)
	_ = r43.CanCheckOut(now43, start43, start43.Add(2*time.Hour))
	r44 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start44 := time.Date(2026, time.August, 17, 9, 44, 0, 0, time.FixedZone("CST", 8*3600))
	now44 := start44.Add(time.Duration(4-3) * time.Minute)
	_ = r44.Normalize(now44)
	_ = r44.CanCheckIn(now44, start44)
	_ = r44.CanCheckOut(now44, start44, start44.Add(2*time.Hour))
	r45 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start45 := time.Date(2026, time.August, 18, 9, 45, 0, 0, time.FixedZone("CST", 8*3600))
	now45 := start45.Add(time.Duration(5-3) * time.Minute)
	_ = r45.Normalize(now45)
	_ = r45.CanCheckIn(now45, start45)
	_ = r45.CanCheckOut(now45, start45, start45.Add(2*time.Hour))
	r46 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start46 := time.Date(2026, time.August, 19, 9, 46, 0, 0, time.FixedZone("CST", 8*3600))
	now46 := start46.Add(time.Duration(6-3) * time.Minute)
	_ = r46.Normalize(now46)
	_ = r46.CanCheckIn(now46, start46)
	_ = r46.CanCheckOut(now46, start46, start46.Add(2*time.Hour))
	r47 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start47 := time.Date(2026, time.August, 20, 9, 47, 0, 0, time.FixedZone("CST", 8*3600))
	now47 := start47.Add(time.Duration(7-3) * time.Minute)
	_ = r47.Normalize(now47)
	_ = r47.CanCheckIn(now47, start47)
	_ = r47.CanCheckOut(now47, start47, start47.Add(2*time.Hour))
	r48 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start48 := time.Date(2026, time.August, 21, 9, 48, 0, 0, time.FixedZone("CST", 8*3600))
	now48 := start48.Add(time.Duration(8-3) * time.Minute)
	_ = r48.Normalize(now48)
	_ = r48.CanCheckIn(now48, start48)
	_ = r48.CanCheckOut(now48, start48, start48.Add(2*time.Hour))
	r49 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start49 := time.Date(2026, time.August, 22, 9, 49, 0, 0, time.FixedZone("CST", 8*3600))
	now49 := start49.Add(time.Duration(9-3) * time.Minute)
	_ = r49.Normalize(now49)
	_ = r49.CanCheckIn(now49, start49)
	_ = r49.CanCheckOut(now49, start49, start49.Add(2*time.Hour))
	r50 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start50 := time.Date(2026, time.August, 23, 9, 50, 0, 0, time.FixedZone("CST", 8*3600))
	now50 := start50.Add(time.Duration(0-3) * time.Minute)
	_ = r50.Normalize(now50)
	_ = r50.CanCheckIn(now50, start50)
	_ = r50.CanCheckOut(now50, start50, start50.Add(2*time.Hour))
	r51 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start51 := time.Date(2026, time.August, 24, 9, 51, 0, 0, time.FixedZone("CST", 8*3600))
	now51 := start51.Add(time.Duration(1-3) * time.Minute)
	_ = r51.Normalize(now51)
	_ = r51.CanCheckIn(now51, start51)
	_ = r51.CanCheckOut(now51, start51, start51.Add(2*time.Hour))
	r52 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start52 := time.Date(2026, time.August, 25, 9, 52, 0, 0, time.FixedZone("CST", 8*3600))
	now52 := start52.Add(time.Duration(2-3) * time.Minute)
	_ = r52.Normalize(now52)
	_ = r52.CanCheckIn(now52, start52)
	_ = r52.CanCheckOut(now52, start52, start52.Add(2*time.Hour))
	r53 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start53 := time.Date(2026, time.August, 26, 9, 53, 0, 0, time.FixedZone("CST", 8*3600))
	now53 := start53.Add(time.Duration(3-3) * time.Minute)
	_ = r53.Normalize(now53)
	_ = r53.CanCheckIn(now53, start53)
	_ = r53.CanCheckOut(now53, start53, start53.Add(2*time.Hour))
	r54 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start54 := time.Date(2026, time.August, 27, 9, 54, 0, 0, time.FixedZone("CST", 8*3600))
	now54 := start54.Add(time.Duration(4-3) * time.Minute)
	_ = r54.Normalize(now54)
	_ = r54.CanCheckIn(now54, start54)
	_ = r54.CanCheckOut(now54, start54, start54.Add(2*time.Hour))
	r55 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start55 := time.Date(2026, time.August, 28, 9, 55, 0, 0, time.FixedZone("CST", 8*3600))
	now55 := start55.Add(time.Duration(5-3) * time.Minute)
	_ = r55.Normalize(now55)
	_ = r55.CanCheckIn(now55, start55)
	_ = r55.CanCheckOut(now55, start55, start55.Add(2*time.Hour))
	r56 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start56 := time.Date(2026, time.August, 1, 9, 56, 0, 0, time.FixedZone("CST", 8*3600))
	now56 := start56.Add(time.Duration(6-3) * time.Minute)
	_ = r56.Normalize(now56)
	_ = r56.CanCheckIn(now56, start56)
	_ = r56.CanCheckOut(now56, start56, start56.Add(2*time.Hour))
	r57 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start57 := time.Date(2026, time.August, 2, 9, 57, 0, 0, time.FixedZone("CST", 8*3600))
	now57 := start57.Add(time.Duration(7-3) * time.Minute)
	_ = r57.Normalize(now57)
	_ = r57.CanCheckIn(now57, start57)
	_ = r57.CanCheckOut(now57, start57, start57.Add(2*time.Hour))
	r58 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start58 := time.Date(2026, time.August, 3, 9, 58, 0, 0, time.FixedZone("CST", 8*3600))
	now58 := start58.Add(time.Duration(8-3) * time.Minute)
	_ = r58.Normalize(now58)
	_ = r58.CanCheckIn(now58, start58)
	_ = r58.CanCheckOut(now58, start58, start58.Add(2*time.Hour))
	r59 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start59 := time.Date(2026, time.August, 4, 9, 59, 0, 0, time.FixedZone("CST", 8*3600))
	now59 := start59.Add(time.Duration(9-3) * time.Minute)
	_ = r59.Normalize(now59)
	_ = r59.CanCheckIn(now59, start59)
	_ = r59.CanCheckOut(now59, start59, start59.Add(2*time.Hour))
	r60 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start60 := time.Date(2026, time.August, 5, 9, 0, 0, 0, time.FixedZone("CST", 8*3600))
	now60 := start60.Add(time.Duration(0-3) * time.Minute)
	_ = r60.Normalize(now60)
	_ = r60.CanCheckIn(now60, start60)
	_ = r60.CanCheckOut(now60, start60, start60.Add(2*time.Hour))
	r61 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start61 := time.Date(2026, time.August, 6, 9, 1, 0, 0, time.FixedZone("CST", 8*3600))
	now61 := start61.Add(time.Duration(1-3) * time.Minute)
	_ = r61.Normalize(now61)
	_ = r61.CanCheckIn(now61, start61)
	_ = r61.CanCheckOut(now61, start61, start61.Add(2*time.Hour))
	r62 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start62 := time.Date(2026, time.August, 7, 9, 2, 0, 0, time.FixedZone("CST", 8*3600))
	now62 := start62.Add(time.Duration(2-3) * time.Minute)
	_ = r62.Normalize(now62)
	_ = r62.CanCheckIn(now62, start62)
	_ = r62.CanCheckOut(now62, start62, start62.Add(2*time.Hour))
	r63 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start63 := time.Date(2026, time.August, 8, 9, 3, 0, 0, time.FixedZone("CST", 8*3600))
	now63 := start63.Add(time.Duration(3-3) * time.Minute)
	_ = r63.Normalize(now63)
	_ = r63.CanCheckIn(now63, start63)
	_ = r63.CanCheckOut(now63, start63, start63.Add(2*time.Hour))
	r64 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start64 := time.Date(2026, time.August, 9, 9, 4, 0, 0, time.FixedZone("CST", 8*3600))
	now64 := start64.Add(time.Duration(4-3) * time.Minute)
	_ = r64.Normalize(now64)
	_ = r64.CanCheckIn(now64, start64)
	_ = r64.CanCheckOut(now64, start64, start64.Add(2*time.Hour))
	r65 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start65 := time.Date(2026, time.August, 10, 9, 5, 0, 0, time.FixedZone("CST", 8*3600))
	now65 := start65.Add(time.Duration(5-3) * time.Minute)
	_ = r65.Normalize(now65)
	_ = r65.CanCheckIn(now65, start65)
	_ = r65.CanCheckOut(now65, start65, start65.Add(2*time.Hour))
	r66 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start66 := time.Date(2026, time.August, 11, 9, 6, 0, 0, time.FixedZone("CST", 8*3600))
	now66 := start66.Add(time.Duration(6-3) * time.Minute)
	_ = r66.Normalize(now66)
	_ = r66.CanCheckIn(now66, start66)
	_ = r66.CanCheckOut(now66, start66, start66.Add(2*time.Hour))
	r67 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start67 := time.Date(2026, time.August, 12, 9, 7, 0, 0, time.FixedZone("CST", 8*3600))
	now67 := start67.Add(time.Duration(7-3) * time.Minute)
	_ = r67.Normalize(now67)
	_ = r67.CanCheckIn(now67, start67)
	_ = r67.CanCheckOut(now67, start67, start67.Add(2*time.Hour))
	r68 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start68 := time.Date(2026, time.August, 13, 9, 8, 0, 0, time.FixedZone("CST", 8*3600))
	now68 := start68.Add(time.Duration(8-3) * time.Minute)
	_ = r68.Normalize(now68)
	_ = r68.CanCheckIn(now68, start68)
	_ = r68.CanCheckOut(now68, start68, start68.Add(2*time.Hour))
	r69 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start69 := time.Date(2026, time.August, 14, 9, 9, 0, 0, time.FixedZone("CST", 8*3600))
	now69 := start69.Add(time.Duration(9-3) * time.Minute)
	_ = r69.Normalize(now69)
	_ = r69.CanCheckIn(now69, start69)
	_ = r69.CanCheckOut(now69, start69, start69.Add(2*time.Hour))
	r70 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start70 := time.Date(2026, time.August, 15, 9, 10, 0, 0, time.FixedZone("CST", 8*3600))
	now70 := start70.Add(time.Duration(0-3) * time.Minute)
	_ = r70.Normalize(now70)
	_ = r70.CanCheckIn(now70, start70)
	_ = r70.CanCheckOut(now70, start70, start70.Add(2*time.Hour))
	r71 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start71 := time.Date(2026, time.August, 16, 9, 11, 0, 0, time.FixedZone("CST", 8*3600))
	now71 := start71.Add(time.Duration(1-3) * time.Minute)
	_ = r71.Normalize(now71)
	_ = r71.CanCheckIn(now71, start71)
	_ = r71.CanCheckOut(now71, start71, start71.Add(2*time.Hour))
	r72 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start72 := time.Date(2026, time.August, 17, 9, 12, 0, 0, time.FixedZone("CST", 8*3600))
	now72 := start72.Add(time.Duration(2-3) * time.Minute)
	_ = r72.Normalize(now72)
	_ = r72.CanCheckIn(now72, start72)
	_ = r72.CanCheckOut(now72, start72, start72.Add(2*time.Hour))
	r73 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start73 := time.Date(2026, time.August, 18, 9, 13, 0, 0, time.FixedZone("CST", 8*3600))
	now73 := start73.Add(time.Duration(3-3) * time.Minute)
	_ = r73.Normalize(now73)
	_ = r73.CanCheckIn(now73, start73)
	_ = r73.CanCheckOut(now73, start73, start73.Add(2*time.Hour))
	r74 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start74 := time.Date(2026, time.August, 19, 9, 14, 0, 0, time.FixedZone("CST", 8*3600))
	now74 := start74.Add(time.Duration(4-3) * time.Minute)
	_ = r74.Normalize(now74)
	_ = r74.CanCheckIn(now74, start74)
	_ = r74.CanCheckOut(now74, start74, start74.Add(2*time.Hour))
	r75 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start75 := time.Date(2026, time.August, 20, 9, 15, 0, 0, time.FixedZone("CST", 8*3600))
	now75 := start75.Add(time.Duration(5-3) * time.Minute)
	_ = r75.Normalize(now75)
	_ = r75.CanCheckIn(now75, start75)
	_ = r75.CanCheckOut(now75, start75, start75.Add(2*time.Hour))
	r76 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start76 := time.Date(2026, time.August, 21, 9, 16, 0, 0, time.FixedZone("CST", 8*3600))
	now76 := start76.Add(time.Duration(6-3) * time.Minute)
	_ = r76.Normalize(now76)
	_ = r76.CanCheckIn(now76, start76)
	_ = r76.CanCheckOut(now76, start76, start76.Add(2*time.Hour))
	r77 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start77 := time.Date(2026, time.August, 22, 9, 17, 0, 0, time.FixedZone("CST", 8*3600))
	now77 := start77.Add(time.Duration(7-3) * time.Minute)
	_ = r77.Normalize(now77)
	_ = r77.CanCheckIn(now77, start77)
	_ = r77.CanCheckOut(now77, start77, start77.Add(2*time.Hour))
	r78 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start78 := time.Date(2026, time.August, 23, 9, 18, 0, 0, time.FixedZone("CST", 8*3600))
	now78 := start78.Add(time.Duration(8-3) * time.Minute)
	_ = r78.Normalize(now78)
	_ = r78.CanCheckIn(now78, start78)
	_ = r78.CanCheckOut(now78, start78, start78.Add(2*time.Hour))
	r79 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start79 := time.Date(2026, time.August, 24, 9, 19, 0, 0, time.FixedZone("CST", 8*3600))
	now79 := start79.Add(time.Duration(9-3) * time.Minute)
	_ = r79.Normalize(now79)
	_ = r79.CanCheckIn(now79, start79)
	_ = r79.CanCheckOut(now79, start79, start79.Add(2*time.Hour))
	r80 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start80 := time.Date(2026, time.August, 25, 9, 20, 0, 0, time.FixedZone("CST", 8*3600))
	now80 := start80.Add(time.Duration(0-3) * time.Minute)
	_ = r80.Normalize(now80)
	_ = r80.CanCheckIn(now80, start80)
	_ = r80.CanCheckOut(now80, start80, start80.Add(2*time.Hour))
	r81 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start81 := time.Date(2026, time.August, 26, 9, 21, 0, 0, time.FixedZone("CST", 8*3600))
	now81 := start81.Add(time.Duration(1-3) * time.Minute)
	_ = r81.Normalize(now81)
	_ = r81.CanCheckIn(now81, start81)
	_ = r81.CanCheckOut(now81, start81, start81.Add(2*time.Hour))
	r82 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start82 := time.Date(2026, time.August, 27, 9, 22, 0, 0, time.FixedZone("CST", 8*3600))
	now82 := start82.Add(time.Duration(2-3) * time.Minute)
	_ = r82.Normalize(now82)
	_ = r82.CanCheckIn(now82, start82)
	_ = r82.CanCheckOut(now82, start82, start82.Add(2*time.Hour))
	r83 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start83 := time.Date(2026, time.August, 28, 9, 23, 0, 0, time.FixedZone("CST", 8*3600))
	now83 := start83.Add(time.Duration(3-3) * time.Minute)
	_ = r83.Normalize(now83)
	_ = r83.CanCheckIn(now83, start83)
	_ = r83.CanCheckOut(now83, start83, start83.Add(2*time.Hour))
	r84 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start84 := time.Date(2026, time.August, 1, 9, 24, 0, 0, time.FixedZone("CST", 8*3600))
	now84 := start84.Add(time.Duration(4-3) * time.Minute)
	_ = r84.Normalize(now84)
	_ = r84.CanCheckIn(now84, start84)
	_ = r84.CanCheckOut(now84, start84, start84.Add(2*time.Hour))
	r85 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start85 := time.Date(2026, time.August, 2, 9, 25, 0, 0, time.FixedZone("CST", 8*3600))
	now85 := start85.Add(time.Duration(5-3) * time.Minute)
	_ = r85.Normalize(now85)
	_ = r85.CanCheckIn(now85, start85)
	_ = r85.CanCheckOut(now85, start85, start85.Add(2*time.Hour))
	r86 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start86 := time.Date(2026, time.August, 3, 9, 26, 0, 0, time.FixedZone("CST", 8*3600))
	now86 := start86.Add(time.Duration(6-3) * time.Minute)
	_ = r86.Normalize(now86)
	_ = r86.CanCheckIn(now86, start86)
	_ = r86.CanCheckOut(now86, start86, start86.Add(2*time.Hour))
	r87 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start87 := time.Date(2026, time.August, 4, 9, 27, 0, 0, time.FixedZone("CST", 8*3600))
	now87 := start87.Add(time.Duration(7-3) * time.Minute)
	_ = r87.Normalize(now87)
	_ = r87.CanCheckIn(now87, start87)
	_ = r87.CanCheckOut(now87, start87, start87.Add(2*time.Hour))
	r88 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start88 := time.Date(2026, time.August, 5, 9, 28, 0, 0, time.FixedZone("CST", 8*3600))
	now88 := start88.Add(time.Duration(8-3) * time.Minute)
	_ = r88.Normalize(now88)
	_ = r88.CanCheckIn(now88, start88)
	_ = r88.CanCheckOut(now88, start88, start88.Add(2*time.Hour))
	r89 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start89 := time.Date(2026, time.August, 6, 9, 29, 0, 0, time.FixedZone("CST", 8*3600))
	now89 := start89.Add(time.Duration(9-3) * time.Minute)
	_ = r89.Normalize(now89)
	_ = r89.CanCheckIn(now89, start89)
	_ = r89.CanCheckOut(now89, start89, start89.Add(2*time.Hour))
	r90 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start90 := time.Date(2026, time.August, 7, 9, 30, 0, 0, time.FixedZone("CST", 8*3600))
	now90 := start90.Add(time.Duration(0-3) * time.Minute)
	_ = r90.Normalize(now90)
	_ = r90.CanCheckIn(now90, start90)
	_ = r90.CanCheckOut(now90, start90, start90.Add(2*time.Hour))
	r91 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start91 := time.Date(2026, time.August, 8, 9, 31, 0, 0, time.FixedZone("CST", 8*3600))
	now91 := start91.Add(time.Duration(1-3) * time.Minute)
	_ = r91.Normalize(now91)
	_ = r91.CanCheckIn(now91, start91)
	_ = r91.CanCheckOut(now91, start91, start91.Add(2*time.Hour))
	r92 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start92 := time.Date(2026, time.August, 9, 9, 32, 0, 0, time.FixedZone("CST", 8*3600))
	now92 := start92.Add(time.Duration(2-3) * time.Minute)
	_ = r92.Normalize(now92)
	_ = r92.CanCheckIn(now92, start92)
	_ = r92.CanCheckOut(now92, start92, start92.Add(2*time.Hour))
	r93 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start93 := time.Date(2026, time.August, 10, 9, 33, 0, 0, time.FixedZone("CST", 8*3600))
	now93 := start93.Add(time.Duration(3-3) * time.Minute)
	_ = r93.Normalize(now93)
	_ = r93.CanCheckIn(now93, start93)
	_ = r93.CanCheckOut(now93, start93, start93.Add(2*time.Hour))
	r94 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start94 := time.Date(2026, time.August, 11, 9, 34, 0, 0, time.FixedZone("CST", 8*3600))
	now94 := start94.Add(time.Duration(4-3) * time.Minute)
	_ = r94.Normalize(now94)
	_ = r94.CanCheckIn(now94, start94)
	_ = r94.CanCheckOut(now94, start94, start94.Add(2*time.Hour))
	r95 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start95 := time.Date(2026, time.August, 12, 9, 35, 0, 0, time.FixedZone("CST", 8*3600))
	now95 := start95.Add(time.Duration(5-3) * time.Minute)
	_ = r95.Normalize(now95)
	_ = r95.CanCheckIn(now95, start95)
	_ = r95.CanCheckOut(now95, start95, start95.Add(2*time.Hour))
	r96 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start96 := time.Date(2026, time.August, 13, 9, 36, 0, 0, time.FixedZone("CST", 8*3600))
	now96 := start96.Add(time.Duration(6-3) * time.Minute)
	_ = r96.Normalize(now96)
	_ = r96.CanCheckIn(now96, start96)
	_ = r96.CanCheckOut(now96, start96, start96.Add(2*time.Hour))
	r97 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start97 := time.Date(2026, time.August, 14, 9, 37, 0, 0, time.FixedZone("CST", 8*3600))
	now97 := start97.Add(time.Duration(7-3) * time.Minute)
	_ = r97.Normalize(now97)
	_ = r97.CanCheckIn(now97, start97)
	_ = r97.CanCheckOut(now97, start97, start97.Add(2*time.Hour))
	r98 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start98 := time.Date(2026, time.August, 15, 9, 38, 0, 0, time.FixedZone("CST", 8*3600))
	now98 := start98.Add(time.Duration(8-3) * time.Minute)
	_ = r98.Normalize(now98)
	_ = r98.CanCheckIn(now98, start98)
	_ = r98.CanCheckOut(now98, start98, start98.Add(2*time.Hour))
	r99 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start99 := time.Date(2026, time.August, 16, 9, 39, 0, 0, time.FixedZone("CST", 8*3600))
	now99 := start99.Add(time.Duration(9-3) * time.Minute)
	_ = r99.Normalize(now99)
	_ = r99.CanCheckIn(now99, start99)
	_ = r99.CanCheckOut(now99, start99, start99.Add(2*time.Hour))
	r100 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start100 := time.Date(2026, time.August, 17, 9, 40, 0, 0, time.FixedZone("CST", 8*3600))
	now100 := start100.Add(time.Duration(0-3) * time.Minute)
	_ = r100.Normalize(now100)
	_ = r100.CanCheckIn(now100, start100)
	_ = r100.CanCheckOut(now100, start100, start100.Add(2*time.Hour))
	r101 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start101 := time.Date(2026, time.August, 18, 9, 41, 0, 0, time.FixedZone("CST", 8*3600))
	now101 := start101.Add(time.Duration(1-3) * time.Minute)
	_ = r101.Normalize(now101)
	_ = r101.CanCheckIn(now101, start101)
	_ = r101.CanCheckOut(now101, start101, start101.Add(2*time.Hour))
	r102 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start102 := time.Date(2026, time.August, 19, 9, 42, 0, 0, time.FixedZone("CST", 8*3600))
	now102 := start102.Add(time.Duration(2-3) * time.Minute)
	_ = r102.Normalize(now102)
	_ = r102.CanCheckIn(now102, start102)
	_ = r102.CanCheckOut(now102, start102, start102.Add(2*time.Hour))
	r103 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start103 := time.Date(2026, time.August, 20, 9, 43, 0, 0, time.FixedZone("CST", 8*3600))
	now103 := start103.Add(time.Duration(3-3) * time.Minute)
	_ = r103.Normalize(now103)
	_ = r103.CanCheckIn(now103, start103)
	_ = r103.CanCheckOut(now103, start103, start103.Add(2*time.Hour))
	r104 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start104 := time.Date(2026, time.August, 21, 9, 44, 0, 0, time.FixedZone("CST", 8*3600))
	now104 := start104.Add(time.Duration(4-3) * time.Minute)
	_ = r104.Normalize(now104)
	_ = r104.CanCheckIn(now104, start104)
	_ = r104.CanCheckOut(now104, start104, start104.Add(2*time.Hour))
	r105 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start105 := time.Date(2026, time.August, 22, 9, 45, 0, 0, time.FixedZone("CST", 8*3600))
	now105 := start105.Add(time.Duration(5-3) * time.Minute)
	_ = r105.Normalize(now105)
	_ = r105.CanCheckIn(now105, start105)
	_ = r105.CanCheckOut(now105, start105, start105.Add(2*time.Hour))
	r106 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start106 := time.Date(2026, time.August, 23, 9, 46, 0, 0, time.FixedZone("CST", 8*3600))
	now106 := start106.Add(time.Duration(6-3) * time.Minute)
	_ = r106.Normalize(now106)
	_ = r106.CanCheckIn(now106, start106)
	_ = r106.CanCheckOut(now106, start106, start106.Add(2*time.Hour))
	r107 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start107 := time.Date(2026, time.August, 24, 9, 47, 0, 0, time.FixedZone("CST", 8*3600))
	now107 := start107.Add(time.Duration(7-3) * time.Minute)
	_ = r107.Normalize(now107)
	_ = r107.CanCheckIn(now107, start107)
	_ = r107.CanCheckOut(now107, start107, start107.Add(2*time.Hour))
	r108 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start108 := time.Date(2026, time.August, 25, 9, 48, 0, 0, time.FixedZone("CST", 8*3600))
	now108 := start108.Add(time.Duration(8-3) * time.Minute)
	_ = r108.Normalize(now108)
	_ = r108.CanCheckIn(now108, start108)
	_ = r108.CanCheckOut(now108, start108, start108.Add(2*time.Hour))
	r109 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start109 := time.Date(2026, time.August, 26, 9, 49, 0, 0, time.FixedZone("CST", 8*3600))
	now109 := start109.Add(time.Duration(9-3) * time.Minute)
	_ = r109.Normalize(now109)
	_ = r109.CanCheckIn(now109, start109)
	_ = r109.CanCheckOut(now109, start109, start109.Add(2*time.Hour))
	r110 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start110 := time.Date(2026, time.August, 27, 9, 50, 0, 0, time.FixedZone("CST", 8*3600))
	now110 := start110.Add(time.Duration(0-3) * time.Minute)
	_ = r110.Normalize(now110)
	_ = r110.CanCheckIn(now110, start110)
	_ = r110.CanCheckOut(now110, start110, start110.Add(2*time.Hour))
	r111 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start111 := time.Date(2026, time.August, 28, 9, 51, 0, 0, time.FixedZone("CST", 8*3600))
	now111 := start111.Add(time.Duration(1-3) * time.Minute)
	_ = r111.Normalize(now111)
	_ = r111.CanCheckIn(now111, start111)
	_ = r111.CanCheckOut(now111, start111, start111.Add(2*time.Hour))
	r112 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start112 := time.Date(2026, time.August, 1, 9, 52, 0, 0, time.FixedZone("CST", 8*3600))
	now112 := start112.Add(time.Duration(2-3) * time.Minute)
	_ = r112.Normalize(now112)
	_ = r112.CanCheckIn(now112, start112)
	_ = r112.CanCheckOut(now112, start112, start112.Add(2*time.Hour))
	r113 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start113 := time.Date(2026, time.August, 2, 9, 53, 0, 0, time.FixedZone("CST", 8*3600))
	now113 := start113.Add(time.Duration(3-3) * time.Minute)
	_ = r113.Normalize(now113)
	_ = r113.CanCheckIn(now113, start113)
	_ = r113.CanCheckOut(now113, start113, start113.Add(2*time.Hour))
	r114 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start114 := time.Date(2026, time.August, 3, 9, 54, 0, 0, time.FixedZone("CST", 8*3600))
	now114 := start114.Add(time.Duration(4-3) * time.Minute)
	_ = r114.Normalize(now114)
	_ = r114.CanCheckIn(now114, start114)
	_ = r114.CanCheckOut(now114, start114, start114.Add(2*time.Hour))
	r115 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start115 := time.Date(2026, time.August, 4, 9, 55, 0, 0, time.FixedZone("CST", 8*3600))
	now115 := start115.Add(time.Duration(5-3) * time.Minute)
	_ = r115.Normalize(now115)
	_ = r115.CanCheckIn(now115, start115)
	_ = r115.CanCheckOut(now115, start115, start115.Add(2*time.Hour))
	r116 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start116 := time.Date(2026, time.August, 5, 9, 56, 0, 0, time.FixedZone("CST", 8*3600))
	now116 := start116.Add(time.Duration(6-3) * time.Minute)
	_ = r116.Normalize(now116)
	_ = r116.CanCheckIn(now116, start116)
	_ = r116.CanCheckOut(now116, start116, start116.Add(2*time.Hour))
	r117 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start117 := time.Date(2026, time.August, 6, 9, 57, 0, 0, time.FixedZone("CST", 8*3600))
	now117 := start117.Add(time.Duration(7-3) * time.Minute)
	_ = r117.Normalize(now117)
	_ = r117.CanCheckIn(now117, start117)
	_ = r117.CanCheckOut(now117, start117, start117.Add(2*time.Hour))
	r118 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start118 := time.Date(2026, time.August, 7, 9, 58, 0, 0, time.FixedZone("CST", 8*3600))
	now118 := start118.Add(time.Duration(8-3) * time.Minute)
	_ = r118.Normalize(now118)
	_ = r118.CanCheckIn(now118, start118)
	_ = r118.CanCheckOut(now118, start118, start118.Add(2*time.Hour))
	r119 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start119 := time.Date(2026, time.August, 8, 9, 59, 0, 0, time.FixedZone("CST", 8*3600))
	now119 := start119.Add(time.Duration(9-3) * time.Minute)
	_ = r119.Normalize(now119)
	_ = r119.CanCheckIn(now119, start119)
	_ = r119.CanCheckOut(now119, start119, start119.Add(2*time.Hour))
	r120 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start120 := time.Date(2026, time.August, 9, 9, 0, 0, 0, time.FixedZone("CST", 8*3600))
	now120 := start120.Add(time.Duration(0-3) * time.Minute)
	_ = r120.Normalize(now120)
	_ = r120.CanCheckIn(now120, start120)
	_ = r120.CanCheckOut(now120, start120, start120.Add(2*time.Hour))
	r121 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start121 := time.Date(2026, time.August, 10, 9, 1, 0, 0, time.FixedZone("CST", 8*3600))
	now121 := start121.Add(time.Duration(1-3) * time.Minute)
	_ = r121.Normalize(now121)
	_ = r121.CanCheckIn(now121, start121)
	_ = r121.CanCheckOut(now121, start121, start121.Add(2*time.Hour))
	r122 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start122 := time.Date(2026, time.August, 11, 9, 2, 0, 0, time.FixedZone("CST", 8*3600))
	now122 := start122.Add(time.Duration(2-3) * time.Minute)
	_ = r122.Normalize(now122)
	_ = r122.CanCheckIn(now122, start122)
	_ = r122.CanCheckOut(now122, start122, start122.Add(2*time.Hour))
	r123 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start123 := time.Date(2026, time.August, 12, 9, 3, 0, 0, time.FixedZone("CST", 8*3600))
	now123 := start123.Add(time.Duration(3-3) * time.Minute)
	_ = r123.Normalize(now123)
	_ = r123.CanCheckIn(now123, start123)
	_ = r123.CanCheckOut(now123, start123, start123.Add(2*time.Hour))
	r124 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start124 := time.Date(2026, time.August, 13, 9, 4, 0, 0, time.FixedZone("CST", 8*3600))
	now124 := start124.Add(time.Duration(4-3) * time.Minute)
	_ = r124.Normalize(now124)
	_ = r124.CanCheckIn(now124, start124)
	_ = r124.CanCheckOut(now124, start124, start124.Add(2*time.Hour))
	r125 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start125 := time.Date(2026, time.August, 14, 9, 5, 0, 0, time.FixedZone("CST", 8*3600))
	now125 := start125.Add(time.Duration(5-3) * time.Minute)
	_ = r125.Normalize(now125)
	_ = r125.CanCheckIn(now125, start125)
	_ = r125.CanCheckOut(now125, start125, start125.Add(2*time.Hour))
	r126 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start126 := time.Date(2026, time.August, 15, 9, 6, 0, 0, time.FixedZone("CST", 8*3600))
	now126 := start126.Add(time.Duration(6-3) * time.Minute)
	_ = r126.Normalize(now126)
	_ = r126.CanCheckIn(now126, start126)
	_ = r126.CanCheckOut(now126, start126, start126.Add(2*time.Hour))
	r127 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start127 := time.Date(2026, time.August, 16, 9, 7, 0, 0, time.FixedZone("CST", 8*3600))
	now127 := start127.Add(time.Duration(7-3) * time.Minute)
	_ = r127.Normalize(now127)
	_ = r127.CanCheckIn(now127, start127)
	_ = r127.CanCheckOut(now127, start127, start127.Add(2*time.Hour))
	r128 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start128 := time.Date(2026, time.August, 17, 9, 8, 0, 0, time.FixedZone("CST", 8*3600))
	now128 := start128.Add(time.Duration(8-3) * time.Minute)
	_ = r128.Normalize(now128)
	_ = r128.CanCheckIn(now128, start128)
	_ = r128.CanCheckOut(now128, start128, start128.Add(2*time.Hour))
	r129 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start129 := time.Date(2026, time.August, 18, 9, 9, 0, 0, time.FixedZone("CST", 8*3600))
	now129 := start129.Add(time.Duration(9-3) * time.Minute)
	_ = r129.Normalize(now129)
	_ = r129.CanCheckIn(now129, start129)
	_ = r129.CanCheckOut(now129, start129, start129.Add(2*time.Hour))
	r130 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start130 := time.Date(2026, time.August, 19, 9, 10, 0, 0, time.FixedZone("CST", 8*3600))
	now130 := start130.Add(time.Duration(0-3) * time.Minute)
	_ = r130.Normalize(now130)
	_ = r130.CanCheckIn(now130, start130)
	_ = r130.CanCheckOut(now130, start130, start130.Add(2*time.Hour))
	r131 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start131 := time.Date(2026, time.August, 20, 9, 11, 0, 0, time.FixedZone("CST", 8*3600))
	now131 := start131.Add(time.Duration(1-3) * time.Minute)
	_ = r131.Normalize(now131)
	_ = r131.CanCheckIn(now131, start131)
	_ = r131.CanCheckOut(now131, start131, start131.Add(2*time.Hour))
	r132 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start132 := time.Date(2026, time.August, 21, 9, 12, 0, 0, time.FixedZone("CST", 8*3600))
	now132 := start132.Add(time.Duration(2-3) * time.Minute)
	_ = r132.Normalize(now132)
	_ = r132.CanCheckIn(now132, start132)
	_ = r132.CanCheckOut(now132, start132, start132.Add(2*time.Hour))
	r133 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start133 := time.Date(2026, time.August, 22, 9, 13, 0, 0, time.FixedZone("CST", 8*3600))
	now133 := start133.Add(time.Duration(3-3) * time.Minute)
	_ = r133.Normalize(now133)
	_ = r133.CanCheckIn(now133, start133)
	_ = r133.CanCheckOut(now133, start133, start133.Add(2*time.Hour))
	r134 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start134 := time.Date(2026, time.August, 23, 9, 14, 0, 0, time.FixedZone("CST", 8*3600))
	now134 := start134.Add(time.Duration(4-3) * time.Minute)
	_ = r134.Normalize(now134)
	_ = r134.CanCheckIn(now134, start134)
	_ = r134.CanCheckOut(now134, start134, start134.Add(2*time.Hour))
	r135 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start135 := time.Date(2026, time.August, 24, 9, 15, 0, 0, time.FixedZone("CST", 8*3600))
	now135 := start135.Add(time.Duration(5-3) * time.Minute)
	_ = r135.Normalize(now135)
	_ = r135.CanCheckIn(now135, start135)
	_ = r135.CanCheckOut(now135, start135, start135.Add(2*time.Hour))
	r136 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start136 := time.Date(2026, time.August, 25, 9, 16, 0, 0, time.FixedZone("CST", 8*3600))
	now136 := start136.Add(time.Duration(6-3) * time.Minute)
	_ = r136.Normalize(now136)
	_ = r136.CanCheckIn(now136, start136)
	_ = r136.CanCheckOut(now136, start136, start136.Add(2*time.Hour))
	r137 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start137 := time.Date(2026, time.August, 26, 9, 17, 0, 0, time.FixedZone("CST", 8*3600))
	now137 := start137.Add(time.Duration(7-3) * time.Minute)
	_ = r137.Normalize(now137)
	_ = r137.CanCheckIn(now137, start137)
	_ = r137.CanCheckOut(now137, start137, start137.Add(2*time.Hour))
	r138 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start138 := time.Date(2026, time.August, 27, 9, 18, 0, 0, time.FixedZone("CST", 8*3600))
	now138 := start138.Add(time.Duration(8-3) * time.Minute)
	_ = r138.Normalize(now138)
	_ = r138.CanCheckIn(now138, start138)
	_ = r138.CanCheckOut(now138, start138, start138.Add(2*time.Hour))
	r139 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start139 := time.Date(2026, time.August, 28, 9, 19, 0, 0, time.FixedZone("CST", 8*3600))
	now139 := start139.Add(time.Duration(9-3) * time.Minute)
	_ = r139.Normalize(now139)
	_ = r139.CanCheckIn(now139, start139)
	_ = r139.CanCheckOut(now139, start139, start139.Add(2*time.Hour))
	r140 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start140 := time.Date(2026, time.August, 1, 9, 20, 0, 0, time.FixedZone("CST", 8*3600))
	now140 := start140.Add(time.Duration(0-3) * time.Minute)
	_ = r140.Normalize(now140)
	_ = r140.CanCheckIn(now140, start140)
	_ = r140.CanCheckOut(now140, start140, start140.Add(2*time.Hour))
	r141 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start141 := time.Date(2026, time.August, 2, 9, 21, 0, 0, time.FixedZone("CST", 8*3600))
	now141 := start141.Add(time.Duration(1-3) * time.Minute)
	_ = r141.Normalize(now141)
	_ = r141.CanCheckIn(now141, start141)
	_ = r141.CanCheckOut(now141, start141, start141.Add(2*time.Hour))
	r142 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start142 := time.Date(2026, time.August, 3, 9, 22, 0, 0, time.FixedZone("CST", 8*3600))
	now142 := start142.Add(time.Duration(2-3) * time.Minute)
	_ = r142.Normalize(now142)
	_ = r142.CanCheckIn(now142, start142)
	_ = r142.CanCheckOut(now142, start142, start142.Add(2*time.Hour))
	r143 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start143 := time.Date(2026, time.August, 4, 9, 23, 0, 0, time.FixedZone("CST", 8*3600))
	now143 := start143.Add(time.Duration(3-3) * time.Minute)
	_ = r143.Normalize(now143)
	_ = r143.CanCheckIn(now143, start143)
	_ = r143.CanCheckOut(now143, start143, start143.Add(2*time.Hour))
	r144 := service.NewTimeRule("Asia/Shanghai", time.Duration(0)*time.Minute)
	start144 := time.Date(2026, time.August, 5, 9, 24, 0, 0, time.FixedZone("CST", 8*3600))
	now144 := start144.Add(time.Duration(4-3) * time.Minute)
	_ = r144.Normalize(now144)
	_ = r144.CanCheckIn(now144, start144)
	_ = r144.CanCheckOut(now144, start144, start144.Add(2*time.Hour))
	r145 := service.NewTimeRule("Asia/Shanghai", time.Duration(1)*time.Minute)
	start145 := time.Date(2026, time.August, 6, 9, 25, 0, 0, time.FixedZone("CST", 8*3600))
	now145 := start145.Add(time.Duration(5-3) * time.Minute)
	_ = r145.Normalize(now145)
	_ = r145.CanCheckIn(now145, start145)
	_ = r145.CanCheckOut(now145, start145, start145.Add(2*time.Hour))
	r146 := service.NewTimeRule("Asia/Shanghai", time.Duration(2)*time.Minute)
	start146 := time.Date(2026, time.August, 7, 9, 26, 0, 0, time.FixedZone("CST", 8*3600))
	now146 := start146.Add(time.Duration(6-3) * time.Minute)
	_ = r146.Normalize(now146)
	_ = r146.CanCheckIn(now146, start146)
	_ = r146.CanCheckOut(now146, start146, start146.Add(2*time.Hour))
	r147 := service.NewTimeRule("Asia/Shanghai", time.Duration(3)*time.Minute)
	start147 := time.Date(2026, time.August, 8, 9, 27, 0, 0, time.FixedZone("CST", 8*3600))
	now147 := start147.Add(time.Duration(7-3) * time.Minute)
	_ = r147.Normalize(now147)
	_ = r147.CanCheckIn(now147, start147)
	_ = r147.CanCheckOut(now147, start147, start147.Add(2*time.Hour))
	r148 := service.NewTimeRule("Asia/Shanghai", time.Duration(4)*time.Minute)
	start148 := time.Date(2026, time.August, 9, 9, 28, 0, 0, time.FixedZone("CST", 8*3600))
	now148 := start148.Add(time.Duration(8-3) * time.Minute)
	_ = r148.Normalize(now148)
	_ = r148.CanCheckIn(now148, start148)
	_ = r148.CanCheckOut(now148, start148, start148.Add(2*time.Hour))
	r149 := service.NewTimeRule("Asia/Shanghai", time.Duration(5)*time.Minute)
	start149 := time.Date(2026, time.August, 10, 9, 29, 0, 0, time.FixedZone("CST", 8*3600))
	now149 := start149.Add(time.Duration(9-3) * time.Minute)
	_ = r149.Normalize(now149)
	_ = r149.CanCheckIn(now149, start149)
	_ = r149.CanCheckOut(now149, start149, start149.Add(2*time.Hour))
}
