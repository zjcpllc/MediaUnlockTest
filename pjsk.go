package mediaunlocktest

import "net/http"

// Project Sekai: Colorful Stage
func PJSK(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://game-version.sekai.colorfulpalette.org/1.8.1/3ed70b6a-8352-4532-b819-108837926ff5", nil)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	req.Header.Set("User-Agent", "pjsekai/48 CFNetwork/1240.0.4 Darwin/20.6.0")

	resp, err := cdo(c, req)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Status: StatusOK}
	case 403:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
