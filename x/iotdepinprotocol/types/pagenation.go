package types

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	db "github.com/tendermint/tm-db"
)

var IteratorNext = errors.New("iterator Next")

// Paginate does pagination of all the results in the PrefixStore based on the
// provided PageRequest. onResult should be used to do actual unmarshaling.
func Paginate(
	prefixStore types.KVStore,
	pageRequest *query.PageRequest,
	onResult func(key []byte, value []byte) (bool, error),
) (*query.PageResponse, error) {
	// if the PageRequest is nil, use default PageRequest
	if pageRequest == nil {
		pageRequest = &query.PageRequest{}
	}

	offset := pageRequest.Offset
	key := pageRequest.Key
	limit := pageRequest.Limit
	countTotal := pageRequest.CountTotal
	reverse := pageRequest.Reverse

	if offset > 0 && key != nil {
		return nil, fmt.Errorf("invalid request, either offset or key is expected, got both")
	}

	if limit == 0 {
		limit = query.DefaultLimit

		// count total results when the limit is zero/not supplied
		countTotal = true
	}

	if len(key) != 0 {
		iterator := getIterator(prefixStore, key, reverse)
		defer iterator.Close()

		var count uint64
		var nextKey []byte

		for ; iterator.Valid(); iterator.Next() {

			if count == limit {
				nextKey = iterator.Key()
				break
			}
			if iterator.Error() != nil {
				return nil, iterator.Error()
			}
			ok, err := onResult(iterator.Key(), iterator.Value())
			if err != nil {
				return nil, err
			}
			// false && err == nil 跳过这条数据
			if !ok {
				continue
			}

			count++
		}

		return &query.PageResponse{
			NextKey: nextKey,
		}, nil
	}

	iterator := getIterator(prefixStore, nil, reverse)
	defer iterator.Close()

	end := offset + limit

	var count uint64
	var nextKey []byte

	for ; iterator.Valid(); iterator.Next() {
		count++

		if count <= offset {
			continue
		}
		if count <= end {
			ok, err := onResult(iterator.Key(), iterator.Value())
			if err != nil {
				return nil, err
			}
			// false && err == nil 跳过这条数据
			if !ok {
				count--
				continue
			}
		} else if count == end+1 {
			nextKey = iterator.Key()

			if !countTotal {
				break
			}
		}
		if iterator.Error() != nil {
			return nil, iterator.Error()
		}
	}

	res := &query.PageResponse{NextKey: nextKey}
	if countTotal {
		res.Total = count
	}

	return res, nil
}

func getIterator(prefixStore types.KVStore, start []byte, reverse bool) db.Iterator {
	if reverse {
		var end []byte
		if start != nil {
			itr := prefixStore.Iterator(start, nil)
			defer itr.Close()
			if itr.Valid() {
				itr.Next()
				end = itr.Key()
			}
		}
		return prefixStore.ReverseIterator(nil, end)
	}
	return prefixStore.Iterator(start, nil)
}
