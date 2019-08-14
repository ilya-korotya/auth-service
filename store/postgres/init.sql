/* File for initial postgres database */

CREATE DATABASE auth;

CREATE USER auth_user_role;

/* Set custom access to auth_user_role role */
GRANT ALL PRIVILEGES ON DATABASE auth TO auth_user_role;

REVOKE connect ON DATABASE auth FROM PUBLIC;

GRANT connect ON DATABASE auth TO auth_user_role;
