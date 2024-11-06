package auth

import (
	"fmt"
	"strings"

	"github.com/kmarkela/duffman/internal/internalTypes"
)

func ResolveVars(env, vars []internalTypes.KeyValue, auth *Auth) error {

	allVars := append(vars, env...)

	for _, v := range allVars {

		vk := fmt.Sprintf("{{%s}}", v.Key)

		for kd, vd := range auth.Details {
			auth.Details[kd] = strings.ReplaceAll(vd, vk, v.Value)
		}
	}

	return nil
}
