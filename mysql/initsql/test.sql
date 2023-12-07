CREATE TABLE tests (
    event_id INT NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id),
    rate TINYINT UNSIGNED NOT NULL,
)
