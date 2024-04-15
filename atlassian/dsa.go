// Imagine you are given a stream of data points consisting of <timestamp, commodityPrice> you are supposed to return the maxCommodityPrice at any point in time.
// The timestamps in the stream can be out of order, or there can be duplicate timestamps, we need to update the commodityPrice at that particular timestamp if an entry for the timestamp already exists.
// Create an in-memory solution tailored to prioritize frequent reads and writes for the given problem statement.

// interface RunningCommodityPrice {

// 	void upsertCommodityPrice(int timestamp, int commodityPrice);

// 	int getMaxCommodityPrice();

// }

// RunningCommodityPrice r = new RunningCommodityPrice();

// r.upsertCommodityPrice(4, 27);

// r.upsertCommodityPrice(6, 26);

// r.upsertCommodityPrice(9, 25);

// r.getMaxCommodityPrice();         // output should be 27 which is at timestamp 4

// r.upsertCommodityPrice(4, 28);    // timestamps can come out of order and there can be duplicates

// the commodity price at timestamp 4 got updated to 28, so the max commodity price is 28

// r.getMaxCommodityPrice();         // output should be 28 from timestamp 4
// RunningCommodityPrice r = new RunningCommodityPrice();
// r.upsertCommodityPrice(4, 27);
// r.upsertCommodityPrice(6, 26);
// r.upsertCommodityPrice(9, 25);
// r.getMaxCommodityPrice();         // output should be 27 which is at timestamp 4
// r.upsertCommodityPrice(4, 24);    // timestamps can come out of order and there can be duplicates
// // the commodity price at timestamp 4 got updated to 24 so it is no longer the max commodity price
// r.getMaxCommodityPrice();         // output should be 26 from timestamp 6
// r.upsertCommodityPrice(6, 21);
// r.getMaxCommodityPrice();         // output should be 25 from timestamp 9
package atlassian

import "errors"

type RunningCommodityPrice interface {
	upsertCommodityPrice(timestamp, commodityPrice int) error
	getMaxCommodityPrice() (int, error)
}

type ImplRunningCommodityPrice struct {
	maxCommodityPrice int
	maxCommodityPriceTimestamp int
	commodityPricesMap map[int]int // key: timestamp,
	// commodityPricesTimestamps map[int]int // key: price
	commodityPrices []int // 28, 26,...
	commodityPricesTimestamps map[int]int // 4:0, 6,... 
}

func NewRunningCommodityPrice() RunningCommodityPrice {
	return &ImplRunningCommodityPrice{
		maxCommodityPrice: 0,
		commodityPricesMap: map[int]int{},
	}
}

func (p *ImplRunningCommodityPrice) upsertCommodityPrice(timestamp, commodityPrice int) error {
	if len(p.commodityPrices) == 0 {
		p.commodityPrices = append(p.commodityPrices, commodityPrice)
		p.commodityPricesTimestamps = append(p.commodityPricesTimestamps, timestamp)
	} else {
		/*
		1. insert if there's no array
		2. 
		*/
		for i := 0; i < len(p.commodityPricesTimestamps) - 1; i++ {
			if p.commodityPrices[i] < commodityPrice && commodityPrice < p.commodityPrices[i+1] {

			}
		}
	}
	return nil
}

func (p *ImplRunningCommodityPrice) upsertCommodityPrice_old(timestamp, commodityPrice int) error {
	p.commodityPricesMap[timestamp] = commodityPrice
	if p.maxCommodityPrice < commodityPrice && p.maxCommodityPriceTimestamp != timestamp {
		p.maxCommodityPrice = commodityPrice
		p.maxCommodityPriceTimestamp = timestamp
	} else if p.maxCommodityPriceTimestamp == timestamp && p.maxCommodityPrice > commodityPrice {
		p.maxCommodityPrice = commodityPrice
		for ts, cp := range p.commodityPricesMap {
			if p.maxCommodityPrice < cp {
				p.maxCommodityPrice = cp
				p.maxCommodityPriceTimestamp = ts
			}
		}
	}
	return nil
}

func (p *ImplRunningCommodityPrice) getMaxCommodityPrice() (int, error) {
	if (len(p.commodityPricesMap) == 0) {
		return 0, errors.New("no commodity prices have been inserted")
	}
	return p.maxCommodityPrice, nil
}
