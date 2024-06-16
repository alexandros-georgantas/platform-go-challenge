-- create required ENUM types (https://stackoverflow.com/questions/7624919/check-if-a-user-defined-type-already-exists-in-postgresql)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typename = 'gender_type') THEN
        CREATE TYPE gender_type AS ENUM ('MALE', 'FEMALE');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typename = 'age_group') THEN
        CREATE TYPE age_group AS ENUM ('CHILDREN', 'YOUTH', 'ADULTS', 'SENIORS');
    END IF;
END$$;