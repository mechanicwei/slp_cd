CREATE TABLE deploy_repos (
    id serial PRIMARY KEY,
    name VARCHAR (50),
    github_url VARCHAR (255),
    webhook_secret VARCHAR (255),
    openids text,
    created_at timestamp without time zone
);

ALTER TABLE deploy_servers ADD deploy_repo_id BIGINT;
CREATE INDEX index_deploy_servers_on_deploy_repo_id on deploy_servers (deploy_repo_id);
