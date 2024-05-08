CREATE TABLE logs (
    time TIMESTAMP,
    level VARCHAR(32),
    message TEXT
);

CREATE INDEX logs_time ON logs(time);
