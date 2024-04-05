package swagger

import (
	"context"
	"fmt"
)

func {{FUNCTION_NAME}}(body Body) {{FUNCTION_RETURN}} {
	cfg := NewConfiguration()
	apiclient := NewAPIClient(cfg)
	s, _, err := apiclient.DefaultApi.RootPost(context.Background(), body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}