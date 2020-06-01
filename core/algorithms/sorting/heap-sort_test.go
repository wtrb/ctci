package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.args.arr))
			copy(input, tt.args.arr)

			HeapSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("HeapSort(%+v) = %v, want %v", input, tt.args.arr, tt.want)
			}
		})
	}
}
