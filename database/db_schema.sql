create database workhorse;

CREATE OR REPLACE FUNCTION notify_event()
    RETURNS trigger
    LANGUAGE plpgsql
AS
$function$

DECLARE
    data         json;
    notification json;

BEGIN

    -- Convert the old or new row to JSON, based on the kind of action.
    -- Action = DELETE?             -> OLD row
    -- Action = INSERT or UPDATE?   -> NEW row
    IF (TG_OP = 'DELETE') THEN
        data = row_to_json(OLD);
    ELSE
        data = row_to_json(NEW);
    END IF;

    -- Contruct the notification as a JSON string.
    notification = json_build_object(
            'table', TG_TABLE_NAME,
            'action', TG_OP,
            'data', data);


    -- Execute pg_notify(channel, notification)
    PERFORM pg_notify('events', notification::text);

    -- Result is ignored since this is an AFTER trigger
    RETURN NULL;
END;
$function$
;

CREATE TABLE project (
    id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
    "name" varchar(200) NOT NULL,
    private_key text NOT NULL,
    clone_url varchar NOT NULL,
    CONSTRAINT project_un UNIQUE (name),
    CONSTRAINT project_pk PRIMARY KEY (id)
);


CREATE TABLE project_branches (
  id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
  "name" varchar NOT NULL,
  project_id int4 NOT NULL,
  CONSTRAINT project_branches_pk PRIMARY KEY (id),
  CONSTRAINT project_branches_fk FOREIGN KEY (project_id) REFERENCES project(id)
);


CREATE TABLE IF NOT EXISTS build
(
    id         int4        NOT NULL GENERATED ALWAYS AS IDENTITY,
    status     varchar(30) NULL,
    project_id int4        NULL,
    created_ts timestamp   NULL,
    start_ts   timestamp   NULL,
    end_ts     timestamp   NULL
);

-- Table Triggers
drop trigger if exists build_notify_event on build;
create trigger build_notify_event
    after
        insert
        or
        delete
        or
        update
    on
        build
    for each row
execute function notify_event();


CREATE TABLE IF NOT EXISTS build_steps
(
    id         int4         NOT NULL GENERATED ALWAYS AS IDENTITY,
    build_id   int4,
    name       varchar(255) NOT NULL,
    image      varchar(255) NOT NULL,
    status     varchar(30)  NULL,
    created_ts timestamp    NULL,
    start_ts   timestamp    NULL,
    end_ts     timestamp    NULL,
    log_info jsonb NULL,
    node_id int4 NULL
);

drop trigger if exists build_steps_notify_event on build_steps;
create trigger build_steps_notify_event
    after
        insert
        or
        delete
        or
        update
    on
        build_steps
    for each row
execute function notify_event();

CREATE TABLE IF NOT EXISTS build_steps_command
(
    id      int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
    step_id int4,
    command text
);

CREATE TABLE IF NOT EXISTS build_node_binding
(
    id       int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
    build_id int4,
    node_id  integer
)

drop trigger if exists build_node_binding_notify_event on build_node_binding;
create trigger build_node_binding_notify_event
    after
        insert
        or
        delete
        or
        update
    on
        build_node_binding
    for each row
execute function notify_event();

CREATE TABLE IF NOT EXISTS build_step_node_binding
(
    id         int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
    step_id    int4,
    ip_address varchar(255)
)

drop trigger if exists build_step_node_binding_notify_event on build_step_node_binding;
create trigger build_step_node_binding_notify_event
    after
        insert
        or
        delete
        or
        update
    on
        build_step_node_binding
    for each row
execute function notify_event();


CREATE TABLE IF NOT EXISTS node_info
(
    id              integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    name            varchar(255) unique,
    last_heart_beat timestamp
)
