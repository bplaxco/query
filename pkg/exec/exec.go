package exec

func ExecList(cmds []*Command) (eCtx ExecCtx, err error) {
	eCtx = NewExecCtx()

	for _, cmd := range cmds {
		eCtx, err = cmd.Exec(eCtx)

		if err != nil {
			break
		}
	}

	return eCtx, err
}
