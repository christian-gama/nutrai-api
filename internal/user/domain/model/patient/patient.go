package patient

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Patient is the model for a patient.
type Patient struct {
	ID       sharedvalue.ID `faker:"uint"`
	User     *user.User     `faker:"-"`
	WeightKG value.WeightKG `faker:"boundary_start=1, boundary_end=999"`
	HeightM  value.HeightM  `faker:"boundary_start=1, boundary_end=3"`
	Age      value.Age      `faker:"boundary_start=1, boundary_end=100"`
}

// New returns a new Patient instance.
func New(patient *Patient) (*Patient, error) {
	if err := patient.Validate(); err != nil {
		return nil, err
	}

	return patient, nil
}

func (p *Patient) Validate() error {
	var errs *errutil.Error

	if err := p.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.WeightKG.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.HeightM.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.Age.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
