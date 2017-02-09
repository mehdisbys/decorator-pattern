package client

import (
	"net/http"
)

func Auth(token string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {

			if ( token == "12345") {
				return c.Do(r)
			}

			panic("Could not authenticate user")
		})
	}
}
