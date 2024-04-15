package atlassian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CommodityPrice(t *testing.T) {
	rcp := NewRunningCommodityPrice()
	tests := []struct{
		timestamp int
		commodityPrice int
		maxCommodityPrice int
		exp error
	}{
		{
			timestamp: 4,
			commodityPrice: 27,
			maxCommodityPrice: 27,
			exp: nil,
		},
		{
			timestamp: 6,
			commodityPrice: 26,
			maxCommodityPrice: 27,
			exp: nil,
		},
		{
			timestamp: 9,
			commodityPrice: 25,
			maxCommodityPrice: 27,
			exp: nil,
		},
		{
			timestamp: 4,
			commodityPrice: 24,
			maxCommodityPrice: 26,
			exp: nil,
		},
		{
			timestamp: 6,
			commodityPrice: 21,
			maxCommodityPrice: 25,
			exp: nil,
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run("commodity price", func(t *testing.T) {
			e := rcp.upsertCommodityPrice(tt.timestamp, tt.commodityPrice)
			assert.Nil(t, e)
			max, _ := rcp.getMaxCommodityPrice()
			assert.Equal(t, tt.maxCommodityPrice, max)
		})
	}
}

func TestGetMaxCommodityPrice_SadPath_NoHashMapCreated(t *testing.T) {
	rcp := NewRunningCommodityPrice()
	_, e := rcp.getMaxCommodityPrice()
	assert.EqualError(t, e, "no commodity prices have been inserted")
}
