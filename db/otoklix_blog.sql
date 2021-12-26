CREATE TABLE "blog_post" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"title"	TEXT,
	"content"	TEXT,
	"published_at"	TEXT,
	"created_at"	TEXT,
	"updated_at"	TEXT
);

INSERT INTO blog_post (title, content, published_at, created_at, updated_at) VALUES ('post-01','content post-01','2021-12-25 10:10','2021-12-25 10:10','2021-12-25 10:10');
INSERT INTO blog_post (title, content, published_at, created_at, updated_at) VALUES ('post-02','content post-02','2021-12-26 10:10','2021-12-26 10:10','2021-12-26 10:10');
INSERT INTO blog_post (title, content, published_at, created_at, updated_at) VALUES ('post-03','content post-03','2021-12-27 10:10','2021-12-27 10:10','2021-12-27 10:10');
