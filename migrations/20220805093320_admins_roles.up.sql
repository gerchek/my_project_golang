CREATE TABLE IF NOT EXISTS public.admins_roles (
  admin_id int NOT NULL,
  role_id int NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (admin_id, role_id),
  FOREIGN KEY (admin_id) REFERENCES admins(id) ON UPDATE CASCADE,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE
);

