package models

import "encoding/xml"

type Response struct {
	XMLName  xml.Name  `xml:"body"`
	Vehicles []Vehicle `xml:"vehicle"`
}

type Vehicle struct {
	ID                 int    `gorm:"primary_key"`
	RouteTag           string `xml:"routeTag,attr"`
	DirectionTag       string `xml:"dirTag,attr"`
	Latitude           string `xml:"lat,attr"`
	Longitude          string `xml:"lon,attr"`
	SecondsSinceReport string `xml:"secsSinceReport,attr"`
	Predictable        string `xml:"predictable,attr"`
	Heading            string `xml:"heading,attr"`
	SpeedKmHr          string `xml:"speedKmHr,attr"`
}
