CREATE TABLE IF NOT EXISTS public.admins
(
    id bigserial NOT NULL,
    username character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE(username)
);

INSERT INTO admins (username,first_name,last_name,password) VALUES ('admin', 'admin', 'admin', '$2a$10$Z69F.NaAsKtMN/5WWgVXAu8GIcXRIUWQj4.Zdn5T8i94dSKKsFlFW');