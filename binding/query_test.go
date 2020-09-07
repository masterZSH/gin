// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	var qb queryBinding
	assert.Equal(t, qb.Name(), "query")
}

type TestQuery struct {
	Test string `form:"test"`
}

func TestBind(t *testing.T) {
	var qb queryBinding
	var obj TestQuery
	r, err := http.NewRequest("get", "/", nil)
	assert.Nil(t, err)
	q := r.URL.Query()
	q.Set("test", "foo")
	r.URL.RawQuery = q.Encode()
	err = qb.Bind(r, &obj)
	assert.Nil(t, err)
	assert.Equal(t, "foo", obj.Test)
}
