package main

import "testing"

func TestAdd(t *testing.T) {

	cases := []struct {
		Name    string
		A, B, C int
	}{
		{"a1", 1, 2, 3},
		{"a2", 4, 5, 6},
		{"a3", 7, 8, 9},
	}
	for _, s := range cases {
		t.Run(s.Name, func(t *testing.T) {
			if Add(s.A, s.B) != s.C {
				t.Errorf("测试失败")
				return
			}
			t.Logf("测试成功")
		})
	}

	//子测试
	//t.Run("add1", func(t *testing.T) {
	//	if Add(1, 2) != 0 {
	//		t.Errorf("测试失败")
	//		return
	//	}
	//})
	//
	//t.Run("add2", func(t *testing.T) {
	//	if Add(1, 2) != 1 {
	//		t.Errorf("测试失败")
	//		return
	//	}
	//
	//})
	//
	//t.Run("add3", func(t *testing.T) {
	//	if Add(1, 2) == 3 {
	//		t.Logf("测试成功")
	//		return
	//	}
	//})

}
