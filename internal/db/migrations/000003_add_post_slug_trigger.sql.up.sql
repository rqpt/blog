CREATE OR REPLACE FUNCTION posts_set_slug()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.id IS NULL THEN
        NEW.id := nextval(
            pg_get_serial_sequence('posts', 'id')::regclass
        );
    END IF;

    NEW.slug := trim(BOTH '-' FROM
        regexp_replace(
            lower(NEW.title),
            '[^a-z0-9]+',
            '-',
            'g'
        )
    ) || '-' || NEW.id::text;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER posts_set_slug_trigger
BEFORE INSERT ON posts
FOR EACH ROW
EXECUTE FUNCTION posts_set_slug();
