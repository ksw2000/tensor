// Code generated by genlib2. DO NOT EDIT.

package tensor

import (
	"github.com/pkg/errors"
	"gorgonia.org/tensor/internal/storage"
)

// Gt performs a > b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) Gt(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Gt failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Gt")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.GtSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.GtSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.GtIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.GtSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.GtSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Gt(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// Gte performs a ≥ b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) Gte(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Gte failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Gte")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.GteSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.GteSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.GteIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.GteSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.GteSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Gte(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// Lt performs a < b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) Lt(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Lt failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Lt")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.LtSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.LtSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.LtIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.LtSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.LtSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Lt(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// Lte performs a ≤ b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) Lte(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Lte failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Lte")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.LteSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.LteSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.LteIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.LteSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.LteSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Lte(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// ElEq performs a == b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) ElEq(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, eqTypes); err != nil {
		return nil, errors.Wrapf(err, "Eq failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Eq")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.EqSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.EqSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.EqIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.EqSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.EqSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Eq(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// ElNe performs a ≠ b elementwise. Both a and b must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) ElNe(a Tensor, b Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = binaryCheck(a, b, eqTypes); err != nil {
		return nil, errors.Wrapf(err, "Ne failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	typ := a.Dtype().Type
	var dataA, dataB, dataReuse *storage.Header
	var ait, bit, iit Iterator
	var useIter, swap bool
	if dataA, dataB, dataReuse, ait, bit, iit, useIter, swap, err = prepDataVV(a, b, reuse); err != nil {
		return nil, errors.Wrapf(err, "StdEng.Ne")
	}
	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		if swap {
			reuse = NewDense(b.Dtype(), b.Shape().Clone(), WithEngine(e))
		} else {
			reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		}
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.NeSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.NeSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.NeIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}

		return
	}

	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.NeSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.NeSame(typ, dataReuse, dataB)
		retVal = reuse
	default:
		err = e.E.Ne(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}

	return
}

// GtScalar performs t > s elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in s
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) GtScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Gt failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Gt failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Gt")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Gt")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.GtSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.GtSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.GtSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.GtIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.GtSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.LtSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.GtSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.GtSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.GtSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Gt(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}

// GteScalar performs t ≥ s elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in s
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) GteScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Gte failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Gte failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Gte")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Gte")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.GteSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.GteSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.GteSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.GteIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.GteSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.LteSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.GteSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.GteSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.GteSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Gte(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}

// LtScalar performs t < s elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in s
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) LtScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Lt failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Lt failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Lt")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Lt")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.LtSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.LtSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.LtSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.LtIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.LtSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.GtSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.LtSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.LtSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.LtSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Lt(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}

// LteScalar performs t ≤ s elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in s
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (e StdEng) LteScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, ordTypes); err != nil {
		return nil, errors.Wrapf(err, "Lte failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Lte failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Lte")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Lte")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.LteSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.LteSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.LteSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.LteIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.LteSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.GteSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.LteSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.LteSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.LteSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Lte(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}

func (e StdEng) EqScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, eqTypes); err != nil {
		return nil, errors.Wrapf(err, "Eq failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Eq failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Eq")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Eq")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.EqSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.EqSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.EqSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.EqIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.EqSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.EqSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.EqSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.EqSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.EqSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Eq(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}

func (e StdEng) NeScalar(t Tensor, s interface{}, leftTensor bool, opts ...FuncOpt) (retVal Tensor, err error) {
	if err = unaryCheck(t, eqTypes); err != nil {
		return nil, errors.Wrapf(err, "Ne failed")
	}

	if err = scalarDtypeCheck(t, s); err != nil {
		return nil, errors.Wrap(err, "Ne failed")
	}

	var reuse DenseTensor
	var safe, same bool
	if reuse, safe, _, _, same, err = handleFuncOpts(t.Shape(), t.Dtype(), t.DataOrder(), false, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}
	if !safe {
		same = true
	}
	a := t
	typ := t.Dtype().Type
	var ait, bit, iit Iterator
	var dataA, dataB, dataReuse, scalarHeader *storage.Header
	var useIter bool

	if leftTensor {
		if dataA, dataB, dataReuse, ait, iit, useIter, err = prepDataVS(t, s, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Ne")
		}
		scalarHeader = dataB
	} else {
		if dataA, dataB, dataReuse, bit, iit, useIter, err = prepDataSV(s, t, reuse); err != nil {
			return nil, errors.Wrapf(err, opFail, "StdEng.Ne")
		}
		scalarHeader = dataA
	}

	// check to see if anything needs to be created
	switch {
	case same && safe && reuse == nil:
		reuse = NewDense(a.Dtype(), a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	case !same && safe && reuse == nil:
		reuse = NewDense(Bool, a.Shape().Clone(), WithEngine(e))
		dataReuse = reuse.hdr()
		if useIter {
			iit = IteratorFromDense(reuse)
		}
	}

	if useIter {
		switch {
		case !safe && same && reuse == nil:
			err = e.E.NeSameIter(typ, dataA, dataB, ait, bit)
			retVal = a
		case same && safe && reuse != nil && !leftTensor:
			storage.CopyIter(typ, dataReuse, dataB, iit, bit)
			bit.Reset()
			iit.Reset()
			err = e.E.NeSameIter(typ, dataA, dataReuse, ait, bit)
			retVal = reuse
		case same && safe && reuse != nil && leftTensor:
			storage.CopyIter(typ, dataReuse, dataA, iit, ait)
			ait.Reset()
			iit.Reset()
			err = e.E.NeSameIter(typ, dataReuse, dataB, iit, bit)
			retVal = reuse
		default: // safe && bool
			err = e.E.NeIter(typ, dataA, dataB, dataReuse, ait, bit, iit)
			retVal = reuse
		}
		returnHeader(scalarHeader)
		return
	}

	// handle special case where A and B have both len 1
	if len(dataA.Raw) == int(typ.Size()) && len(dataB.Raw) == int(typ.Size()) {
		switch {
		case same && safe && reuse != nil && leftTensor:
			storage.Copy(typ, dataReuse, dataA)
			err = e.E.NeSame(typ, dataReuse, dataB)
			retVal = reuse
			return
		case same && safe && reuse != nil && !leftTensor:
			storage.Copy(typ, dataReuse, dataB)
			err = e.E.NeSame(typ, dataReuse, dataA)
			retVal = reuse
			return
		}
	}
	// standard
	switch {
	case !safe && same && reuse == nil:
		err = e.E.NeSame(typ, dataA, dataB)
		retVal = a
	case same && safe && reuse != nil && leftTensor:
		storage.Copy(typ, dataReuse, dataA)
		err = e.E.NeSame(typ, dataReuse, dataB)
		retVal = reuse
	case same && safe && reuse != nil && !leftTensor:
		storage.Copy(typ, dataReuse, dataB)
		err = e.E.NeSame(typ, dataA, dataReuse)
		retVal = reuse
	default:
		err = e.E.Ne(typ, dataA, dataB, dataReuse)
		retVal = reuse
	}
	returnHeader(scalarHeader)
	return
}
