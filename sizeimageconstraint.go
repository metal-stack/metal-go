package metalgo

import (
	"github.com/metal-stack/metal-go/api/client/sizeimageconstraint"
	"github.com/metal-stack/metal-go/api/models"
)

// TrySizeImageConstraint try if size and image can be used for a allocation
func (d *Driver) TrySizeImageConstraint(size, image string) error {
	params := sizeimageconstraint.NewTrySizeImageConstraintParams().WithBody(&models.V1SizeImageConstraintTryRequest{
		Size:  &size,
		Image: &image,
	})

	_, err := d.Sizeimageconstraint().TrySizeImageConstraint(params, nil)
	if err != nil {
		return err
	}
	return nil
}
