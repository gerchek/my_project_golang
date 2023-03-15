CREATE TABLE IF NOT EXISTS public.roles
(
    id bigserial NOT NULL,
    name character varying(70) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);