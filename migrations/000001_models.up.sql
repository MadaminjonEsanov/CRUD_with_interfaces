CREATE TABLE users (

    userId       UUID        PRIMARY KEY, 
    fullname     VARCHAR,
    username     VARCHAR   NOT NULL,
    passwords    VARCHAR   NOT NULL,
    gmail        VARCHAR   NOT NULL
);


CREATE TABLE todos(
    todoId      UUID        PRIMARY KEY,
    task        VARCHAR     NOT NULL,
    created_at  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    is_completed BOOLEAN    DEFAULT false
);