//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Route struct {
	Model
	ControllerID 	uint		`sql:"index" json:"controllerID"`
	Name			string		`json:"name"`
	Description		string		`sql:"size:1000" json:"description"`
	Route 			string		`json:"route"`
	Content 		string		`json:"content" sql:"size:1000"`
	Responses		[]Response	`json:"responses,omitempty"`
	Variables		[]Variable	`sql:"gorm:many2many:route_variables;" json:"variables,omitempty"`
}
