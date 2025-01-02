CREATE TABLE IF NOT EXISTS "folder"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       VARCHAR  NOT NULL,
    "pid"        INTEGER  NOT NULL,
    "cover_fid"  INTEGER,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "folder_index_0"
    ON "folder" ("id");


CREATE TABLE IF NOT EXISTS "file"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "folder_id"  INTEGER,
    "name"       TEXT,
    "desc"       TEXT,
    "md5"        VARCHAR,
    "updated_at" DATETIME,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "file_index_0"
    ON "file" ("id");

CREATE TABLE IF NOT EXISTS "tag"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       TEXT     NOT NULL,
    "pid"        INTEGER,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "tag_index_0"
    ON "tag" ("id");

CREATE TABLE IF NOT EXISTS "file_tag"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "file_id"    INTEGER,
    "tag_id"     INTEGER,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "label"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "key"        TEXT     NOT NULL,
    "value"      TEXT,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "label_index_0"
    ON "tag" ("id");


CREATE TABLE IF NOT EXISTS "label"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "fid"        INTEGER  not null,
    "key"        TEXT     NOT NULL,
    "value"      TEXT,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "label_index_0"
    ON "label" ("id");