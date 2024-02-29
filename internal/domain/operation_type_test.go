package domain

import "testing"

// TestOperationType_IsAmountValid tests the IsAmountValid method from the OperationType struct
func TestOperationType_IsAmountValid(t *testing.T) {
	type fields struct {
		ID          int
		Description string
		IsDebit     bool
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{

		{
			"IsAmountValid - IsDebit true and amount is 1",
			fields{
				ID:          1,
				Description: "COMPRA A VISTA",
				IsDebit:     true,
			},
			args{
				amount: 1.0,
			},
			false,
		},
		{
			"IsAmountValid - IsDebit true and amount is 0",
			fields{
				ID:          2,
				Description: "COMPRA PARCELADA",
				IsDebit:     true,
			},
			args{
				amount: 0.0,
			},
			false,
		},
		{
			"IsAmountValid - IsDebit true and amount is -1",
			fields{
				ID:          3,
				Description: "SAQUE",
				IsDebit:     true,
			},
			args{
				amount: -1.0,
			},
			true,
		},
		{
			"IsAmountValid - IsDebit false and amount is 0",
			fields{
				ID:          4,
				Description: "PAGAMENTO",
				IsDebit:     false,
			},
			args{
				amount: 0.0,
			},
			false,
		},
		{
			"IsAmountValid - IsDebit false and amount is 1",
			fields{
				ID:          4,
				Description: "PAGAMENTO",
				IsDebit:     false,
			},
			args{
				amount: 1.0,
			},
			true,
		},
		{
			"IsAmountValid - IsDebit false and amount is -1",
			fields{
				ID:          4,
				Description: "PAGAMENTO",
				IsDebit:     false,
			},
			args{
				amount: -1.0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperationType{
				ID:          tt.fields.ID,
				Description: tt.fields.Description,
				IsDebit:     tt.fields.IsDebit,
			}
			if got := o.IsAmountValid(tt.args.amount); got != tt.want {
				t.Errorf("OperationType.IsAmountValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
