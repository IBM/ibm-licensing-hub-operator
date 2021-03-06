//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	operatorv1alpha1 "github.com/ibm/ibm-licensing-hub-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const APISecretTokenKeyName = "token"

func GetAPISecretToken(instance *operatorv1alpha1.IBMLicensingHub) *corev1.Secret {
	metaLabels := LabelsForLicensingHubMeta(instance)
	expectedSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Spec.APISecretToken,
			Namespace: instance.GetNamespace(),
			Labels:    metaLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{APISecretTokenKeyName: RandString(24)},
	}
	return expectedSecret
}

func GetDatabaseSecret(instance *operatorv1alpha1.IBMLicensingHub) *corev1.Secret {
	metaLabels := LabelsForLicensingHubMeta(instance)
	expectedSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DatabaseConfigSecretName,
			Namespace: instance.GetNamespace(),
			Labels:    metaLabels,
		},
		Type: corev1.SecretTypeOpaque,
		StringData: map[string]string{
			PostgresPasswordKey:     RandString(8),
			PostgresUserKey:         DatabaseUser,
			PostgresDatabaseNameKey: DatabaseName,
			PostgresPgDataKey:       PgData,
		},
	}
	return expectedSecret
}
