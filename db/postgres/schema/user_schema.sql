/*
    Database Table Relationship

table: task


table: task_user
+++ task_seqno <-> task.seqno

*/

CREATE TABLE
    task (
        seqno serial NOT NULL PRIMARY KEY,
        task_group_no int NOT NULL,
        task_name varchar(60) NOT NULL,
        task_desc varchar(255) NOT NULL,
        start_time timestamptz NOT NULL,
        end_time timestamptz NOT NULL
    );

CREATE TABLE
    user_task (
        task_seqno int NOT NULL,
        wallet_address varchar(255) NOT NULL,
        total_amount bigint NOT NULL,
        point int NOT NULL,
        status varchar(30) NOT NULL,
        create_time timestamptz NOT NULL,
        update_time timestamptz NOT NULL
    );
