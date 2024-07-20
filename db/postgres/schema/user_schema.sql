
CREATE TABLE
    user_task (
        wallet_address varchar(80) NOT NULL PRIMARY KEY,
        amount bigint NOT NULL,
        point int NOT NULL,
        event_name varchar(60) NOT NULL,
        create_time timestamptz NOT NULL,
        update_time timestamptz NOT NULL
    );
