package exec

import (
	"fmt"
	"io"
	"net/http"
)

func fetchURL(eCtx ExecCtx, args map[string]string) (ExecCtx, error) {
	url, hasURL := args["url"]
	if !hasURL {
		return eCtx, fmt.Errorf("no url provided for fetch with a source of url")
	}

	resp, err := http.Get(url)
	if err != nil {
		return eCtx, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return eCtx, fmt.Errorf("http GET failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return eCtx, err
	}

	fieldName, customNameDefined := args["name"]
	if !customNameDefined {
		fieldName = "result"
	}

	return NewExecCtx(eCtx, ExecCtx{fieldName: string(body)}), err
}

func fetch(eCtx ExecCtx, args map[string]string) (ExecCtx, error) {
	source, hasSource := args["source"]

	if !hasSource {
		return eCtx, fmt.Errorf("no source provided for fetch")
	}

	switch source {
	case "url":
		return fetchURL(eCtx, args)

	}

	return eCtx, fmt.Errorf("%q is not a valid source for fetch", source)
}
