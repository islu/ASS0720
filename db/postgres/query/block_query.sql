
-- Create Uniswap USDC/ETH pair swap log
-- name: CreateUniswapUSDCETHPairSwapLog :one
INSERT INTO uniswap_usdc_eth_pair_swap_log (
    block_sender, block_number, block_time, tx_hash, amount0_in, amount0_out, amount1_in, amount1_out
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- List Uniswap USDC/ETH pair swap log by sender
-- name: ListUniswapUSDCETHPairSwapLogBySender :many
SELECT * FROM uniswap_usdc_eth_pair_swap_log
WHERE block_sender = $1
ORDER BY block_number;

-- List Uniswap USDC/ETH pair swap log between timestamp
-- name: ListUniswapUSDCETHPairSwapLogByTimestamp :many
SELECT * FROM uniswap_usdc_eth_pair_swap_log
WHERE block_time >= $1 AND block_time <= $2
ORDER BY block_number;