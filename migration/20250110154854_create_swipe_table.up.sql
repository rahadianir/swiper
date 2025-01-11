CREATE TABLE IF NOT EXISTS "swipes" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "target_id" INTEGER NOT NULL,
    "is_matched" BOOLEAN NOT NULL DEFAULT false,
    "created_at" TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) 
	  REFERENCES users(id)
	  ON DELETE CASCADE,
    CONSTRAINT fk_target_id FOREIGN KEY(target_id) 
	  REFERENCES users(id)
	  ON DELETE CASCADE
);