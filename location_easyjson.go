// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package valhalla

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson14b80819DecodeGithubComLittlemonkeyltdValhallaGoBindings(in *jlexer.Lexer, out *Location) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lat":
			out.Lat = float64(in.Float64())
		case "lon":
			out.Lon = float64(in.Float64())
		case "street":
			out.Street = string(in.String())
		case "Type":
			out.Type = string(in.String())
		case "Heading":
			out.Heading = int(in.Int())
		case "HeadingTolerance":
			out.HeadingTolerance = int(in.Int())
		case "MinimumReachability":
			out.MinimumReachability = int(in.Int())
		case "Radius":
			out.Radius = int(in.Int())
		case "RankCandidates":
			out.RankCandidates = bool(in.Bool())
		case "Name":
			out.Name = string(in.String())
		case "City":
			out.City = string(in.String())
		case "State":
			out.State = string(in.String())
		case "PostalCode":
			out.PostalCode = string(in.String())
		case "Country":
			out.Country = string(in.String())
		case "Phone":
			out.Phone = string(in.String())
		case "URL":
			out.URL = string(in.String())
		case "SideOfStreet":
			out.SideOfStreet = string(in.String())
		case "DateTime":
			out.DateTime = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson14b80819EncodeGithubComLittlemonkeyltdValhallaGoBindings(out *jwriter.Writer, in Location) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Lat != 0 {
		const prefix string = ",\"lat\":"
		first = false
		out.RawString(prefix[1:])
		out.Float64(float64(in.Lat))
	}
	if in.Lon != 0 {
		const prefix string = ",\"lon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Lon))
	}
	if in.Street != "" {
		const prefix string = ",\"street\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Heading\":"
		out.RawString(prefix)
		out.Int(int(in.Heading))
	}
	{
		const prefix string = ",\"HeadingTolerance\":"
		out.RawString(prefix)
		out.Int(int(in.HeadingTolerance))
	}
	{
		const prefix string = ",\"MinimumReachability\":"
		out.RawString(prefix)
		out.Int(int(in.MinimumReachability))
	}
	{
		const prefix string = ",\"Radius\":"
		out.RawString(prefix)
		out.Int(int(in.Radius))
	}
	{
		const prefix string = ",\"RankCandidates\":"
		out.RawString(prefix)
		out.Bool(bool(in.RankCandidates))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"State\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"PostalCode\":"
		out.RawString(prefix)
		out.String(string(in.PostalCode))
	}
	{
		const prefix string = ",\"Country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"Phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"URL\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"SideOfStreet\":"
		out.RawString(prefix)
		out.String(string(in.SideOfStreet))
	}
	{
		const prefix string = ",\"DateTime\":"
		out.RawString(prefix)
		out.String(string(in.DateTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Location) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson14b80819EncodeGithubComLittlemonkeyltdValhallaGoBindings(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Location) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson14b80819EncodeGithubComLittlemonkeyltdValhallaGoBindings(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Location) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson14b80819DecodeGithubComLittlemonkeyltdValhallaGoBindings(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Location) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson14b80819DecodeGithubComLittlemonkeyltdValhallaGoBindings(l, v)
}
