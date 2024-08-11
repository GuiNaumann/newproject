CREATE TABLE user (
                         id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         status_code TINYINT NOT NULL DEFAULT 0,
                         created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         modified_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)