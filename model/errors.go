package model

import (
	"fmt"
)

type ServiceNotAvailableError interface {
	GetServiceEndpoint() string
	Error() string
}

type DefaultServiceNotAvailableError struct {
	ServiceEndpoint string
	ErrorMsg        string
}

func (e *DefaultServiceNotAvailableError) GetServiceEndpoint() string {
	return e.ServiceEndpoint
}

func (e *DefaultServiceNotAvailableError) Error() string {
	return fmt.Sprintf("Results not available from %s. %s", e.ServiceEndpoint, e.ErrorMsg)
}

func LbErrorItems(e ServiceNotAvailableError) LBItems {
	var lbItems LBItems
	errEndpointConnectItem := LBErrorItem(e)
	lbItems.Items = append(lbItems.Items, *errEndpointConnectItem)

	return lbItems
}

func LbErrorItem(e ServiceNotAvailableError) LBItems {
	var lbItems LBItems
	errEndpointConnectItem := LBErrorItem(e)
	lbItems.Items = append(lbItems.Items, *errEndpointConnectItem)

	return lbItems
}
