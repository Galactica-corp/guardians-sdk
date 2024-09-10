// Copyright © 2024 Galactica Network
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package zkcertificate

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func Test_dollarsToCentsTruncated(t *testing.T) {
	type args struct {
		dollars string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "simple",
			args: args{dollars: "1.50"},
			want: big.NewInt(150),
		},
		{
			name: "no decimals",
			args: args{dollars: "1"},
			want: big.NewInt(100),
		},
		{
			name: "too many digits",
			args: args{dollars: "1.239"},
			want: big.NewInt(123),
		},
		{
			name: "never round up",
			args: args{dollars: "1.999"},
			want: big.NewInt(199),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dollars, err := decimal.NewFromString(tt.args.dollars)
			require.NoError(t, err)

			if got := dollarsToCentsTruncated(dollars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dollarsToCentsTruncated() = %v, want %v", got, tt.want)
			}
		})
	}
}
