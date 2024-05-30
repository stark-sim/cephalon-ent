-- create "income_manages" table
CREATE TABLE "income_manages" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "phone" character varying NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT 'unknown', "amount" bigint NOT NULL DEFAULT 0, "reason" character varying NOT NULL DEFAULT '', "current_balance" bigint NOT NULL DEFAULT 0, "last_edited_at" timestamptz NOT NULL, "reject_reason" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'pending', "user_id" bigint NOT NULL DEFAULT 0, "approve_user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "incomemanage_user_id" to table: "income_manages"
CREATE INDEX "incomemanage_user_id" ON "income_manages" ("user_id");
-- create index "incomemanage_approve_user_id" to table: "income_manages"
CREATE INDEX "incomemanage_approve_user_id" ON "income_manages" ("approve_user_id");
-- set comment to table: "income_manages"
COMMENT ON TABLE "income_manages" IS '收益补发记录，需要系统补发的收益，记录到这个表里';
-- set comment to column: "id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "phone" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."phone" IS '用户手机号';
-- set comment to column: "type" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."type" IS '类型';
-- set comment to column: "amount" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."amount" IS '金额，单位：厘';
-- set comment to column: "reason" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."reason" IS '操作该用户收益钱包的原因';
-- set comment to column: "current_balance" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."current_balance" IS '当前余额（在生成这条记录时刻的余额），单位：厘';
-- set comment to column: "last_edited_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."last_edited_at" IS '审批前最后一次编辑的时间';
-- set comment to column: "reject_reason" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."reject_reason" IS '拒绝此条记录原因';
-- set comment to column: "status" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."status" IS '状态';
-- set comment to column: "user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."user_id" IS '用戶 id';
-- set comment to column: "approve_user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."approve_user_id" IS '审批人 id';
