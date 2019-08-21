CREATE TABLE IF NOT EXISTS coin_info (
    id bigserial PRIMARY KEY,
	code text,
  	symbol text,
  	name text,
  	decimals integer,
  	coin_type integer,
  	min_confirms integer,
  	withdrawal integer,
  	receivable integer,
  	path text,
  	deleted_at timestamp,
  	created_at timestamp,
	updated_at timestamp
);

CREATE TABLE IF NOT EXISTS light_wallet (
	id bigserial PRIMARY KEY,
	wallet_name text,
	user_id bigserial,
	coin_id bigserial,
	address text,
	created_at timestamp,
	updated_at timestamp
);

CREATE TABLE IF NOT EXISTS platform_wallet (
	id bigserial PRIMARY KEY,
	user_id bigserial,
	coin_id bigserial,
	wallet_name text,
	account text,
	address text,
	created_at timestamp,
	updated_at timestamp
);

CREATE TABLE IF NOT EXISTS tx_cache (
	id bigserial PRIMARY KEY,
	coin_id bigserial,
	coin_symbol text,
	block_hash text,
	account text,
	address text,
	tx_type text,
	block_height integer,
	block_index bigserial,
	fee float,
	amount float,
	confirmations integer,
	finished bool,
	created_at timestamp,
	updated_at timestamp,
	finished_at timestamp
);