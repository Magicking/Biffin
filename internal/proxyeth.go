package internal

import (
)

func InitContext(raw_url, private_key string) (*Context, error) {
	ctx := &Context{RawURL: raw_url, PrivateKey: private_key}
	return ctx, nil
}
