package exec

type ExecCtx map[any]any

func NewExecCtx(eCtxs ...ExecCtx) ExecCtx {
	eCtx := make(ExecCtx)

	for _, e := range eCtxs {
		eCtx.Update(e)
	}

	return eCtx
}

func (c ExecCtx) Update(eCtx ExecCtx) {
	for k, v := range eCtx {
		c[k] = v
	}
}
