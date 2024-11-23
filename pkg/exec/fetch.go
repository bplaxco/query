package exec

import (
	"fmt"
	"io"
	"net/http"
)

func fetchURL(eCtx ExecCtx, args map[string]string) (ExecCtx, error) {
	resp, err := http.Get(args["url"])
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
		fieldName = "fetched_url"
	}

	return NewExecCtx(eCtx, ExecCtx{fieldName: string(body)}), err
}

func fetch(eCtx ExecCtx, args map[string]string) (ExecCtx, error) {
	if _, hasURL := args["url"]; hasURL {
		return fetchURL(eCtx, args)
	}

	return eCtx, fmt.Errorf("not enough options provided to fetch")
}
