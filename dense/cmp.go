// Code generated by genlib3. DO NOT EDIT

package dense

import (
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/internal/errors"
)

func (t *Dense[DT]) basicCmpPrep(u *Dense[DT], opts ...FuncOpt) (e Engine, newAPT, newAPU *tensor.AP, retVal DescWithStorage, fo Option, err error) {
	e = getEngine[DT](t, u)
	if err = check(checkFlags(e, t, u)); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	tShp := t.Shape()
	uShp := u.Shape()
	expShape := largestShape(tShp, uShp)

	var prepper tensor.DescFuncOptHandler[DT]
	var ok bool
	if prepper, ok = e.(tensor.DescFuncOptHandler[DT]); !ok {
		return nil, nil, nil, nil, fo, errors.Errorf(errors.EngineSupport, e, prepper, errors.ThisFn(1))
	}

	opts = defaultCmpFuncOpt(opts)
	if retVal, fo, err = prepper.HandleFuncOptsDesc(t, expShape, opts...); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedFuncOpt, errors.ThisFn(1))
	}

	newAPT = t.Info()
	newAPU = u.Info()

	// fast path
	if !fo.Broadcast || tShp.TotalSize() == uShp.TotalSize() {
		// no broadcasting necessary
		fo.Broadcast = false
		return
	}

	newAPT, newAPU = tensor.CalcBroadcastShapes(newAPT, newAPU)
	if err = tensor.CheckBroadcastable(newAPT.Shape(), newAPU.Shape()); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	return
}

// Lt performs `t < u`
func (t *Dense[DT]) Lt(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.Ord[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for Lt")
	}

	switch {
	case toBroadcast:
		err = cmper.LtBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.Lt(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}

// Lte performs `t <= u`
func (t *Dense[DT]) Lte(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.Ord[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for Lte")
	}

	switch {
	case toBroadcast:
		err = cmper.LteBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.Lte(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}

// Gt performs `t > u`
func (t *Dense[DT]) Gt(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.FullOrd[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for Gt")
	}

	switch {
	case toBroadcast:
		err = cmper.GtBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.Gt(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}

// Gte performs `t >= u`
func (t *Dense[DT]) Gte(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.FullOrd[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for Gte")
	}

	switch {
	case toBroadcast:
		err = cmper.GteBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.Gte(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}

// ElEq performs `t == u`
func (t *Dense[DT]) ElEq(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.Comparer[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for ElEq")
	}

	switch {
	case toBroadcast:
		err = cmper.ElEqBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.ElEq(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}

// Ne performs `t != u`
func (t *Dense[DT]) Ne(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := t.basicCmpPrep(u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	cmper, ok := e.(tensor.Comparer[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for Ne")
	}

	switch {
	case toBroadcast:
		err = cmper.NeBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		err = cmper.Ne(ctx, t, u, retVal, asSame)

	}
	return retVal, nil
}
