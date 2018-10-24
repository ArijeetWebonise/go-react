
-- +migrate Up

insert into USERS (email, first_name, password, modified_at, created_at) VALUES ('admin@admin', 'admin', '$2y$12$8GkZQ/TXDYo2p1weK.GV8eEYj2xCsqAQC4BFEZnHtNdE/oR6g0YGi', now(), now());

-- +migrate Down
