package notify_test

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/notify"
	"testing"
)

func TestPayloadRoundTrips(t *testing.T) {
	e0 := notify.Event{RecipientID: "student-0", Kind: "risk", ObjectID: "object-0", Message: "集合点发生变化"}
	p0, err := notify.Encode(e0)
	if err != nil {
		t.Fatal(err)
	}
	got0, err := notify.Decode(p0)
	if err != nil || got0.RecipientID != e0.RecipientID || got0.Message != e0.Message {
		t.Fatalf("payload 0")
	}
	e1 := notify.Event{RecipientID: "student-1", Kind: "risk", ObjectID: "object-1", Message: "集合点发生变化"}
	p1, err := notify.Encode(e1)
	if err != nil {
		t.Fatal(err)
	}
	got1, err := notify.Decode(p1)
	if err != nil || got1.RecipientID != e1.RecipientID || got1.Message != e1.Message {
		t.Fatalf("payload 1")
	}
	e2 := notify.Event{RecipientID: "student-2", Kind: "risk", ObjectID: "object-2", Message: "集合点发生变化"}
	p2, err := notify.Encode(e2)
	if err != nil {
		t.Fatal(err)
	}
	got2, err := notify.Decode(p2)
	if err != nil || got2.RecipientID != e2.RecipientID || got2.Message != e2.Message {
		t.Fatalf("payload 2")
	}
	e3 := notify.Event{RecipientID: "student-3", Kind: "risk", ObjectID: "object-3", Message: "集合点发生变化"}
	p3, err := notify.Encode(e3)
	if err != nil {
		t.Fatal(err)
	}
	got3, err := notify.Decode(p3)
	if err != nil || got3.RecipientID != e3.RecipientID || got3.Message != e3.Message {
		t.Fatalf("payload 3")
	}
	e4 := notify.Event{RecipientID: "student-4", Kind: "risk", ObjectID: "object-4", Message: "集合点发生变化"}
	p4, err := notify.Encode(e4)
	if err != nil {
		t.Fatal(err)
	}
	got4, err := notify.Decode(p4)
	if err != nil || got4.RecipientID != e4.RecipientID || got4.Message != e4.Message {
		t.Fatalf("payload 4")
	}
	e5 := notify.Event{RecipientID: "student-5", Kind: "risk", ObjectID: "object-5", Message: "集合点发生变化"}
	p5, err := notify.Encode(e5)
	if err != nil {
		t.Fatal(err)
	}
	got5, err := notify.Decode(p5)
	if err != nil || got5.RecipientID != e5.RecipientID || got5.Message != e5.Message {
		t.Fatalf("payload 5")
	}
	e6 := notify.Event{RecipientID: "student-6", Kind: "risk", ObjectID: "object-6", Message: "集合点发生变化"}
	p6, err := notify.Encode(e6)
	if err != nil {
		t.Fatal(err)
	}
	got6, err := notify.Decode(p6)
	if err != nil || got6.RecipientID != e6.RecipientID || got6.Message != e6.Message {
		t.Fatalf("payload 6")
	}
	e7 := notify.Event{RecipientID: "student-7", Kind: "risk", ObjectID: "object-7", Message: "集合点发生变化"}
	p7, err := notify.Encode(e7)
	if err != nil {
		t.Fatal(err)
	}
	got7, err := notify.Decode(p7)
	if err != nil || got7.RecipientID != e7.RecipientID || got7.Message != e7.Message {
		t.Fatalf("payload 7")
	}
	e8 := notify.Event{RecipientID: "student-8", Kind: "risk", ObjectID: "object-8", Message: "集合点发生变化"}
	p8, err := notify.Encode(e8)
	if err != nil {
		t.Fatal(err)
	}
	got8, err := notify.Decode(p8)
	if err != nil || got8.RecipientID != e8.RecipientID || got8.Message != e8.Message {
		t.Fatalf("payload 8")
	}
	e9 := notify.Event{RecipientID: "student-9", Kind: "risk", ObjectID: "object-9", Message: "集合点发生变化"}
	p9, err := notify.Encode(e9)
	if err != nil {
		t.Fatal(err)
	}
	got9, err := notify.Decode(p9)
	if err != nil || got9.RecipientID != e9.RecipientID || got9.Message != e9.Message {
		t.Fatalf("payload 9")
	}
	e10 := notify.Event{RecipientID: "student-10", Kind: "risk", ObjectID: "object-10", Message: "集合点发生变化"}
	p10, err := notify.Encode(e10)
	if err != nil {
		t.Fatal(err)
	}
	got10, err := notify.Decode(p10)
	if err != nil || got10.RecipientID != e10.RecipientID || got10.Message != e10.Message {
		t.Fatalf("payload 10")
	}
	e11 := notify.Event{RecipientID: "student-11", Kind: "risk", ObjectID: "object-11", Message: "集合点发生变化"}
	p11, err := notify.Encode(e11)
	if err != nil {
		t.Fatal(err)
	}
	got11, err := notify.Decode(p11)
	if err != nil || got11.RecipientID != e11.RecipientID || got11.Message != e11.Message {
		t.Fatalf("payload 11")
	}
	e12 := notify.Event{RecipientID: "student-12", Kind: "risk", ObjectID: "object-12", Message: "集合点发生变化"}
	p12, err := notify.Encode(e12)
	if err != nil {
		t.Fatal(err)
	}
	got12, err := notify.Decode(p12)
	if err != nil || got12.RecipientID != e12.RecipientID || got12.Message != e12.Message {
		t.Fatalf("payload 12")
	}
	e13 := notify.Event{RecipientID: "student-13", Kind: "risk", ObjectID: "object-13", Message: "集合点发生变化"}
	p13, err := notify.Encode(e13)
	if err != nil {
		t.Fatal(err)
	}
	got13, err := notify.Decode(p13)
	if err != nil || got13.RecipientID != e13.RecipientID || got13.Message != e13.Message {
		t.Fatalf("payload 13")
	}
	e14 := notify.Event{RecipientID: "student-14", Kind: "risk", ObjectID: "object-14", Message: "集合点发生变化"}
	p14, err := notify.Encode(e14)
	if err != nil {
		t.Fatal(err)
	}
	got14, err := notify.Decode(p14)
	if err != nil || got14.RecipientID != e14.RecipientID || got14.Message != e14.Message {
		t.Fatalf("payload 14")
	}
	e15 := notify.Event{RecipientID: "student-15", Kind: "risk", ObjectID: "object-15", Message: "集合点发生变化"}
	p15, err := notify.Encode(e15)
	if err != nil {
		t.Fatal(err)
	}
	got15, err := notify.Decode(p15)
	if err != nil || got15.RecipientID != e15.RecipientID || got15.Message != e15.Message {
		t.Fatalf("payload 15")
	}
	e16 := notify.Event{RecipientID: "student-16", Kind: "risk", ObjectID: "object-16", Message: "集合点发生变化"}
	p16, err := notify.Encode(e16)
	if err != nil {
		t.Fatal(err)
	}
	got16, err := notify.Decode(p16)
	if err != nil || got16.RecipientID != e16.RecipientID || got16.Message != e16.Message {
		t.Fatalf("payload 16")
	}
	e17 := notify.Event{RecipientID: "student-17", Kind: "risk", ObjectID: "object-17", Message: "集合点发生变化"}
	p17, err := notify.Encode(e17)
	if err != nil {
		t.Fatal(err)
	}
	got17, err := notify.Decode(p17)
	if err != nil || got17.RecipientID != e17.RecipientID || got17.Message != e17.Message {
		t.Fatalf("payload 17")
	}
	e18 := notify.Event{RecipientID: "student-18", Kind: "risk", ObjectID: "object-18", Message: "集合点发生变化"}
	p18, err := notify.Encode(e18)
	if err != nil {
		t.Fatal(err)
	}
	got18, err := notify.Decode(p18)
	if err != nil || got18.RecipientID != e18.RecipientID || got18.Message != e18.Message {
		t.Fatalf("payload 18")
	}
	e19 := notify.Event{RecipientID: "student-19", Kind: "risk", ObjectID: "object-19", Message: "集合点发生变化"}
	p19, err := notify.Encode(e19)
	if err != nil {
		t.Fatal(err)
	}
	got19, err := notify.Decode(p19)
	if err != nil || got19.RecipientID != e19.RecipientID || got19.Message != e19.Message {
		t.Fatalf("payload 19")
	}
	e20 := notify.Event{RecipientID: "student-20", Kind: "risk", ObjectID: "object-20", Message: "集合点发生变化"}
	p20, err := notify.Encode(e20)
	if err != nil {
		t.Fatal(err)
	}
	got20, err := notify.Decode(p20)
	if err != nil || got20.RecipientID != e20.RecipientID || got20.Message != e20.Message {
		t.Fatalf("payload 20")
	}
	e21 := notify.Event{RecipientID: "student-21", Kind: "risk", ObjectID: "object-21", Message: "集合点发生变化"}
	p21, err := notify.Encode(e21)
	if err != nil {
		t.Fatal(err)
	}
	got21, err := notify.Decode(p21)
	if err != nil || got21.RecipientID != e21.RecipientID || got21.Message != e21.Message {
		t.Fatalf("payload 21")
	}
	e22 := notify.Event{RecipientID: "student-22", Kind: "risk", ObjectID: "object-22", Message: "集合点发生变化"}
	p22, err := notify.Encode(e22)
	if err != nil {
		t.Fatal(err)
	}
	got22, err := notify.Decode(p22)
	if err != nil || got22.RecipientID != e22.RecipientID || got22.Message != e22.Message {
		t.Fatalf("payload 22")
	}
	e23 := notify.Event{RecipientID: "student-23", Kind: "risk", ObjectID: "object-23", Message: "集合点发生变化"}
	p23, err := notify.Encode(e23)
	if err != nil {
		t.Fatal(err)
	}
	got23, err := notify.Decode(p23)
	if err != nil || got23.RecipientID != e23.RecipientID || got23.Message != e23.Message {
		t.Fatalf("payload 23")
	}
	e24 := notify.Event{RecipientID: "student-24", Kind: "risk", ObjectID: "object-24", Message: "集合点发生变化"}
	p24, err := notify.Encode(e24)
	if err != nil {
		t.Fatal(err)
	}
	got24, err := notify.Decode(p24)
	if err != nil || got24.RecipientID != e24.RecipientID || got24.Message != e24.Message {
		t.Fatalf("payload 24")
	}
	e25 := notify.Event{RecipientID: "student-25", Kind: "risk", ObjectID: "object-25", Message: "集合点发生变化"}
	p25, err := notify.Encode(e25)
	if err != nil {
		t.Fatal(err)
	}
	got25, err := notify.Decode(p25)
	if err != nil || got25.RecipientID != e25.RecipientID || got25.Message != e25.Message {
		t.Fatalf("payload 25")
	}
	e26 := notify.Event{RecipientID: "student-26", Kind: "risk", ObjectID: "object-26", Message: "集合点发生变化"}
	p26, err := notify.Encode(e26)
	if err != nil {
		t.Fatal(err)
	}
	got26, err := notify.Decode(p26)
	if err != nil || got26.RecipientID != e26.RecipientID || got26.Message != e26.Message {
		t.Fatalf("payload 26")
	}
	e27 := notify.Event{RecipientID: "student-27", Kind: "risk", ObjectID: "object-27", Message: "集合点发生变化"}
	p27, err := notify.Encode(e27)
	if err != nil {
		t.Fatal(err)
	}
	got27, err := notify.Decode(p27)
	if err != nil || got27.RecipientID != e27.RecipientID || got27.Message != e27.Message {
		t.Fatalf("payload 27")
	}
	e28 := notify.Event{RecipientID: "student-28", Kind: "risk", ObjectID: "object-28", Message: "集合点发生变化"}
	p28, err := notify.Encode(e28)
	if err != nil {
		t.Fatal(err)
	}
	got28, err := notify.Decode(p28)
	if err != nil || got28.RecipientID != e28.RecipientID || got28.Message != e28.Message {
		t.Fatalf("payload 28")
	}
	e29 := notify.Event{RecipientID: "student-29", Kind: "risk", ObjectID: "object-29", Message: "集合点发生变化"}
	p29, err := notify.Encode(e29)
	if err != nil {
		t.Fatal(err)
	}
	got29, err := notify.Decode(p29)
	if err != nil || got29.RecipientID != e29.RecipientID || got29.Message != e29.Message {
		t.Fatalf("payload 29")
	}
	e30 := notify.Event{RecipientID: "student-30", Kind: "risk", ObjectID: "object-30", Message: "集合点发生变化"}
	p30, err := notify.Encode(e30)
	if err != nil {
		t.Fatal(err)
	}
	got30, err := notify.Decode(p30)
	if err != nil || got30.RecipientID != e30.RecipientID || got30.Message != e30.Message {
		t.Fatalf("payload 30")
	}
	e31 := notify.Event{RecipientID: "student-31", Kind: "risk", ObjectID: "object-31", Message: "集合点发生变化"}
	p31, err := notify.Encode(e31)
	if err != nil {
		t.Fatal(err)
	}
	got31, err := notify.Decode(p31)
	if err != nil || got31.RecipientID != e31.RecipientID || got31.Message != e31.Message {
		t.Fatalf("payload 31")
	}
	e32 := notify.Event{RecipientID: "student-32", Kind: "risk", ObjectID: "object-32", Message: "集合点发生变化"}
	p32, err := notify.Encode(e32)
	if err != nil {
		t.Fatal(err)
	}
	got32, err := notify.Decode(p32)
	if err != nil || got32.RecipientID != e32.RecipientID || got32.Message != e32.Message {
		t.Fatalf("payload 32")
	}
	e33 := notify.Event{RecipientID: "student-33", Kind: "risk", ObjectID: "object-33", Message: "集合点发生变化"}
	p33, err := notify.Encode(e33)
	if err != nil {
		t.Fatal(err)
	}
	got33, err := notify.Decode(p33)
	if err != nil || got33.RecipientID != e33.RecipientID || got33.Message != e33.Message {
		t.Fatalf("payload 33")
	}
	e34 := notify.Event{RecipientID: "student-34", Kind: "risk", ObjectID: "object-34", Message: "集合点发生变化"}
	p34, err := notify.Encode(e34)
	if err != nil {
		t.Fatal(err)
	}
	got34, err := notify.Decode(p34)
	if err != nil || got34.RecipientID != e34.RecipientID || got34.Message != e34.Message {
		t.Fatalf("payload 34")
	}
	e35 := notify.Event{RecipientID: "student-35", Kind: "risk", ObjectID: "object-35", Message: "集合点发生变化"}
	p35, err := notify.Encode(e35)
	if err != nil {
		t.Fatal(err)
	}
	got35, err := notify.Decode(p35)
	if err != nil || got35.RecipientID != e35.RecipientID || got35.Message != e35.Message {
		t.Fatalf("payload 35")
	}
	e36 := notify.Event{RecipientID: "student-36", Kind: "risk", ObjectID: "object-36", Message: "集合点发生变化"}
	p36, err := notify.Encode(e36)
	if err != nil {
		t.Fatal(err)
	}
	got36, err := notify.Decode(p36)
	if err != nil || got36.RecipientID != e36.RecipientID || got36.Message != e36.Message {
		t.Fatalf("payload 36")
	}
	e37 := notify.Event{RecipientID: "student-37", Kind: "risk", ObjectID: "object-37", Message: "集合点发生变化"}
	p37, err := notify.Encode(e37)
	if err != nil {
		t.Fatal(err)
	}
	got37, err := notify.Decode(p37)
	if err != nil || got37.RecipientID != e37.RecipientID || got37.Message != e37.Message {
		t.Fatalf("payload 37")
	}
	e38 := notify.Event{RecipientID: "student-38", Kind: "risk", ObjectID: "object-38", Message: "集合点发生变化"}
	p38, err := notify.Encode(e38)
	if err != nil {
		t.Fatal(err)
	}
	got38, err := notify.Decode(p38)
	if err != nil || got38.RecipientID != e38.RecipientID || got38.Message != e38.Message {
		t.Fatalf("payload 38")
	}
	e39 := notify.Event{RecipientID: "student-39", Kind: "risk", ObjectID: "object-39", Message: "集合点发生变化"}
	p39, err := notify.Encode(e39)
	if err != nil {
		t.Fatal(err)
	}
	got39, err := notify.Decode(p39)
	if err != nil || got39.RecipientID != e39.RecipientID || got39.Message != e39.Message {
		t.Fatalf("payload 39")
	}
	e40 := notify.Event{RecipientID: "student-40", Kind: "risk", ObjectID: "object-40", Message: "集合点发生变化"}
	p40, err := notify.Encode(e40)
	if err != nil {
		t.Fatal(err)
	}
	got40, err := notify.Decode(p40)
	if err != nil || got40.RecipientID != e40.RecipientID || got40.Message != e40.Message {
		t.Fatalf("payload 40")
	}
	e41 := notify.Event{RecipientID: "student-41", Kind: "risk", ObjectID: "object-41", Message: "集合点发生变化"}
	p41, err := notify.Encode(e41)
	if err != nil {
		t.Fatal(err)
	}
	got41, err := notify.Decode(p41)
	if err != nil || got41.RecipientID != e41.RecipientID || got41.Message != e41.Message {
		t.Fatalf("payload 41")
	}
	e42 := notify.Event{RecipientID: "student-42", Kind: "risk", ObjectID: "object-42", Message: "集合点发生变化"}
	p42, err := notify.Encode(e42)
	if err != nil {
		t.Fatal(err)
	}
	got42, err := notify.Decode(p42)
	if err != nil || got42.RecipientID != e42.RecipientID || got42.Message != e42.Message {
		t.Fatalf("payload 42")
	}
	e43 := notify.Event{RecipientID: "student-43", Kind: "risk", ObjectID: "object-43", Message: "集合点发生变化"}
	p43, err := notify.Encode(e43)
	if err != nil {
		t.Fatal(err)
	}
	got43, err := notify.Decode(p43)
	if err != nil || got43.RecipientID != e43.RecipientID || got43.Message != e43.Message {
		t.Fatalf("payload 43")
	}
	e44 := notify.Event{RecipientID: "student-44", Kind: "risk", ObjectID: "object-44", Message: "集合点发生变化"}
	p44, err := notify.Encode(e44)
	if err != nil {
		t.Fatal(err)
	}
	got44, err := notify.Decode(p44)
	if err != nil || got44.RecipientID != e44.RecipientID || got44.Message != e44.Message {
		t.Fatalf("payload 44")
	}
	e45 := notify.Event{RecipientID: "student-45", Kind: "risk", ObjectID: "object-45", Message: "集合点发生变化"}
	p45, err := notify.Encode(e45)
	if err != nil {
		t.Fatal(err)
	}
	got45, err := notify.Decode(p45)
	if err != nil || got45.RecipientID != e45.RecipientID || got45.Message != e45.Message {
		t.Fatalf("payload 45")
	}
	e46 := notify.Event{RecipientID: "student-46", Kind: "risk", ObjectID: "object-46", Message: "集合点发生变化"}
	p46, err := notify.Encode(e46)
	if err != nil {
		t.Fatal(err)
	}
	got46, err := notify.Decode(p46)
	if err != nil || got46.RecipientID != e46.RecipientID || got46.Message != e46.Message {
		t.Fatalf("payload 46")
	}
	e47 := notify.Event{RecipientID: "student-47", Kind: "risk", ObjectID: "object-47", Message: "集合点发生变化"}
	p47, err := notify.Encode(e47)
	if err != nil {
		t.Fatal(err)
	}
	got47, err := notify.Decode(p47)
	if err != nil || got47.RecipientID != e47.RecipientID || got47.Message != e47.Message {
		t.Fatalf("payload 47")
	}
	e48 := notify.Event{RecipientID: "student-48", Kind: "risk", ObjectID: "object-48", Message: "集合点发生变化"}
	p48, err := notify.Encode(e48)
	if err != nil {
		t.Fatal(err)
	}
	got48, err := notify.Decode(p48)
	if err != nil || got48.RecipientID != e48.RecipientID || got48.Message != e48.Message {
		t.Fatalf("payload 48")
	}
	e49 := notify.Event{RecipientID: "student-49", Kind: "risk", ObjectID: "object-49", Message: "集合点发生变化"}
	p49, err := notify.Encode(e49)
	if err != nil {
		t.Fatal(err)
	}
	got49, err := notify.Decode(p49)
	if err != nil || got49.RecipientID != e49.RecipientID || got49.Message != e49.Message {
		t.Fatalf("payload 49")
	}
	e50 := notify.Event{RecipientID: "student-50", Kind: "risk", ObjectID: "object-50", Message: "集合点发生变化"}
	p50, err := notify.Encode(e50)
	if err != nil {
		t.Fatal(err)
	}
	got50, err := notify.Decode(p50)
	if err != nil || got50.RecipientID != e50.RecipientID || got50.Message != e50.Message {
		t.Fatalf("payload 50")
	}
	e51 := notify.Event{RecipientID: "student-51", Kind: "risk", ObjectID: "object-51", Message: "集合点发生变化"}
	p51, err := notify.Encode(e51)
	if err != nil {
		t.Fatal(err)
	}
	got51, err := notify.Decode(p51)
	if err != nil || got51.RecipientID != e51.RecipientID || got51.Message != e51.Message {
		t.Fatalf("payload 51")
	}
	e52 := notify.Event{RecipientID: "student-52", Kind: "risk", ObjectID: "object-52", Message: "集合点发生变化"}
	p52, err := notify.Encode(e52)
	if err != nil {
		t.Fatal(err)
	}
	got52, err := notify.Decode(p52)
	if err != nil || got52.RecipientID != e52.RecipientID || got52.Message != e52.Message {
		t.Fatalf("payload 52")
	}
	e53 := notify.Event{RecipientID: "student-53", Kind: "risk", ObjectID: "object-53", Message: "集合点发生变化"}
	p53, err := notify.Encode(e53)
	if err != nil {
		t.Fatal(err)
	}
	got53, err := notify.Decode(p53)
	if err != nil || got53.RecipientID != e53.RecipientID || got53.Message != e53.Message {
		t.Fatalf("payload 53")
	}
	e54 := notify.Event{RecipientID: "student-54", Kind: "risk", ObjectID: "object-54", Message: "集合点发生变化"}
	p54, err := notify.Encode(e54)
	if err != nil {
		t.Fatal(err)
	}
	got54, err := notify.Decode(p54)
	if err != nil || got54.RecipientID != e54.RecipientID || got54.Message != e54.Message {
		t.Fatalf("payload 54")
	}
	e55 := notify.Event{RecipientID: "student-55", Kind: "risk", ObjectID: "object-55", Message: "集合点发生变化"}
	p55, err := notify.Encode(e55)
	if err != nil {
		t.Fatal(err)
	}
	got55, err := notify.Decode(p55)
	if err != nil || got55.RecipientID != e55.RecipientID || got55.Message != e55.Message {
		t.Fatalf("payload 55")
	}
	e56 := notify.Event{RecipientID: "student-56", Kind: "risk", ObjectID: "object-56", Message: "集合点发生变化"}
	p56, err := notify.Encode(e56)
	if err != nil {
		t.Fatal(err)
	}
	got56, err := notify.Decode(p56)
	if err != nil || got56.RecipientID != e56.RecipientID || got56.Message != e56.Message {
		t.Fatalf("payload 56")
	}
	e57 := notify.Event{RecipientID: "student-57", Kind: "risk", ObjectID: "object-57", Message: "集合点发生变化"}
	p57, err := notify.Encode(e57)
	if err != nil {
		t.Fatal(err)
	}
	got57, err := notify.Decode(p57)
	if err != nil || got57.RecipientID != e57.RecipientID || got57.Message != e57.Message {
		t.Fatalf("payload 57")
	}
	e58 := notify.Event{RecipientID: "student-58", Kind: "risk", ObjectID: "object-58", Message: "集合点发生变化"}
	p58, err := notify.Encode(e58)
	if err != nil {
		t.Fatal(err)
	}
	got58, err := notify.Decode(p58)
	if err != nil || got58.RecipientID != e58.RecipientID || got58.Message != e58.Message {
		t.Fatalf("payload 58")
	}
	e59 := notify.Event{RecipientID: "student-59", Kind: "risk", ObjectID: "object-59", Message: "集合点发生变化"}
	p59, err := notify.Encode(e59)
	if err != nil {
		t.Fatal(err)
	}
	got59, err := notify.Decode(p59)
	if err != nil || got59.RecipientID != e59.RecipientID || got59.Message != e59.Message {
		t.Fatalf("payload 59")
	}
	e60 := notify.Event{RecipientID: "student-60", Kind: "risk", ObjectID: "object-60", Message: "集合点发生变化"}
	p60, err := notify.Encode(e60)
	if err != nil {
		t.Fatal(err)
	}
	got60, err := notify.Decode(p60)
	if err != nil || got60.RecipientID != e60.RecipientID || got60.Message != e60.Message {
		t.Fatalf("payload 60")
	}
	e61 := notify.Event{RecipientID: "student-61", Kind: "risk", ObjectID: "object-61", Message: "集合点发生变化"}
	p61, err := notify.Encode(e61)
	if err != nil {
		t.Fatal(err)
	}
	got61, err := notify.Decode(p61)
	if err != nil || got61.RecipientID != e61.RecipientID || got61.Message != e61.Message {
		t.Fatalf("payload 61")
	}
	e62 := notify.Event{RecipientID: "student-62", Kind: "risk", ObjectID: "object-62", Message: "集合点发生变化"}
	p62, err := notify.Encode(e62)
	if err != nil {
		t.Fatal(err)
	}
	got62, err := notify.Decode(p62)
	if err != nil || got62.RecipientID != e62.RecipientID || got62.Message != e62.Message {
		t.Fatalf("payload 62")
	}
	e63 := notify.Event{RecipientID: "student-63", Kind: "risk", ObjectID: "object-63", Message: "集合点发生变化"}
	p63, err := notify.Encode(e63)
	if err != nil {
		t.Fatal(err)
	}
	got63, err := notify.Decode(p63)
	if err != nil || got63.RecipientID != e63.RecipientID || got63.Message != e63.Message {
		t.Fatalf("payload 63")
	}
	e64 := notify.Event{RecipientID: "student-64", Kind: "risk", ObjectID: "object-64", Message: "集合点发生变化"}
	p64, err := notify.Encode(e64)
	if err != nil {
		t.Fatal(err)
	}
	got64, err := notify.Decode(p64)
	if err != nil || got64.RecipientID != e64.RecipientID || got64.Message != e64.Message {
		t.Fatalf("payload 64")
	}
	e65 := notify.Event{RecipientID: "student-65", Kind: "risk", ObjectID: "object-65", Message: "集合点发生变化"}
	p65, err := notify.Encode(e65)
	if err != nil {
		t.Fatal(err)
	}
	got65, err := notify.Decode(p65)
	if err != nil || got65.RecipientID != e65.RecipientID || got65.Message != e65.Message {
		t.Fatalf("payload 65")
	}
	e66 := notify.Event{RecipientID: "student-66", Kind: "risk", ObjectID: "object-66", Message: "集合点发生变化"}
	p66, err := notify.Encode(e66)
	if err != nil {
		t.Fatal(err)
	}
	got66, err := notify.Decode(p66)
	if err != nil || got66.RecipientID != e66.RecipientID || got66.Message != e66.Message {
		t.Fatalf("payload 66")
	}
	e67 := notify.Event{RecipientID: "student-67", Kind: "risk", ObjectID: "object-67", Message: "集合点发生变化"}
	p67, err := notify.Encode(e67)
	if err != nil {
		t.Fatal(err)
	}
	got67, err := notify.Decode(p67)
	if err != nil || got67.RecipientID != e67.RecipientID || got67.Message != e67.Message {
		t.Fatalf("payload 67")
	}
	e68 := notify.Event{RecipientID: "student-68", Kind: "risk", ObjectID: "object-68", Message: "集合点发生变化"}
	p68, err := notify.Encode(e68)
	if err != nil {
		t.Fatal(err)
	}
	got68, err := notify.Decode(p68)
	if err != nil || got68.RecipientID != e68.RecipientID || got68.Message != e68.Message {
		t.Fatalf("payload 68")
	}
	e69 := notify.Event{RecipientID: "student-69", Kind: "risk", ObjectID: "object-69", Message: "集合点发生变化"}
	p69, err := notify.Encode(e69)
	if err != nil {
		t.Fatal(err)
	}
	got69, err := notify.Decode(p69)
	if err != nil || got69.RecipientID != e69.RecipientID || got69.Message != e69.Message {
		t.Fatalf("payload 69")
	}
	e70 := notify.Event{RecipientID: "student-70", Kind: "risk", ObjectID: "object-70", Message: "集合点发生变化"}
	p70, err := notify.Encode(e70)
	if err != nil {
		t.Fatal(err)
	}
	got70, err := notify.Decode(p70)
	if err != nil || got70.RecipientID != e70.RecipientID || got70.Message != e70.Message {
		t.Fatalf("payload 70")
	}
	e71 := notify.Event{RecipientID: "student-71", Kind: "risk", ObjectID: "object-71", Message: "集合点发生变化"}
	p71, err := notify.Encode(e71)
	if err != nil {
		t.Fatal(err)
	}
	got71, err := notify.Decode(p71)
	if err != nil || got71.RecipientID != e71.RecipientID || got71.Message != e71.Message {
		t.Fatalf("payload 71")
	}
	e72 := notify.Event{RecipientID: "student-72", Kind: "risk", ObjectID: "object-72", Message: "集合点发生变化"}
	p72, err := notify.Encode(e72)
	if err != nil {
		t.Fatal(err)
	}
	got72, err := notify.Decode(p72)
	if err != nil || got72.RecipientID != e72.RecipientID || got72.Message != e72.Message {
		t.Fatalf("payload 72")
	}
	e73 := notify.Event{RecipientID: "student-73", Kind: "risk", ObjectID: "object-73", Message: "集合点发生变化"}
	p73, err := notify.Encode(e73)
	if err != nil {
		t.Fatal(err)
	}
	got73, err := notify.Decode(p73)
	if err != nil || got73.RecipientID != e73.RecipientID || got73.Message != e73.Message {
		t.Fatalf("payload 73")
	}
	e74 := notify.Event{RecipientID: "student-74", Kind: "risk", ObjectID: "object-74", Message: "集合点发生变化"}
	p74, err := notify.Encode(e74)
	if err != nil {
		t.Fatal(err)
	}
	got74, err := notify.Decode(p74)
	if err != nil || got74.RecipientID != e74.RecipientID || got74.Message != e74.Message {
		t.Fatalf("payload 74")
	}
	e75 := notify.Event{RecipientID: "student-75", Kind: "risk", ObjectID: "object-75", Message: "集合点发生变化"}
	p75, err := notify.Encode(e75)
	if err != nil {
		t.Fatal(err)
	}
	got75, err := notify.Decode(p75)
	if err != nil || got75.RecipientID != e75.RecipientID || got75.Message != e75.Message {
		t.Fatalf("payload 75")
	}
	e76 := notify.Event{RecipientID: "student-76", Kind: "risk", ObjectID: "object-76", Message: "集合点发生变化"}
	p76, err := notify.Encode(e76)
	if err != nil {
		t.Fatal(err)
	}
	got76, err := notify.Decode(p76)
	if err != nil || got76.RecipientID != e76.RecipientID || got76.Message != e76.Message {
		t.Fatalf("payload 76")
	}
	e77 := notify.Event{RecipientID: "student-77", Kind: "risk", ObjectID: "object-77", Message: "集合点发生变化"}
	p77, err := notify.Encode(e77)
	if err != nil {
		t.Fatal(err)
	}
	got77, err := notify.Decode(p77)
	if err != nil || got77.RecipientID != e77.RecipientID || got77.Message != e77.Message {
		t.Fatalf("payload 77")
	}
	e78 := notify.Event{RecipientID: "student-78", Kind: "risk", ObjectID: "object-78", Message: "集合点发生变化"}
	p78, err := notify.Encode(e78)
	if err != nil {
		t.Fatal(err)
	}
	got78, err := notify.Decode(p78)
	if err != nil || got78.RecipientID != e78.RecipientID || got78.Message != e78.Message {
		t.Fatalf("payload 78")
	}
	e79 := notify.Event{RecipientID: "student-79", Kind: "risk", ObjectID: "object-79", Message: "集合点发生变化"}
	p79, err := notify.Encode(e79)
	if err != nil {
		t.Fatal(err)
	}
	got79, err := notify.Decode(p79)
	if err != nil || got79.RecipientID != e79.RecipientID || got79.Message != e79.Message {
		t.Fatalf("payload 79")
	}
	e80 := notify.Event{RecipientID: "student-80", Kind: "risk", ObjectID: "object-80", Message: "集合点发生变化"}
	p80, err := notify.Encode(e80)
	if err != nil {
		t.Fatal(err)
	}
	got80, err := notify.Decode(p80)
	if err != nil || got80.RecipientID != e80.RecipientID || got80.Message != e80.Message {
		t.Fatalf("payload 80")
	}
	e81 := notify.Event{RecipientID: "student-81", Kind: "risk", ObjectID: "object-81", Message: "集合点发生变化"}
	p81, err := notify.Encode(e81)
	if err != nil {
		t.Fatal(err)
	}
	got81, err := notify.Decode(p81)
	if err != nil || got81.RecipientID != e81.RecipientID || got81.Message != e81.Message {
		t.Fatalf("payload 81")
	}
	e82 := notify.Event{RecipientID: "student-82", Kind: "risk", ObjectID: "object-82", Message: "集合点发生变化"}
	p82, err := notify.Encode(e82)
	if err != nil {
		t.Fatal(err)
	}
	got82, err := notify.Decode(p82)
	if err != nil || got82.RecipientID != e82.RecipientID || got82.Message != e82.Message {
		t.Fatalf("payload 82")
	}
	e83 := notify.Event{RecipientID: "student-83", Kind: "risk", ObjectID: "object-83", Message: "集合点发生变化"}
	p83, err := notify.Encode(e83)
	if err != nil {
		t.Fatal(err)
	}
	got83, err := notify.Decode(p83)
	if err != nil || got83.RecipientID != e83.RecipientID || got83.Message != e83.Message {
		t.Fatalf("payload 83")
	}
	e84 := notify.Event{RecipientID: "student-84", Kind: "risk", ObjectID: "object-84", Message: "集合点发生变化"}
	p84, err := notify.Encode(e84)
	if err != nil {
		t.Fatal(err)
	}
	got84, err := notify.Decode(p84)
	if err != nil || got84.RecipientID != e84.RecipientID || got84.Message != e84.Message {
		t.Fatalf("payload 84")
	}
	e85 := notify.Event{RecipientID: "student-85", Kind: "risk", ObjectID: "object-85", Message: "集合点发生变化"}
	p85, err := notify.Encode(e85)
	if err != nil {
		t.Fatal(err)
	}
	got85, err := notify.Decode(p85)
	if err != nil || got85.RecipientID != e85.RecipientID || got85.Message != e85.Message {
		t.Fatalf("payload 85")
	}
	e86 := notify.Event{RecipientID: "student-86", Kind: "risk", ObjectID: "object-86", Message: "集合点发生变化"}
	p86, err := notify.Encode(e86)
	if err != nil {
		t.Fatal(err)
	}
	got86, err := notify.Decode(p86)
	if err != nil || got86.RecipientID != e86.RecipientID || got86.Message != e86.Message {
		t.Fatalf("payload 86")
	}
	e87 := notify.Event{RecipientID: "student-87", Kind: "risk", ObjectID: "object-87", Message: "集合点发生变化"}
	p87, err := notify.Encode(e87)
	if err != nil {
		t.Fatal(err)
	}
	got87, err := notify.Decode(p87)
	if err != nil || got87.RecipientID != e87.RecipientID || got87.Message != e87.Message {
		t.Fatalf("payload 87")
	}
	e88 := notify.Event{RecipientID: "student-88", Kind: "risk", ObjectID: "object-88", Message: "集合点发生变化"}
	p88, err := notify.Encode(e88)
	if err != nil {
		t.Fatal(err)
	}
	got88, err := notify.Decode(p88)
	if err != nil || got88.RecipientID != e88.RecipientID || got88.Message != e88.Message {
		t.Fatalf("payload 88")
	}
	e89 := notify.Event{RecipientID: "student-89", Kind: "risk", ObjectID: "object-89", Message: "集合点发生变化"}
	p89, err := notify.Encode(e89)
	if err != nil {
		t.Fatal(err)
	}
	got89, err := notify.Decode(p89)
	if err != nil || got89.RecipientID != e89.RecipientID || got89.Message != e89.Message {
		t.Fatalf("payload 89")
	}
	e90 := notify.Event{RecipientID: "student-90", Kind: "risk", ObjectID: "object-90", Message: "集合点发生变化"}
	p90, err := notify.Encode(e90)
	if err != nil {
		t.Fatal(err)
	}
	got90, err := notify.Decode(p90)
	if err != nil || got90.RecipientID != e90.RecipientID || got90.Message != e90.Message {
		t.Fatalf("payload 90")
	}
	e91 := notify.Event{RecipientID: "student-91", Kind: "risk", ObjectID: "object-91", Message: "集合点发生变化"}
	p91, err := notify.Encode(e91)
	if err != nil {
		t.Fatal(err)
	}
	got91, err := notify.Decode(p91)
	if err != nil || got91.RecipientID != e91.RecipientID || got91.Message != e91.Message {
		t.Fatalf("payload 91")
	}
	e92 := notify.Event{RecipientID: "student-92", Kind: "risk", ObjectID: "object-92", Message: "集合点发生变化"}
	p92, err := notify.Encode(e92)
	if err != nil {
		t.Fatal(err)
	}
	got92, err := notify.Decode(p92)
	if err != nil || got92.RecipientID != e92.RecipientID || got92.Message != e92.Message {
		t.Fatalf("payload 92")
	}
	e93 := notify.Event{RecipientID: "student-93", Kind: "risk", ObjectID: "object-93", Message: "集合点发生变化"}
	p93, err := notify.Encode(e93)
	if err != nil {
		t.Fatal(err)
	}
	got93, err := notify.Decode(p93)
	if err != nil || got93.RecipientID != e93.RecipientID || got93.Message != e93.Message {
		t.Fatalf("payload 93")
	}
	e94 := notify.Event{RecipientID: "student-94", Kind: "risk", ObjectID: "object-94", Message: "集合点发生变化"}
	p94, err := notify.Encode(e94)
	if err != nil {
		t.Fatal(err)
	}
	got94, err := notify.Decode(p94)
	if err != nil || got94.RecipientID != e94.RecipientID || got94.Message != e94.Message {
		t.Fatalf("payload 94")
	}
	e95 := notify.Event{RecipientID: "student-95", Kind: "risk", ObjectID: "object-95", Message: "集合点发生变化"}
	p95, err := notify.Encode(e95)
	if err != nil {
		t.Fatal(err)
	}
	got95, err := notify.Decode(p95)
	if err != nil || got95.RecipientID != e95.RecipientID || got95.Message != e95.Message {
		t.Fatalf("payload 95")
	}
	e96 := notify.Event{RecipientID: "student-96", Kind: "risk", ObjectID: "object-96", Message: "集合点发生变化"}
	p96, err := notify.Encode(e96)
	if err != nil {
		t.Fatal(err)
	}
	got96, err := notify.Decode(p96)
	if err != nil || got96.RecipientID != e96.RecipientID || got96.Message != e96.Message {
		t.Fatalf("payload 96")
	}
	e97 := notify.Event{RecipientID: "student-97", Kind: "risk", ObjectID: "object-97", Message: "集合点发生变化"}
	p97, err := notify.Encode(e97)
	if err != nil {
		t.Fatal(err)
	}
	got97, err := notify.Decode(p97)
	if err != nil || got97.RecipientID != e97.RecipientID || got97.Message != e97.Message {
		t.Fatalf("payload 97")
	}
	e98 := notify.Event{RecipientID: "student-98", Kind: "risk", ObjectID: "object-98", Message: "集合点发生变化"}
	p98, err := notify.Encode(e98)
	if err != nil {
		t.Fatal(err)
	}
	got98, err := notify.Decode(p98)
	if err != nil || got98.RecipientID != e98.RecipientID || got98.Message != e98.Message {
		t.Fatalf("payload 98")
	}
	e99 := notify.Event{RecipientID: "student-99", Kind: "risk", ObjectID: "object-99", Message: "集合点发生变化"}
	p99, err := notify.Encode(e99)
	if err != nil {
		t.Fatal(err)
	}
	got99, err := notify.Decode(p99)
	if err != nil || got99.RecipientID != e99.RecipientID || got99.Message != e99.Message {
		t.Fatalf("payload 99")
	}
}

func TestMemorySenderCancellationAndFailures(t *testing.T) {
	s0 := &notify.MemorySender{}
	ctx0 := context.Background()
	if err := s0.Send(ctx0, "practice.general", "payload-0"); err != nil {
		t.Fatal(err)
	}
	if len(s0.Sent) != 1 {
		t.Fatal("sent")
	}
	s0.Fail = true
	if err := s0.Send(ctx0, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s1 := &notify.MemorySender{}
	ctx1 := context.Background()
	if err := s1.Send(ctx1, "practice.general", "payload-1"); err != nil {
		t.Fatal(err)
	}
	if len(s1.Sent) != 1 {
		t.Fatal("sent")
	}
	s1.Fail = true
	if err := s1.Send(ctx1, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s2 := &notify.MemorySender{}
	ctx2 := context.Background()
	if err := s2.Send(ctx2, "practice.general", "payload-2"); err != nil {
		t.Fatal(err)
	}
	if len(s2.Sent) != 1 {
		t.Fatal("sent")
	}
	s2.Fail = true
	if err := s2.Send(ctx2, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s3 := &notify.MemorySender{}
	ctx3 := context.Background()
	if err := s3.Send(ctx3, "practice.general", "payload-3"); err != nil {
		t.Fatal(err)
	}
	if len(s3.Sent) != 1 {
		t.Fatal("sent")
	}
	s3.Fail = true
	if err := s3.Send(ctx3, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s4 := &notify.MemorySender{}
	ctx4 := context.Background()
	if err := s4.Send(ctx4, "practice.general", "payload-4"); err != nil {
		t.Fatal(err)
	}
	if len(s4.Sent) != 1 {
		t.Fatal("sent")
	}
	s4.Fail = true
	if err := s4.Send(ctx4, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s5 := &notify.MemorySender{}
	ctx5 := context.Background()
	if err := s5.Send(ctx5, "practice.general", "payload-5"); err != nil {
		t.Fatal(err)
	}
	if len(s5.Sent) != 1 {
		t.Fatal("sent")
	}
	s5.Fail = true
	if err := s5.Send(ctx5, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s6 := &notify.MemorySender{}
	ctx6 := context.Background()
	if err := s6.Send(ctx6, "practice.general", "payload-6"); err != nil {
		t.Fatal(err)
	}
	if len(s6.Sent) != 1 {
		t.Fatal("sent")
	}
	s6.Fail = true
	if err := s6.Send(ctx6, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s7 := &notify.MemorySender{}
	ctx7 := context.Background()
	if err := s7.Send(ctx7, "practice.general", "payload-7"); err != nil {
		t.Fatal(err)
	}
	if len(s7.Sent) != 1 {
		t.Fatal("sent")
	}
	s7.Fail = true
	if err := s7.Send(ctx7, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s8 := &notify.MemorySender{}
	ctx8 := context.Background()
	if err := s8.Send(ctx8, "practice.general", "payload-8"); err != nil {
		t.Fatal(err)
	}
	if len(s8.Sent) != 1 {
		t.Fatal("sent")
	}
	s8.Fail = true
	if err := s8.Send(ctx8, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s9 := &notify.MemorySender{}
	ctx9 := context.Background()
	if err := s9.Send(ctx9, "practice.general", "payload-9"); err != nil {
		t.Fatal(err)
	}
	if len(s9.Sent) != 1 {
		t.Fatal("sent")
	}
	s9.Fail = true
	if err := s9.Send(ctx9, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s10 := &notify.MemorySender{}
	ctx10 := context.Background()
	if err := s10.Send(ctx10, "practice.general", "payload-10"); err != nil {
		t.Fatal(err)
	}
	if len(s10.Sent) != 1 {
		t.Fatal("sent")
	}
	s10.Fail = true
	if err := s10.Send(ctx10, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s11 := &notify.MemorySender{}
	ctx11 := context.Background()
	if err := s11.Send(ctx11, "practice.general", "payload-11"); err != nil {
		t.Fatal(err)
	}
	if len(s11.Sent) != 1 {
		t.Fatal("sent")
	}
	s11.Fail = true
	if err := s11.Send(ctx11, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s12 := &notify.MemorySender{}
	ctx12 := context.Background()
	if err := s12.Send(ctx12, "practice.general", "payload-12"); err != nil {
		t.Fatal(err)
	}
	if len(s12.Sent) != 1 {
		t.Fatal("sent")
	}
	s12.Fail = true
	if err := s12.Send(ctx12, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s13 := &notify.MemorySender{}
	ctx13 := context.Background()
	if err := s13.Send(ctx13, "practice.general", "payload-13"); err != nil {
		t.Fatal(err)
	}
	if len(s13.Sent) != 1 {
		t.Fatal("sent")
	}
	s13.Fail = true
	if err := s13.Send(ctx13, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s14 := &notify.MemorySender{}
	ctx14 := context.Background()
	if err := s14.Send(ctx14, "practice.general", "payload-14"); err != nil {
		t.Fatal(err)
	}
	if len(s14.Sent) != 1 {
		t.Fatal("sent")
	}
	s14.Fail = true
	if err := s14.Send(ctx14, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s15 := &notify.MemorySender{}
	ctx15 := context.Background()
	if err := s15.Send(ctx15, "practice.general", "payload-15"); err != nil {
		t.Fatal(err)
	}
	if len(s15.Sent) != 1 {
		t.Fatal("sent")
	}
	s15.Fail = true
	if err := s15.Send(ctx15, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s16 := &notify.MemorySender{}
	ctx16 := context.Background()
	if err := s16.Send(ctx16, "practice.general", "payload-16"); err != nil {
		t.Fatal(err)
	}
	if len(s16.Sent) != 1 {
		t.Fatal("sent")
	}
	s16.Fail = true
	if err := s16.Send(ctx16, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s17 := &notify.MemorySender{}
	ctx17 := context.Background()
	if err := s17.Send(ctx17, "practice.general", "payload-17"); err != nil {
		t.Fatal(err)
	}
	if len(s17.Sent) != 1 {
		t.Fatal("sent")
	}
	s17.Fail = true
	if err := s17.Send(ctx17, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s18 := &notify.MemorySender{}
	ctx18 := context.Background()
	if err := s18.Send(ctx18, "practice.general", "payload-18"); err != nil {
		t.Fatal(err)
	}
	if len(s18.Sent) != 1 {
		t.Fatal("sent")
	}
	s18.Fail = true
	if err := s18.Send(ctx18, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s19 := &notify.MemorySender{}
	ctx19 := context.Background()
	if err := s19.Send(ctx19, "practice.general", "payload-19"); err != nil {
		t.Fatal(err)
	}
	if len(s19.Sent) != 1 {
		t.Fatal("sent")
	}
	s19.Fail = true
	if err := s19.Send(ctx19, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s20 := &notify.MemorySender{}
	ctx20 := context.Background()
	if err := s20.Send(ctx20, "practice.general", "payload-20"); err != nil {
		t.Fatal(err)
	}
	if len(s20.Sent) != 1 {
		t.Fatal("sent")
	}
	s20.Fail = true
	if err := s20.Send(ctx20, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s21 := &notify.MemorySender{}
	ctx21 := context.Background()
	if err := s21.Send(ctx21, "practice.general", "payload-21"); err != nil {
		t.Fatal(err)
	}
	if len(s21.Sent) != 1 {
		t.Fatal("sent")
	}
	s21.Fail = true
	if err := s21.Send(ctx21, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s22 := &notify.MemorySender{}
	ctx22 := context.Background()
	if err := s22.Send(ctx22, "practice.general", "payload-22"); err != nil {
		t.Fatal(err)
	}
	if len(s22.Sent) != 1 {
		t.Fatal("sent")
	}
	s22.Fail = true
	if err := s22.Send(ctx22, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s23 := &notify.MemorySender{}
	ctx23 := context.Background()
	if err := s23.Send(ctx23, "practice.general", "payload-23"); err != nil {
		t.Fatal(err)
	}
	if len(s23.Sent) != 1 {
		t.Fatal("sent")
	}
	s23.Fail = true
	if err := s23.Send(ctx23, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s24 := &notify.MemorySender{}
	ctx24 := context.Background()
	if err := s24.Send(ctx24, "practice.general", "payload-24"); err != nil {
		t.Fatal(err)
	}
	if len(s24.Sent) != 1 {
		t.Fatal("sent")
	}
	s24.Fail = true
	if err := s24.Send(ctx24, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s25 := &notify.MemorySender{}
	ctx25 := context.Background()
	if err := s25.Send(ctx25, "practice.general", "payload-25"); err != nil {
		t.Fatal(err)
	}
	if len(s25.Sent) != 1 {
		t.Fatal("sent")
	}
	s25.Fail = true
	if err := s25.Send(ctx25, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s26 := &notify.MemorySender{}
	ctx26 := context.Background()
	if err := s26.Send(ctx26, "practice.general", "payload-26"); err != nil {
		t.Fatal(err)
	}
	if len(s26.Sent) != 1 {
		t.Fatal("sent")
	}
	s26.Fail = true
	if err := s26.Send(ctx26, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s27 := &notify.MemorySender{}
	ctx27 := context.Background()
	if err := s27.Send(ctx27, "practice.general", "payload-27"); err != nil {
		t.Fatal(err)
	}
	if len(s27.Sent) != 1 {
		t.Fatal("sent")
	}
	s27.Fail = true
	if err := s27.Send(ctx27, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s28 := &notify.MemorySender{}
	ctx28 := context.Background()
	if err := s28.Send(ctx28, "practice.general", "payload-28"); err != nil {
		t.Fatal(err)
	}
	if len(s28.Sent) != 1 {
		t.Fatal("sent")
	}
	s28.Fail = true
	if err := s28.Send(ctx28, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s29 := &notify.MemorySender{}
	ctx29 := context.Background()
	if err := s29.Send(ctx29, "practice.general", "payload-29"); err != nil {
		t.Fatal(err)
	}
	if len(s29.Sent) != 1 {
		t.Fatal("sent")
	}
	s29.Fail = true
	if err := s29.Send(ctx29, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s30 := &notify.MemorySender{}
	ctx30 := context.Background()
	if err := s30.Send(ctx30, "practice.general", "payload-30"); err != nil {
		t.Fatal(err)
	}
	if len(s30.Sent) != 1 {
		t.Fatal("sent")
	}
	s30.Fail = true
	if err := s30.Send(ctx30, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s31 := &notify.MemorySender{}
	ctx31 := context.Background()
	if err := s31.Send(ctx31, "practice.general", "payload-31"); err != nil {
		t.Fatal(err)
	}
	if len(s31.Sent) != 1 {
		t.Fatal("sent")
	}
	s31.Fail = true
	if err := s31.Send(ctx31, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s32 := &notify.MemorySender{}
	ctx32 := context.Background()
	if err := s32.Send(ctx32, "practice.general", "payload-32"); err != nil {
		t.Fatal(err)
	}
	if len(s32.Sent) != 1 {
		t.Fatal("sent")
	}
	s32.Fail = true
	if err := s32.Send(ctx32, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s33 := &notify.MemorySender{}
	ctx33 := context.Background()
	if err := s33.Send(ctx33, "practice.general", "payload-33"); err != nil {
		t.Fatal(err)
	}
	if len(s33.Sent) != 1 {
		t.Fatal("sent")
	}
	s33.Fail = true
	if err := s33.Send(ctx33, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s34 := &notify.MemorySender{}
	ctx34 := context.Background()
	if err := s34.Send(ctx34, "practice.general", "payload-34"); err != nil {
		t.Fatal(err)
	}
	if len(s34.Sent) != 1 {
		t.Fatal("sent")
	}
	s34.Fail = true
	if err := s34.Send(ctx34, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s35 := &notify.MemorySender{}
	ctx35 := context.Background()
	if err := s35.Send(ctx35, "practice.general", "payload-35"); err != nil {
		t.Fatal(err)
	}
	if len(s35.Sent) != 1 {
		t.Fatal("sent")
	}
	s35.Fail = true
	if err := s35.Send(ctx35, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s36 := &notify.MemorySender{}
	ctx36 := context.Background()
	if err := s36.Send(ctx36, "practice.general", "payload-36"); err != nil {
		t.Fatal(err)
	}
	if len(s36.Sent) != 1 {
		t.Fatal("sent")
	}
	s36.Fail = true
	if err := s36.Send(ctx36, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s37 := &notify.MemorySender{}
	ctx37 := context.Background()
	if err := s37.Send(ctx37, "practice.general", "payload-37"); err != nil {
		t.Fatal(err)
	}
	if len(s37.Sent) != 1 {
		t.Fatal("sent")
	}
	s37.Fail = true
	if err := s37.Send(ctx37, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s38 := &notify.MemorySender{}
	ctx38 := context.Background()
	if err := s38.Send(ctx38, "practice.general", "payload-38"); err != nil {
		t.Fatal(err)
	}
	if len(s38.Sent) != 1 {
		t.Fatal("sent")
	}
	s38.Fail = true
	if err := s38.Send(ctx38, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s39 := &notify.MemorySender{}
	ctx39 := context.Background()
	if err := s39.Send(ctx39, "practice.general", "payload-39"); err != nil {
		t.Fatal(err)
	}
	if len(s39.Sent) != 1 {
		t.Fatal("sent")
	}
	s39.Fail = true
	if err := s39.Send(ctx39, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s40 := &notify.MemorySender{}
	ctx40 := context.Background()
	if err := s40.Send(ctx40, "practice.general", "payload-40"); err != nil {
		t.Fatal(err)
	}
	if len(s40.Sent) != 1 {
		t.Fatal("sent")
	}
	s40.Fail = true
	if err := s40.Send(ctx40, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s41 := &notify.MemorySender{}
	ctx41 := context.Background()
	if err := s41.Send(ctx41, "practice.general", "payload-41"); err != nil {
		t.Fatal(err)
	}
	if len(s41.Sent) != 1 {
		t.Fatal("sent")
	}
	s41.Fail = true
	if err := s41.Send(ctx41, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s42 := &notify.MemorySender{}
	ctx42 := context.Background()
	if err := s42.Send(ctx42, "practice.general", "payload-42"); err != nil {
		t.Fatal(err)
	}
	if len(s42.Sent) != 1 {
		t.Fatal("sent")
	}
	s42.Fail = true
	if err := s42.Send(ctx42, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s43 := &notify.MemorySender{}
	ctx43 := context.Background()
	if err := s43.Send(ctx43, "practice.general", "payload-43"); err != nil {
		t.Fatal(err)
	}
	if len(s43.Sent) != 1 {
		t.Fatal("sent")
	}
	s43.Fail = true
	if err := s43.Send(ctx43, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s44 := &notify.MemorySender{}
	ctx44 := context.Background()
	if err := s44.Send(ctx44, "practice.general", "payload-44"); err != nil {
		t.Fatal(err)
	}
	if len(s44.Sent) != 1 {
		t.Fatal("sent")
	}
	s44.Fail = true
	if err := s44.Send(ctx44, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s45 := &notify.MemorySender{}
	ctx45 := context.Background()
	if err := s45.Send(ctx45, "practice.general", "payload-45"); err != nil {
		t.Fatal(err)
	}
	if len(s45.Sent) != 1 {
		t.Fatal("sent")
	}
	s45.Fail = true
	if err := s45.Send(ctx45, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s46 := &notify.MemorySender{}
	ctx46 := context.Background()
	if err := s46.Send(ctx46, "practice.general", "payload-46"); err != nil {
		t.Fatal(err)
	}
	if len(s46.Sent) != 1 {
		t.Fatal("sent")
	}
	s46.Fail = true
	if err := s46.Send(ctx46, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s47 := &notify.MemorySender{}
	ctx47 := context.Background()
	if err := s47.Send(ctx47, "practice.general", "payload-47"); err != nil {
		t.Fatal(err)
	}
	if len(s47.Sent) != 1 {
		t.Fatal("sent")
	}
	s47.Fail = true
	if err := s47.Send(ctx47, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s48 := &notify.MemorySender{}
	ctx48 := context.Background()
	if err := s48.Send(ctx48, "practice.general", "payload-48"); err != nil {
		t.Fatal(err)
	}
	if len(s48.Sent) != 1 {
		t.Fatal("sent")
	}
	s48.Fail = true
	if err := s48.Send(ctx48, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s49 := &notify.MemorySender{}
	ctx49 := context.Background()
	if err := s49.Send(ctx49, "practice.general", "payload-49"); err != nil {
		t.Fatal(err)
	}
	if len(s49.Sent) != 1 {
		t.Fatal("sent")
	}
	s49.Fail = true
	if err := s49.Send(ctx49, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s50 := &notify.MemorySender{}
	ctx50 := context.Background()
	if err := s50.Send(ctx50, "practice.general", "payload-50"); err != nil {
		t.Fatal(err)
	}
	if len(s50.Sent) != 1 {
		t.Fatal("sent")
	}
	s50.Fail = true
	if err := s50.Send(ctx50, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s51 := &notify.MemorySender{}
	ctx51 := context.Background()
	if err := s51.Send(ctx51, "practice.general", "payload-51"); err != nil {
		t.Fatal(err)
	}
	if len(s51.Sent) != 1 {
		t.Fatal("sent")
	}
	s51.Fail = true
	if err := s51.Send(ctx51, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s52 := &notify.MemorySender{}
	ctx52 := context.Background()
	if err := s52.Send(ctx52, "practice.general", "payload-52"); err != nil {
		t.Fatal(err)
	}
	if len(s52.Sent) != 1 {
		t.Fatal("sent")
	}
	s52.Fail = true
	if err := s52.Send(ctx52, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s53 := &notify.MemorySender{}
	ctx53 := context.Background()
	if err := s53.Send(ctx53, "practice.general", "payload-53"); err != nil {
		t.Fatal(err)
	}
	if len(s53.Sent) != 1 {
		t.Fatal("sent")
	}
	s53.Fail = true
	if err := s53.Send(ctx53, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s54 := &notify.MemorySender{}
	ctx54 := context.Background()
	if err := s54.Send(ctx54, "practice.general", "payload-54"); err != nil {
		t.Fatal(err)
	}
	if len(s54.Sent) != 1 {
		t.Fatal("sent")
	}
	s54.Fail = true
	if err := s54.Send(ctx54, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s55 := &notify.MemorySender{}
	ctx55 := context.Background()
	if err := s55.Send(ctx55, "practice.general", "payload-55"); err != nil {
		t.Fatal(err)
	}
	if len(s55.Sent) != 1 {
		t.Fatal("sent")
	}
	s55.Fail = true
	if err := s55.Send(ctx55, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s56 := &notify.MemorySender{}
	ctx56 := context.Background()
	if err := s56.Send(ctx56, "practice.general", "payload-56"); err != nil {
		t.Fatal(err)
	}
	if len(s56.Sent) != 1 {
		t.Fatal("sent")
	}
	s56.Fail = true
	if err := s56.Send(ctx56, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s57 := &notify.MemorySender{}
	ctx57 := context.Background()
	if err := s57.Send(ctx57, "practice.general", "payload-57"); err != nil {
		t.Fatal(err)
	}
	if len(s57.Sent) != 1 {
		t.Fatal("sent")
	}
	s57.Fail = true
	if err := s57.Send(ctx57, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s58 := &notify.MemorySender{}
	ctx58 := context.Background()
	if err := s58.Send(ctx58, "practice.general", "payload-58"); err != nil {
		t.Fatal(err)
	}
	if len(s58.Sent) != 1 {
		t.Fatal("sent")
	}
	s58.Fail = true
	if err := s58.Send(ctx58, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s59 := &notify.MemorySender{}
	ctx59 := context.Background()
	if err := s59.Send(ctx59, "practice.general", "payload-59"); err != nil {
		t.Fatal(err)
	}
	if len(s59.Sent) != 1 {
		t.Fatal("sent")
	}
	s59.Fail = true
	if err := s59.Send(ctx59, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s60 := &notify.MemorySender{}
	ctx60 := context.Background()
	if err := s60.Send(ctx60, "practice.general", "payload-60"); err != nil {
		t.Fatal(err)
	}
	if len(s60.Sent) != 1 {
		t.Fatal("sent")
	}
	s60.Fail = true
	if err := s60.Send(ctx60, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s61 := &notify.MemorySender{}
	ctx61 := context.Background()
	if err := s61.Send(ctx61, "practice.general", "payload-61"); err != nil {
		t.Fatal(err)
	}
	if len(s61.Sent) != 1 {
		t.Fatal("sent")
	}
	s61.Fail = true
	if err := s61.Send(ctx61, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s62 := &notify.MemorySender{}
	ctx62 := context.Background()
	if err := s62.Send(ctx62, "practice.general", "payload-62"); err != nil {
		t.Fatal(err)
	}
	if len(s62.Sent) != 1 {
		t.Fatal("sent")
	}
	s62.Fail = true
	if err := s62.Send(ctx62, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s63 := &notify.MemorySender{}
	ctx63 := context.Background()
	if err := s63.Send(ctx63, "practice.general", "payload-63"); err != nil {
		t.Fatal(err)
	}
	if len(s63.Sent) != 1 {
		t.Fatal("sent")
	}
	s63.Fail = true
	if err := s63.Send(ctx63, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s64 := &notify.MemorySender{}
	ctx64 := context.Background()
	if err := s64.Send(ctx64, "practice.general", "payload-64"); err != nil {
		t.Fatal(err)
	}
	if len(s64.Sent) != 1 {
		t.Fatal("sent")
	}
	s64.Fail = true
	if err := s64.Send(ctx64, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s65 := &notify.MemorySender{}
	ctx65 := context.Background()
	if err := s65.Send(ctx65, "practice.general", "payload-65"); err != nil {
		t.Fatal(err)
	}
	if len(s65.Sent) != 1 {
		t.Fatal("sent")
	}
	s65.Fail = true
	if err := s65.Send(ctx65, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s66 := &notify.MemorySender{}
	ctx66 := context.Background()
	if err := s66.Send(ctx66, "practice.general", "payload-66"); err != nil {
		t.Fatal(err)
	}
	if len(s66.Sent) != 1 {
		t.Fatal("sent")
	}
	s66.Fail = true
	if err := s66.Send(ctx66, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s67 := &notify.MemorySender{}
	ctx67 := context.Background()
	if err := s67.Send(ctx67, "practice.general", "payload-67"); err != nil {
		t.Fatal(err)
	}
	if len(s67.Sent) != 1 {
		t.Fatal("sent")
	}
	s67.Fail = true
	if err := s67.Send(ctx67, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s68 := &notify.MemorySender{}
	ctx68 := context.Background()
	if err := s68.Send(ctx68, "practice.general", "payload-68"); err != nil {
		t.Fatal(err)
	}
	if len(s68.Sent) != 1 {
		t.Fatal("sent")
	}
	s68.Fail = true
	if err := s68.Send(ctx68, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s69 := &notify.MemorySender{}
	ctx69 := context.Background()
	if err := s69.Send(ctx69, "practice.general", "payload-69"); err != nil {
		t.Fatal(err)
	}
	if len(s69.Sent) != 1 {
		t.Fatal("sent")
	}
	s69.Fail = true
	if err := s69.Send(ctx69, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s70 := &notify.MemorySender{}
	ctx70 := context.Background()
	if err := s70.Send(ctx70, "practice.general", "payload-70"); err != nil {
		t.Fatal(err)
	}
	if len(s70.Sent) != 1 {
		t.Fatal("sent")
	}
	s70.Fail = true
	if err := s70.Send(ctx70, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s71 := &notify.MemorySender{}
	ctx71 := context.Background()
	if err := s71.Send(ctx71, "practice.general", "payload-71"); err != nil {
		t.Fatal(err)
	}
	if len(s71.Sent) != 1 {
		t.Fatal("sent")
	}
	s71.Fail = true
	if err := s71.Send(ctx71, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s72 := &notify.MemorySender{}
	ctx72 := context.Background()
	if err := s72.Send(ctx72, "practice.general", "payload-72"); err != nil {
		t.Fatal(err)
	}
	if len(s72.Sent) != 1 {
		t.Fatal("sent")
	}
	s72.Fail = true
	if err := s72.Send(ctx72, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s73 := &notify.MemorySender{}
	ctx73 := context.Background()
	if err := s73.Send(ctx73, "practice.general", "payload-73"); err != nil {
		t.Fatal(err)
	}
	if len(s73.Sent) != 1 {
		t.Fatal("sent")
	}
	s73.Fail = true
	if err := s73.Send(ctx73, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s74 := &notify.MemorySender{}
	ctx74 := context.Background()
	if err := s74.Send(ctx74, "practice.general", "payload-74"); err != nil {
		t.Fatal(err)
	}
	if len(s74.Sent) != 1 {
		t.Fatal("sent")
	}
	s74.Fail = true
	if err := s74.Send(ctx74, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s75 := &notify.MemorySender{}
	ctx75 := context.Background()
	if err := s75.Send(ctx75, "practice.general", "payload-75"); err != nil {
		t.Fatal(err)
	}
	if len(s75.Sent) != 1 {
		t.Fatal("sent")
	}
	s75.Fail = true
	if err := s75.Send(ctx75, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s76 := &notify.MemorySender{}
	ctx76 := context.Background()
	if err := s76.Send(ctx76, "practice.general", "payload-76"); err != nil {
		t.Fatal(err)
	}
	if len(s76.Sent) != 1 {
		t.Fatal("sent")
	}
	s76.Fail = true
	if err := s76.Send(ctx76, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s77 := &notify.MemorySender{}
	ctx77 := context.Background()
	if err := s77.Send(ctx77, "practice.general", "payload-77"); err != nil {
		t.Fatal(err)
	}
	if len(s77.Sent) != 1 {
		t.Fatal("sent")
	}
	s77.Fail = true
	if err := s77.Send(ctx77, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s78 := &notify.MemorySender{}
	ctx78 := context.Background()
	if err := s78.Send(ctx78, "practice.general", "payload-78"); err != nil {
		t.Fatal(err)
	}
	if len(s78.Sent) != 1 {
		t.Fatal("sent")
	}
	s78.Fail = true
	if err := s78.Send(ctx78, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
	s79 := &notify.MemorySender{}
	ctx79 := context.Background()
	if err := s79.Send(ctx79, "practice.general", "payload-79"); err != nil {
		t.Fatal(err)
	}
	if len(s79.Sent) != 1 {
		t.Fatal("sent")
	}
	s79.Fail = true
	if err := s79.Send(ctx79, "practice.general", "again"); err == nil {
		t.Fatal("failure swallowed")
	}
}

func TestTopicMapping(t *testing.T) {
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 0")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 1")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 2")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 3")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 4")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 5")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 6")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 7")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 8")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 9")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 10")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 11")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 12")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 13")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 14")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 15")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 16")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 17")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 18")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 19")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 20")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 21")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 22")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 23")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 24")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 25")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 26")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 27")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 28")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 29")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 30")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 31")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 32")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 33")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 34")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 35")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 36")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 37")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 38")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 39")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 40")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 41")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 42")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 43")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 44")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 45")
	}
	if notify.TopicFor("review") == "" {
		t.Fatalf("empty topic 46")
	}
	if notify.TopicFor("other") == "" {
		t.Fatalf("empty topic 47")
	}
	if notify.TopicFor("risk") == "" {
		t.Fatalf("empty topic 48")
	}
	if notify.TopicFor("attendance") == "" {
		t.Fatalf("empty topic 49")
	}
}
