--  ID        uint64    `json:"id,omitempty"`
-- 	Name      string    `json:"name,omitempty"`
-- 	Email     string    `json:"email,omitempty"`
-- 	Skill     string    `json:"skill,omitempty"`
-- 	Sentiment string    `json:"sentiment,omitempty"`
-- 	CreatedAt time.Time `json:"createdAt,omitempty"`
-- 	UpdateAt  time.Time `json:"updatedAt,omitempty"`

CREATE TABLE IF NOT EXISTS "analysts" (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "skill" VARCHAR(255),
    "sentiment" VARCHAR(255),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
