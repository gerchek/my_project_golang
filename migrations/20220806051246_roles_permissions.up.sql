CREATE TABLE IF NOT EXISTS public.roles_permissions (
  role_id int NOT NULL,
  permission_id int NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (role_id, permission_id),
  FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE,
  FOREIGN KEY (permission_id) REFERENCES permissions(id) ON UPDATE CASCADE
);

