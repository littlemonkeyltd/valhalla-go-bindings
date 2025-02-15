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

func easyjson7152e075DecodeGithubComLittlemonkeyltdValhallaGoBindings(in *jlexer.Lexer, out *DirectionsOptions) {
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
		case "units":
			out.Units = string(in.String())
		case "narrative":
			out.Narrative = bool(in.Bool())
		case "language":
			out.Language = string(in.String())
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
func easyjson7152e075EncodeGithubComLittlemonkeyltdValhallaGoBindings(out *jwriter.Writer, in DirectionsOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Units != "" {
		const prefix string = ",\"units\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Units))
	}
	if in.Narrative {
		const prefix string = ",\"narrative\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Narrative))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DirectionsOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7152e075EncodeGithubComLittlemonkeyltdValhallaGoBindings(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DirectionsOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7152e075EncodeGithubComLittlemonkeyltdValhallaGoBindings(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DirectionsOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7152e075DecodeGithubComLittlemonkeyltdValhallaGoBindings(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DirectionsOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7152e075DecodeGithubComLittlemonkeyltdValhallaGoBindings(l, v)
}
