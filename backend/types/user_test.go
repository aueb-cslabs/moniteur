package types

import "testing"

func TestUser_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"No Data Provided",
			fields{},
			true,
		},
		{
			"Empty Username",
			fields{
				"",
				"password123",
			},
			true,
		},
		{
			"Correct Data Provided",
			fields{
				"username123",
				"password123",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
