package main

import (
	"reflect"
	"testing"
)

func Test1SortAscending(t *testing.T) {
	arr := []int{2, 1, 3}
	expected := []int{1, 2, 3}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}

}

func Test2SortAscending(t *testing.T) {
	arr := []int{12, 6, 3}
	expected := []int{3, 6, 12}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}
}

func Test3SortAscending(t *testing.T) {
	arr := []int{6, 3, 9}
	expected := []int{3, 6, 9}
	t.Logf("initial arr= %d", arr)

	ascArr:=SortAscending(arr)
	t.Logf("sort asc= %d", ascArr)
	if !reflect.DeepEqual(ascArr, expected) {
		t.Errorf("expected sort asc %d but got %d", expected, ascArr)
	}else{
		t.Logf("sorting ok")
	}
}

