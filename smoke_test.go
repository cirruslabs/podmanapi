package podmanapi

import (
	"github.com/cirruslabs/podmanapi/pkg/swagger"
	"testing"
)

func TestSmoke(t *testing.T) {
	_ = swagger.NewAPIClient(&swagger.Configuration{})
}
