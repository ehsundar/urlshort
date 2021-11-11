package composer

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
)

type md5Base64 struct {
}

func NewMd5Base64() Composer {
	return md5Base64{}
}

func (c md5Base64) Compose(long, nonce string) (short string) {
	hash := md5.Sum([]byte(long + nonce))
	buff := bytes.NewBuffer([]byte{})
	enc := base64.NewEncoder(base64.URLEncoding, buff)
	_, _ = enc.Write(hash[:])

	return buff.String()
}
