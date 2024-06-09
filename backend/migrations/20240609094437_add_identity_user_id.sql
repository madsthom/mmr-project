-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "identity_user_id" text NULL, ADD CONSTRAINT "uni_users_identity_user_id" UNIQUE ("identity_user_id");
