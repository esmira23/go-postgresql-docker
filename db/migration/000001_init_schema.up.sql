CREATE TABLE example (
    "transaction_id" int PRIMARY KEY NOT NULL,
    "request_id" int NOT NULL,
    "terminal_id" int NOT NULL,
    "partner_object_id" int NOT NULL,
    "amount_total" float NOT NULL,
    "amount_original" float NOT NULL,
    "commission_ps" float NOT NULL,
    "commission_client" float NOT NULL,
    "commission_provider" float NOT NULL,
    "date_input" timestamp NOT NULL,
    "date_post" timestamp NOT NULL,
    "status" varchar(255) NOT NULL,
    "payment_type" varchar(255) NOT NULL,
    "payment_number" varchar(255) NOT NULL,
    "service_id" int NOT NULL,
    "service" varchar(255) NOT NULL,
    "payee_id" bigint NOT NULL,
    "payee_name" varchar(255) NOT NULL,
    "payee_bank_mfo" bigint NOT NULL,
    "payee_bank_account" varchar(255) NOT NULL,
    "payment_narrative" varchar(255) NOT NULL
);