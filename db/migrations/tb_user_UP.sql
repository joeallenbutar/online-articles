CREATE TABLE IF NOT EXISTS public.tb_userrr (
  username varchar(100) NOT NULL,
  password varchar(200) DEFAULT NULL,
  is_active int2 NULL,
  PRIMARY KEY (username)
)