package domain

import "testing"

func TestAccount_IsDocumentNumberValid(t *testing.T) {
	type fields struct {
		ID             int64
		DocumentNumber string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid cpf",
			fields: fields{
				DocumentNumber: "80502674075",
			},
			want: true,
		},
		{
			name: "valid cnpj",
			fields: fields{
				DocumentNumber: "19931301000167",
			},
			want: true,
		},
		{
			name: "invalid document number",
			fields: fields{
				DocumentNumber: "123456789012345",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				ID:             tt.fields.ID,
				DocumentNumber: tt.fields.DocumentNumber,
			}
			if got := a.IsDocumentNumberValid(); got != tt.want {
				t.Errorf("Account.IsDocumentNumberValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
