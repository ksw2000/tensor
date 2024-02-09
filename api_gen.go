// Code generated by genlib3. DO NOT EDIT

package tensor

import "gorgonia.org/tensor/internal/errors"

// Add performs `t + u`.
func Add[DT Num](t, u Basic[DT], opts ...FuncOpt) (Basic[DT], error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBasicBinOpCis[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	ctx := fo.Ctx
	toIncr := fo.Incr
	toBroadcast := fo.Broadcast

	adder, ok := e.(Adder[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, adder, errors.ThisFn())
	}
	switch {
	case toBroadcast:
		err = adder.AddBroadcastable(ctx, t, u, retVal, newAPT, newAPU, toIncr)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = adder.Add(ctx, t, u, retVal, toIncr)

	}
	return retVal, err
}

// Sub performs `t - u`.
func Sub[DT Num](t, u Basic[DT], opts ...FuncOpt) (Basic[DT], error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBasicBinOpCis[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	ctx := fo.Ctx
	toIncr := fo.Incr
	toBroadcast := fo.Broadcast

	basicarither, ok := e.(BasicArither[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, basicarither, errors.ThisFn())
	}
	switch {
	case toBroadcast:
		err = basicarither.SubBroadcastable(ctx, t, u, retVal, newAPT, newAPU, toIncr)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = basicarither.Sub(ctx, t, u, retVal, toIncr)

	}
	return retVal, err
}

// Mul performs `t * u`.
func Mul[DT Num](t, u Basic[DT], opts ...FuncOpt) (Basic[DT], error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBasicBinOpCis[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	ctx := fo.Ctx
	toIncr := fo.Incr
	toBroadcast := fo.Broadcast

	basicarither, ok := e.(BasicArither[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, basicarither, errors.ThisFn())
	}
	switch {
	case toBroadcast:
		err = basicarither.MulBroadcastable(ctx, t, u, retVal, newAPT, newAPU, toIncr)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = basicarither.Mul(ctx, t, u, retVal, toIncr)

	}
	return retVal, err
}

// Div performs `t / u`.
func Div[DT Num](t, u Basic[DT], opts ...FuncOpt) (Basic[DT], error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBasicBinOpCis[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	ctx := fo.Ctx
	toIncr := fo.Incr
	toBroadcast := fo.Broadcast

	basicarither, ok := e.(BasicArither[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, basicarither, errors.ThisFn())
	}
	switch {
	case toBroadcast:
		err = basicarither.DivBroadcastable(ctx, t, u, retVal, newAPT, newAPU, toIncr)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = basicarither.Div(ctx, t, u, retVal, toIncr)

	}
	return retVal, err
}

// Lt performs `t < u`.
func Lt[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	ord, ok := e.(Ord[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, ord, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for Lt")
	}
	switch {
	case toBroadcast:
		err = ord.LtBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = ord.Lt(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}

// Lte performs `t <= u`.
func Lte[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	ord, ok := e.(Ord[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, ord, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for Lte")
	}
	switch {
	case toBroadcast:
		err = ord.LteBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = ord.Lte(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}

// Gt performs `t > u`.
func Gt[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	fullord, ok := e.(FullOrd[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, fullord, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for Gt")
	}
	switch {
	case toBroadcast:
		err = fullord.GtBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = fullord.Gt(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}

// Gte performs `t >= u`.
func Gte[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	fullord, ok := e.(FullOrd[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, fullord, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for Gte")
	}
	switch {
	case toBroadcast:
		err = fullord.GteBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = fullord.Gte(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}

// ElEq performs `t == u`.
func ElEq[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	comparer, ok := e.(Comparer[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, comparer, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for ElEq")
	}
	switch {
	case toBroadcast:
		err = comparer.ElEqBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = comparer.ElEq(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}

// ElNe performs `t != u`.
func ElNe[DT Num](t, u Basic[DT], opts ...FuncOpt) (DescWithStorage, error) {
	e, newAPT, newAPU, retVal, fo, err := PrepBinOpTrans[DT](t, u, opts...)
	if err != nil {
		return nil, err
	}

	asSame := fo.AsType == t.Dtype()
	ctx := fo.Ctx
	toBroadcast := fo.Broadcast

	comparer, ok := e.(Comparer[DT, Basic[DT]])
	if !ok {
		return nil, errors.Errorf(errors.EngineSupport, e, comparer, errors.ThisFn())
	}
	if fo.Incr {
		return nil, errors.Errorf("Unable to perform Incr for ElNe")
	}
	switch {
	case toBroadcast:
		err = comparer.ElNeBroadcastable(ctx, t, u, retVal, asSame, newAPT, newAPU)
	default:
		if err := checkCompatibleShape(t.Shape(), u.Shape())(); err != nil {
			return nil, err
		}
		err = comparer.ElNe(ctx, t, u, retVal, asSame)

	}
	return retVal, err
}