
CREATE TABLE
    uniswap_usdc_eth_pair_swap_log (
        block_sender varchar(255) NOT NULL, -- wallet address
        block_number bigint NOT NULL,
        block_time timestamptz NOT NULL,
        tx_hash varchar(255) NOT NULL,
        amount0_in bigint NOT NULL,
        amount0_out bigint NOT NULL,
        amount1_in bigint NOT NULL,
        amount1_out bigint NOT NULL
    );