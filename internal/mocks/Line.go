// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	internal "github.com/farishadibrata/maroto/internal"
	mock "github.com/stretchr/testify/mock"

	props "github.com/farishadibrata/maroto/pkg/props"
)

// Line is an autogenerated mock type for the Line type
type Line struct {
	mock.Mock
}

// Draw provides a mock function with given fields: cell, lineProp
func (_m *Line) Draw(cell internal.Cell, lineProp props.Line) {
	_m.Called(cell, lineProp)
}
