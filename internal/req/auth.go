package req

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func ResolveAuthVars(env, vars []pcollection.KeyValue, auth *pcollection.Auth) error {

	allVars := append(vars, env...)

	at, det, err := auth.Get()
	if err != nil {
		return err
	}

	for _, v := range allVars {

		vk := fmt.Sprintf("{{%s}}", v.Key)

		for i, kv := range det {
			str, ok := kv.Value.(string)
			if !ok {
				continue
			}

			det[i].Value = strings.ReplaceAll(str, vk, v.Value)
		}
	}

	if at == "Oath2" {
		auth.Oauth2 = det
	}

	return nil
}
