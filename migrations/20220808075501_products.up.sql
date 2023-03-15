CREATE TABLE IF NOT EXISTS public.products
(
    id bigserial NOT NULL,
    admin_id int NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id,admin_id),
    UNIQUE(id),
    FOREIGN KEY (admin_id) REFERENCES admins(id) ON DELETE CASCADE
);