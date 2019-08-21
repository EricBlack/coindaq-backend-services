CREATE TABLE IF NOT EXISTS projects (
    id bigserial PRIMARY KEY,
  	admin_id bigserial,
  	summary text,
  	description text,
  	company text,
  	images jsonb,
  	target_value integer,
  	stage_count integer,
  	min_value integer,
  	classify text,
  	status integer,
	content_items jsonb,
  	create_time timestamp,
  	end_time timestamp,
  	update_time timestamp
);
CREATE INDEX IF NOT EXISTS index_projects_on_admin_id ON projects (admin_id);
COMMENT ON COLUMN projects.target_value IS '投资总额';
COMMENT ON COLUMN projects.stage_count IS '投资阶段';
COMMENT ON COLUMN projects.min_value IS '最小投资额';
COMMENT ON COLUMN projects.classify IS '投资类别';
COMMENT ON COLUMN projects.status IS '投资状态';


CREATE TABLE IF NOT EXISTS notice_news (
	id bigserial PRIMARY KEY,
	project_id bigserial,
	title text,
	description text,
	status integer,
	create_time timestamp,
	expire_time timestamp,
	update_time timestamp
);
CREATE INDEX IF NOT EXISTS index_notice_news_on_project_id ON notice_news (project_id);


CREATE TABLE IF NOT EXISTS members_info (
	id bigserial PRIMARY KEY,
	project_id bigserial,
	user_id bigserial,
	description jsonb,
	member_type integer,
	join_time timestamp,
	update_time timestamp
);
CREATE INDEX IF NOT EXISTS index_members_info_on_project_id ON members_info (project_id);


CREATE TABLE IF NOT EXISTS stage_info (
	id bigserial PRIMARY KEY,
	project_id bigserial,
	stage_number integer,
	target_value integer,
	target_type integer,
	min_value integer,
	max_value integer,
	complete_value integer,
	stage_status integer,
	begin_time timestamp,
	end_time timestamp,
	create_time timestamp,
	update_time timestamp
);
CREATE INDEX IF NOT EXISTS index_stage_info_on_project_id ON stage_info (project_id);
COMMENT ON COLUMN stage_info.stage_number IS '投资阶段';
COMMENT ON COLUMN stage_info.target_value IS '阶段目标投资';
COMMENT ON COLUMN stage_info.target_type IS '阶段目标投资类别';
COMMENT ON COLUMN stage_info.complete_value IS '阶段实现投资';
COMMENT ON COLUMN stage_info.stage_status IS '阶段投资状况';


CREATE TABLE IF NOT EXISTS investment_info (
	id bigserial PRIMARY KEY,
	user_id bigserial,
	project_id bigserial,
	stage_number integer,
	investment_value integer,
	investment_type integer,
	wallet_path text,
	investment_status integer,
	create_time timestamp,
	update_time timestamp
);
CREATE INDEX IF NOT EXISTS index_investment_info_on_user_id ON investment_info (user_id);
COMMENT ON COLUMN investment_info.stage_number IS '投资阶段';
COMMENT ON COLUMN investment_info.investment_value IS '用户投资金额';
COMMENT ON COLUMN investment_info.investment_type IS '用户投资类别';
COMMENT ON COLUMN investment_info.investment_status IS '用户投资状态';