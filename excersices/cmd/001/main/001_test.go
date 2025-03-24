package main

import "testing"

/**
Go test --cover = Test coverage = line code called test / tong so lines * 100 = trong truong hop nay = 5/7*100 = 71.4%
Phạm vi câu lệnh (Statement Coverage):
Đảm bảo rằng tất cả các dòng trong mã nguồn đã được kiểm tra ít nhất một lần.
Mức độ bao phủ của câu lệnh = (Số câu lệnh mã được kiểm tra / Tổng số câu lệnh mã) x100
Ví dụ: Có 100 lệnh, 87 lệnh có thể kiểm tra độ phủ của câu lệnh = 87%. Đối với ví dụ này, để bao phủ 100% câu lệnh dc, mỗi lệnh bắt buộc phải được kiểm tra ít nhất một lần.
**/

func TestTask(t *testing.T) {
	expected := "112,119,126,133,147,154,161,168,182,189,196"
	actual := cal(100, 200)

	if actual != expected {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}
