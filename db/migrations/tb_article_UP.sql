CREATE TABLE IF NOT EXISTS public.tb_articlesss (
  article_id varchar(36) NOT NULL,
  article_title varchar(100) DEFAULT NULL,
  article_body text,
  category varchar(100) DEFAULT NULL,
  is_active int2 NULL,
  thumbnail text,
  created_by varchar(100) DEFAULT NULL,
  created_at timestamp without time zone,
  updated_by varchar(100) DEFAULT NULL,
  updated_at timestamp without time zone,
  deleted_by varchar(100) DEFAULT NULL,
  deleted_at timestamp without time zone,
  PRIMARY KEY (article_id)
);