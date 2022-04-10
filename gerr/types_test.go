package gerr

import "testing"

func TestTypes(t *testing.T) {
	var (
		input = []error{
			BadRequest("domain", "reason_400"),
			Unauthorized("domain", "reason_401"),
			Forbidden("domain", "reason_403"),
			NotFound("domain", "reason_404"),
			Conflict("domain", "reason_409"),
			InternalServer("domain", "reason_500"),
			ServiceUnavailable("domain", "reason_503"),
		}
		output = []func(error) bool{
			IsBadRequest,
			IsUnauthorized,
			IsForbidden,
			IsNotFound,
			IsConflict,
			IsInternalServer,
			IsServiceUnavailable,
		}
	)

	for i, in := range input {
		if !output[i](in) {
			t.Errorf("not expect: %v", in)
		}
	}
}
