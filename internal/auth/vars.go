package auth

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/internal/internalTypes"
)

func ResolveVars(env, vars []internalTypes.KeyValue, auth *Auth) Auth {

	var result = Auth{Type: auth.Type, Details: map[string]string{}}

	for kd, vd := range auth.Details {
		result.Details[kd] = vd
	}

	allVars := append(vars, env...)

	for _, v := range allVars {

		vk := fmt.Sprintf("{{%s}}", v.Key)

		for kd, vd := range auth.Details {

			if _, ok := result.Details[kd]; ok {
				result.Details[kd] = strings.ReplaceAll(result.Details[kd], vk, v.Value)
				continue
			}

			result.Details[kd] = strings.ReplaceAll(vd, vk, v.Value)

		}

	}
	return result
}
