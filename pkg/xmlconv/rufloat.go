package xmlconv

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type RuFloat float64

func (rf *RuFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	if s == "" {
		*rf = RuFloat(0)
		return nil
	}
	s = strings.Replace(s, ",", ".", 1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*rf = RuFloat(f)
	return nil
}
