package tensor

import (
	"gorgonia.org/dtype"
	"gorgonia.org/tensor/internal/errors"
)

// GetEngine gets the workhorse engine from the given list of tensors.
func GetEngine(ts ...Engineer) Engine {
	// TODO: get highest capability engine
	for _, t := range ts {
		if e := t.Engine(); e != nil {
			return e.Workhorse()
		}
	}
	return nil
}

// defaultCmpFuncOpt handles the FuncOpts of a comparison function. It adds a `As(dtype.Bool)` in the head of the func opts as by default
// comparison operations return bools.
func defaultCmpFuncOpt(opts []FuncOpt) []FuncOpt {
	opts = append(opts, nil)
	copy(opts[1:], opts[0:])
	opts[0] = As(dtype.Bool) // default
	return opts
}

// PrepBinOpCis is a function that preps two basic tensors for a elementwise binary operation that returns the a tensor of the same datatype as its inputs.
func PrepBinOpCis[DT any, T Tensor[DT, T]](a, b T, opts ...FuncOpt) (e Engine, newAPA, newAPB *AP, retVal T, fo Option, err error) {
	e = GetEngine(a, b)
	if err = check(checkFlags(e, a, b)); err != nil {
		return nil, nil, nil, retVal, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	aShp := a.Shape()
	bShp := b.Shape()
	cShp := getLargestShape(aShp, bShp)

	var prepper SpecializedFuncOptHandler[DT, T]
	var ok bool

	if prepper, ok = e.(SpecializedFuncOptHandler[DT, T]); !ok {
		return nil, nil, nil, retVal, fo, errors.Errorf(errors.EngineSupport, e, prepper, errors.ThisFn(1))
	}
	if retVal, fo, err = prepper.HandleFuncOptsSpecialized(a, cShp, opts...); err != nil {
		return nil, nil, nil, retVal, fo, errors.Wrapf(err, errors.FailedFuncOpt, errors.ThisFn(1))
	}

	newAPA = a.Info()
	newAPB = b.Info()

	// fast path
	if !fo.Broadcast || aShp.TotalSize() == bShp.TotalSize() {
		// no broadcasting necessary
		fo.Broadcast = false
		return
	}

	newAPA, newAPB = CalcBroadcastShapes(newAPA, newAPB)
	if err = CheckBroadcastable(newAPA.Shape(), newAPB.Shape()); err != nil {
		return nil, nil, nil, retVal, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	return
}

// PrepBasicBinOpCis is a function that preps two basic tensors for a elementwise binary operation that returns the a tensor of the same datatype as its inputs. It is like PrepBinOpCis except the input and output types are Basic[DT].
func PrepBasicBinOpCis[DT any](a, b Basic[DT], opts ...FuncOpt) (e Engine, newAPA, newAPB *AP, retVal Basic[DT], fo Option, err error) {
	e = GetEngine(a, b)
	aShp := a.Shape()
	bShp := b.Shape()
	cShp := getLargestShape(aShp, bShp)

	var prepper FuncOptHandler[DT]
	var ok bool
	if prepper, ok = e.(FuncOptHandler[DT]); !ok {
		return nil, nil, nil, nil, fo, errors.Errorf(errors.EngineSupport, e, prepper, errors.ThisFn(1))
	}

	if retVal, fo, err = prepper.HandleFuncOpts(a, cShp, opts...); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedFuncOpt, errors.ThisFn(1))
	}

	newAPA = a.Info()
	newAPB = b.Info()

	// fast path
	if !fo.Broadcast || aShp.TotalSize() == bShp.TotalSize() {
		// no broadcasting necessary
		fo.Broadcast = false
		return
	}

	newAPA, newAPB = CalcBroadcastShapes(newAPA, newAPB)
	if err = CheckBroadcastable(newAPA.Shape(), newAPB.Shape()); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	return
}

// PrepBinOpTrans is a function that preps two basic tensors for a comparison based binary operation.
func PrepBinOpTrans[DT any](a, b Basic[DT], opts ...FuncOpt) (e Engine, newAPA, newAPB *AP, retVal DescWithStorage, fo Option, err error) {
	e = GetEngine(a, b)
	if err = check(checkFlags(e, a, b)); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	aShp := a.Shape()
	bShp := b.Shape()
	cShp := getLargestShape(aShp, bShp)

	var prepper DescFuncOptHandler[DT]
	var ok bool
	if prepper, ok = e.(DescFuncOptHandler[DT]); !ok {
		return nil, nil, nil, nil, fo, errors.Errorf(errors.EngineSupport, e, prepper, errors.ThisFn(1))
	}

	opts = defaultCmpFuncOpt(opts)
	if retVal, fo, err = prepper.HandleFuncOptsDesc(a, cShp, opts...); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedFuncOpt, errors.ThisFn(1))
	}

	newAPA = a.Info()
	newAPB = b.Info()

	// fast path
	if !fo.Broadcast || aShp.TotalSize() == bShp.TotalSize() {
		// no broadcasting necessary
		fo.Broadcast = false
		return
	}

	newAPA, newAPB = CalcBroadcastShapes(newAPA, newAPB)
	if err = CheckBroadcastable(newAPA.Shape(), newAPB.Shape()); err != nil {
		return nil, nil, nil, nil, fo, errors.Wrapf(err, errors.FailedSanity, errors.ThisFn(1))
	}
	return

}