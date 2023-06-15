CREATE ROLE user;
CREATE ROLE admin;

GRANT SELECT ON movies TO user;
GRANT SELECT ON actors TO user;
GRANT SELECT ON awards TO user;
GRANT SELECT ON nominations TO user;
GRANT SELECT ON performances TO user;
GRANT SELECT ON nominated_performances TO user;

GRANT ALL PRIVILEGES ON movies TO admin;
GRANT ALL PRIVILEGES ON actors TO admin;
GRANT ALL PRIVILEGES ON awards TO admin;
GRANT ALL PRIVILEGES ON nominations TO admin;
GRANT ALL PRIVILEGES ON performances TO admin;
GRANT ALL PRIVILEGES ON nominated_performances TO admin;