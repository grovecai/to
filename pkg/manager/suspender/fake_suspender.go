// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package suspender

import (
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
)

type FakeSuspender struct {
	SuspendComponentFunc func(v1alpha1.Cluster, v1alpha1.MemberType) (bool, error)
}

func NewFakeSuspender() *FakeSuspender {
	return &FakeSuspender{}
}

func (s *FakeSuspender) SuspendComponent(cluster v1alpha1.Cluster, comp v1alpha1.MemberType) (bool, error) {
	if s.SuspendComponentFunc == nil {
		return false, nil
	}

	return s.SuspendComponentFunc(cluster, comp)
}
