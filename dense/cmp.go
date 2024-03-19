// Code generated by genlib3. DO NOT EDIT

package dense

import (
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/internal/errors"
)

// Lt performs `t < u`
func (t *Dense[DT]) Lt(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

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
		if err = cmper.Lt(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}

// Lte performs `t <= u`
func (t *Dense[DT]) Lte(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

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
		if err = cmper.Lte(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}

// Gt performs `t > u`
func (t *Dense[DT]) Gt(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

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
		if err = cmper.Gt(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}

// Gte performs `t >= u`
func (t *Dense[DT]) Gte(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

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
		if err = cmper.Gte(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}

// ElEq performs `t == u`
func (t *Dense[DT]) ElEq(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

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
		if err = cmper.ElEq(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}

// ElNe performs `t != u`
func (t *Dense[DT]) ElNe(u *Dense[DT], opts ...FuncOpt) (retVal DescWithStorage, err error) {
	e, newAPT, newAPU, retVal, fo, err := tensor.PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}
	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast.BroadcastData()

	cmper, ok := e.(tensor.Comparer[DT, *Dense[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, cmper, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to Incr for ElNe")
	}

	switch {
	case toBroadcast:
		err = cmper.ElNeBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return retVal, err
		}
		if err = cmper.ElNe(ctx, t, u, retVal, asSame); err != nil {
			return nil, err
		}
		err = postOpBroadcastReshape(fo.Broadcast, t, u, retVal)
	}
	return retVal, nil
}
