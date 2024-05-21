DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'dbadmin') THEN
   ELSE
      CREATE ROLE dbadmin LOGIN PASSWORD 'dbadmin';
   END IF;
END
$do$;

CREATE TABLE user_tb (
         id serial PRIMARY KEY,
         name VARCHAR(255),
         detail VARCHAR(255),
         role VARCHAR(255)
      );

ALTER TABLE user_tb OWNER TO dbadmin;

INSERT INTO user_tb (name, detail, role) VALUES ('kirisagi chitoge', 'real heroine', 'mc');
INSERT INTO user_tb (name, detail, role) VALUES ('onodera kosaki' , 'unknow', 'sc');
-- SELECT 'CREATE DATABASE labdb'
-- WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'labdb')\gexec

-- ALTER DATABASE labdb OWNER TO dbadmin;
-- GRANT ALL PRIVILEGES ON DATABASE labdb TO dbadmin;

-- BEGIN
--    IF EXISTS (
--       SELECT FROM pg_catalog.pg_tables
--       WHERE schemaname = 'public'
--       AND 
--    ) THEN
--    ELSE
--       CREATE TABLE labdb.public.user_tb (
--          id serial PRIMARY KEY,
--          name VARCHAR(255),
--          detail VARCHAR(255),
--          role VARCHAR(255)
--       );
--    END IF;
-- END
