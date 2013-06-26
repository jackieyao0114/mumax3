package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var addexchange_code cu.Function

type addexchange_args struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_aLUT2d  unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_wx      float32
	arg_wy      float32
	arg_wz      float32
	arg_N0      int
	arg_N1      int
	arg_N2      int
	argptr      [14]unsafe.Pointer
}

// Wrapper for addexchange CUDA kernel, asynchronous.
func k_addexchange_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, aLUT2d unsafe.Pointer, regions unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if addexchange_code == 0 {
		addexchange_code = fatbinLoad(addexchange_map, "addexchange")
	}

	var _a_ addexchange_args

	_a_.arg_Bx = Bx
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_Bx)
	_a_.arg_By = By
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_By)
	_a_.arg_Bz = Bz
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Bz)
	_a_.arg_mx = mx
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_mx)
	_a_.arg_my = my
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_my)
	_a_.arg_mz = mz
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_mz)
	_a_.arg_aLUT2d = aLUT2d
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_aLUT2d)
	_a_.arg_regions = regions
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_regions)
	_a_.arg_wx = wx
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_wx)
	_a_.arg_wy = wy
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_wy)
	_a_.arg_wz = wz
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_wz)
	_a_.arg_N0 = N0
	_a_.argptr[11] = unsafe.Pointer(&_a_.arg_N0)
	_a_.arg_N1 = N1
	_a_.argptr[12] = unsafe.Pointer(&_a_.arg_N1)
	_a_.arg_N2 = N2
	_a_.argptr[13] = unsafe.Pointer(&_a_.arg_N2)

	args := _a_.argptr[:]
	cu.LaunchKernel(addexchange_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for addexchange CUDA kernel, synchronized.
func k_addexchange(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, aLUT2d unsafe.Pointer, regions unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_addexchange_async(Bx, By, Bz, mx, my, mz, aLUT2d, regions, wx, wy, wz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var addexchange_map = map[int]string{0: "",
	20: addexchange_ptx_20,
	30: addexchange_ptx_30,
	35: addexchange_ptx_35}

const (
	addexchange_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<13>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<258>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<106>;


	ld.param.u64 	%rd8, [addexchange_param_0];
	ld.param.u64 	%rd9, [addexchange_param_1];
	ld.param.u64 	%rd10, [addexchange_param_2];
	ld.param.u64 	%rd11, [addexchange_param_3];
	ld.param.u64 	%rd12, [addexchange_param_4];
	ld.param.u64 	%rd13, [addexchange_param_5];
	ld.param.u64 	%rd14, [addexchange_param_6];
	ld.param.u64 	%rd15, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r33, [addexchange_param_11];
	ld.param.u32 	%r34, [addexchange_param_12];
	ld.param.u32 	%r35, [addexchange_param_13];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd15;
	cvta.to.global.u64 	%rd5, %rd13;
	cvta.to.global.u64 	%rd6, %rd12;
	cvta.to.global.u64 	%rd7, %rd11;
	.loc 2 16 1
	mov.u32 	%r36, %ntid.z;
	mov.u32 	%r37, %ctaid.z;
	mov.u32 	%r38, %tid.z;
	mad.lo.s32 	%r1, %r36, %r37, %r38;
	.loc 2 17 1
	mov.u32 	%r39, %ntid.y;
	mov.u32 	%r40, %ctaid.y;
	mov.u32 	%r41, %tid.y;
	mad.lo.s32 	%r2, %r39, %r40, %r41;
	.loc 2 18 1
	mov.u32 	%r42, %ntid.x;
	mov.u32 	%r43, %ctaid.x;
	mov.u32 	%r44, %tid.x;
	mad.lo.s32 	%r3, %r42, %r43, %r44;
	.loc 2 20 1
	setp.ge.s32 	%p1, %r2, %r34;
	setp.ge.s32 	%p2, %r1, %r33;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r3, %r35;
	or.pred  	%p5, %p3, %p4;
	@%p5 bra 	BB0_22;

	.loc 2 25 1
	mad.lo.s32 	%r45, %r1, %r34, %r2;
	mad.lo.s32 	%r46, %r45, %r35, %r3;
	.loc 2 26 1
	cvt.s64.s32 	%rd16, %r46;
	mul.wide.s32 	%rd17, %r46, 4;
	add.s64 	%rd18, %rd7, %rd17;
	ld.global.f32 	%f1, [%rd18];
	add.s64 	%rd19, %rd6, %rd17;
	ld.global.f32 	%f2, [%rd19];
	add.s64 	%rd20, %rd5, %rd17;
	ld.global.f32 	%f3, [%rd20];
	.loc 2 27 1
	add.s64 	%rd21, %rd4, %rd16;
	.loc 2 28 1
	add.s64 	%rd22, %rd3, %rd17;
	ld.global.f32 	%f4, [%rd22];
	add.s64 	%rd23, %rd2, %rd17;
	ld.global.f32 	%f5, [%rd23];
	add.s64 	%rd24, %rd1, %rd17;
	ld.global.f32 	%f6, [%rd24];
	.loc 2 35 1
	add.s32 	%r53, %r3, -1;
	mov.u32 	%r54, 0;
	.loc 3 238 5
	max.s32 	%r55, %r53, %r54;
	.loc 2 35 1
	mad.lo.s32 	%r56, %r45, %r35, %r55;
	.loc 2 36 1
	cvt.s64.s32 	%rd25, %r56;
	mul.wide.s32 	%rd26, %r56, 4;
	add.s64 	%rd27, %rd7, %rd26;
	ld.global.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd6, %rd26;
	ld.global.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd5, %rd26;
	ld.global.f32 	%f9, [%rd29];
	.loc 2 37 1
	add.s64 	%rd30, %rd4, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	.loc 2 27 1
	ld.global.u8 	%rc1, [%rd21];
	.loc 2 37 1
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r4, %rc2;
	@%p6 bra 	BB0_3;

	cvt.s32.s8 	%r62, %rc1;
	add.s32 	%r63, %r62, 1;
	mul.lo.s32 	%r64, %r63, %r62;
	shr.u32 	%r65, %r64, 31;
	mad.lo.s32 	%r66, %r63, %r62, %r65;
	shr.s32 	%r67, %r66, 1;
	add.s32 	%r252, %r4, %r67;
	bra.uni 	BB0_4;

BB0_3:
	.loc 2 37 1
	add.s32 	%r68, %r4, 1;
	mul.lo.s32 	%r69, %r68, %r4;
	shr.u32 	%r70, %r69, 31;
	mad.lo.s32 	%r71, %r68, %r4, %r70;
	shr.s32 	%r72, %r71, 1;
	cvt.s32.s8 	%r73, %rc1;
	add.s32 	%r252, %r72, %r73;

BB0_4:
	cvta.to.global.u64 	%rd32, %rd14;
	.loc 2 37 1
	mul.wide.s32 	%rd33, %r252, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 2 38 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 2 38 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 2 41 1
	add.s32 	%r79, %r3, 1;
	add.s32 	%r80, %r35, -1;
	.loc 3 210 5
	min.s32 	%r81, %r79, %r80;
	.loc 2 41 1
	mad.lo.s32 	%r91, %r45, %r35, %r81;
	.loc 2 42 1
	cvt.s64.s32 	%rd35, %r91;
	mul.wide.s32 	%rd37, %r91, 4;
	add.s64 	%rd38, %rd7, %rd37;
	ld.global.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd6, %rd37;
	ld.global.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd5, %rd37;
	ld.global.f32 	%f15, [%rd42];
	.loc 2 43 1
	add.s64 	%rd43, %rd4, %rd35;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r8, %rc1;
	cvt.s32.s8 	%r9, %rc3;
	@%p7 bra 	BB0_6;

	add.s32 	%r96, %r8, 1;
	mul.lo.s32 	%r97, %r96, %r8;
	shr.u32 	%r98, %r97, 31;
	mad.lo.s32 	%r99, %r96, %r8, %r98;
	shr.s32 	%r100, %r99, 1;
	add.s32 	%r253, %r9, %r100;
	bra.uni 	BB0_7;

BB0_6:
	.loc 2 43 1
	add.s32 	%r101, %r9, 1;
	mul.lo.s32 	%r102, %r101, %r9;
	shr.u32 	%r103, %r102, 31;
	mad.lo.s32 	%r104, %r101, %r9, %r103;
	shr.s32 	%r105, %r104, 1;
	add.s32 	%r253, %r105, %r8;

BB0_7:
	mul.wide.s32 	%rd46, %r253, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 2 44 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 2 44 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 2 47 1
	add.s32 	%r111, %r2, -1;
	.loc 3 238 5
	max.s32 	%r113, %r111, %r54;
	.loc 2 47 1
	mad.lo.s32 	%r118, %r1, %r34, %r113;
	mad.lo.s32 	%r123, %r118, %r35, %r3;
	.loc 2 48 1
	cvt.s64.s32 	%rd48, %r123;
	mul.wide.s32 	%rd50, %r123, 4;
	add.s64 	%rd51, %rd7, %rd50;
	ld.global.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd6, %rd50;
	ld.global.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd5, %rd50;
	ld.global.f32 	%f21, [%rd55];
	.loc 2 49 1
	add.s64 	%rd56, %rd4, %rd48;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r14, %rc4;
	@%p8 bra 	BB0_9;

	add.s32 	%r128, %r8, 1;
	mul.lo.s32 	%r129, %r128, %r8;
	shr.u32 	%r130, %r129, 31;
	mad.lo.s32 	%r131, %r128, %r8, %r130;
	shr.s32 	%r132, %r131, 1;
	add.s32 	%r254, %r14, %r132;
	bra.uni 	BB0_10;

BB0_9:
	.loc 2 49 1
	add.s32 	%r133, %r14, 1;
	mul.lo.s32 	%r134, %r133, %r14;
	shr.u32 	%r135, %r134, 31;
	mad.lo.s32 	%r136, %r133, %r14, %r135;
	shr.s32 	%r137, %r136, 1;
	add.s32 	%r254, %r137, %r8;

BB0_10:
	mul.wide.s32 	%rd59, %r254, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 2 50 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 2 50 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 2 53 1
	add.s32 	%r143, %r2, 1;
	add.s32 	%r144, %r34, -1;
	.loc 3 210 5
	min.s32 	%r145, %r143, %r144;
	.loc 2 53 1
	mad.lo.s32 	%r150, %r1, %r34, %r145;
	mad.lo.s32 	%r155, %r150, %r35, %r3;
	.loc 2 54 1
	cvt.s64.s32 	%rd61, %r155;
	mul.wide.s32 	%rd63, %r155, 4;
	add.s64 	%rd64, %rd7, %rd63;
	ld.global.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd6, %rd63;
	ld.global.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd5, %rd63;
	ld.global.f32 	%f27, [%rd68];
	.loc 2 55 1
	add.s64 	%rd69, %rd4, %rd61;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r19, %rc5;
	@%p9 bra 	BB0_12;

	add.s32 	%r160, %r8, 1;
	mul.lo.s32 	%r161, %r160, %r8;
	shr.u32 	%r162, %r161, 31;
	mad.lo.s32 	%r163, %r160, %r8, %r162;
	shr.s32 	%r164, %r163, 1;
	add.s32 	%r255, %r19, %r164;
	bra.uni 	BB0_13;

BB0_12:
	.loc 2 55 1
	add.s32 	%r165, %r19, 1;
	mul.lo.s32 	%r166, %r165, %r19;
	shr.u32 	%r167, %r166, 31;
	mad.lo.s32 	%r168, %r165, %r19, %r167;
	shr.s32 	%r169, %r168, 1;
	add.s32 	%r255, %r169, %r8;

BB0_13:
	mul.wide.s32 	%rd71, %r255, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 2 56 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 2 56 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	.loc 2 59 1
	setp.eq.s32 	%p10, %r33, 1;
	@%p10 bra 	BB0_21;

	.loc 2 61 1
	add.s32 	%r175, %r1, -1;
	.loc 3 238 5
	max.s32 	%r177, %r175, %r54;
	.loc 2 61 1
	mad.lo.s32 	%r182, %r177, %r34, %r2;
	mad.lo.s32 	%r187, %r182, %r35, %r3;
	.loc 2 62 1
	cvt.s64.s32 	%rd74, %r187;
	mul.wide.s32 	%rd76, %r187, 4;
	add.s64 	%rd77, %rd7, %rd76;
	ld.global.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd6, %rd76;
	ld.global.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd5, %rd76;
	ld.global.f32 	%f33, [%rd81];
	.loc 2 63 1
	add.s64 	%rd82, %rd4, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r24, %rc6;
	@%p11 bra 	BB0_16;

	add.s32 	%r192, %r8, 1;
	mul.lo.s32 	%r193, %r192, %r8;
	shr.u32 	%r194, %r193, 31;
	mad.lo.s32 	%r195, %r192, %r8, %r194;
	shr.s32 	%r196, %r195, 1;
	add.s32 	%r256, %r24, %r196;
	bra.uni 	BB0_17;

BB0_16:
	.loc 2 63 1
	add.s32 	%r197, %r24, 1;
	mul.lo.s32 	%r198, %r197, %r24;
	shr.u32 	%r199, %r198, 31;
	mad.lo.s32 	%r200, %r197, %r24, %r199;
	shr.s32 	%r201, %r200, 1;
	add.s32 	%r256, %r201, %r8;

BB0_17:
	mul.wide.s32 	%rd85, %r256, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 2 64 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 2 64 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 2 67 1
	add.s32 	%r207, %r1, 1;
	add.s32 	%r208, %r33, -1;
	.loc 3 210 5
	min.s32 	%r209, %r207, %r208;
	.loc 2 67 1
	mad.lo.s32 	%r214, %r209, %r34, %r2;
	mad.lo.s32 	%r219, %r214, %r35, %r3;
	.loc 2 68 1
	cvt.s64.s32 	%rd87, %r219;
	mul.wide.s32 	%rd89, %r219, 4;
	add.s64 	%rd90, %rd7, %rd89;
	ld.global.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd6, %rd89;
	ld.global.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd5, %rd89;
	ld.global.f32 	%f39, [%rd94];
	.loc 2 69 1
	add.s64 	%rd95, %rd4, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r29, %rc7;
	@%p12 bra 	BB0_19;

	add.s32 	%r224, %r8, 1;
	mul.lo.s32 	%r225, %r224, %r8;
	shr.u32 	%r226, %r225, 31;
	mad.lo.s32 	%r227, %r224, %r8, %r226;
	shr.s32 	%r228, %r227, 1;
	add.s32 	%r257, %r29, %r228;
	bra.uni 	BB0_20;

BB0_19:
	.loc 2 69 1
	add.s32 	%r229, %r29, 1;
	mul.lo.s32 	%r230, %r229, %r29;
	shr.u32 	%r231, %r230, 31;
	mad.lo.s32 	%r232, %r229, %r29, %r231;
	shr.s32 	%r233, %r232, 1;
	add.s32 	%r257, %r233, %r8;

BB0_20:
	mul.wide.s32 	%rd97, %r257, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 2 70 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 2 70 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB0_21:
	.loc 2 28 1
	mul.wide.s32 	%rd100, %r46, 4;
	add.s64 	%rd101, %rd3, %rd100;
	.loc 2 73 1
	st.global.f32 	[%rd101], %f79;
	.loc 2 28 1
	add.s64 	%rd103, %rd2, %rd100;
	.loc 2 74 1
	st.global.f32 	[%rd103], %f80;
	.loc 2 28 1
	add.s64 	%rd105, %rd1, %rd100;
	.loc 2 75 1
	st.global.f32 	[%rd105], %f81;

BB0_22:
	.loc 2 76 2
	ret;
}


`
	addexchange_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<13>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<185>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<106>;


	ld.param.u64 	%rd9, [addexchange_param_0];
	ld.param.u64 	%rd10, [addexchange_param_1];
	ld.param.u64 	%rd11, [addexchange_param_2];
	ld.param.u64 	%rd12, [addexchange_param_3];
	ld.param.u64 	%rd13, [addexchange_param_4];
	ld.param.u64 	%rd14, [addexchange_param_5];
	ld.param.u64 	%rd15, [addexchange_param_6];
	ld.param.u64 	%rd16, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r35, [addexchange_param_11];
	ld.param.u32 	%r36, [addexchange_param_12];
	ld.param.u32 	%r37, [addexchange_param_13];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd16;
	cvta.to.global.u64 	%rd5, %rd14;
	cvta.to.global.u64 	%rd6, %rd13;
	cvta.to.global.u64 	%rd7, %rd12;
	.loc 2 16 1
	mov.u32 	%r38, %ntid.z;
	mov.u32 	%r39, %ctaid.z;
	mov.u32 	%r40, %tid.z;
	mad.lo.s32 	%r1, %r38, %r39, %r40;
	.loc 2 17 1
	mov.u32 	%r41, %ntid.y;
	mov.u32 	%r42, %ctaid.y;
	mov.u32 	%r43, %tid.y;
	mad.lo.s32 	%r2, %r41, %r42, %r43;
	.loc 2 18 1
	mov.u32 	%r44, %ntid.x;
	mov.u32 	%r45, %ctaid.x;
	mov.u32 	%r46, %tid.x;
	mad.lo.s32 	%r3, %r44, %r45, %r46;
	.loc 2 20 1
	setp.ge.s32 	%p1, %r2, %r36;
	setp.ge.s32 	%p2, %r1, %r35;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r3, %r37;
	or.pred  	%p5, %p3, %p4;
	@%p5 bra 	BB0_22;

	.loc 2 25 1
	mul.lo.s32 	%r4, %r1, %r36;
	mad.lo.s32 	%r47, %r1, %r36, %r2;
	mul.lo.s32 	%r5, %r47, %r37;
	mad.lo.s32 	%r48, %r47, %r37, %r3;
	.loc 2 26 1
	cvt.s64.s32 	%rd8, %r48;
	mul.wide.s32 	%rd17, %r48, 4;
	add.s64 	%rd18, %rd7, %rd17;
	ld.global.f32 	%f1, [%rd18];
	add.s64 	%rd19, %rd6, %rd17;
	ld.global.f32 	%f2, [%rd19];
	add.s64 	%rd20, %rd5, %rd17;
	ld.global.f32 	%f3, [%rd20];
	.loc 2 27 1
	add.s64 	%rd21, %rd4, %rd8;
	.loc 2 28 1
	add.s64 	%rd22, %rd3, %rd17;
	ld.global.f32 	%f4, [%rd22];
	add.s64 	%rd23, %rd2, %rd17;
	ld.global.f32 	%f5, [%rd23];
	add.s64 	%rd24, %rd1, %rd17;
	ld.global.f32 	%f6, [%rd24];
	.loc 2 35 1
	add.s32 	%r55, %r3, -1;
	mov.u32 	%r56, 0;
	.loc 3 238 5
	max.s32 	%r57, %r55, %r56;
	.loc 2 35 1
	mad.lo.s32 	%r58, %r47, %r37, %r57;
	.loc 2 36 1
	cvt.s64.s32 	%rd25, %r58;
	mul.wide.s32 	%rd26, %r58, 4;
	add.s64 	%rd27, %rd7, %rd26;
	ld.global.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd6, %rd26;
	ld.global.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd5, %rd26;
	ld.global.f32 	%f9, [%rd29];
	.loc 2 37 1
	add.s64 	%rd30, %rd4, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	.loc 2 27 1
	ld.global.u8 	%rc1, [%rd21];
	.loc 2 37 1
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r6, %rc2;
	@%p6 bra 	BB0_3;

	cvt.s32.s8 	%r64, %rc1;
	add.s32 	%r65, %r64, 1;
	mul.lo.s32 	%r66, %r65, %r64;
	shr.u32 	%r67, %r66, 31;
	mad.lo.s32 	%r68, %r65, %r64, %r67;
	shr.s32 	%r69, %r68, 1;
	add.s32 	%r179, %r6, %r69;
	bra.uni 	BB0_4;

BB0_3:
	.loc 2 37 1
	add.s32 	%r70, %r6, 1;
	mul.lo.s32 	%r71, %r70, %r6;
	shr.u32 	%r72, %r71, 31;
	mad.lo.s32 	%r73, %r70, %r6, %r72;
	shr.s32 	%r74, %r73, 1;
	cvt.s32.s8 	%r75, %rc1;
	add.s32 	%r179, %r74, %r75;

BB0_4:
	cvta.to.global.u64 	%rd32, %rd15;
	.loc 2 37 1
	mul.wide.s32 	%rd33, %r179, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 2 38 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 2 38 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 2 41 1
	add.s32 	%r77, %r37, -1;
	add.s32 	%r78, %r3, 1;
	.loc 3 210 5
	min.s32 	%r79, %r78, %r77;
	.loc 2 41 1
	add.s32 	%r80, %r79, %r5;
	.loc 2 42 1
	cvt.s64.s32 	%rd35, %r80;
	mul.wide.s32 	%rd37, %r80, 4;
	add.s64 	%rd38, %rd7, %rd37;
	ld.global.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd6, %rd37;
	ld.global.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd5, %rd37;
	ld.global.f32 	%f15, [%rd42];
	.loc 2 43 1
	add.s64 	%rd43, %rd4, %rd35;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r10, %rc1;
	cvt.s32.s8 	%r11, %rc3;
	@%p7 bra 	BB0_6;

	add.s32 	%r85, %r10, 1;
	mul.lo.s32 	%r86, %r85, %r10;
	shr.u32 	%r87, %r86, 31;
	mad.lo.s32 	%r88, %r85, %r10, %r87;
	shr.s32 	%r89, %r88, 1;
	add.s32 	%r180, %r11, %r89;
	bra.uni 	BB0_7;

BB0_6:
	.loc 2 43 1
	add.s32 	%r90, %r11, 1;
	mul.lo.s32 	%r91, %r90, %r11;
	shr.u32 	%r92, %r91, 31;
	mad.lo.s32 	%r93, %r90, %r11, %r92;
	shr.s32 	%r94, %r93, 1;
	add.s32 	%r180, %r94, %r10;

BB0_7:
	mul.wide.s32 	%rd46, %r180, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 2 44 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 2 44 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 2 47 1
	add.s32 	%r96, %r2, -1;
	.loc 3 238 5
	max.s32 	%r98, %r96, %r56;
	.loc 2 47 1
	add.s32 	%r99, %r98, %r4;
	mad.lo.s32 	%r100, %r99, %r37, %r3;
	.loc 2 48 1
	cvt.s64.s32 	%rd48, %r100;
	mul.wide.s32 	%rd50, %r100, 4;
	add.s64 	%rd51, %rd7, %rd50;
	ld.global.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd6, %rd50;
	ld.global.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd5, %rd50;
	ld.global.f32 	%f21, [%rd55];
	.loc 2 49 1
	add.s64 	%rd56, %rd4, %rd48;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r16, %rc4;
	@%p8 bra 	BB0_9;

	add.s32 	%r105, %r10, 1;
	mul.lo.s32 	%r106, %r105, %r10;
	shr.u32 	%r107, %r106, 31;
	mad.lo.s32 	%r108, %r105, %r10, %r107;
	shr.s32 	%r109, %r108, 1;
	add.s32 	%r181, %r16, %r109;
	bra.uni 	BB0_10;

BB0_9:
	.loc 2 49 1
	add.s32 	%r110, %r16, 1;
	mul.lo.s32 	%r111, %r110, %r16;
	shr.u32 	%r112, %r111, 31;
	mad.lo.s32 	%r113, %r110, %r16, %r112;
	shr.s32 	%r114, %r113, 1;
	add.s32 	%r181, %r114, %r10;

BB0_10:
	mul.wide.s32 	%rd59, %r181, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 2 50 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 2 50 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 2 53 1
	add.s32 	%r116, %r36, -1;
	add.s32 	%r117, %r2, 1;
	.loc 3 210 5
	min.s32 	%r118, %r117, %r116;
	.loc 2 53 1
	add.s32 	%r119, %r118, %r4;
	mad.lo.s32 	%r120, %r119, %r37, %r3;
	.loc 2 54 1
	cvt.s64.s32 	%rd61, %r120;
	mul.wide.s32 	%rd63, %r120, 4;
	add.s64 	%rd64, %rd7, %rd63;
	ld.global.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd6, %rd63;
	ld.global.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd5, %rd63;
	ld.global.f32 	%f27, [%rd68];
	.loc 2 55 1
	add.s64 	%rd69, %rd4, %rd61;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r21, %rc5;
	@%p9 bra 	BB0_12;

	add.s32 	%r125, %r10, 1;
	mul.lo.s32 	%r126, %r125, %r10;
	shr.u32 	%r127, %r126, 31;
	mad.lo.s32 	%r128, %r125, %r10, %r127;
	shr.s32 	%r129, %r128, 1;
	add.s32 	%r182, %r21, %r129;
	bra.uni 	BB0_13;

BB0_12:
	.loc 2 55 1
	add.s32 	%r130, %r21, 1;
	mul.lo.s32 	%r131, %r130, %r21;
	shr.u32 	%r132, %r131, 31;
	mad.lo.s32 	%r133, %r130, %r21, %r132;
	shr.s32 	%r134, %r133, 1;
	add.s32 	%r182, %r134, %r10;

BB0_13:
	mul.wide.s32 	%rd71, %r182, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 2 56 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 2 56 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	.loc 2 59 1
	setp.eq.s32 	%p10, %r35, 1;
	@%p10 bra 	BB0_21;

	.loc 2 61 1
	add.s32 	%r136, %r1, -1;
	.loc 3 238 5
	max.s32 	%r138, %r136, %r56;
	.loc 2 61 1
	mad.lo.s32 	%r139, %r138, %r36, %r2;
	mad.lo.s32 	%r140, %r139, %r37, %r3;
	.loc 2 62 1
	cvt.s64.s32 	%rd74, %r140;
	mul.wide.s32 	%rd76, %r140, 4;
	add.s64 	%rd77, %rd7, %rd76;
	ld.global.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd6, %rd76;
	ld.global.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd5, %rd76;
	ld.global.f32 	%f33, [%rd81];
	.loc 2 63 1
	add.s64 	%rd82, %rd4, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r26, %rc6;
	@%p11 bra 	BB0_16;

	add.s32 	%r145, %r10, 1;
	mul.lo.s32 	%r146, %r145, %r10;
	shr.u32 	%r147, %r146, 31;
	mad.lo.s32 	%r148, %r145, %r10, %r147;
	shr.s32 	%r149, %r148, 1;
	add.s32 	%r183, %r26, %r149;
	bra.uni 	BB0_17;

BB0_16:
	.loc 2 63 1
	add.s32 	%r150, %r26, 1;
	mul.lo.s32 	%r151, %r150, %r26;
	shr.u32 	%r152, %r151, 31;
	mad.lo.s32 	%r153, %r150, %r26, %r152;
	shr.s32 	%r154, %r153, 1;
	add.s32 	%r183, %r154, %r10;

BB0_17:
	mul.wide.s32 	%rd85, %r183, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 2 64 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 2 64 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 2 67 1
	add.s32 	%r156, %r35, -1;
	add.s32 	%r157, %r1, 1;
	.loc 3 210 5
	min.s32 	%r158, %r157, %r156;
	.loc 2 67 1
	mad.lo.s32 	%r159, %r158, %r36, %r2;
	mad.lo.s32 	%r160, %r159, %r37, %r3;
	.loc 2 68 1
	cvt.s64.s32 	%rd87, %r160;
	mul.wide.s32 	%rd89, %r160, 4;
	add.s64 	%rd90, %rd7, %rd89;
	ld.global.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd6, %rd89;
	ld.global.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd5, %rd89;
	ld.global.f32 	%f39, [%rd94];
	.loc 2 69 1
	add.s64 	%rd95, %rd4, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r31, %rc7;
	@%p12 bra 	BB0_19;

	add.s32 	%r165, %r10, 1;
	mul.lo.s32 	%r166, %r165, %r10;
	shr.u32 	%r167, %r166, 31;
	mad.lo.s32 	%r168, %r165, %r10, %r167;
	shr.s32 	%r169, %r168, 1;
	add.s32 	%r184, %r31, %r169;
	bra.uni 	BB0_20;

BB0_19:
	.loc 2 69 1
	add.s32 	%r170, %r31, 1;
	mul.lo.s32 	%r171, %r170, %r31;
	shr.u32 	%r172, %r171, 31;
	mad.lo.s32 	%r173, %r170, %r31, %r172;
	shr.s32 	%r174, %r173, 1;
	add.s32 	%r184, %r174, %r10;

BB0_20:
	mul.wide.s32 	%rd97, %r184, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 2 70 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 2 70 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB0_21:
	.loc 2 28 1
	shl.b64 	%rd100, %rd8, 2;
	add.s64 	%rd101, %rd3, %rd100;
	.loc 2 73 1
	st.global.f32 	[%rd101], %f79;
	.loc 2 28 1
	add.s64 	%rd103, %rd2, %rd100;
	.loc 2 74 1
	st.global.f32 	[%rd103], %f80;
	.loc 2 28 1
	add.s64 	%rd105, %rd1, %rd100;
	.loc 2 75 1
	st.global.f32 	[%rd105], %f81;

BB0_22:
	.loc 2 76 2
	ret;
}


`
	addexchange_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<13>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<164>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<106>;


	ld.param.u64 	%rd9, [addexchange_param_0];
	ld.param.u64 	%rd10, [addexchange_param_1];
	ld.param.u64 	%rd11, [addexchange_param_2];
	ld.param.u64 	%rd12, [addexchange_param_3];
	ld.param.u64 	%rd13, [addexchange_param_4];
	ld.param.u64 	%rd14, [addexchange_param_5];
	ld.param.u64 	%rd15, [addexchange_param_6];
	ld.param.u64 	%rd16, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r35, [addexchange_param_11];
	ld.param.u32 	%r36, [addexchange_param_12];
	ld.param.u32 	%r37, [addexchange_param_13];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd16;
	cvta.to.global.u64 	%rd5, %rd14;
	cvta.to.global.u64 	%rd6, %rd13;
	cvta.to.global.u64 	%rd7, %rd12;
	.loc 3 16 1
	mov.u32 	%r38, %ntid.z;
	mov.u32 	%r39, %ctaid.z;
	mov.u32 	%r40, %tid.z;
	mad.lo.s32 	%r1, %r38, %r39, %r40;
	.loc 3 17 1
	mov.u32 	%r41, %ntid.y;
	mov.u32 	%r42, %ctaid.y;
	mov.u32 	%r43, %tid.y;
	mad.lo.s32 	%r2, %r41, %r42, %r43;
	.loc 3 18 1
	mov.u32 	%r44, %ntid.x;
	mov.u32 	%r45, %ctaid.x;
	mov.u32 	%r46, %tid.x;
	mad.lo.s32 	%r3, %r44, %r45, %r46;
	.loc 3 20 1
	setp.ge.s32 	%p1, %r2, %r36;
	setp.ge.s32 	%p2, %r1, %r35;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r3, %r37;
	or.pred  	%p5, %p3, %p4;
	@%p5 bra 	BB2_22;

	.loc 3 25 1
	mul.lo.s32 	%r4, %r1, %r36;
	mad.lo.s32 	%r47, %r1, %r36, %r2;
	mul.lo.s32 	%r5, %r47, %r37;
	mad.lo.s32 	%r48, %r47, %r37, %r3;
	.loc 3 26 1
	cvt.s64.s32 	%rd8, %r48;
	mul.wide.s32 	%rd17, %r48, 4;
	add.s64 	%rd18, %rd7, %rd17;
	ld.global.nc.f32 	%f1, [%rd18];
	add.s64 	%rd19, %rd6, %rd17;
	ld.global.nc.f32 	%f2, [%rd19];
	add.s64 	%rd20, %rd5, %rd17;
	ld.global.nc.f32 	%f3, [%rd20];
	.loc 3 27 1
	add.s64 	%rd21, %rd4, %rd8;
	ld.global.u8 	%rc1, [%rd21];
	.loc 3 28 1
	add.s64 	%rd22, %rd3, %rd17;
	ld.global.f32 	%f4, [%rd22];
	add.s64 	%rd23, %rd2, %rd17;
	ld.global.f32 	%f5, [%rd23];
	add.s64 	%rd24, %rd1, %rd17;
	ld.global.f32 	%f6, [%rd24];
	.loc 3 35 1
	add.s32 	%r53, %r3, -1;
	mov.u32 	%r54, 0;
	.loc 4 238 5
	max.s32 	%r55, %r53, %r54;
	.loc 3 35 1
	mad.lo.s32 	%r56, %r47, %r37, %r55;
	.loc 3 36 1
	cvt.s64.s32 	%rd25, %r56;
	mul.wide.s32 	%rd26, %r56, 4;
	add.s64 	%rd27, %rd7, %rd26;
	ld.global.nc.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd6, %rd26;
	ld.global.nc.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd5, %rd26;
	ld.global.nc.f32 	%f9, [%rd29];
	.loc 3 37 1
	add.s64 	%rd30, %rd4, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r6, %rc2;
	@%p6 bra 	BB2_3;

	cvt.s32.s8 	%r58, %rc1;
	add.s32 	%r59, %r58, 1;
	mul.lo.s32 	%r60, %r59, %r58;
	shr.u32 	%r61, %r60, 31;
	mad.lo.s32 	%r62, %r59, %r58, %r61;
	shr.s32 	%r63, %r62, 1;
	add.s32 	%r158, %r6, %r63;
	bra.uni 	BB2_4;

BB2_3:
	.loc 3 37 1
	add.s32 	%r64, %r6, 1;
	mul.lo.s32 	%r65, %r64, %r6;
	shr.u32 	%r66, %r65, 31;
	mad.lo.s32 	%r67, %r64, %r6, %r66;
	shr.s32 	%r68, %r67, 1;
	cvt.s32.s8 	%r69, %rc1;
	add.s32 	%r158, %r68, %r69;

BB2_4:
	cvta.to.global.u64 	%rd32, %rd15;
	.loc 3 37 1
	mul.wide.s32 	%rd33, %r158, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 3 38 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 3 38 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 3 41 1
	add.s32 	%r71, %r37, -1;
	add.s32 	%r72, %r3, 1;
	.loc 4 210 5
	min.s32 	%r73, %r72, %r71;
	.loc 3 41 1
	add.s32 	%r74, %r73, %r5;
	.loc 3 42 1
	cvt.s64.s32 	%rd35, %r74;
	mul.wide.s32 	%rd37, %r74, 4;
	add.s64 	%rd38, %rd7, %rd37;
	ld.global.nc.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd6, %rd37;
	ld.global.nc.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd5, %rd37;
	ld.global.nc.f32 	%f15, [%rd42];
	.loc 3 43 1
	add.s64 	%rd43, %rd4, %rd35;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r10, %rc1;
	cvt.s32.s8 	%r11, %rc3;
	@%p7 bra 	BB2_6;

	add.s32 	%r76, %r10, 1;
	mul.lo.s32 	%r77, %r76, %r10;
	shr.u32 	%r78, %r77, 31;
	mad.lo.s32 	%r79, %r76, %r10, %r78;
	shr.s32 	%r80, %r79, 1;
	add.s32 	%r159, %r11, %r80;
	bra.uni 	BB2_7;

BB2_6:
	.loc 3 43 1
	add.s32 	%r81, %r11, 1;
	mul.lo.s32 	%r82, %r81, %r11;
	shr.u32 	%r83, %r82, 31;
	mad.lo.s32 	%r84, %r81, %r11, %r83;
	shr.s32 	%r85, %r84, 1;
	add.s32 	%r159, %r85, %r10;

BB2_7:
	mul.wide.s32 	%rd46, %r159, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 3 44 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 3 44 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 3 47 1
	add.s32 	%r87, %r2, -1;
	.loc 4 238 5
	max.s32 	%r89, %r87, %r54;
	.loc 3 47 1
	add.s32 	%r90, %r89, %r4;
	mad.lo.s32 	%r91, %r90, %r37, %r3;
	.loc 3 48 1
	cvt.s64.s32 	%rd48, %r91;
	mul.wide.s32 	%rd50, %r91, 4;
	add.s64 	%rd51, %rd7, %rd50;
	ld.global.nc.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd6, %rd50;
	ld.global.nc.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd5, %rd50;
	ld.global.nc.f32 	%f21, [%rd55];
	.loc 3 49 1
	add.s64 	%rd56, %rd4, %rd48;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r16, %rc4;
	@%p8 bra 	BB2_9;

	add.s32 	%r93, %r10, 1;
	mul.lo.s32 	%r94, %r93, %r10;
	shr.u32 	%r95, %r94, 31;
	mad.lo.s32 	%r96, %r93, %r10, %r95;
	shr.s32 	%r97, %r96, 1;
	add.s32 	%r160, %r16, %r97;
	bra.uni 	BB2_10;

BB2_9:
	.loc 3 49 1
	add.s32 	%r98, %r16, 1;
	mul.lo.s32 	%r99, %r98, %r16;
	shr.u32 	%r100, %r99, 31;
	mad.lo.s32 	%r101, %r98, %r16, %r100;
	shr.s32 	%r102, %r101, 1;
	add.s32 	%r160, %r102, %r10;

BB2_10:
	mul.wide.s32 	%rd59, %r160, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 3 50 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 3 50 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 3 53 1
	add.s32 	%r104, %r36, -1;
	add.s32 	%r105, %r2, 1;
	.loc 4 210 5
	min.s32 	%r106, %r105, %r104;
	.loc 3 53 1
	add.s32 	%r107, %r106, %r4;
	mad.lo.s32 	%r108, %r107, %r37, %r3;
	.loc 3 54 1
	cvt.s64.s32 	%rd61, %r108;
	mul.wide.s32 	%rd63, %r108, 4;
	add.s64 	%rd64, %rd7, %rd63;
	ld.global.nc.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd6, %rd63;
	ld.global.nc.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd5, %rd63;
	ld.global.nc.f32 	%f27, [%rd68];
	.loc 3 55 1
	add.s64 	%rd69, %rd4, %rd61;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r21, %rc5;
	@%p9 bra 	BB2_12;

	add.s32 	%r110, %r10, 1;
	mul.lo.s32 	%r111, %r110, %r10;
	shr.u32 	%r112, %r111, 31;
	mad.lo.s32 	%r113, %r110, %r10, %r112;
	shr.s32 	%r114, %r113, 1;
	add.s32 	%r161, %r21, %r114;
	bra.uni 	BB2_13;

BB2_12:
	.loc 3 55 1
	add.s32 	%r115, %r21, 1;
	mul.lo.s32 	%r116, %r115, %r21;
	shr.u32 	%r117, %r116, 31;
	mad.lo.s32 	%r118, %r115, %r21, %r117;
	shr.s32 	%r119, %r118, 1;
	add.s32 	%r161, %r119, %r10;

BB2_13:
	mul.wide.s32 	%rd71, %r161, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 3 56 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 3 56 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	.loc 3 59 1
	setp.eq.s32 	%p10, %r35, 1;
	@%p10 bra 	BB2_21;

	.loc 3 61 1
	add.s32 	%r121, %r1, -1;
	.loc 4 238 5
	max.s32 	%r123, %r121, %r54;
	.loc 3 61 1
	mad.lo.s32 	%r124, %r123, %r36, %r2;
	mad.lo.s32 	%r125, %r124, %r37, %r3;
	.loc 3 62 1
	cvt.s64.s32 	%rd74, %r125;
	mul.wide.s32 	%rd76, %r125, 4;
	add.s64 	%rd77, %rd7, %rd76;
	ld.global.nc.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd6, %rd76;
	ld.global.nc.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd5, %rd76;
	ld.global.nc.f32 	%f33, [%rd81];
	.loc 3 63 1
	add.s64 	%rd82, %rd4, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r26, %rc6;
	@%p11 bra 	BB2_16;

	add.s32 	%r127, %r10, 1;
	mul.lo.s32 	%r128, %r127, %r10;
	shr.u32 	%r129, %r128, 31;
	mad.lo.s32 	%r130, %r127, %r10, %r129;
	shr.s32 	%r131, %r130, 1;
	add.s32 	%r162, %r26, %r131;
	bra.uni 	BB2_17;

BB2_16:
	.loc 3 63 1
	add.s32 	%r132, %r26, 1;
	mul.lo.s32 	%r133, %r132, %r26;
	shr.u32 	%r134, %r133, 31;
	mad.lo.s32 	%r135, %r132, %r26, %r134;
	shr.s32 	%r136, %r135, 1;
	add.s32 	%r162, %r136, %r10;

BB2_17:
	mul.wide.s32 	%rd85, %r162, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 3 64 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 3 64 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 3 67 1
	add.s32 	%r138, %r35, -1;
	add.s32 	%r139, %r1, 1;
	.loc 4 210 5
	min.s32 	%r140, %r139, %r138;
	.loc 3 67 1
	mad.lo.s32 	%r141, %r140, %r36, %r2;
	mad.lo.s32 	%r142, %r141, %r37, %r3;
	.loc 3 68 1
	cvt.s64.s32 	%rd87, %r142;
	mul.wide.s32 	%rd89, %r142, 4;
	add.s64 	%rd90, %rd7, %rd89;
	ld.global.nc.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd6, %rd89;
	ld.global.nc.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd5, %rd89;
	ld.global.nc.f32 	%f39, [%rd94];
	.loc 3 69 1
	add.s64 	%rd95, %rd4, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r31, %rc7;
	@%p12 bra 	BB2_19;

	add.s32 	%r144, %r10, 1;
	mul.lo.s32 	%r145, %r144, %r10;
	shr.u32 	%r146, %r145, 31;
	mad.lo.s32 	%r147, %r144, %r10, %r146;
	shr.s32 	%r148, %r147, 1;
	add.s32 	%r163, %r31, %r148;
	bra.uni 	BB2_20;

BB2_19:
	.loc 3 69 1
	add.s32 	%r149, %r31, 1;
	mul.lo.s32 	%r150, %r149, %r31;
	shr.u32 	%r151, %r150, 31;
	mad.lo.s32 	%r152, %r149, %r31, %r151;
	shr.s32 	%r153, %r152, 1;
	add.s32 	%r163, %r153, %r10;

BB2_20:
	mul.wide.s32 	%rd97, %r163, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 3 70 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 3 70 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB2_21:
	.loc 3 28 1
	shl.b64 	%rd100, %rd8, 2;
	add.s64 	%rd101, %rd3, %rd100;
	.loc 3 73 1
	st.global.f32 	[%rd101], %f79;
	.loc 3 28 1
	add.s64 	%rd103, %rd2, %rd100;
	.loc 3 74 1
	st.global.f32 	[%rd103], %f80;
	.loc 3 28 1
	add.s64 	%rd105, %rd1, %rd100;
	.loc 3 75 1
	st.global.f32 	[%rd105], %f81;

BB2_22:
	.loc 3 76 2
	ret;
}


`
)
