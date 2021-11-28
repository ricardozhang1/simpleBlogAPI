CREATE TABLE "blog_tag" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "created_id" int NOT NULL,
  "modified_on" timestamptz DEFAULT (now()),
  "modified_id" int NOT NULL,
  "deleted_on" timestamptz DEFAULT ('0001-01-01 00:00:00Z'),
  "state" smallint DEFAULT 1
);

CREATE TABLE "blog_article" (
  "id" SERIAL PRIMARY KEY,
  "tag_id" bigserial NOT NULL,
  "title" varchar(100) NOT NULL,
  "desc" varchar(255) NOT NULL,
  "content" text,
  "created_on" timestamptz DEFAULT (now()),
  "created_id" int NOT NULL,
  "modified_on" timestamptz DEFAULT (now()),
  "modified_id" int NOT NULL,
  "deleted_on" timestamptz DEFAULT ('0001-01-01 00:00:00Z'),
  "state" smallint DEFAULT 1
);

CREATE TABLE "blog_user" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar(50) UNIQUE NOT NULL,
  "username" varchar(50) NOT NULL,
  "name" varchar(100) NOT NULL,
  "password" varchar(50) NOT NULL,
  "created_on" timestamptz DEFAULT (now()),
  "state" smallint DEFAULT 1
);

ALTER TABLE "blog_tag" ADD FOREIGN KEY ("created_id") REFERENCES "blog_user" ("id");

ALTER TABLE "blog_tag" ADD FOREIGN KEY ("modified_id") REFERENCES "blog_user" ("id");

ALTER TABLE "blog_article" ADD FOREIGN KEY ("tag_id") REFERENCES "blog_tag" ("id");

ALTER TABLE "blog_article" ADD FOREIGN KEY ("created_id") REFERENCES "blog_user" ("id");

ALTER TABLE "blog_article" ADD FOREIGN KEY ("modified_id") REFERENCES "blog_user" ("id");

COMMENT ON COLUMN "blog_tag"."name" IS '标签名称';

COMMENT ON COLUMN "blog_tag"."created_id" IS '创建人ID';

COMMENT ON COLUMN "blog_tag"."modified_on" IS '修改时间';

COMMENT ON COLUMN "blog_tag"."modified_id" IS '修改人ID';

COMMENT ON COLUMN "blog_tag"."deleted_on" IS '删除时间';

COMMENT ON COLUMN "blog_tag"."state" IS '状态 0为禁用、1为启用';

COMMENT ON COLUMN "blog_article"."tag_id" IS '标签ID';

COMMENT ON COLUMN "blog_article"."title" IS '文章标题';

COMMENT ON COLUMN "blog_article"."desc" IS '简述';

COMMENT ON COLUMN "blog_article"."created_on" IS '创建时间';

COMMENT ON COLUMN "blog_article"."created_id" IS '创建人ID';

COMMENT ON COLUMN "blog_article"."modified_on" IS '修改时间';

COMMENT ON COLUMN "blog_article"."modified_id" IS '修改人ID';

COMMENT ON COLUMN "blog_article"."deleted_on" IS '删除时间';

COMMENT ON COLUMN "blog_article"."state" IS '状态 0为禁用、1为启用';

COMMENT ON COLUMN "blog_user"."username" IS '账号';

COMMENT ON COLUMN "blog_user"."name" IS '用户名';

COMMENT ON COLUMN "blog_user"."password" IS '密码';

COMMENT ON COLUMN "blog_user"."created_on" IS '创建时间';

COMMENT ON COLUMN "blog_user"."state" IS '状态 0为禁用、1为启用';
