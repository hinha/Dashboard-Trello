package security

import "encoding/json"

func Authorize(cipher *BearerCipher, key string) (interface{}, error) {
	plain, err := cipher.DecryptStringCBC(key)
	if err != nil {
		return nil, err
	}

	var decode map[string]interface{}
	if err := json.Unmarshal([]byte(plain), &decode); err != nil {
		return nil, err
	}

	// TODO: Need validation time expiration need refactor

	return decode, nil
}
