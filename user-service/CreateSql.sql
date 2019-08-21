-------------------------- users ---------------------------
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email text,
    password text,
    salt text,
    display_name text,
    phone_number text,
    kind integer,
    activated integer,
    country_code text,
    identity_id text,
    real_name text,
    disabled integer,
    kyc_photos jsonb,
    kyc_vedios jsonb,
    kyc_stage integer,
    invite_code text,
    register_ip text,
    device_id text,
    created_at timestamp,
    updated_at timestamp,
    activated_at timestamp,
	disabled_at timestamp
);
CREATE INDEX IF NOT EXISTS uni_index_users_on_email ON users (email) WHERE email IS NOT NULL;
CREATE INDEX IF NOT EXISTS uni_index_users_on_phone_number ON users (phone_number) WHERE phone_number IS NOT NULL;
CREATE INDEX IF NOT EXISTS uni_index_users_on_identity_id ON users (identity_id) WHERE identity_id IS NOT NULL;

-------------------------- login_records ----------------------------
CREATE TABLE IF NOT EXISTS login_records (
    id bigserial PRIMARY KEY,
    user_id bigserial,
    login_ip text,
    device_id text,
    login_status integer,
	login_comment text,
    login_time timestamp
);
CREATE INDEX IF NOT EXISTS index_login_records_on_user_id ON login_records (user_id);
COMMENT ON COLUMN login_records.login_status IS '登陆状态';

-------------------------- kyc_requests ---------------------------
CREATE TABLE IF NOT EXISTS kyc_requests (
  id bigserial PRIMARY KEY,
  user_id bigserial,
  stage integer,
  state integer,
  kind integer,
  resource jsonb,
  created_at timestamp,
  updated_at timestamp,
  rejected_at timestamp,
  passed_at timestamp
);
CREATE INDEX IF NOT EXISTS index_kyc_requests_on_user_id ON kyc_requests (user_id);

-------------------------- two_factors ---------------------------
CREATE TABLE IF NOT EXISTS two_factors (
  id bigserial PRIMARY KEY,
  user_id bigserial,
  otp_secret text,
  activated integer,
  verify_type integer,
  last_verify_at timestamp,
  refreshed_at timestamp
);
CREATE INDEX IF NOT EXISTS uni_index_two_factors_on_user_id ON two_factors (user_id, verify_type);
COMMENT ON COLUMN two_factors.otp_secret IS 'password created by server side';
COMMENT ON COLUMN two_factors.activated IS 'used or not';
COMMENT ON COLUMN two_factors.verify_type IS 'verify type';

-------------------------- invite_info ---------------------------
CREATE TABLE IF NOT EXISTS inviter_info (
  id bigserial PRIMARY KEY,
  user_id bigserial,
  inviter_id bigserial,
  created_at timestamp
);
CREATE INDEX IF NOT EXISTS uni_index_inviter_info_on_user_id ON inviter_info (user_id);
COMMENT ON COLUMN inviter_info.user_id IS '邀请人Id';
COMMENT ON COLUMN inviter_info.inviter_id IS '被邀请人Id';
