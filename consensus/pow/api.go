/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package pow

type API struct {
	engine *Engine
}

// GetHashrate returns the current hashrate for local CPU miner and remote miner.
func (api *API) GetHashrate() uint64 {
	return uint64(api.engine.hashrate.Rate1())
}