package secret

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func Create(pass []byte) ([]byte, error) {
	var placeholder = "REPLACE_ME"

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      placeholder,
			Namespace: placeholder,
		},
		Type: "Opaque",
		Data: map[string][]byte{
			"username": []byte(placeholder),
			"password": pass,
		},
	}

	return yaml.Marshal(secret)
}
