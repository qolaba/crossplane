/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"

	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
)

// GetBindingPhase of this MySQLInstance.
func (cm *MySQLInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetCondition of this MySQLInstance.
func (cm *MySQLInstance) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return cm.Status.GetCondition(ct)
}

// GetPortableClassReference of this MySQLInstance.
func (cm *MySQLInstance) GetPortableClassReference() *corev1.LocalObjectReference {
	return cm.Spec.PortableClassReference
}

// GetResourceReference of this MySQLInstance.
func (cm *MySQLInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this MySQLInstance.
func (cm *MySQLInstance) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this MySQLInstance.
func (cm *MySQLInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetConditions of this MySQLInstance.
func (cm *MySQLInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetPortableClassReference of this MySQLInstance.
func (cm *MySQLInstance) SetPortableClassReference(r *corev1.LocalObjectReference) {
	cm.Spec.PortableClassReference = r
}

// SetResourceReference of this MySQLInstance.
func (cm *MySQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this MySQLInstance.
func (cm *MySQLInstance) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetCondition of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return cm.Status.GetCondition(ct)
}

// GetPortableClassReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetPortableClassReference() *corev1.LocalObjectReference {
	return cm.Spec.PortableClassReference
}

// GetResourceReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetConditions of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetPortableClassReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetPortableClassReference(r *corev1.LocalObjectReference) {
	cm.Spec.PortableClassReference = r
}

// SetResourceReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}
