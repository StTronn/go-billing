CREATE TABLE billing_metric (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    aggregation ENUM('SUM', 'COUNT') NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);