// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudsearchdomainiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain"
	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudsearchdomainiface.CloudSearchDomainAPI)(nil), cloudsearchdomain.New(nil))
}