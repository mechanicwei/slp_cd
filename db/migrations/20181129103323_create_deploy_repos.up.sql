CREATE TABLE deploy_repos (
    id serial PRIMARY KEY,
    name VARCHAR (50),
    github_url VARCHAR (255),
    webhook_secret VARCHAR (255),
    openids text[] DEFAULT '{}'::text[],
    created_at timestamp without time zone
);
