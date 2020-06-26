package types

import "testing"

func TestAnnouncement_Validate(t *testing.T) {
	type fields struct {
		End int64
		Msg string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"No Data Provided",
			fields{

			},
			true,
		},
		{
			"Empty Message String",
			fields{
				123,
				"",
			},
			true,
		},
		{
			"Correct Data Provided",
			fields{
				123,
				"Test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Announcement{
				End: tt.fields.End,
				Msg: tt.fields.Msg,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
