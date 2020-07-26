package types

import "testing"

func TestHoliday_Validate(t *testing.T) {
	type fields struct {
		Date    string
		Name    string
		IntName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"No Data Provided",
			fields{ },
			true,
		},
		{
			"Empty Strings",
			fields{
				"",
				"",
				"",
			},
			true,
		},
		{
			"Wrong Date Type Type 1",
			fields{
				"This is gonna break",
				"Test",
				"Test2",
			},
			true,
		},
		{
			"Wrong Date Type Type 2",
			fields{
				"20/03/2019",
				"Test",
				"Test2",
			},
			true,
		},
		{
			"Correct Data Provided",
			fields{
				"2006-01-02",
				"Test",
				"Test2",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Holiday{
				Date:    tt.fields.Date,
				Name:    tt.fields.Name,
				IntName: tt.fields.IntName,
			}
			if err := h.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
