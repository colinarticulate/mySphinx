/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: xyz.i

package xyz

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _goslice_ swig_type_1;
typedef _goslice_ swig_type_2;
typedef _gostring_ swig_type_3;
extern void _wrap_Swig_free_xyz_cf6b1f9fea9c086b(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_xyz_cf6b1f9fea9c086b(swig_intgo arg1);
extern swig_intgo _wrap_passing_bytes_xyz_cf6b1f9fea9c086b(swig_type_1 arg1);
extern swig_intgo _wrap_create_file_params_nofilename_xyz_cf6b1f9fea9c086b(swig_type_2 arg1);
extern swig_intgo _wrap_check_string_xyz_cf6b1f9fea9c086b(swig_type_3 arg1);
extern swig_intgo _wrap_ps_call_xyz_cf6b1f9fea9c086b(swig_type_1 arg1, swig_type_1 arg2, swig_type_3 arg3);
#undef intgo
*/
import "C"

import (
	_ "runtime/cgo"
	"sync"
	"unsafe"
)

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_xyz_cf6b1f9fea9c086b(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_xyz_cf6b1f9fea9c086b(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Passing_bytes(arg1 []byte) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_passing_bytes_xyz_cf6b1f9fea9c086b(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Create_file_params_nofilename(arg1 []string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_create_file_params_nofilename_xyz_cf6b1f9fea9c086b(*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Check_string(arg1 string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_check_string_xyz_cf6b1f9fea9c086b(*(*C.swig_type_3)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Ps_call(arg1 []byte, arg2 []byte, arg3 []string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_ps_call_xyz_cf6b1f9fea9c086b((*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))), (*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1))), (*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_2)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}
