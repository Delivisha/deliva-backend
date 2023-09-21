-- +goose Up
CREATE TABLE departments (
    id                  text NOT NULL,
    name                text NOT NULL,
    head_of_department  text NOT NULL,
    number_of_employees text NULL,
    budget              numeric NULL,
    created_at          timestamptz   NOT NULL DEFAULT NOW(),
    updated_at          timestamptz   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TABLE employees (
    id              text NOT NULL,
    dept_id         text NOT NULL,
    first_name      text NOT NULL,
    middle_name     text NOT NULL,
    last_name       text NOT NULL,
    gender          text NOT NULL,
    email           text NOT NULL,
    date_of_birth   date NOT NULL,
    home_address    text NOT NULL,
    phone_number    text NOT NULL,
    bank            text NOT NULL,
    account_number  text NOT NULL,
    gross_salary    numeric NOT NULL,
    next_of_kin_name text NOT NULL,
    next_of_kin_phone text NOT NULL,
    reference_name   text NOT NULL,
    reference_phone text NOT NULL,
    country         text NOT NULL,
    suspended       bool NOT NULL,
    suspension_reason text NULL,
    suspension_duration text NULL,
    sacked          bool NOT NULL,
    sacked_reason   text NULL,
    password        text NOT NULL,
    created_at      timestamptz NOT NULL DEFAULT NOW(),
    updated_at      timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id),
    CONSTRAINT FK_employee_department FOREIGN KEY(dept_id)
        REFERENCES departments(id)
);
CREATE TRIGGER created_at_employees_trgr
    BEFORE UPDATE
    ON employees
    FOR EACH ROW
    EXECUTE PROCEDURE created_at_trigger();
CREATE TRIGGER updated_at_employees_trgr
    BEFORE UPDATE
    ON employees
    FOR EACH ROW
    EXECUTE PROCEDURE updated_at_trigger();


CREATE TABLE projects (
    id              text NOT NULL,
    dept_id         text NOT NULL,
    name            text NOT NULL,
    duration        text NOT NULL,
    employee_id     text NOT NULL,
    budget          numeric NULL,
    resource        text NOT NULL,
    created_at  timestamptz   NOT NULL DEFAULT NOW(),
    updated_at  timestamptz   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT FK_project_department FOREIGN KEY(dept_id)
        REFERENCES departments(id)
);


CREATE TABLE events (
    stream_id      text        NOT NULL,
    stream_name    text        NOT NULL,
    stream_version int         NOT NULL,
    event_id       text        NOT NULL,
    event_name     text        NOT NULL,
    event_data     bytea       NOT NULL,
    occurred_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (stream_id, stream_name, stream_version)
);
CREATE TABLE snapshots (
    stream_id      text        NOT NULL,
    stream_name    text        NOT NULL,
    stream_version int         NOT NULL,
    snapshot_name  text        NOT NULL,
    snapshot_data  bytea       NOT NULL,
    updated_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (stream_id, stream_name)
);

CREATE TRIGGER updated_at_snapshots_trgr
    BEFORE UPDATE
    ON snapshots
    FOR EACH ROW
EXECUTE PROCEDURE updated_at_trigger();

CREATE TABLE inbox (
    id          text        NOT NULL,
    name        text        NOT NULL,
    subject     text        NOT NULL,
    data        bytea       NOT NULL,
    metadata    bytea       NOT NULL,
    sent_at     timestamptz NOT NULL,
    received_at timestamptz NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE outbox (
    id           text        NOT NULL,
    name         text        NOT NULL,
    subject      text        NOT NULL,
    data         bytea       NOT NULL,
    metadata     bytea       NOT NULL,
    sent_at      timestamptz NOT NULL,
    published_at timestamptz,
    PRIMARY KEY (id)
);
CREATE INDEX unpublished_idx ON outbox (published_at) WHERE published_at IS NULL;

CREATE TABLE sagas (
    id           text        NOT NULL,
    name         text        NOT NULL,
    data         bytea       NOT NULL,
    step         int         NOT NULL,
    done         bool        NOT NULL,
    compensating bool        NOT NULL,
    updated_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, name)
);
CREATE TRIGGER updated_at_sagas_trgr
    BEFORE UPDATE
    ON sagas
    FOR EACH ROW
EXECUTE PROCEDURE updated_at_trigger();

-- +goose Down
DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS employees;
DROP TABLE IF EXISTS projects;

DROP TABLE IF EXISTS sagas;
DROP TABLE IF EXISTS outbox;
DROP TABLE IF EXISTS inbox;
DROP TABLE IF EXISTS snapshots;
DROP TABLE IF EXISTS events;