CREATE TABLE IF NOT EXISTS deploy_servers (
    id serial PRIMARY KEY,
    name VARCHAR (50),
    branch VARCHAR (50),
    dir VARCHAR (255),
    cmd VARCHAR (255),
    created_at timestamp without time zone
);

INSERT INTO deploy_servers
    (name, branch, dir, cmd, created_at)
SELECT 'skylark', 'master', 'test', 'test', '2018-10-02 18:40:25'
WHERE
    NOT EXISTS (
        SELECT name FROM deploy_servers WHERE name = 'skylark'
    );
