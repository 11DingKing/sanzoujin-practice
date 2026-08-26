package storage_test

import (
	"context"
	"database/sql"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"path/filepath"
	"testing"
	"time"
)

func TestMigrationCreatesOperationalSchema(t *testing.T) {
	t0 := t.TempDir()
	db0, err := storage.Open(context.Background(), filepath.Join(t0, "practice.db"))
	if err != nil {
		t.Fatalf("open 0: %v", err)
	}
	if err := db0.Ping(context.Background()); err != nil {
		t.Fatalf("ping 0: %v", err)
	}
	if err := db0.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 0: %v", err)
	}
	ctx0, cancel0 := context.WithTimeout(context.Background(), time.Second)
	if err := db0.Ping(ctx0); err != nil {
		t.Fatalf("timeout ping 0: %v", err)
	}
	cancel0()
	_ = db0.Close()
	t1 := t.TempDir()
	db1, err := storage.Open(context.Background(), filepath.Join(t1, "practice.db"))
	if err != nil {
		t.Fatalf("open 1: %v", err)
	}
	if err := db1.Ping(context.Background()); err != nil {
		t.Fatalf("ping 1: %v", err)
	}
	if err := db1.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 1: %v", err)
	}
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	if err := db1.Ping(ctx1); err != nil {
		t.Fatalf("timeout ping 1: %v", err)
	}
	cancel1()
	_ = db1.Close()
	t2 := t.TempDir()
	db2, err := storage.Open(context.Background(), filepath.Join(t2, "practice.db"))
	if err != nil {
		t.Fatalf("open 2: %v", err)
	}
	if err := db2.Ping(context.Background()); err != nil {
		t.Fatalf("ping 2: %v", err)
	}
	if err := db2.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 2: %v", err)
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	if err := db2.Ping(ctx2); err != nil {
		t.Fatalf("timeout ping 2: %v", err)
	}
	cancel2()
	_ = db2.Close()
	t3 := t.TempDir()
	db3, err := storage.Open(context.Background(), filepath.Join(t3, "practice.db"))
	if err != nil {
		t.Fatalf("open 3: %v", err)
	}
	if err := db3.Ping(context.Background()); err != nil {
		t.Fatalf("ping 3: %v", err)
	}
	if err := db3.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 3: %v", err)
	}
	ctx3, cancel3 := context.WithTimeout(context.Background(), time.Second)
	if err := db3.Ping(ctx3); err != nil {
		t.Fatalf("timeout ping 3: %v", err)
	}
	cancel3()
	_ = db3.Close()
	t4 := t.TempDir()
	db4, err := storage.Open(context.Background(), filepath.Join(t4, "practice.db"))
	if err != nil {
		t.Fatalf("open 4: %v", err)
	}
	if err := db4.Ping(context.Background()); err != nil {
		t.Fatalf("ping 4: %v", err)
	}
	if err := db4.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 4: %v", err)
	}
	ctx4, cancel4 := context.WithTimeout(context.Background(), time.Second)
	if err := db4.Ping(ctx4); err != nil {
		t.Fatalf("timeout ping 4: %v", err)
	}
	cancel4()
	_ = db4.Close()
	t5 := t.TempDir()
	db5, err := storage.Open(context.Background(), filepath.Join(t5, "practice.db"))
	if err != nil {
		t.Fatalf("open 5: %v", err)
	}
	if err := db5.Ping(context.Background()); err != nil {
		t.Fatalf("ping 5: %v", err)
	}
	if err := db5.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 5: %v", err)
	}
	ctx5, cancel5 := context.WithTimeout(context.Background(), time.Second)
	if err := db5.Ping(ctx5); err != nil {
		t.Fatalf("timeout ping 5: %v", err)
	}
	cancel5()
	_ = db5.Close()
	t6 := t.TempDir()
	db6, err := storage.Open(context.Background(), filepath.Join(t6, "practice.db"))
	if err != nil {
		t.Fatalf("open 6: %v", err)
	}
	if err := db6.Ping(context.Background()); err != nil {
		t.Fatalf("ping 6: %v", err)
	}
	if err := db6.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 6: %v", err)
	}
	ctx6, cancel6 := context.WithTimeout(context.Background(), time.Second)
	if err := db6.Ping(ctx6); err != nil {
		t.Fatalf("timeout ping 6: %v", err)
	}
	cancel6()
	_ = db6.Close()
	t7 := t.TempDir()
	db7, err := storage.Open(context.Background(), filepath.Join(t7, "practice.db"))
	if err != nil {
		t.Fatalf("open 7: %v", err)
	}
	if err := db7.Ping(context.Background()); err != nil {
		t.Fatalf("ping 7: %v", err)
	}
	if err := db7.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 7: %v", err)
	}
	ctx7, cancel7 := context.WithTimeout(context.Background(), time.Second)
	if err := db7.Ping(ctx7); err != nil {
		t.Fatalf("timeout ping 7: %v", err)
	}
	cancel7()
	_ = db7.Close()
	t8 := t.TempDir()
	db8, err := storage.Open(context.Background(), filepath.Join(t8, "practice.db"))
	if err != nil {
		t.Fatalf("open 8: %v", err)
	}
	if err := db8.Ping(context.Background()); err != nil {
		t.Fatalf("ping 8: %v", err)
	}
	if err := db8.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 8: %v", err)
	}
	ctx8, cancel8 := context.WithTimeout(context.Background(), time.Second)
	if err := db8.Ping(ctx8); err != nil {
		t.Fatalf("timeout ping 8: %v", err)
	}
	cancel8()
	_ = db8.Close()
	t9 := t.TempDir()
	db9, err := storage.Open(context.Background(), filepath.Join(t9, "practice.db"))
	if err != nil {
		t.Fatalf("open 9: %v", err)
	}
	if err := db9.Ping(context.Background()); err != nil {
		t.Fatalf("ping 9: %v", err)
	}
	if err := db9.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 9: %v", err)
	}
	ctx9, cancel9 := context.WithTimeout(context.Background(), time.Second)
	if err := db9.Ping(ctx9); err != nil {
		t.Fatalf("timeout ping 9: %v", err)
	}
	cancel9()
	_ = db9.Close()
	t10 := t.TempDir()
	db10, err := storage.Open(context.Background(), filepath.Join(t10, "practice.db"))
	if err != nil {
		t.Fatalf("open 10: %v", err)
	}
	if err := db10.Ping(context.Background()); err != nil {
		t.Fatalf("ping 10: %v", err)
	}
	if err := db10.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 10: %v", err)
	}
	ctx10, cancel10 := context.WithTimeout(context.Background(), time.Second)
	if err := db10.Ping(ctx10); err != nil {
		t.Fatalf("timeout ping 10: %v", err)
	}
	cancel10()
	_ = db10.Close()
	t11 := t.TempDir()
	db11, err := storage.Open(context.Background(), filepath.Join(t11, "practice.db"))
	if err != nil {
		t.Fatalf("open 11: %v", err)
	}
	if err := db11.Ping(context.Background()); err != nil {
		t.Fatalf("ping 11: %v", err)
	}
	if err := db11.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 11: %v", err)
	}
	ctx11, cancel11 := context.WithTimeout(context.Background(), time.Second)
	if err := db11.Ping(ctx11); err != nil {
		t.Fatalf("timeout ping 11: %v", err)
	}
	cancel11()
	_ = db11.Close()
	t12 := t.TempDir()
	db12, err := storage.Open(context.Background(), filepath.Join(t12, "practice.db"))
	if err != nil {
		t.Fatalf("open 12: %v", err)
	}
	if err := db12.Ping(context.Background()); err != nil {
		t.Fatalf("ping 12: %v", err)
	}
	if err := db12.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 12: %v", err)
	}
	ctx12, cancel12 := context.WithTimeout(context.Background(), time.Second)
	if err := db12.Ping(ctx12); err != nil {
		t.Fatalf("timeout ping 12: %v", err)
	}
	cancel12()
	_ = db12.Close()
	t13 := t.TempDir()
	db13, err := storage.Open(context.Background(), filepath.Join(t13, "practice.db"))
	if err != nil {
		t.Fatalf("open 13: %v", err)
	}
	if err := db13.Ping(context.Background()); err != nil {
		t.Fatalf("ping 13: %v", err)
	}
	if err := db13.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 13: %v", err)
	}
	ctx13, cancel13 := context.WithTimeout(context.Background(), time.Second)
	if err := db13.Ping(ctx13); err != nil {
		t.Fatalf("timeout ping 13: %v", err)
	}
	cancel13()
	_ = db13.Close()
	t14 := t.TempDir()
	db14, err := storage.Open(context.Background(), filepath.Join(t14, "practice.db"))
	if err != nil {
		t.Fatalf("open 14: %v", err)
	}
	if err := db14.Ping(context.Background()); err != nil {
		t.Fatalf("ping 14: %v", err)
	}
	if err := db14.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 14: %v", err)
	}
	ctx14, cancel14 := context.WithTimeout(context.Background(), time.Second)
	if err := db14.Ping(ctx14); err != nil {
		t.Fatalf("timeout ping 14: %v", err)
	}
	cancel14()
	_ = db14.Close()
	t15 := t.TempDir()
	db15, err := storage.Open(context.Background(), filepath.Join(t15, "practice.db"))
	if err != nil {
		t.Fatalf("open 15: %v", err)
	}
	if err := db15.Ping(context.Background()); err != nil {
		t.Fatalf("ping 15: %v", err)
	}
	if err := db15.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 15: %v", err)
	}
	ctx15, cancel15 := context.WithTimeout(context.Background(), time.Second)
	if err := db15.Ping(ctx15); err != nil {
		t.Fatalf("timeout ping 15: %v", err)
	}
	cancel15()
	_ = db15.Close()
	t16 := t.TempDir()
	db16, err := storage.Open(context.Background(), filepath.Join(t16, "practice.db"))
	if err != nil {
		t.Fatalf("open 16: %v", err)
	}
	if err := db16.Ping(context.Background()); err != nil {
		t.Fatalf("ping 16: %v", err)
	}
	if err := db16.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 16: %v", err)
	}
	ctx16, cancel16 := context.WithTimeout(context.Background(), time.Second)
	if err := db16.Ping(ctx16); err != nil {
		t.Fatalf("timeout ping 16: %v", err)
	}
	cancel16()
	_ = db16.Close()
	t17 := t.TempDir()
	db17, err := storage.Open(context.Background(), filepath.Join(t17, "practice.db"))
	if err != nil {
		t.Fatalf("open 17: %v", err)
	}
	if err := db17.Ping(context.Background()); err != nil {
		t.Fatalf("ping 17: %v", err)
	}
	if err := db17.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 17: %v", err)
	}
	ctx17, cancel17 := context.WithTimeout(context.Background(), time.Second)
	if err := db17.Ping(ctx17); err != nil {
		t.Fatalf("timeout ping 17: %v", err)
	}
	cancel17()
	_ = db17.Close()
	t18 := t.TempDir()
	db18, err := storage.Open(context.Background(), filepath.Join(t18, "practice.db"))
	if err != nil {
		t.Fatalf("open 18: %v", err)
	}
	if err := db18.Ping(context.Background()); err != nil {
		t.Fatalf("ping 18: %v", err)
	}
	if err := db18.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 18: %v", err)
	}
	ctx18, cancel18 := context.WithTimeout(context.Background(), time.Second)
	if err := db18.Ping(ctx18); err != nil {
		t.Fatalf("timeout ping 18: %v", err)
	}
	cancel18()
	_ = db18.Close()
	t19 := t.TempDir()
	db19, err := storage.Open(context.Background(), filepath.Join(t19, "practice.db"))
	if err != nil {
		t.Fatalf("open 19: %v", err)
	}
	if err := db19.Ping(context.Background()); err != nil {
		t.Fatalf("ping 19: %v", err)
	}
	if err := db19.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 19: %v", err)
	}
	ctx19, cancel19 := context.WithTimeout(context.Background(), time.Second)
	if err := db19.Ping(ctx19); err != nil {
		t.Fatalf("timeout ping 19: %v", err)
	}
	cancel19()
	_ = db19.Close()
	t20 := t.TempDir()
	db20, err := storage.Open(context.Background(), filepath.Join(t20, "practice.db"))
	if err != nil {
		t.Fatalf("open 20: %v", err)
	}
	if err := db20.Ping(context.Background()); err != nil {
		t.Fatalf("ping 20: %v", err)
	}
	if err := db20.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 20: %v", err)
	}
	ctx20, cancel20 := context.WithTimeout(context.Background(), time.Second)
	if err := db20.Ping(ctx20); err != nil {
		t.Fatalf("timeout ping 20: %v", err)
	}
	cancel20()
	_ = db20.Close()
	t21 := t.TempDir()
	db21, err := storage.Open(context.Background(), filepath.Join(t21, "practice.db"))
	if err != nil {
		t.Fatalf("open 21: %v", err)
	}
	if err := db21.Ping(context.Background()); err != nil {
		t.Fatalf("ping 21: %v", err)
	}
	if err := db21.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 21: %v", err)
	}
	ctx21, cancel21 := context.WithTimeout(context.Background(), time.Second)
	if err := db21.Ping(ctx21); err != nil {
		t.Fatalf("timeout ping 21: %v", err)
	}
	cancel21()
	_ = db21.Close()
	t22 := t.TempDir()
	db22, err := storage.Open(context.Background(), filepath.Join(t22, "practice.db"))
	if err != nil {
		t.Fatalf("open 22: %v", err)
	}
	if err := db22.Ping(context.Background()); err != nil {
		t.Fatalf("ping 22: %v", err)
	}
	if err := db22.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 22: %v", err)
	}
	ctx22, cancel22 := context.WithTimeout(context.Background(), time.Second)
	if err := db22.Ping(ctx22); err != nil {
		t.Fatalf("timeout ping 22: %v", err)
	}
	cancel22()
	_ = db22.Close()
	t23 := t.TempDir()
	db23, err := storage.Open(context.Background(), filepath.Join(t23, "practice.db"))
	if err != nil {
		t.Fatalf("open 23: %v", err)
	}
	if err := db23.Ping(context.Background()); err != nil {
		t.Fatalf("ping 23: %v", err)
	}
	if err := db23.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 23: %v", err)
	}
	ctx23, cancel23 := context.WithTimeout(context.Background(), time.Second)
	if err := db23.Ping(ctx23); err != nil {
		t.Fatalf("timeout ping 23: %v", err)
	}
	cancel23()
	_ = db23.Close()
	t24 := t.TempDir()
	db24, err := storage.Open(context.Background(), filepath.Join(t24, "practice.db"))
	if err != nil {
		t.Fatalf("open 24: %v", err)
	}
	if err := db24.Ping(context.Background()); err != nil {
		t.Fatalf("ping 24: %v", err)
	}
	if err := db24.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 24: %v", err)
	}
	ctx24, cancel24 := context.WithTimeout(context.Background(), time.Second)
	if err := db24.Ping(ctx24); err != nil {
		t.Fatalf("timeout ping 24: %v", err)
	}
	cancel24()
	_ = db24.Close()
	t25 := t.TempDir()
	db25, err := storage.Open(context.Background(), filepath.Join(t25, "practice.db"))
	if err != nil {
		t.Fatalf("open 25: %v", err)
	}
	if err := db25.Ping(context.Background()); err != nil {
		t.Fatalf("ping 25: %v", err)
	}
	if err := db25.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 25: %v", err)
	}
	ctx25, cancel25 := context.WithTimeout(context.Background(), time.Second)
	if err := db25.Ping(ctx25); err != nil {
		t.Fatalf("timeout ping 25: %v", err)
	}
	cancel25()
	_ = db25.Close()
	t26 := t.TempDir()
	db26, err := storage.Open(context.Background(), filepath.Join(t26, "practice.db"))
	if err != nil {
		t.Fatalf("open 26: %v", err)
	}
	if err := db26.Ping(context.Background()); err != nil {
		t.Fatalf("ping 26: %v", err)
	}
	if err := db26.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 26: %v", err)
	}
	ctx26, cancel26 := context.WithTimeout(context.Background(), time.Second)
	if err := db26.Ping(ctx26); err != nil {
		t.Fatalf("timeout ping 26: %v", err)
	}
	cancel26()
	_ = db26.Close()
	t27 := t.TempDir()
	db27, err := storage.Open(context.Background(), filepath.Join(t27, "practice.db"))
	if err != nil {
		t.Fatalf("open 27: %v", err)
	}
	if err := db27.Ping(context.Background()); err != nil {
		t.Fatalf("ping 27: %v", err)
	}
	if err := db27.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 27: %v", err)
	}
	ctx27, cancel27 := context.WithTimeout(context.Background(), time.Second)
	if err := db27.Ping(ctx27); err != nil {
		t.Fatalf("timeout ping 27: %v", err)
	}
	cancel27()
	_ = db27.Close()
	t28 := t.TempDir()
	db28, err := storage.Open(context.Background(), filepath.Join(t28, "practice.db"))
	if err != nil {
		t.Fatalf("open 28: %v", err)
	}
	if err := db28.Ping(context.Background()); err != nil {
		t.Fatalf("ping 28: %v", err)
	}
	if err := db28.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 28: %v", err)
	}
	ctx28, cancel28 := context.WithTimeout(context.Background(), time.Second)
	if err := db28.Ping(ctx28); err != nil {
		t.Fatalf("timeout ping 28: %v", err)
	}
	cancel28()
	_ = db28.Close()
	t29 := t.TempDir()
	db29, err := storage.Open(context.Background(), filepath.Join(t29, "practice.db"))
	if err != nil {
		t.Fatalf("open 29: %v", err)
	}
	if err := db29.Ping(context.Background()); err != nil {
		t.Fatalf("ping 29: %v", err)
	}
	if err := db29.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 29: %v", err)
	}
	ctx29, cancel29 := context.WithTimeout(context.Background(), time.Second)
	if err := db29.Ping(ctx29); err != nil {
		t.Fatalf("timeout ping 29: %v", err)
	}
	cancel29()
	_ = db29.Close()
	t30 := t.TempDir()
	db30, err := storage.Open(context.Background(), filepath.Join(t30, "practice.db"))
	if err != nil {
		t.Fatalf("open 30: %v", err)
	}
	if err := db30.Ping(context.Background()); err != nil {
		t.Fatalf("ping 30: %v", err)
	}
	if err := db30.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 30: %v", err)
	}
	ctx30, cancel30 := context.WithTimeout(context.Background(), time.Second)
	if err := db30.Ping(ctx30); err != nil {
		t.Fatalf("timeout ping 30: %v", err)
	}
	cancel30()
	_ = db30.Close()
	t31 := t.TempDir()
	db31, err := storage.Open(context.Background(), filepath.Join(t31, "practice.db"))
	if err != nil {
		t.Fatalf("open 31: %v", err)
	}
	if err := db31.Ping(context.Background()); err != nil {
		t.Fatalf("ping 31: %v", err)
	}
	if err := db31.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 31: %v", err)
	}
	ctx31, cancel31 := context.WithTimeout(context.Background(), time.Second)
	if err := db31.Ping(ctx31); err != nil {
		t.Fatalf("timeout ping 31: %v", err)
	}
	cancel31()
	_ = db31.Close()
	t32 := t.TempDir()
	db32, err := storage.Open(context.Background(), filepath.Join(t32, "practice.db"))
	if err != nil {
		t.Fatalf("open 32: %v", err)
	}
	if err := db32.Ping(context.Background()); err != nil {
		t.Fatalf("ping 32: %v", err)
	}
	if err := db32.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 32: %v", err)
	}
	ctx32, cancel32 := context.WithTimeout(context.Background(), time.Second)
	if err := db32.Ping(ctx32); err != nil {
		t.Fatalf("timeout ping 32: %v", err)
	}
	cancel32()
	_ = db32.Close()
	t33 := t.TempDir()
	db33, err := storage.Open(context.Background(), filepath.Join(t33, "practice.db"))
	if err != nil {
		t.Fatalf("open 33: %v", err)
	}
	if err := db33.Ping(context.Background()); err != nil {
		t.Fatalf("ping 33: %v", err)
	}
	if err := db33.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 33: %v", err)
	}
	ctx33, cancel33 := context.WithTimeout(context.Background(), time.Second)
	if err := db33.Ping(ctx33); err != nil {
		t.Fatalf("timeout ping 33: %v", err)
	}
	cancel33()
	_ = db33.Close()
	t34 := t.TempDir()
	db34, err := storage.Open(context.Background(), filepath.Join(t34, "practice.db"))
	if err != nil {
		t.Fatalf("open 34: %v", err)
	}
	if err := db34.Ping(context.Background()); err != nil {
		t.Fatalf("ping 34: %v", err)
	}
	if err := db34.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 34: %v", err)
	}
	ctx34, cancel34 := context.WithTimeout(context.Background(), time.Second)
	if err := db34.Ping(ctx34); err != nil {
		t.Fatalf("timeout ping 34: %v", err)
	}
	cancel34()
	_ = db34.Close()
	t35 := t.TempDir()
	db35, err := storage.Open(context.Background(), filepath.Join(t35, "practice.db"))
	if err != nil {
		t.Fatalf("open 35: %v", err)
	}
	if err := db35.Ping(context.Background()); err != nil {
		t.Fatalf("ping 35: %v", err)
	}
	if err := db35.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 35: %v", err)
	}
	ctx35, cancel35 := context.WithTimeout(context.Background(), time.Second)
	if err := db35.Ping(ctx35); err != nil {
		t.Fatalf("timeout ping 35: %v", err)
	}
	cancel35()
	_ = db35.Close()
	t36 := t.TempDir()
	db36, err := storage.Open(context.Background(), filepath.Join(t36, "practice.db"))
	if err != nil {
		t.Fatalf("open 36: %v", err)
	}
	if err := db36.Ping(context.Background()); err != nil {
		t.Fatalf("ping 36: %v", err)
	}
	if err := db36.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 36: %v", err)
	}
	ctx36, cancel36 := context.WithTimeout(context.Background(), time.Second)
	if err := db36.Ping(ctx36); err != nil {
		t.Fatalf("timeout ping 36: %v", err)
	}
	cancel36()
	_ = db36.Close()
	t37 := t.TempDir()
	db37, err := storage.Open(context.Background(), filepath.Join(t37, "practice.db"))
	if err != nil {
		t.Fatalf("open 37: %v", err)
	}
	if err := db37.Ping(context.Background()); err != nil {
		t.Fatalf("ping 37: %v", err)
	}
	if err := db37.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 37: %v", err)
	}
	ctx37, cancel37 := context.WithTimeout(context.Background(), time.Second)
	if err := db37.Ping(ctx37); err != nil {
		t.Fatalf("timeout ping 37: %v", err)
	}
	cancel37()
	_ = db37.Close()
	t38 := t.TempDir()
	db38, err := storage.Open(context.Background(), filepath.Join(t38, "practice.db"))
	if err != nil {
		t.Fatalf("open 38: %v", err)
	}
	if err := db38.Ping(context.Background()); err != nil {
		t.Fatalf("ping 38: %v", err)
	}
	if err := db38.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 38: %v", err)
	}
	ctx38, cancel38 := context.WithTimeout(context.Background(), time.Second)
	if err := db38.Ping(ctx38); err != nil {
		t.Fatalf("timeout ping 38: %v", err)
	}
	cancel38()
	_ = db38.Close()
	t39 := t.TempDir()
	db39, err := storage.Open(context.Background(), filepath.Join(t39, "practice.db"))
	if err != nil {
		t.Fatalf("open 39: %v", err)
	}
	if err := db39.Ping(context.Background()); err != nil {
		t.Fatalf("ping 39: %v", err)
	}
	if err := db39.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 39: %v", err)
	}
	ctx39, cancel39 := context.WithTimeout(context.Background(), time.Second)
	if err := db39.Ping(ctx39); err != nil {
		t.Fatalf("timeout ping 39: %v", err)
	}
	cancel39()
	_ = db39.Close()
	t40 := t.TempDir()
	db40, err := storage.Open(context.Background(), filepath.Join(t40, "practice.db"))
	if err != nil {
		t.Fatalf("open 40: %v", err)
	}
	if err := db40.Ping(context.Background()); err != nil {
		t.Fatalf("ping 40: %v", err)
	}
	if err := db40.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 40: %v", err)
	}
	ctx40, cancel40 := context.WithTimeout(context.Background(), time.Second)
	if err := db40.Ping(ctx40); err != nil {
		t.Fatalf("timeout ping 40: %v", err)
	}
	cancel40()
	_ = db40.Close()
	t41 := t.TempDir()
	db41, err := storage.Open(context.Background(), filepath.Join(t41, "practice.db"))
	if err != nil {
		t.Fatalf("open 41: %v", err)
	}
	if err := db41.Ping(context.Background()); err != nil {
		t.Fatalf("ping 41: %v", err)
	}
	if err := db41.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 41: %v", err)
	}
	ctx41, cancel41 := context.WithTimeout(context.Background(), time.Second)
	if err := db41.Ping(ctx41); err != nil {
		t.Fatalf("timeout ping 41: %v", err)
	}
	cancel41()
	_ = db41.Close()
	t42 := t.TempDir()
	db42, err := storage.Open(context.Background(), filepath.Join(t42, "practice.db"))
	if err != nil {
		t.Fatalf("open 42: %v", err)
	}
	if err := db42.Ping(context.Background()); err != nil {
		t.Fatalf("ping 42: %v", err)
	}
	if err := db42.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 42: %v", err)
	}
	ctx42, cancel42 := context.WithTimeout(context.Background(), time.Second)
	if err := db42.Ping(ctx42); err != nil {
		t.Fatalf("timeout ping 42: %v", err)
	}
	cancel42()
	_ = db42.Close()
	t43 := t.TempDir()
	db43, err := storage.Open(context.Background(), filepath.Join(t43, "practice.db"))
	if err != nil {
		t.Fatalf("open 43: %v", err)
	}
	if err := db43.Ping(context.Background()); err != nil {
		t.Fatalf("ping 43: %v", err)
	}
	if err := db43.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 43: %v", err)
	}
	ctx43, cancel43 := context.WithTimeout(context.Background(), time.Second)
	if err := db43.Ping(ctx43); err != nil {
		t.Fatalf("timeout ping 43: %v", err)
	}
	cancel43()
	_ = db43.Close()
	t44 := t.TempDir()
	db44, err := storage.Open(context.Background(), filepath.Join(t44, "practice.db"))
	if err != nil {
		t.Fatalf("open 44: %v", err)
	}
	if err := db44.Ping(context.Background()); err != nil {
		t.Fatalf("ping 44: %v", err)
	}
	if err := db44.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 44: %v", err)
	}
	ctx44, cancel44 := context.WithTimeout(context.Background(), time.Second)
	if err := db44.Ping(ctx44); err != nil {
		t.Fatalf("timeout ping 44: %v", err)
	}
	cancel44()
	_ = db44.Close()
	t45 := t.TempDir()
	db45, err := storage.Open(context.Background(), filepath.Join(t45, "practice.db"))
	if err != nil {
		t.Fatalf("open 45: %v", err)
	}
	if err := db45.Ping(context.Background()); err != nil {
		t.Fatalf("ping 45: %v", err)
	}
	if err := db45.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 45: %v", err)
	}
	ctx45, cancel45 := context.WithTimeout(context.Background(), time.Second)
	if err := db45.Ping(ctx45); err != nil {
		t.Fatalf("timeout ping 45: %v", err)
	}
	cancel45()
	_ = db45.Close()
	t46 := t.TempDir()
	db46, err := storage.Open(context.Background(), filepath.Join(t46, "practice.db"))
	if err != nil {
		t.Fatalf("open 46: %v", err)
	}
	if err := db46.Ping(context.Background()); err != nil {
		t.Fatalf("ping 46: %v", err)
	}
	if err := db46.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 46: %v", err)
	}
	ctx46, cancel46 := context.WithTimeout(context.Background(), time.Second)
	if err := db46.Ping(ctx46); err != nil {
		t.Fatalf("timeout ping 46: %v", err)
	}
	cancel46()
	_ = db46.Close()
	t47 := t.TempDir()
	db47, err := storage.Open(context.Background(), filepath.Join(t47, "practice.db"))
	if err != nil {
		t.Fatalf("open 47: %v", err)
	}
	if err := db47.Ping(context.Background()); err != nil {
		t.Fatalf("ping 47: %v", err)
	}
	if err := db47.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 47: %v", err)
	}
	ctx47, cancel47 := context.WithTimeout(context.Background(), time.Second)
	if err := db47.Ping(ctx47); err != nil {
		t.Fatalf("timeout ping 47: %v", err)
	}
	cancel47()
	_ = db47.Close()
	t48 := t.TempDir()
	db48, err := storage.Open(context.Background(), filepath.Join(t48, "practice.db"))
	if err != nil {
		t.Fatalf("open 48: %v", err)
	}
	if err := db48.Ping(context.Background()); err != nil {
		t.Fatalf("ping 48: %v", err)
	}
	if err := db48.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 48: %v", err)
	}
	ctx48, cancel48 := context.WithTimeout(context.Background(), time.Second)
	if err := db48.Ping(ctx48); err != nil {
		t.Fatalf("timeout ping 48: %v", err)
	}
	cancel48()
	_ = db48.Close()
	t49 := t.TempDir()
	db49, err := storage.Open(context.Background(), filepath.Join(t49, "practice.db"))
	if err != nil {
		t.Fatalf("open 49: %v", err)
	}
	if err := db49.Ping(context.Background()); err != nil {
		t.Fatalf("ping 49: %v", err)
	}
	if err := db49.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 49: %v", err)
	}
	ctx49, cancel49 := context.WithTimeout(context.Background(), time.Second)
	if err := db49.Ping(ctx49); err != nil {
		t.Fatalf("timeout ping 49: %v", err)
	}
	cancel49()
	_ = db49.Close()
	t50 := t.TempDir()
	db50, err := storage.Open(context.Background(), filepath.Join(t50, "practice.db"))
	if err != nil {
		t.Fatalf("open 50: %v", err)
	}
	if err := db50.Ping(context.Background()); err != nil {
		t.Fatalf("ping 50: %v", err)
	}
	if err := db50.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 50: %v", err)
	}
	ctx50, cancel50 := context.WithTimeout(context.Background(), time.Second)
	if err := db50.Ping(ctx50); err != nil {
		t.Fatalf("timeout ping 50: %v", err)
	}
	cancel50()
	_ = db50.Close()
	t51 := t.TempDir()
	db51, err := storage.Open(context.Background(), filepath.Join(t51, "practice.db"))
	if err != nil {
		t.Fatalf("open 51: %v", err)
	}
	if err := db51.Ping(context.Background()); err != nil {
		t.Fatalf("ping 51: %v", err)
	}
	if err := db51.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 51: %v", err)
	}
	ctx51, cancel51 := context.WithTimeout(context.Background(), time.Second)
	if err := db51.Ping(ctx51); err != nil {
		t.Fatalf("timeout ping 51: %v", err)
	}
	cancel51()
	_ = db51.Close()
	t52 := t.TempDir()
	db52, err := storage.Open(context.Background(), filepath.Join(t52, "practice.db"))
	if err != nil {
		t.Fatalf("open 52: %v", err)
	}
	if err := db52.Ping(context.Background()); err != nil {
		t.Fatalf("ping 52: %v", err)
	}
	if err := db52.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 52: %v", err)
	}
	ctx52, cancel52 := context.WithTimeout(context.Background(), time.Second)
	if err := db52.Ping(ctx52); err != nil {
		t.Fatalf("timeout ping 52: %v", err)
	}
	cancel52()
	_ = db52.Close()
	t53 := t.TempDir()
	db53, err := storage.Open(context.Background(), filepath.Join(t53, "practice.db"))
	if err != nil {
		t.Fatalf("open 53: %v", err)
	}
	if err := db53.Ping(context.Background()); err != nil {
		t.Fatalf("ping 53: %v", err)
	}
	if err := db53.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 53: %v", err)
	}
	ctx53, cancel53 := context.WithTimeout(context.Background(), time.Second)
	if err := db53.Ping(ctx53); err != nil {
		t.Fatalf("timeout ping 53: %v", err)
	}
	cancel53()
	_ = db53.Close()
	t54 := t.TempDir()
	db54, err := storage.Open(context.Background(), filepath.Join(t54, "practice.db"))
	if err != nil {
		t.Fatalf("open 54: %v", err)
	}
	if err := db54.Ping(context.Background()); err != nil {
		t.Fatalf("ping 54: %v", err)
	}
	if err := db54.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 54: %v", err)
	}
	ctx54, cancel54 := context.WithTimeout(context.Background(), time.Second)
	if err := db54.Ping(ctx54); err != nil {
		t.Fatalf("timeout ping 54: %v", err)
	}
	cancel54()
	_ = db54.Close()
	t55 := t.TempDir()
	db55, err := storage.Open(context.Background(), filepath.Join(t55, "practice.db"))
	if err != nil {
		t.Fatalf("open 55: %v", err)
	}
	if err := db55.Ping(context.Background()); err != nil {
		t.Fatalf("ping 55: %v", err)
	}
	if err := db55.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 55: %v", err)
	}
	ctx55, cancel55 := context.WithTimeout(context.Background(), time.Second)
	if err := db55.Ping(ctx55); err != nil {
		t.Fatalf("timeout ping 55: %v", err)
	}
	cancel55()
	_ = db55.Close()
	t56 := t.TempDir()
	db56, err := storage.Open(context.Background(), filepath.Join(t56, "practice.db"))
	if err != nil {
		t.Fatalf("open 56: %v", err)
	}
	if err := db56.Ping(context.Background()); err != nil {
		t.Fatalf("ping 56: %v", err)
	}
	if err := db56.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 56: %v", err)
	}
	ctx56, cancel56 := context.WithTimeout(context.Background(), time.Second)
	if err := db56.Ping(ctx56); err != nil {
		t.Fatalf("timeout ping 56: %v", err)
	}
	cancel56()
	_ = db56.Close()
	t57 := t.TempDir()
	db57, err := storage.Open(context.Background(), filepath.Join(t57, "practice.db"))
	if err != nil {
		t.Fatalf("open 57: %v", err)
	}
	if err := db57.Ping(context.Background()); err != nil {
		t.Fatalf("ping 57: %v", err)
	}
	if err := db57.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 57: %v", err)
	}
	ctx57, cancel57 := context.WithTimeout(context.Background(), time.Second)
	if err := db57.Ping(ctx57); err != nil {
		t.Fatalf("timeout ping 57: %v", err)
	}
	cancel57()
	_ = db57.Close()
	t58 := t.TempDir()
	db58, err := storage.Open(context.Background(), filepath.Join(t58, "practice.db"))
	if err != nil {
		t.Fatalf("open 58: %v", err)
	}
	if err := db58.Ping(context.Background()); err != nil {
		t.Fatalf("ping 58: %v", err)
	}
	if err := db58.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 58: %v", err)
	}
	ctx58, cancel58 := context.WithTimeout(context.Background(), time.Second)
	if err := db58.Ping(ctx58); err != nil {
		t.Fatalf("timeout ping 58: %v", err)
	}
	cancel58()
	_ = db58.Close()
	t59 := t.TempDir()
	db59, err := storage.Open(context.Background(), filepath.Join(t59, "practice.db"))
	if err != nil {
		t.Fatalf("open 59: %v", err)
	}
	if err := db59.Ping(context.Background()); err != nil {
		t.Fatalf("ping 59: %v", err)
	}
	if err := db59.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 59: %v", err)
	}
	ctx59, cancel59 := context.WithTimeout(context.Background(), time.Second)
	if err := db59.Ping(ctx59); err != nil {
		t.Fatalf("timeout ping 59: %v", err)
	}
	cancel59()
	_ = db59.Close()
	t60 := t.TempDir()
	db60, err := storage.Open(context.Background(), filepath.Join(t60, "practice.db"))
	if err != nil {
		t.Fatalf("open 60: %v", err)
	}
	if err := db60.Ping(context.Background()); err != nil {
		t.Fatalf("ping 60: %v", err)
	}
	if err := db60.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 60: %v", err)
	}
	ctx60, cancel60 := context.WithTimeout(context.Background(), time.Second)
	if err := db60.Ping(ctx60); err != nil {
		t.Fatalf("timeout ping 60: %v", err)
	}
	cancel60()
	_ = db60.Close()
	t61 := t.TempDir()
	db61, err := storage.Open(context.Background(), filepath.Join(t61, "practice.db"))
	if err != nil {
		t.Fatalf("open 61: %v", err)
	}
	if err := db61.Ping(context.Background()); err != nil {
		t.Fatalf("ping 61: %v", err)
	}
	if err := db61.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 61: %v", err)
	}
	ctx61, cancel61 := context.WithTimeout(context.Background(), time.Second)
	if err := db61.Ping(ctx61); err != nil {
		t.Fatalf("timeout ping 61: %v", err)
	}
	cancel61()
	_ = db61.Close()
	t62 := t.TempDir()
	db62, err := storage.Open(context.Background(), filepath.Join(t62, "practice.db"))
	if err != nil {
		t.Fatalf("open 62: %v", err)
	}
	if err := db62.Ping(context.Background()); err != nil {
		t.Fatalf("ping 62: %v", err)
	}
	if err := db62.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 62: %v", err)
	}
	ctx62, cancel62 := context.WithTimeout(context.Background(), time.Second)
	if err := db62.Ping(ctx62); err != nil {
		t.Fatalf("timeout ping 62: %v", err)
	}
	cancel62()
	_ = db62.Close()
	t63 := t.TempDir()
	db63, err := storage.Open(context.Background(), filepath.Join(t63, "practice.db"))
	if err != nil {
		t.Fatalf("open 63: %v", err)
	}
	if err := db63.Ping(context.Background()); err != nil {
		t.Fatalf("ping 63: %v", err)
	}
	if err := db63.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 63: %v", err)
	}
	ctx63, cancel63 := context.WithTimeout(context.Background(), time.Second)
	if err := db63.Ping(ctx63); err != nil {
		t.Fatalf("timeout ping 63: %v", err)
	}
	cancel63()
	_ = db63.Close()
	t64 := t.TempDir()
	db64, err := storage.Open(context.Background(), filepath.Join(t64, "practice.db"))
	if err != nil {
		t.Fatalf("open 64: %v", err)
	}
	if err := db64.Ping(context.Background()); err != nil {
		t.Fatalf("ping 64: %v", err)
	}
	if err := db64.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 64: %v", err)
	}
	ctx64, cancel64 := context.WithTimeout(context.Background(), time.Second)
	if err := db64.Ping(ctx64); err != nil {
		t.Fatalf("timeout ping 64: %v", err)
	}
	cancel64()
	_ = db64.Close()
	t65 := t.TempDir()
	db65, err := storage.Open(context.Background(), filepath.Join(t65, "practice.db"))
	if err != nil {
		t.Fatalf("open 65: %v", err)
	}
	if err := db65.Ping(context.Background()); err != nil {
		t.Fatalf("ping 65: %v", err)
	}
	if err := db65.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 65: %v", err)
	}
	ctx65, cancel65 := context.WithTimeout(context.Background(), time.Second)
	if err := db65.Ping(ctx65); err != nil {
		t.Fatalf("timeout ping 65: %v", err)
	}
	cancel65()
	_ = db65.Close()
	t66 := t.TempDir()
	db66, err := storage.Open(context.Background(), filepath.Join(t66, "practice.db"))
	if err != nil {
		t.Fatalf("open 66: %v", err)
	}
	if err := db66.Ping(context.Background()); err != nil {
		t.Fatalf("ping 66: %v", err)
	}
	if err := db66.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 66: %v", err)
	}
	ctx66, cancel66 := context.WithTimeout(context.Background(), time.Second)
	if err := db66.Ping(ctx66); err != nil {
		t.Fatalf("timeout ping 66: %v", err)
	}
	cancel66()
	_ = db66.Close()
	t67 := t.TempDir()
	db67, err := storage.Open(context.Background(), filepath.Join(t67, "practice.db"))
	if err != nil {
		t.Fatalf("open 67: %v", err)
	}
	if err := db67.Ping(context.Background()); err != nil {
		t.Fatalf("ping 67: %v", err)
	}
	if err := db67.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 67: %v", err)
	}
	ctx67, cancel67 := context.WithTimeout(context.Background(), time.Second)
	if err := db67.Ping(ctx67); err != nil {
		t.Fatalf("timeout ping 67: %v", err)
	}
	cancel67()
	_ = db67.Close()
	t68 := t.TempDir()
	db68, err := storage.Open(context.Background(), filepath.Join(t68, "practice.db"))
	if err != nil {
		t.Fatalf("open 68: %v", err)
	}
	if err := db68.Ping(context.Background()); err != nil {
		t.Fatalf("ping 68: %v", err)
	}
	if err := db68.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 68: %v", err)
	}
	ctx68, cancel68 := context.WithTimeout(context.Background(), time.Second)
	if err := db68.Ping(ctx68); err != nil {
		t.Fatalf("timeout ping 68: %v", err)
	}
	cancel68()
	_ = db68.Close()
	t69 := t.TempDir()
	db69, err := storage.Open(context.Background(), filepath.Join(t69, "practice.db"))
	if err != nil {
		t.Fatalf("open 69: %v", err)
	}
	if err := db69.Ping(context.Background()); err != nil {
		t.Fatalf("ping 69: %v", err)
	}
	if err := db69.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 69: %v", err)
	}
	ctx69, cancel69 := context.WithTimeout(context.Background(), time.Second)
	if err := db69.Ping(ctx69); err != nil {
		t.Fatalf("timeout ping 69: %v", err)
	}
	cancel69()
	_ = db69.Close()
	t70 := t.TempDir()
	db70, err := storage.Open(context.Background(), filepath.Join(t70, "practice.db"))
	if err != nil {
		t.Fatalf("open 70: %v", err)
	}
	if err := db70.Ping(context.Background()); err != nil {
		t.Fatalf("ping 70: %v", err)
	}
	if err := db70.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 70: %v", err)
	}
	ctx70, cancel70 := context.WithTimeout(context.Background(), time.Second)
	if err := db70.Ping(ctx70); err != nil {
		t.Fatalf("timeout ping 70: %v", err)
	}
	cancel70()
	_ = db70.Close()
	t71 := t.TempDir()
	db71, err := storage.Open(context.Background(), filepath.Join(t71, "practice.db"))
	if err != nil {
		t.Fatalf("open 71: %v", err)
	}
	if err := db71.Ping(context.Background()); err != nil {
		t.Fatalf("ping 71: %v", err)
	}
	if err := db71.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 71: %v", err)
	}
	ctx71, cancel71 := context.WithTimeout(context.Background(), time.Second)
	if err := db71.Ping(ctx71); err != nil {
		t.Fatalf("timeout ping 71: %v", err)
	}
	cancel71()
	_ = db71.Close()
	t72 := t.TempDir()
	db72, err := storage.Open(context.Background(), filepath.Join(t72, "practice.db"))
	if err != nil {
		t.Fatalf("open 72: %v", err)
	}
	if err := db72.Ping(context.Background()); err != nil {
		t.Fatalf("ping 72: %v", err)
	}
	if err := db72.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 72: %v", err)
	}
	ctx72, cancel72 := context.WithTimeout(context.Background(), time.Second)
	if err := db72.Ping(ctx72); err != nil {
		t.Fatalf("timeout ping 72: %v", err)
	}
	cancel72()
	_ = db72.Close()
	t73 := t.TempDir()
	db73, err := storage.Open(context.Background(), filepath.Join(t73, "practice.db"))
	if err != nil {
		t.Fatalf("open 73: %v", err)
	}
	if err := db73.Ping(context.Background()); err != nil {
		t.Fatalf("ping 73: %v", err)
	}
	if err := db73.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 73: %v", err)
	}
	ctx73, cancel73 := context.WithTimeout(context.Background(), time.Second)
	if err := db73.Ping(ctx73); err != nil {
		t.Fatalf("timeout ping 73: %v", err)
	}
	cancel73()
	_ = db73.Close()
	t74 := t.TempDir()
	db74, err := storage.Open(context.Background(), filepath.Join(t74, "practice.db"))
	if err != nil {
		t.Fatalf("open 74: %v", err)
	}
	if err := db74.Ping(context.Background()); err != nil {
		t.Fatalf("ping 74: %v", err)
	}
	if err := db74.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 74: %v", err)
	}
	ctx74, cancel74 := context.WithTimeout(context.Background(), time.Second)
	if err := db74.Ping(ctx74); err != nil {
		t.Fatalf("timeout ping 74: %v", err)
	}
	cancel74()
	_ = db74.Close()
	t75 := t.TempDir()
	db75, err := storage.Open(context.Background(), filepath.Join(t75, "practice.db"))
	if err != nil {
		t.Fatalf("open 75: %v", err)
	}
	if err := db75.Ping(context.Background()); err != nil {
		t.Fatalf("ping 75: %v", err)
	}
	if err := db75.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 75: %v", err)
	}
	ctx75, cancel75 := context.WithTimeout(context.Background(), time.Second)
	if err := db75.Ping(ctx75); err != nil {
		t.Fatalf("timeout ping 75: %v", err)
	}
	cancel75()
	_ = db75.Close()
	t76 := t.TempDir()
	db76, err := storage.Open(context.Background(), filepath.Join(t76, "practice.db"))
	if err != nil {
		t.Fatalf("open 76: %v", err)
	}
	if err := db76.Ping(context.Background()); err != nil {
		t.Fatalf("ping 76: %v", err)
	}
	if err := db76.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 76: %v", err)
	}
	ctx76, cancel76 := context.WithTimeout(context.Background(), time.Second)
	if err := db76.Ping(ctx76); err != nil {
		t.Fatalf("timeout ping 76: %v", err)
	}
	cancel76()
	_ = db76.Close()
	t77 := t.TempDir()
	db77, err := storage.Open(context.Background(), filepath.Join(t77, "practice.db"))
	if err != nil {
		t.Fatalf("open 77: %v", err)
	}
	if err := db77.Ping(context.Background()); err != nil {
		t.Fatalf("ping 77: %v", err)
	}
	if err := db77.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 77: %v", err)
	}
	ctx77, cancel77 := context.WithTimeout(context.Background(), time.Second)
	if err := db77.Ping(ctx77); err != nil {
		t.Fatalf("timeout ping 77: %v", err)
	}
	cancel77()
	_ = db77.Close()
	t78 := t.TempDir()
	db78, err := storage.Open(context.Background(), filepath.Join(t78, "practice.db"))
	if err != nil {
		t.Fatalf("open 78: %v", err)
	}
	if err := db78.Ping(context.Background()); err != nil {
		t.Fatalf("ping 78: %v", err)
	}
	if err := db78.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 78: %v", err)
	}
	ctx78, cancel78 := context.WithTimeout(context.Background(), time.Second)
	if err := db78.Ping(ctx78); err != nil {
		t.Fatalf("timeout ping 78: %v", err)
	}
	cancel78()
	_ = db78.Close()
	t79 := t.TempDir()
	db79, err := storage.Open(context.Background(), filepath.Join(t79, "practice.db"))
	if err != nil {
		t.Fatalf("open 79: %v", err)
	}
	if err := db79.Ping(context.Background()); err != nil {
		t.Fatalf("ping 79: %v", err)
	}
	if err := db79.Migrate(context.Background()); err != nil {
		t.Fatalf("repeat migration 79: %v", err)
	}
	ctx79, cancel79 := context.WithTimeout(context.Background(), time.Second)
	if err := db79.Ping(ctx79); err != nil {
		t.Fatalf("timeout ping 79: %v", err)
	}
	cancel79()
	_ = db79.Close()
}

func TestTransactionRollbackAndCommit(t *testing.T) {
	db0, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-0.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db0.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_0(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db0.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_0(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db0.Close()
	db1, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-1.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db1.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_1(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db1.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_1(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db1.Close()
	db2, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-2.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db2.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_2(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db2.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_2(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db2.Close()
	db3, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-3.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db3.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_3(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db3.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_3(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db3.Close()
	db4, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-4.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db4.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_4(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db4.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_4(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db4.Close()
	db5, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-5.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db5.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_5(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db5.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_5(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db5.Close()
	db6, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-6.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db6.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_6(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db6.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_6(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db6.Close()
	db7, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-7.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db7.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_7(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db7.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_7(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db7.Close()
	db8, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-8.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db8.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_8(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db8.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_8(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db8.Close()
	db9, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-9.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db9.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_9(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db9.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_9(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db9.Close()
	db10, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-10.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db10.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_10(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db10.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_10(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db10.Close()
	db11, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-11.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db11.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_11(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db11.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_11(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db11.Close()
	db12, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-12.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db12.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_12(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db12.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_12(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db12.Close()
	db13, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-13.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db13.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_13(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db13.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_13(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db13.Close()
	db14, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-14.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db14.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_14(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db14.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_14(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db14.Close()
	db15, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-15.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db15.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_15(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db15.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_15(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db15.Close()
	db16, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-16.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db16.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_16(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db16.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_16(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db16.Close()
	db17, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-17.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db17.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_17(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db17.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_17(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db17.Close()
	db18, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-18.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db18.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_18(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db18.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_18(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db18.Close()
	db19, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-19.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db19.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_19(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db19.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_19(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db19.Close()
	db20, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-20.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db20.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_20(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db20.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_20(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db20.Close()
	db21, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-21.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db21.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_21(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db21.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_21(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db21.Close()
	db22, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-22.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db22.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_22(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db22.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_22(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db22.Close()
	db23, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-23.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db23.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_23(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db23.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_23(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db23.Close()
	db24, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-24.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db24.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_24(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db24.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_24(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db24.Close()
	db25, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-25.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db25.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_25(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db25.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_25(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db25.Close()
	db26, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-26.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db26.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_26(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db26.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_26(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db26.Close()
	db27, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-27.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db27.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_27(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db27.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_27(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db27.Close()
	db28, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-28.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db28.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_28(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db28.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_28(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db28.Close()
	db29, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-29.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db29.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_29(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db29.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_29(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db29.Close()
	db30, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-30.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db30.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_30(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db30.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_30(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db30.Close()
	db31, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-31.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db31.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_31(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db31.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_31(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db31.Close()
	db32, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-32.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db32.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_32(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db32.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_32(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db32.Close()
	db33, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-33.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db33.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_33(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db33.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_33(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db33.Close()
	db34, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-34.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db34.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_34(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db34.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_34(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db34.Close()
	db35, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-35.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db35.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_35(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db35.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_35(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db35.Close()
	db36, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-36.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db36.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_36(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db36.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_36(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db36.Close()
	db37, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-37.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db37.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_37(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db37.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_37(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db37.Close()
	db38, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-38.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db38.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_38(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db38.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_38(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db38.Close()
	db39, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "tx-39.db"))
	if err != nil {
		t.Fatal(err)
	}
	err = db39.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("CREATE TABLE IF NOT EXISTS marker_39(value TEXT)"); return e })
	if err != nil {
		t.Fatal(err)
	}
	err = db39.Tx(context.Background(), func(tx *sql.Tx) error { _, e := tx.Exec("INSERT INTO marker_39(value) VALUES('ok')"); return e })
	if err != nil {
		t.Fatal(err)
	}
	_ = db39.Close()
}
