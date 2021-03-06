package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*DoctorInfo doctor info

swagger:model doctorInfo
*/
type DoctorInfo struct {

	/* namespaces

	Required: true
	*/
	Namespaces []*Namespace `json:"namespaces"`

	/* nodes

	Required: true
	*/
	Nodes []*K8sResource `json:"nodes"`

	/* workflow

	Required: true
	*/
	Workflow *Cluster `json:"workflow"`
}

// Validate validates this doctor info
func (m *DoctorInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkflow(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DoctorInfo) validateNamespaces(formats strfmt.Registry) error {

	if err := validate.Required("namespaces", "body", m.Namespaces); err != nil {
		return err
	}

	for i := 0; i < len(m.Namespaces); i++ {

		if m.Namespaces[i] != nil {

			if err := m.Namespaces[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DoctorInfo) validateNodes(formats strfmt.Registry) error {

	if err := validate.Required("nodes", "body", m.Nodes); err != nil {
		return err
	}

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {

			if err := m.Nodes[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DoctorInfo) validateWorkflow(formats strfmt.Registry) error {

	if m.Workflow != nil {

		if err := m.Workflow.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
